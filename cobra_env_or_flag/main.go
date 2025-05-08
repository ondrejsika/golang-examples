package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var FlagToken string

var Cmd = &cobra.Command{
	Use:   "cobra_env_or_flag",
	Short: "Example of using environment variables or required flags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Token: %s\n", FlagToken)
	},
}

func init() {
	token := os.Getenv("TOKEN")
	Cmd.Flags().StringVarP(
		&FlagToken,
		"token",
		"t",
		token,
		"Token",
	)
	if token == "" {
		Cmd.MarkFlagRequired("token")
	}
}

func main() {
	Cmd.Execute()
}
