package infra

import (
	"database/sql"
	"log"
	"time"
)

type MySQLConnector struct {
	Conn *sql.DB
}

func NewMySQLConnector() *MySQLConnector {
	var err error
	db, err := sql.Open("mysql", "user:password@tcp(mysql)/myapp")
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	return &MySQLConnector{
		Conn: db,
	}

}
