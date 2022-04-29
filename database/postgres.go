package database

import (
	"fmt"
	"sesi-final-project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host    string = "localhost"
	port    int    = 5432
	user    string = "postgres"
	pass    string = "koinworks2022*" // diisi password masing-masing
	dbname  string = "db-go-sql"
	sslmode string = "disabled"
	db      *gorm.DB
	err     error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbname, sslmode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	fmt.Println("db connection success")
}

func GetDB() *gorm.DB {
	return db
}
