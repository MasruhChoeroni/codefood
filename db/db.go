package db

import (
	"codefood/config"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {

	conf := config.GetConfig()

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	if host != "" || port != "" || user != "" || pass != "" || dbname != "" {
		conf.DB_PORT = port
		conf.DB_USERNAME = user
		conf.DB_PASSWORD = pass
		conf.DB_NAME = dbname
		conf.DB_HOST = host
	}

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?parseTime=true"

	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("connectionString error...!!!")
	}

	err = db.Ping()

	if err != nil {
		panic("DSN Invalid : " + connectionString)
	}
}

func CreateCon() *sql.DB {
	return db
}
