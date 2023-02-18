## SSH tunnel manager
Save SSH tunnel configurations and start a tunnel using one of the saved configurations.

## Features
- Save an SSH tunnel configuration with a description
- Start an SSH tunnel with port forwarding using a configuration. Same as `ssh -L [LOCAL_IP:]LOCAL_PORT:DESTINATION:DESTINATION_PORT [USER@]SSH_SERVER`
- List all configurations
- Remove a configuration

## Install 
`go install github.com/magdyamr542/ssh-tunnel-manager@latest`

## Usage
Run `ssh-tunnel-manager`
