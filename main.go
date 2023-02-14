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
		return c.String(http.StatusOK, "Welcome my app!!")
	})

	e.GET("/users", userHandler.FindAllUser())
	e.GET("/user/:id", userHandler.FindUserByID())
	e.POST("/user", userHandler.CreateUser())
	e.Logger.Fatal(e.Start(":8080"))
}
