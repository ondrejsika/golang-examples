package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/alash3al/go-smtpsrv/v3"
)

func main() {
	fmt.Println(smtpsrv.ListenAndServe(&smtpsrv.ServerConfig{
		ListenAddr:   ":25",
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
		Handler: func(c *smtpsrv.Context) error {
			var message string
			b := make([]byte, 8)
			for {
				_, err := c.Read(b)
				if err == io.EOF {
					break
				}
				message = message + string(b)
			}

			email, _ := smtpsrv.ParseEmail(strings.NewReader(message))

			fmt.Println("smtp from:", c.From())
			fmt.Println("smtp to:", c.To())

			fmt.Println("---")

			fmt.Println("from:", email.From)
			fmt.Println("to:", email.To)
			fmt.Println("subject:", email.Subject)
			fmt.Println("text body:", email.TextBody)

			fmt.Println("---")

			fmt.Println("raw message:")
			fmt.Println("")
			fmt.Println(message)

			fmt.Println("==========")
			return nil
		},
	}))
}
