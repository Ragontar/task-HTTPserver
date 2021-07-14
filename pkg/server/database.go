package server

import (
	"context"
	"fmt"
	"time"
	"github.com/jackc/pgx/v4"
)

/*

	Database-related stuff:

	type Database struct --- structure of a DSN (data source name string) parameters. Self-explained.
		Used to form the dsn-string and to parse the json config file.

	Init() *pgx.Conn --- tries to initialize a connection to PostgreSQL database. (driver - pgx)
		Returns nil in case of failure.

	GetDB() *pgx.Conn --- returns db connection if it's initialized, otherwise runs Init(). 
		Also has a progressive delay between attempts.


*/

type Database struct {
	Host string `json:"host"`
	User string `json:"user"`
	Password string `json:"password"`
	Dbname string `json:"dbname"`
	Port string `json:"port"`
	Sslmode string `json:"sslmode"`
}

var db *pgx.Conn

func Init() *pgx.Conn {
	dbConf := getConfig().Db

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", 
						dbConf.Host, dbConf.User, dbConf.Password, dbConf.Dbname, dbConf.Port, 
						dbConf.Sslmode)

	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Println(err)
	}

	return db
}

func GetDB() *pgx.Conn {
	if db == nil {
		db = Init()
		sleep := time.Duration(1)

		for db == nil {
			sleep *= 2
			fmt.Printf("DB is unavailable. Sleeping %ds....", sleep)
			time.Sleep(sleep * time.Second)
			db = Init()
		}
	}

	return db
}

