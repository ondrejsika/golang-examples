package cmd

import (
	_ "foo_example/cmd/hello"
	"foo_example/cmd/root"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
