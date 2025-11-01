package cmd

import (
	"bar_example/cmd/root"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
