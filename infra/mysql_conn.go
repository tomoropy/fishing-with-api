package infra

import (
	"log"
	"time"

	"database/sql"

	"github.com/tomoropy/fishing-with-api/config"
	"github.com/tomoropy/fishing-with-api/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConnector struct {
	Conn *gorm.DB
}

func NewMySQLConnector() *MySQLConnector {

	config, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Printf("config: %v", config.DB)

	dsn := config.DB.User + ":" + config.DB.Password + "@tcp(mysql)/" + config.DB.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDB, err := sql.Open("mysql", dsn)

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

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		log.Fatal(err)
	}

	// Userテーブルを作成
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Invitation{})

	return &MySQLConnector{
		Conn: db,
	}

}
