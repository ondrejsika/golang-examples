package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("mon")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.SetEnvPrefix("mon")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.BindEnv("notifications.mail.smtp_host")
	viper.BindEnv("notifications.mail.smtp_port")
	viper.BindEnv("notifications.telegram.token")

	fmt.Println("notifications.mail.smtp_host", viper.GetString("notifications.mail.smtp_host"))
	fmt.Println("notifications.mail.smtp_port", viper.GetString("notifications.mail.smtp_port"))
	fmt.Println("notifications.telegram.token", viper.GetString("notifications.telegram.token"))
}
