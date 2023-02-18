package remove

import (
	"fmt"

	"github.com/magdyamr542/ssh-tunnel-manager/cmd/add"
	"github.com/magdyamr542/ssh-tunnel-manager/configmanager"
	"github.com/magdyamr542/ssh-tunnel-manager/utils"
	"github.com/urfave/cli/v2"
)

var Cmd cli.Command = cli.Command{
	Name:    "remove",
	Usage:   "Remove a configuration",
	Aliases: []string{"rm"},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Value:    "",
			Usage:    "Name of the configuration to remove",
			Required: true,
		},
	},
	Action: func(cCtx *cli.Context) error {
		configdir, err := utils.ResolveDir(cCtx.String(add.ConfigDirFlagName))
		if err != nil {
			return err
		}
		entryName := cCtx.String("name")
		err = configmanager.NewManager(configdir).RemoveConfiguration(entryName)
		if err != nil {
			return fmt.Errorf("couln't remove configuration %s: %v", entryName, err)
		}
		return nil
	},
}
