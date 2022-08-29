package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"os"
	"postgres-without-orm/internal/godb"
	_ "postgres-without-orm/internal/migrations"
	"postgres-without-orm/pkg/helpers/pg"
)

func main() {
	cfg := &pg.Config{
		Host:     "localhost",
		Username: "root",
		Password: "rootroot",
		Port:     "54320",
		DbName:   "db_test",
		Timeout:  5,
	}

	poolConfig, err := pg.NewPoolConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Pool config error: %v\n", err)
		os.Exit(1)
	}

	poolConfig.MaxConns = 5

	c, err := pg.NewConnection(poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK!")

	_, err = c.Exec(context.Background(), ";")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Ping OK!")

	ins := &godb.Instance{Db: c}
	ins.Start()

	mdb, _ := sql.Open("postgres", poolConfig.ConnString())
	err = mdb.Ping()
	if err != nil {
		panic(err)
	}
	err = goose.Up(mdb, "/var")
	if err != nil {
		panic(err)
	}
}
