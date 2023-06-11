package autocomplete

import (
	"fmt"
	"os"

	"github.com/posener/complete/v2/install"
	"github.com/urfave/cli/v2"
)

var Cmd cli.Command = cli.Command{
	Name:  "install-autocomplete",
	Usage: "Install the autocompletion for your shell",
	Action: func(cCtx *cli.Context) error {
		err := install.Install("ssh-tunnel-manager")
		if err != nil {
			fmt.Fprintf(os.Stdout, "Installed autocompletion successfully. Please reload your shell and check if the changes made make sense!\n")
		}
		return err
	},
}
