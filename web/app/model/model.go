package model

import (
	"database/sql"
	"fmt"
	"os"
	"stocker-quant/util"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var dbConnection *sql.DB
var err error

func Connect() {
	username := os.Getenv("dbUsername")
	password := os.Getenv("dbPassword")
	port := os.Getenv("dbPort")
	host := os.Getenv("dbHost")
	dbname := os.Getenv("dbDatabase")
	option := os.Getenv("dbOption")

	connInfo := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + option
	dbConnection, err = sql.Open("postgres", connInfo)
	util.HandleError(err, "Database Connection Fail")

	err = dbConnection.Ping()
	util.HandleError(err, "Database Connection Fail")

	log.Info("Database Connected.")
}

func TestQuery() {
	rows, err := dbConnection.Query("SELECT * FROM schema_migrations")
	util.HandleError(err, "Test Query Fail")

	for rows.Next() {
		var (
			version int64
			dirty   bool
		)
		if err := rows.Scan(&version, &dirty); err != nil {
			log.Fatal(err)
		}
		fmt.Println(version, dirty)
	}
}

func Close() {
	err := dbConnection.Close()
	util.HandleError(err, "Database Disconnection Fail")

	log.Info("Database Disconnected.")
}

func GetDB() *sql.DB {
	return dbConnection
}
