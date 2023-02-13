package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/tomoropy/clean-arc-go/adapter"
	"github.com/tomoropy/clean-arc-go/infra"
	"github.com/tomoropy/clean-arc-go/usecase"
)

func main() {

	mySQLConn := infra.NewMySQLConnector()
	userRepository := infra.NewUserRepository(mySQLConn.Conn)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := adapter.NewUserHandler(userUsecase)

	
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	_, err := mySQLConn.Conn.Exec("CREATE TABLE IF NOT EXISTS user (id INT, username VARCHAR(100), email VARCHAR(20), password VARCHAR(40), age INT);")
	if err != nil {
		e.Logger.Fatal(err)
	}

	_, err = mySQLConn.Conn.Exec("INSERT INTO user (id, username, email, password, age) VALUES ('1', 'tomo', 'ucchi', 'password', 23);")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/users", userHandler.FindAllUser())
	e.GET("/user/id", userHandler.FindUserByID())
	e.Logger.Fatal(e.Start(":8080"))
}
