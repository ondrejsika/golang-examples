package root

import (
	"bar_example/cmd/hello"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "bar_example",
	Short: "bar_example",
}

func init() {
	Cmd.AddCommand(hello.Cmd)
}
