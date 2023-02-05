package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Doggo struct {
	gorm.Model
	Name   string
	Age    int
	Weight int
	Breed  string
}

func main() {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Doggo{})

	// Create
	db.Create(&Doggo{
		Name:   "Dela",
		Age:    2,
		Weight: 15,
		Breed:  "Labrador+Wippet",
	})

	db.Create(&Doggo{
		Name:   "Nela",
		Age:    6,
		Weight: 25,
		Breed:  "Labrador+SwissShepherd",
	})

	db.Create(&Doggo{
		Name:   "Debora",
		Age:    4,
		Weight: 30,
		Breed:  "Wolfdog",
	})

	// Read
	var doggos []Doggo
	db.Order("created_at DESC").Find(&doggos, "age > ?", "1")

	for _, doggo := range doggos {
		fmt.Printf("id=%d name=%s age=%d\n", doggo.ID, doggo.Name, doggo.Age)
	}

	// Update
	fmt.Println()

	for _, doggo := range doggos {
		db.Model(&doggo).Update("Age", doggo.Age+1)
	}

	db.Order("created_at DESC").Find(&doggos, "age > ?", "1")
	for _, doggo := range doggos {
		fmt.Printf("id=%d name=%s age=%d\n", doggo.ID, doggo.Name, doggo.Age)
	}
}
