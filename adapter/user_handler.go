package adapter

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tomoropy/clean-arc-go/usecase"
)

type userHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *userHandler {
	return &userHandler{
		usecase: uu,
	}
}

func (uh *userHandler) FindAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		users, err := uh.usecase.FindAllUser(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, users)
	}
}

func (uh *userHandler) FindUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		user, err := uh.usecase.FindUserByID(ctx, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	}
}
