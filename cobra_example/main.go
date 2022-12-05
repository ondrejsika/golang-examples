package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use:   "cobra_example",
	Short: "Hello World example for Cobra",
	Run: func(cmd *cobra.Command, args []string) {
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

func main() {
	Cmd.Execute()
}
