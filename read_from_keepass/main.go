package main

import (
	"fmt"
	"log"
	"os"
	"path"

	kp "github.com/tobischo/gokeepasslib/v3"
)

func main() {
	f, err := os.Open("example.kdbx")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	db := kp.NewDatabase()
	db.Credentials = kp.NewPasswordCredentials("asdfasdf")

	if err := kp.NewDecoder(f).Decode(db); err != nil {
		log.Fatal("decode:", err)
	}

	db.UnlockProtectedEntries()

	root := db.Content.Root
	for _, g := range root.Groups {
		walkGroup(g, "/")
	}
}

// walkGroup recursively prints all subgroups and their entries.
func walkGroup(g kp.Group, parent string) {
	cur := path.Join(parent, g.Name)
	fmt.Println("Group:", cur)

	// Entries in this group
	for _, e := range g.Entries {
		fmt.Println("  Entry:", e.GetTitle())
		fmt.Println("    Username:", e.GetContent("UserName"))
		fmt.Println("    Password:", e.GetPassword())
	}

	// Recurse into subgroups
	for _, sg := range g.Groups {
		walkGroup(sg, cur)
	}
}
