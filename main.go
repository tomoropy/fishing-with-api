package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tomoropy/fishing-with-api/adapter"
	"github.com/tomoropy/fishing-with-api/auth"
	"github.com/tomoropy/fishing-with-api/infra"
	"github.com/tomoropy/fishing-with-api/usecase"
)

func main() {

	// DI
	mySQLConn := infra.NewMySQLConnector()
	userRepository := infra.NewUserRepository(mySQLConn.Conn)
	invRepository := infra.NewInvRepostitory(mySQLConn.Conn)
	usecase := usecase.NewUsecase(userRepository, invRepository)
	handler := adapter.NewHandler(usecase)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome my app!!")
	})

	// user
	userRoute := e.Group("/user")
	userRoute.Use(auth.AuthMiddleware)

	userRoute.GET("", handler.FindAllUser())
	userRoute.GET(":id", handler.FindUserByID())
	userRoute.POST("", handler.CreateUser())
	userRoute.PUT(":id", handler.UpdateUser())
	userRoute.DELETE(":id", handler.DeleteUser())

	// invitation
	invRoute := e.Group("/invitation")
	invRoute.Use(auth.AuthMiddleware)

	invRoute.GET("", handler.FindAllInv())
	invRoute.GET(":id", handler.FindInv())
	invRoute.GET("user/:id", handler.UserInv())
	invRoute.POST("user/:id", handler.CreateInv())
	invRoute.PUT(":id", handler.UpdateInv())
	invRoute.DELETE(":id", handler.DeleteInv())

	e.Logger.Fatal(e.Start(":8080"))
}
