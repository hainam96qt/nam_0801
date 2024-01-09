package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	Username     string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         string `yaml:"port"`
	DatabaseName string `yaml:"database"`
	Host         string `yaml:"host"`
}

func ConnectDatabase(args DatabaseConfig) (*sql.DB, error) {
	var con = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable\"", args.Host, args.Port, args.Username, args.DatabaseName, args.Password)
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return db, nil
}
