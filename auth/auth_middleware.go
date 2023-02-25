package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tomoropy/fishing-with-api/auth/token"
	"github.com/tomoropy/fishing-with-api/config"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(authorizationHeaderKey)

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, errors.New("authraization header is not provided"))
		}

		field := strings.Fields(authHeader)
		if len(field) < 2 {
			return c.JSON(http.StatusUnauthorized, errors.New("invalid authraization header format"))
		}

		authorizationType := strings.ToLower(field[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsuported authorization type %s", authorizationType)
			return c.JSON(http.StatusUnauthorized, err)
		}

		// load config
		config, err := config.Load()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}

		accessToken := field[1]
		tokenMaker, err := token.NewJWTMaker(config.Auth.SecretKey)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		c.Set(authorizationPayloadKey, payload)
		return next(c)
	}
}
