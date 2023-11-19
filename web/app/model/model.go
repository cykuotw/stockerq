package model

import (
	"database/sql"
	"fmt"
	"stocker-quant/util"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"go.deanishe.net/env"
)

var dbConnection *sql.DB
var err error

func Connect() {
	username := env.Get("dbUsername")
	password := env.Get("dbPassword")
	port := env.Get("dbPort")
	host := env.Get("dbHost")
	dbname := env.Get("dbDatabase")
	option := env.Get("dbOption")

	connInfo := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?" + option
	dbConnection, err = sql.Open("mysql", connInfo)
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
