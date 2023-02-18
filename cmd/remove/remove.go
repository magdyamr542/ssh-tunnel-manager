package remove

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Cmd cli.Command = cli.Command{
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
}