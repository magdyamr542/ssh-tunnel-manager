package list

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/magdyamr542/ssh-tunnel-manager/cmd/add"
	"github.com/magdyamr542/ssh-tunnel-manager/configmanager"
	"github.com/magdyamr542/ssh-tunnel-manager/utils"
	"github.com/urfave/cli/v2"
)

var Cmd cli.Command = cli.Command{
	Name:    "list",
	Aliases: []string{"l", "ls"},
	Usage:   "List configurations",
	Action: func(cCtx *cli.Context) error {
		configdir, err := utils.ResolveDir(cCtx.String(add.ConfigDirFlagName))
		if err != nil {
			return err
		}
		cfgs, err := configmanager.NewManager(configdir).GetConfigurations()
		if err != nil {
			return fmt.Errorf("couldn't get saved configurations: %v", err)
		}

		if len(cfgs) == 0 {
			fmt.Println("No configurations found")
			return nil
		}
		for i, cfg := range cfgs {
			printConfig(os.Stdout, cfg)
			if i != len(cfgs)-1 {
				fmt.Println("")
			}
		}
		return nil
	},
}

func printConfig(w io.Writer, entry configmanager.Entry) {
	template := `%s
  - SSH server:  %s
  - Private key: %s
  - Remote:      %s:%d
`
	nameAndDesc := entry.Name
	if strings.TrimSpace(entry.Description) != "" {
		nameAndDesc += " " + "(" + entry.Description + ")"
	}
	w.Write([]byte(
		fmt.Sprintf(
			template,
			nameAndDesc,
			entry.Server,
			entry.KeyFile,
			entry.RemoteHost,
			entry.RemotePort,
		)))
}
