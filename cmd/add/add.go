package add

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Cmd cli.Command = cli.Command{
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
}
