package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "ssh-tool-manager",
		Usage:       "An SSH tunnel manager tool",
		Description: "Save SSH tunnel configurations and start a tunnel using one of the saved configurations.",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all configurations",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("listing all configurations", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Add a configuration",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Name of the configuration",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "description",
						Usage: "Short description of the configuration (optional)",
					},
					&cli.StringFlag{
						Name:     "server",
						Usage:    "Name of the SSH server to connect to",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "user",
						Usage:    "Username to use when connecting",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "keyFile",
						Usage:    "Path to the private key file for SSH authentication",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "localPort",
						Usage:    "The localport to listen on",
						Required: true,
						Value:    1234,
					},
					&cli.StringFlag{
						Name:     "remoteHost",
						Usage:    "The remote host to forward traffic to",
						Required: true,
					},
					&cli.IntFlag{
						Name:     "remotePort",
						Usage:    "The remote port to forward traffic to",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("adding a configuration: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "Remove a configuration",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Value:    "",
						Usage:    "Name of the configuration to remove",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					configName := cCtx.String("name")
					fmt.Printf("removing configuration %s", configName)
					return nil
				},
			},
			{
				Name:  "tunnel",
				Usage: "Start a tunnel using a configuration",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Value:    "",
						Usage:    "Name of the configuration to use",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					configName := cCtx.String("name")
					fmt.Printf("tunneling configuration %s", configName)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
