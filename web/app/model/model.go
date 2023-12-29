package model

import (
	"database/sql"
	"fmt"
	apperror "stocker-hf-data/web/app/app-error"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"go.deanishe.net/env"
)

var dbConnection *sql.DB
var dbAdminConnection *sql.DB
var err error

func Connect() *apperror.ModelError {
	username := env.Get("dbUsername")
	password := env.Get("dbPassword")
	adminuser := env.Get("dbAdmin")
	adminpass := env.Get("dbAdminPassword")
	port := env.Get("dbPort")
	host := env.Get("dbHost")
	dbname := env.Get("dbDatabase")
	option := env.Get("dbOption")

	// user connect
	connInfo := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?" + option
	dbConnection, err = sql.Open("mysql", connInfo)
	if err != nil {
		return apperror.NewModelError(apperror.ErrDbConnectFail)
	}

	err = dbConnection.Ping()
	if err != nil {
		return apperror.NewModelError(apperror.ErrDbConnectFail)
	}

	// admin connect
	connInfo = adminuser + ":" + adminpass + "@tcp(" + host + ":" + port + ")/" + dbname + "?" + option
	dbAdminConnection, err = sql.Open("mysql", connInfo)
	if err != nil {
		return apperror.NewModelError(apperror.ErrDbConnectFail)
	}

	err = dbAdminConnection.Ping()
	if err != nil {
		return apperror.NewModelError(apperror.ErrDbConnectFail)
	}

	log.Info("Database Connected.")
	return nil
}

func TestQuery() *apperror.ModelError {
	rows, err := dbConnection.Query("SELECT * FROM schema_migrations")
	if err != nil {
		return apperror.NewModelError(apperror.ErrDbTestFail)
	}

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
	return nil
}

func Close() *apperror.ModelError {
	err := dbConnection.Close()
	if err != nil {
		return apperror.NewModelError(err)
	}

	log.Info("Database Disconnected.")
	return nil
}

func GetDB() *sql.DB {
	return dbConnection
}

func GetAdminDB() *sql.DB {
	return dbAdminConnection
}
