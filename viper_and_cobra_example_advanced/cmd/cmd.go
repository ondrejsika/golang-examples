package cmd

import (
	_ "github.com/ondrejsika/golang-examples/viper_and_cobra_example_advanced/cmd/cmd/hello"
	"github.com/ondrejsika/golang-examples/viper_and_cobra_example_advanced/cmd/cmd/root"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
