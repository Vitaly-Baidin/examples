package mongo

import (
	"fmt"
	mongo "github.com/globalsign/mgo"
)

type Config struct {
	Host   string
	Port   string
	DbName string
}

func CreateSession(cfg *Config) (*mongo.Session, error) {
	connStr := fmt.Sprintf(
		"%s://%s:%s/%s",
		"mongo",
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	sess, err := mongo.Dial(connStr)
	if err != nil {
		panic(err)
	}

	return sess, nil
}
