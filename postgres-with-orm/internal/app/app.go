package app

import (
	"fmt"
	"gorm.io/gorm"
	"postgres-with-orm/internal/models"
)

func Start(db *gorm.DB) {
	create(db)
	selectAndUpdate(db)
}

func create(db *gorm.DB) {
	u := &models.User{
		Name:     "Artem",
		Age:      30,
		IsVerify: true,
		Cards: []models.Card{{
			Number: "11111111111",
			Type:   "VISA",
		}},
	}
	result := db.Create(u)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("User created with ID: %d", u.Id)
}

func selectAndUpdate(db *gorm.DB) {
	user := &models.User{}
	db.Where("name = ?", "Artem").First(&user)
	if user.Id > 0 {
		user.Name = "Dimitri"
		db.Save(user)
	}
}
