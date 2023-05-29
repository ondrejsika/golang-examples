package hello

import (
	"fmt"

	"viper_and_cobra_example_advanced/cmd/root"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagMessage string

var Cmd = &cobra.Command{
	Use:  "hello",
	Args: cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(FlagMessage)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	viper.BindEnv("MESSAGE")
	message := viper.GetString("MESSAGE")
	Cmd.Flags().StringVarP(
		&FlagMessage,
		"message",
		"m",
		message,
		"Example message",
	)
	if message == "" {
		Cmd.MarkFlagRequired("message")
	}
}
