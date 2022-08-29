package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"postgres-with-orm/internal/app"
	"postgres-with-orm/internal/models"
)

func main() {
	dsn := "host=localhost user=root password=rootroot dbname=gorm port=54320 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = models.InitModels(db)
	if err != nil {
		panic(err)
	}

	app.Start(db)
}
