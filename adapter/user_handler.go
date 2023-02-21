package adapter

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tomoropy/fishing-with-api/usecase"
)

type handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *handler {
	return &handler{
		uc: uc,
	}
}

type userResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Text     string `json:"text"`
	Avater   string `json:"avater"`
	Header   string `json:"header"`
}

func (h *handler) FindAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		users, err := h.uc.FindAllUser(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var userRes []userResponse

		for _, user := range users {
			userRes = append(userRes, userResponse{
				ID:       int(user.ID),
				Username: user.Username,
				Email:    user.Email,
				Text:     user.Text,
				Avater:   user.Avater,
				Header:   user.Header,
			})
		}
		return c.JSON(http.StatusOK, userRes)
	}
}

func (h *handler) FindUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		user, err := h.uc.FindUserByID(ctx, userID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		userRes := userResponse{
			ID:       int(user.ID),
			Username: user.Username,
			Email:    user.Email,
			Text:     user.Text,
			Avater:   user.Avater,
			Header:   user.Header,
		}

		return c.JSON(http.StatusOK, userRes)
	}
}

type userPostRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Text     string `json:"text"`
	Avater   string `json:"avater"`
	Header   string `json:"header"`
}

func (h *handler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		if c.Request().Method != "POST" {
			return c.JSON(http.StatusBadRequest, "Post method only allow")
		}

		params := &userPostRequest{}
		if err := c.Bind(params); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		createdUser, err := h.uc.CreateUser(
			ctx,
			params.Username,
			params.Email,
			params.Password,
			params.Text,
			params.Avater,
			params.Header,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		var userRes userResponse

		userRes.ID = int(createdUser.ID)
		userRes.Username = createdUser.Username
		userRes.Email = createdUser.Email
		userRes.Text = createdUser.Text
		userRes.Avater = createdUser.Avater
		userRes.Header = createdUser.Header

		return c.JSON(http.StatusOK, userRes)
	}
}

func (h *handler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		if c.Request().Method != "PUT" {
			return c.JSON(http.StatusBadRequest, "PUT method only allow")
		}

		params := &userPostRequest{}
		if err := c.Bind(params); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updateUser, err := h.uc.UpdateUser(
			ctx, userID,
			params.Username,
			params.Email,
			params.Password,
			params.Text,
			params.Avater,
			params.Header)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		var userRes userResponse

		userRes.ID = int(updateUser.ID)
		userRes.Username = updateUser.Username
		userRes.Email = updateUser.Email
		userRes.Text = updateUser.Text
		userRes.Avater = updateUser.Avater
		userRes.Header = updateUser.Header

		return c.JSON(http.StatusOK, userRes)

	}
}

func (h *handler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		if c.Request().Method != "DELETE" {
			return c.JSON(http.StatusBadRequest, "DELETE method only allow")
		}

		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = h.uc.DeleteUser(ctx, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "succsess!")
	}
}

// Invitation Handler
// type invHandler struct {
// 	us usecase.IinvUsecase
// }

// func NewInvHandler(iu usecase.IinvUsecase) *invHandler {
// 	return &invHandler{
// 		us: iu,
// 	}
// }

// func (ih *invHandler) FindInv() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, "find Inv WIP...")
// 	}
// }
// func (ih *invHandler) AllInv() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, "all inv WIP...")
// 	}
// }
// func (ih *invHandler) UserInv() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, "user inv WIP...")
// 	}
// }
// func (ih *invHandler) CreateInv() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, "create inv WIP...")
// 	}
// }
// func (ih *invHandler) UpdateInv() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, "delete inv WIP...")
// 	}
// }
// func (ih *invHandler) DeleteInv() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, "delete inv WIP...")
// 	}
// }
