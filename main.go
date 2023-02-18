package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/signal"

	"golang.org/x/crypto/ssh"
)

func main() {
	// The SSH server to connect to
	sshServer := "turntable.xtk-stage.de"

	// The username to use when connecting
	sshUser := "amrm"

	// The private key file to use for authentication
	keyFile := "/home/amr/.ssh/turntable"

	// The local port to listen on
	localPort := "1235"

	// The remote host and port to forward traffic to
	remoteHost := "invoice-manager.cluster-cvlsayhdpvet.eu-central-1.rds.amazonaws.com:5432"

	// Load the private key file
	key, err := ssh.ParsePrivateKey(readPrivateKeyFile(keyFile))
	if err != nil {
		fmt.Printf("Error parsing private key: %s\n", err.Error())
		return
	}

	// Set up the SSH client config
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", sshServer+":22", config)
	if err != nil {
		fmt.Printf("Error connecting to SSH server: %s\n", err.Error())
		return
	}

	// Set up the local listener
	localListener, err := net.Listen("tcp", "localhost:"+localPort)
	if err != nil {
		fmt.Printf("Error setting up local listener: %s\n", err.Error())
		return
	}

	// Handle server shutdown
	stopper := make(chan os.Signal)
	signal.Notify(stopper, os.Kill, os.Interrupt)
	go func(stopper <-chan os.Signal) {
		<-stopper
		fmt.Printf("Stopping the tunnel server. Close existing connections...\n")
		code := 0
		if err := client.Close(); err != nil {
			fmt.Printf("Eror stopping the tunnel server: %v", err)
			code = 1
		}
		os.Exit(code)
	}(stopper)

	// Start accepting connections on the local listener
	fmt.Printf("Tunnel listening on %s:%s\n", "localhost", *&localPort)
	for {
		localConn, err := localListener.Accept()
		if err != nil {
			fmt.Printf("Error accepting local connection: %s\n", err.Error())
			return
		}
		fmt.Printf("New local connections %s\n", localConn.RemoteAddr())

		// Start the SSH tunnel for each incoming connection
		go func(localConn net.Conn) {
			remoteConn, err := client.Dial("tcp", remoteHost)
			if err != nil {
				fmt.Printf("Error connecting to remote host: %s\n", err.Error())
				return
			}

			fmt.Printf("Connection open: %s\n", localConn.RemoteAddr())
			go runTunnel(localConn, remoteConn)
		}(localConn)
	}
}

// Helper function to read the private key file
func readPrivateKeyFile(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading private key file: %s\n", err.Error())
		return nil
	}
	return data
}

// runTunnel runs a tunnel between two connections; as soon as one connection
// reaches EOF or reports an error, both connections are closed and this
// function returns.
func runTunnel(local, remote net.Conn) {
	// Clean up
	defer local.Close()
	defer remote.Close()

	done := make(chan struct{}, 2)
	go func() {
		io.Copy(local, remote)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(remote, local)
		done <- struct{}{}
	}()

	<-done
	fmt.Printf("Connection closed: %s\n", local.RemoteAddr())
}
