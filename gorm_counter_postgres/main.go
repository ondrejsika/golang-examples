package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Counter struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	Counter int
}

func getCounter() int {
	// Connection string for the PostgreSQL database
	dsn := "host=localhost user=postgres password=pg dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Counter{})

	// Ensure the counter exists
	db.Where(Counter{ID: 1}).FirstOrCreate(&Counter{ID: 1, Counter: 0})

	// Increment the counter using Exec to execute SQL directly
	db.Exec("UPDATE counters SET counter = counter + 1 WHERE id = 1")

	// Get the counter value
	var counter Counter
	db.First(&counter, 1)

	return counter.Counter
}

func main() {
	fmt.Println(getCounter())
}
