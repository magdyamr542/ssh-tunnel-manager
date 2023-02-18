package list

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Cmd cli.Command = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "List all configurations",
	Action: func(cCtx *cli.Context) error {
		fmt.Println("listing all configurations", cCtx.Args().First())
		return nil
	},
}
