package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FlagToken string

var Cmd = &cobra.Command{
	Use:  "example",
	Args: cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(FlagToken)
	},
}

func init() {
	viperInit()

	viper.BindEnv("TOKEN")
	token := viper.GetString("TOKEN")
	Cmd.Flags().StringVarP(
		&FlagToken,
		"token",
		"t",
		token,
		"Example token",
	)
	if token == "" {
		Cmd.MarkFlagRequired("token")
	}
}

func viperInit() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("EXAMPLE")
	viper.ReadInConfig()
}

func main() {
	Cmd.Execute()
}
