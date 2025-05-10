package main

import (
	"fmt"
	"os"
	"strings"
)

const DOG = ` /)-_-(\  /
  (o o)
   \o/\__-----.
    \      __  \
     \| /_/  \ /\__/
      ||      \\
      ||      //
      /|     /|`

func bubble(text string) string {
	lines := strings.Split(text, "\n")
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	top := "           " + " " + strings.Repeat("_", maxLen+2)
	bottom := "           " + " " + strings.Repeat("-", maxLen+2)
	var middle []string
	for _, line := range lines {
		padded := line + strings.Repeat(" ", maxLen-len(line))
		middle = append(middle, "           "+fmt.Sprintf("< %s >", padded))
	}

	// Connect bubble to dog (6 spaces indentation + positioning)
	connector := "           /"
	return fmt.Sprintf("%s\n%s\n%s\n%s", top, strings.Join(middle, "\n"), bottom, connector)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go \"I'm pawesome\"")
		return
	}
	text := os.Args[1]
	fmt.Println(bubble(text))
	fmt.Println(DOG)
}
