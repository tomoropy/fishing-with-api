package main

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	generated "github.com/tomoropy/fishing-with-api/graph"
	"github.com/tomoropy/fishing-with-api/infra"
	"github.com/tomoropy/fishing-with-api/usecase/registry"
)

func main() {

	e := echo.New()

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
