package cmd

import (
	"fmt"
	"os"

	"github.com/oam-dev/kubevela/references/cli"
	"github.com/oam-dev/kubevela/version"
	"github.com/oam-dev/velad/pkg/utils"
	veladVersion "github.com/oam-dev/velad/version"
	"github.com/spf13/cobra"
)

// App is entry of all CLI, created by NewApp
type App struct {
	args []string
}

// NewApp create app
func NewApp() App {
	app := App{args: os.Args}
	return app
}

func (a App) Run() {
	if len(a.args) == 0 {
		fmt.Println("No args")
		os.Exit(1)
	}

	var cmd *cobra.Command
	if utils.IsVelaCommand(a.args[0]) {
		utils.SetDefaultKubeConfigEnv()
		cmd = cli.NewCommand()
		// TODO set right gitVersion
		version.VelaVersion = veladVersion.VelaVersion
	} else {
		cmd = NewVeladCommand()
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
