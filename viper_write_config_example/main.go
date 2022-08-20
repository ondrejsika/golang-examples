package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func touch(filePath string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

func main() {
	touch("config.yml")

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.SetEnvPrefix("EXAMPLE")
	viper.BindEnv("TOKEN")

	token := viper.GetString("TOKEN")
	viper.Set("TOKEN", token)

	fmt.Println(token)

	viper.WriteConfig()
}
