package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/tomoropy/fishing-with-api/adapter"
	"github.com/tomoropy/fishing-with-api/infra"
	"github.com/tomoropy/fishing-with-api/usecase"
)

func main() {

	mySQLConn := infra.NewMySQLConnector()
	userRepository := infra.NewUserRepository(mySQLConn.Conn)
	invRepository := infra.NewInvRepostitory(mySQLConn.Conn)
	usecase := usecase.NewUsecase(userRepository, invRepository)
	handler := adapter.NewHandler(usecase)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome my app!!")
	})

	// user
	e.GET("/users", handler.FindAllUser())
	e.GET("/user/:id", handler.FindUserByID())
	e.POST("/user", handler.CreateUser())
	e.PUT("/user/:id", handler.UpdateUser())
	e.DELETE("/user/:id", handler.DeleteUser())

	// invitation
	// e.GET("invitation/:id", invHandler.FindInv())
	e.GET("invitations", handler.FindAllInv())
	e.GET("user/:id/invitations", handler.UserInv())
	e.POST("user/:id/invitation", handler.CreateInv())
	// e.PUT("user/:id/invitation", invHandler.UpdateInv())
	// e.DELETE("user/:id/invitation", invHandler.DeleteInv())

	e.Logger.Fatal(e.Start(":8080"))
}
