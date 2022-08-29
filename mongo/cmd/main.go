package main

import (
	"mongo/internal/app"
	"mongo/internal/models"
	"mongo/pkg/helpers/mongo"
	"time"
)

func main() {
	cfg := &mongo.Config{
		Host:   "localhost",
		Port:   "27017",
		DbName: "db_test",
	}

	mongoSession, err := mongo.CreateSession(cfg)
	if err != nil {
		println(err)
	}

	u := models.User{
		Name:     "Vitaly",
		Age:      24,
		CreateAt: time.Now().Unix(),
	}

	//==============================================

	u.Documents.PassportNumber = "1111111"
	u.Documents.INN = 111111

	app.Store(mongoSession, u)

}
