package hello

import (
	"fmt"

	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use: "hello",
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("Hello %s!\n", FlagName)
	},
}

func init() {
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"World",
		"Name to say hello to",
	)
}
