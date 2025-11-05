package main

import (
	"fmt"
	"log"

	"github.com/sikalabs/sikalabs-crypt-go/pkg/sikalabs_crypt"
)

const PASSWORD = "delinka"
const TEXT = `            __________________
           < Hello, I'm Dela! >
            ------------------
           /
 /)-_-(\  /
  (o o)
   \o/\__-----.
    \      __  \
     \| /_/  \ /\__/
      ||      \\
      ||      //
      /|     /|`
const ENCRYPTED = "rhx7YD+YuXwrUNNBkbYKr/JHq/CS+8KhgXFnmJCLlYYB3gnq5iNTJyd98yommvYmt8kkrs7GJcG2GTy6/bjp5v4f08ePITabKTtdMW6NUei3env/K80FzKkhe0hma8G7pjCEGsQJ01hToBfGLAVWATBLKTguAnmFFfo+1rG0jh0iWz2ofxRT97PR24vse/AyOrgKlpg+385sMxqIoT/3PUcdQoLIuiU+rBUWmj4asM95vrSazPGQAQ6R89LZuQiJSkTxycmVvZN+I8mEatAtjc4Xv27/rOjd4VUXydK2Euwt0cmmpSj1edep2n2cs3ughwpMa/geiNV7UlufV3z9F2Z9IlHI66ygJsLvNfouuCNW6g=="

func main() {
	ecrypted, err := sikalabs_crypt.SikaLabsSymmetricEncryptV1(PASSWORD, TEXT)
	handleError(err)
	fmt.Println("Encrypted")
	fmt.Println(ecrypted)
	fmt.Println()

	decrypted, err := sikalabs_crypt.SikaLabsSymmetricDecryptV1(PASSWORD, ecrypted)
	handleError(err)
	fmt.Println("Decrypted")
	fmt.Println(decrypted)
	fmt.Println()

	decrypted2, err := sikalabs_crypt.SikaLabsSymmetricDecryptV1(PASSWORD, ENCRYPTED)
	handleError(err)
	fmt.Println("Decrypted (predefined)")
	fmt.Println(decrypted2)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
