package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Counter struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	Counter int
}

func getCounter() int {
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Counter{})

	// Ensure the counter exists
	db.Where(Counter{ID: 1}).FirstOrCreate(&Counter{ID: 1, Counter: 0})

	// Increment the counter
	db.Exec("UPDATE counters SET counter = counter + 1 WHERE id = 1")

	// Get the counter
	var counter Counter
	db.First(&counter, 1)

	return counter.Counter
}

func main() {
	fmt.Println(getCounter())
}
