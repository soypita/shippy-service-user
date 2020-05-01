package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	var counter int
	var err error

	for counter <= 3 {
		time.Sleep(3 * time.Second)

		res, err := gorm.Open(
			"postgres",
			fmt.Sprintf(
				"postgres://%s:%s@%s/%s?sslmode=disable",
				user, password, host, DBName,
			),
		)

		if err != nil {
			counter++
			continue
		}

		return res, nil
	}

	return nil, err
}
