package tunnel

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Cmd cli.Command = cli.Command{

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
}
