package postPack

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func dbConnInit() (db *sql.DB) {
	dbHost := "localhost"
	dbName := "golang_beginner"
	dbUser := "golang_beginner"
	dbPass := "ymVz4U6PBmq51HPP"
	dbDriver := "mysql"
	//db, err :=sql.Open(dbDriver, "golang_beginner:ymVz4U6PBmq51HPP@/golang_beginner")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	if err != nil {
		log.Println("dbConnInit ERROR: ", err.Error())
		panic(err.Error())
	}
	return db
}
