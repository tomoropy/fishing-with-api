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

	// user
	e.GET("/users", userHandler.FindAllUser())
	e.GET("/user/:id", userHandler.FindUserByID())
	e.POST("/user", userHandler.CreateUser())
	e.POST("/user/:id", userHandler.UpdateUser())
	e.DELETE("/user/:id", userHandler.DeleteUser())

	// invitation
	e.GET("invitation/:id", invHandler.FindInv())
	e.GET("invitations", invHandler.AllInv())
	e.GET("user/:id/invitations", invHandler.UserInv())
	e.POST("user/:id/invitation", invHandler.CreateInv())
	e.PUT("user/:id/invitation", invHandler.UpdateInv())
	e.DELETE("user/:id/invitation", invHandler.DeleteInv())

	e.Logger.Fatal(e.Start(":8080"))
}
