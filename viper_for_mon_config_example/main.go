package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Notifications struct {
		Mail struct {
			SmtpHost string `mapstructure:"smtp_host"`
			SmtpPort int    `mapstructure:"smtp_port"`
		} `mapstructure:"mail"`
		Telegram struct {
			Token string `mapstructure:"token"`
		} `mapstructure:"telegram"`
	} `mapstructure:"notifications"`
}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("mon")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.SetEnvPrefix("mon")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Unmarshal into struct
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("error unmarshaling config: %v\n", err)
		return
	}

	fmt.Println("Struct:")
	fmt.Printf("  smtp_host: %s\n", cfg.Notifications.Mail.SmtpHost)
	fmt.Printf("  smtp_port: %d\n", cfg.Notifications.Mail.SmtpPort)
	fmt.Printf("  telegram:  %s\n", cfg.Notifications.Telegram.Token)
}
