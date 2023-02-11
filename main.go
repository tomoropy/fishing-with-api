package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(mysql)/myapp")
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

	router := gin.Default()

	router.Run()
}
