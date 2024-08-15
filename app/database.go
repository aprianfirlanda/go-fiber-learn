package app

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dbUser := viper.GetString("DATABASE_USER")
	dbPassword := viper.GetString("DATABASE_PASSWORD")
	dbName := viper.GetString("DATABASE_NAME")
	dbHost := viper.GetString("DATABASE_HOST")
	dbPort := viper.GetInt("DATABASE_PORT")
	dbSSLMode := viper.GetString("DATABASE_SSLMODE")

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode,
	)

	dialect := postgres.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
