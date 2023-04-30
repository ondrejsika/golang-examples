package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("EXAMPLE")
	viper.SetDefault("TOKEN", "default")
	viper.BindEnv("TOKEN")

	token := viper.GetString("TOKEN")
	fmt.Println(token)
}
