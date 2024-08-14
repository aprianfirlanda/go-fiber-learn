package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := postgres.Open("user=user password=password dbname=mydatabase host=localhost port=5432 sslmode=disable")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
