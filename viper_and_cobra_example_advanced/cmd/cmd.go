package cmd

import (
	_ "viper_and_cobra_example_advanced/cmd/hello"
	"viper_and_cobra_example_advanced/cmd/root"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
