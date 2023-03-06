package infra

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tomoropy/fishing-with-api/config"
)

type MySQLConnector struct {
	Conn *sqlx.DB
}

func NewMySQLConnector() *MySQLConnector {

	config, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dsn := config.DB.User + ":" + config.DB.Password + "@tcp(mysql)/" + config.DB.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	// -----------------------------------------------------------------------------------------------------------
	// DBへのアクセスを時間をおいて試みる

	for {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	// -----------------------------------------------------------------------------------------------------------
	// テーブルの作成

	userTabeleGenerateSql := `
	CREATE TABLE IF NOT EXISTS users (
		uid VARCHAR(36) NOT NULL,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		hashed_password VARCHAR(255) NOT NULL,
		text VARCHAR(255) NOT NULL,
		avater VARCHAR(255) NOT NULL,
		header VARCHAR(255) NOT NULL,
		created_at VARCHAR(255) NOT NULL,
		PRIMARY KEY (uid)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`
	_, err = db.Exec(userTabeleGenerateSql)
	if err != nil {
		log.Fatal(err)
	}

	tweetTabeleGenerateSql := `
	CREATE TABLE IF NOT EXISTS tweets (
		uid VARCHAR(36) NOT NULL,
		user_uid VARCHAR(36) NOT NULL,
		body VARCHAR(255) NOT NULL,
		image VARCHAR(255) NOT NULL,
		created_at VARCHAR(255) NOT NULL,
		PRIMARY KEY (uid),
		FOREIGN KEY (user_uid) REFERENCES users(uid)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`
	_, err = db.Exec(tweetTabeleGenerateSql)
	if err != nil {
		log.Fatal(err)
	}

	return &MySQLConnector{
		Conn: db,
	}

}
