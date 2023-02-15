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
	db, err := sql.Open("mysql", "root:password@tcp(mysql)/myapp")
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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user (id INT NOT NULL AUTO_INCREMENT, username VARCHAR(100) NOT NULL, email VARCHAR(20) NOT NULL, password VARCHAR(40) NOT NULL, age INT, PRIMARY KEY (`id`));")
	if err != nil {
		log.Fatal(err)
	}

	return &MySQLConnector{
		Conn: db,
	}

}
