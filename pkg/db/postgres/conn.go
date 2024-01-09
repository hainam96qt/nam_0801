package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Username     string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"database"`
	Host         string `yaml:"host"`
}

func ConnectDatabase(args DatabaseConfig) (*sql.DB, error) {
	var con = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", args.Username, args.Password, args.Host, args.Port, args.DatabaseName)
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return db, nil
}
