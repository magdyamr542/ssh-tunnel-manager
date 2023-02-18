package main

import (
	"log"
	"os"

	"github.com/magdyamr542/ssh-tunnel-manager/cmd/add"
	"github.com/magdyamr542/ssh-tunnel-manager/cmd/list"
	"github.com/magdyamr542/ssh-tunnel-manager/cmd/remove"
	"github.com/magdyamr542/ssh-tunnel-manager/cmd/tunnel"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "ssh-tool-manager",
		Usage:       "An SSH tunnel manager tool",
		Description: "Save SSH tunnel configurations and start a tunnel using one of the saved configurations.",
		Commands: []*cli.Command{
			&list.Cmd,
			&add.Cmd,
			&remove.Cmd,
			&tunnel.Cmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
