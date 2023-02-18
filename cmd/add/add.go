package add

import (
	"github.com/magdyamr542/ssh-tunnel-manager/configmanager"
	"github.com/magdyamr542/ssh-tunnel-manager/utils"
	"github.com/urfave/cli/v2"
)

const ConfigDirFlagName = "config-dir"

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
		configdir, err := utils.ResolveDir(cCtx.String(ConfigDirFlagName))
		if err != nil {
			return err
		}
		e := configmanager.Entry{
			Name:        cCtx.String("name"),
			Description: cCtx.String("description"),
			Server:      cCtx.String("server"),
			User:        cCtx.String("user"),
			KeyFile:     cCtx.String("keyFile"),
			LocalPort:   cCtx.Int("localPort"),
			RemoteHost:  cCtx.String("remoteHost"),
			RemotePort:  cCtx.Int("remotePort"),
		}
		return configmanager.NewManager(configdir).AddConfiguration(e)
	},
}
