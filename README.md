## SSH tunnel manager
Save SSH tunnel configurations and start a tunnel with port forwarding using one of the saved configurations.

## Features
- Save an SSH tunnel configuration with a description
- Start an SSH tunnel with port forwarding using a configuration. Same as `ssh -L [LOCAL_IP:]LOCAL_PORT:DESTINATION:DESTINATION_PORT [USER@]SSH_SERVER`
- List all configurations
- Remove a configuration

## Install 
`go install github.com/magdyamr542/ssh-tunnel-manager@latest`

## Usage
Run `ssh-tunnel-manager`

## Example
- Adding a configuration for a tunnel to a development cluster
    ```
    ssh-tunnel-manager add \
    --name my-dev-tunnel \
    --description "A tunnel to access the db in dev cluster" \
    --server my-ssh.server.com \
    --user me \
    --keyFile ~/.ssh/key \
    --remoteHost my-db.remote.com \
    --remotePort 5432
    ```
- Listing the configurations
    ```
    ssh-tunnel-manager list

    # output
    my-dev-tunnel (A tunnel to access the db in dev cluster)
      - SSH server:  my-ssh.server.com 
      - User:        me
      - Private key: <your-.ssh-path>/key
      - Remote:      my-db.remote.com:5432
    ```
- Start tunneling using a configuration. Now you can access your remote db from `localhost:1234`
    ```
    ssh-tunnel-manager tunnel my-dev-tunnel 1234

    # output
    2023/02/20 09:19:32 Tunneling "localhost:1234" <==> "my-db.remote.com:5432"
    ```
