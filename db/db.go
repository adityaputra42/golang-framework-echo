package db

import (
	"database/sql"
	"golang_framework_echo/config"
	"golang_framework_echo/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func InitDB() {
	conf := config.GetConfiguration()
	configString := conf.DB_Username + ":" + conf.DB_Password + "@tcp(" + conf.DB_Host + ":" + conf.DB_Port + ")/" + conf.Db_Name
	db, err = sql.Open("mysql", configString)

	helper.PanicIfError(err)
	err = db.Ping()
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	// return db

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate -database "mysql://root:@tcp(localhost:3306)/golang_database_migration" -path db/migrations up
	// migrate -database "mysql://root:@tcp(localhost:3306)/golang_database_migration" -path db/migrations down

	// Fixx Dirty State
	// migrate -database "mysql://root:@tcp(localhost:3306)/golang_database_migration" -path db/migrations force 20230919075726
	// migrate -database "mysql://root:@tcp(localhost:3306)/golang_database_migration" -path db/migrations version

}

func CreateCon() *sql.DB {
	return db
}
