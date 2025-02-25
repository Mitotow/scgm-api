package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "scgm"
	password = "scgm"
	dbName   = "scgm"
)

func DatabaseConnection() *gorm.DB {
	// port, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 0)
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
