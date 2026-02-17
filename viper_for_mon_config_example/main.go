package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	// mon.yaml schema
	MON_NOTIFICATIONS_EMAIL_SMTP_HOST = "notifications.mail.smtp_host"
	MON_NOTIFICATIONS_EMAIL_SMTP_PORT = "notifications.mail.smtp_port"
	MON_NOTIFICATIONS_TELEGRAM_TOKEN  = "notifications.telegram.token"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("mon")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.SetEnvPrefix("mon")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_HOST)
	viper.BindEnv(MON_NOTIFICATIONS_EMAIL_SMTP_PORT)
	viper.BindEnv(MON_NOTIFICATIONS_TELEGRAM_TOKEN)

	fmt.Println("notifications.mail.smtp_host")
	fmt.Println(viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_HOST))
	fmt.Println()
	fmt.Println("notifications.mail.smtp_port")
	fmt.Println(viper.GetString(MON_NOTIFICATIONS_EMAIL_SMTP_PORT))
	fmt.Println()
	fmt.Println("notifications.telegram.token")
	fmt.Println(viper.GetString(MON_NOTIFICATIONS_TELEGRAM_TOKEN))
	fmt.Println()
}
