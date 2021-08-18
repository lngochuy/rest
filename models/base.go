package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var db *gorm.DB

func InitializeDB() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		5432,
		"postgres",
		"Postgres#123",
		"tmarest",
	)

	db, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatalln(err)
	}

	sqldb, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	err = sqldb.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func CreateTable(drop bool) {
	if drop {
		err := db.Migrator().DropTable(&User{})
		if err != nil {
			log.Fatalln(err)
		}
	}

	err := db.AutoMigrate(User{})
	if err != nil {
		log.Fatalln(err)
	}
}
