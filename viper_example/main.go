package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.SetEnvPrefix("EXAMPLE")
	viper.BindEnv("TOKEN")

	token := viper.GetString("TOKEN")
	fmt.Println(token)
}
