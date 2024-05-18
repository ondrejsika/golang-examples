package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s folderPath prefix\n", os.Args[0])
	}

	folderPath := os.Args[1]
	prefix := os.Args[2]

	// Get a list of all files in the folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalln(err)
	}

	// Sort the files alphabetically
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// Loop through each file and rename it with the prefix and a number
	for i, file := range files {
		extension := filepath.Ext(file.Name())                                                       // Get the file extension
		newName := fmt.Sprintf("%s%d%s", prefix, i+1, extension)                                     // Create the new file name
		err := os.Rename(filepath.Join(folderPath, file.Name()), filepath.Join(folderPath, newName)) // Rename the file
		if err != nil {
			log.Fatalln(err)
		}
	}
}
