package app

import (
	mongo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"mongo/internal/models"
)

func Store(sess *mongo.Session, user models.User) error {

	err := sess.DB("").C("users").Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func Find(sess *mongo.Session, name string) (models.User, error) {
	var u models.User
	q := bson.M{
		"name": "Vitaly",
	}
	err := sess.DB("").C("users").Find(q).One(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}
