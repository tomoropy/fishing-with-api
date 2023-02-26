package main

import (
	"context"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	generated "github.com/tomoropy/fishing-with-api/graph"
	"github.com/tomoropy/fishing-with-api/infra"
	"github.com/tomoropy/fishing-with-api/usecase/registry"
)

// "log"
// "net/http"
// "os"

// _ "github.com/go-sql-driver/mysql"
// "github.com/labstack/echo/v4"

// "github.com/labstack/echo/v4"

// "github.com/labstack/echo/v4/middleware"

// "github.com/tomoropy/fishing-with-api/adapter"
// "github.com/tomoropy/fishing-with-api/auth"
// "github.com/tomoropy/fishing-with-api/infra"
// "github.com/tomoropy/fishing-with-api/usecase"

// "github.com/99designs/gqlgen/graphql/handler"
// "github.com/99designs/gqlgen/graphql/playground"
// generated "github.com/tomoropy/fishing-with-api/graph"

const defaultPort = "8080"

func main() {

	// DI
	// mySQLConn := infra.NewMySQLConnector()
	// userRepository := infra.NewUserRepository(mySQLConn.Conn)
	// invRepository := infra.NewInvRepostitory(mySQLConn.Conn)
	// usecase := usecase.NewUsecase(userRepository, invRepository)
	// handler := adapter.NewHandler(usecase)

	// e := echo.New()
	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Welcome my app!!")
	// })

	// e.POST("login", handler.Login())
	// e.POST("register", handler.Register())

	// // user
	// e.GET("users", handler.FindAllUser())
	// e.GET("user/:id", handler.FindUserByID())
	// // need auth
	// userRoute := e.Group("user")
	// userRoute.Use(auth.AuthMiddleware)
	// userRoute.PUT("/:id", handler.UpdateUser())
	// userRoute.DELETE("/:id", handler.DeleteUser())

	// // invitation
	// e.GET("/invitation", handler.FindAllInv())
	// e.GET("/invitation/:id", handler.FindInv())
	// // need auth
	// invRoute := e.Group("/invitation")
	// invRoute.Use(auth.AuthMiddleware)
	// invRoute.GET("user/:id", handler.UserInv())
	// invRoute.POST("user/:id", handler.CreateInv())
	// invRoute.PUT(":id", handler.UpdateInv())
	// invRoute.DELETE(":id", handler.DeleteInv())

	// e.Logger.Fatal(e.Start(":8080"))

	// gqlgen
	// const defaultPort = "8080"

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// srv := gqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	// e.Logger.Fatal(e.Start(":8080"))
	// e := echo.New()

	// // GraphQL handler
	// gqlHandler := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &generated.Resolver{}}))
	// e.POST("/query", echo.WrapHandler(gqlHandler))

	// // Playground handler
	// // playgroundHandler := handler.Playground("GraphQL", "/query")
	// playgroundHandler := playground.Handler("GraphQL", "/query")

	// e.GET("/playground", echo.WrapHandler(playgroundHandler))

	// // REST API routes
	// // api := e.Group("/api")

	// e.Logger.Fatal(e.Start(":8080"))

	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx := context.Background()
	register := regist(ctx)
	resolver, err := register.NewResolver(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: resolver}),
	)
	e.POST("/query", echo.WrapHandler(graphqlHandler))

	playgoundHandler := playground.Handler("GraphQL playground", "/query")
	e.GET("/playground", echo.WrapHandler(playgoundHandler))

	e.Logger.Fatal(e.Start(":8080"))
}

func regist(ctx context.Context) registry.Register {
	db := infra.NewMySQLConnector()
	return registry.NewRegister(db.Conn)
}
