package list

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/magdyamr542/ssh-tunnel-manager/cmd/add"
	"github.com/magdyamr542/ssh-tunnel-manager/configmanager"
	"github.com/magdyamr542/ssh-tunnel-manager/utils"
	"github.com/posener/complete/v2"
	"github.com/urfave/cli/v2"
)

var Predictor complete.Predictor = complete.PredictFunc(predictConfigurations)

var Cmd cli.Command = cli.Command{
	Name:    "list",
	Aliases: []string{"l", "ls"},
	Usage:   "List configurations",
	Action: func(cCtx *cli.Context) error {
		cfgs, err := getConfigs()
		if err != nil {
			return err
		}

		output := os.Stdout

		if len(cfgs) == 0 {
			fmt.Fprintf(output, "No configurations found\n")
			return nil
		}
		for i, cfg := range cfgs {
			printConfig(output, cfg)
			if i != len(cfgs)-1 {
				fmt.Fprintf(output, "")
			}
		}
		return nil
	},
}

func getConfigs() ([]configmanager.Entry, error) {
	dirpath := configmanager.DefaultConfigDir
	if value := os.Getenv(add.ConfigDirFlagName); value != "" {
		dirpath = value
	}

	configdir, err := utils.ResolveDir(dirpath)
	if err != nil {
		return nil, err
	}

	cfgs, err := configmanager.NewManager(configdir).GetConfigurations()
	if err != nil {
		return nil, fmt.Errorf("couldn't get saved configurations: %v", err)
	}

	if len(cfgs) == 0 {
		return nil, nil
	}

	return cfgs, nil
}

func predictConfigurations(prefix string) []string {
	cfgs, err := getConfigs()
	if err != nil {
		return nil
	}
	configs := make([]string, len(cfgs))
	for _, cfg := range cfgs {
		configs = append(configs, cfg.Name)
	}
	return configs
}

func printConfig(w io.Writer, entry configmanager.Entry) {
	template := `%s
  - SSH server:  %s
  - User:        %s
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
			entry.User,
			entry.KeyFile,
			entry.RemoteHost,
			entry.RemotePort,
		)))
}
