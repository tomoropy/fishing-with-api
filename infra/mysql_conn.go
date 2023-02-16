package infra

import (
	"log"
	"time"

	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnector struct {
	Conn *gorm.DB
}

func NewMySQLConnector() *MySQLConnector {
	dsn := "root:passwrod@tcp(mysql)/myapp"
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		log.Fatal(err)
	}

	// DBへのアクセスを時間をおいて試みる
	for {
		err = sqlDB.Ping()
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS user (id INT NOT NULL AUTO_INCREMENT, username VARCHAR(100) NOT NULL, email VARCHAR(20) NOT NULL, password VARCHAR(40) NOT NULL, age INT, PRIMARY KEY (`id`));")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return &MySQLConnector{
		Conn: db,
	}

}
