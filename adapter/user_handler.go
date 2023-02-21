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

type userRes struct {
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

		var usersRes []userRes

		for _, user := range users {
			usersRes = append(usersRes, userRes{
				ID:       int(user.ID),
				Username: user.Username,
				Email:    user.Email,
				Text:     user.Text,
				Avater:   user.Avater,
				Header:   user.Header,
			})
		}
		return c.JSON(http.StatusOK, usersRes)
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

		userRes := userRes{
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

type userPostReq struct {
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

		params := &userPostReq{}
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

		var userRes userRes

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

		params := &userPostReq{}
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

		var userRes userRes

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

// Invitaion
type invRes struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Comment string `json:"comment"`
	Place   string `json:"place"`
}

func (h *handler) FindAllInv() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		invs, err := h.uc.FindAllInv(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var invsRes []invRes

		for _, inv := range invs {
			invsRes = append(invsRes, invRes{
				ID:      int(inv.ID),
				UserID:  inv.UserID,
				Comment: inv.Comment,
				Place:   inv.Place,
			})
		}
		return c.JSON(http.StatusOK, invsRes)
	}
}

type invPostReq struct {
	Comment string `json:"comment" validate:"required"`
	Place   string `json:"place" validate:"required"`
}

func (h *handler) CreateInv() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		if c.Request().Method != "POST" {
			return c.JSON(http.StatusBadRequest, "Post method only allow")
		}

		params := &invPostReq{}
		if err := c.Bind(params); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdInv, err := h.uc.CreateInv(
			ctx,
			userID,
			params.Comment,
			params.Place,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		var res invRes

		res.ID = int(createdInv.ID)
		res.UserID = createdInv.UserID
		res.Comment = createdInv.Comment
		res.Place = createdInv.Place

		return c.JSON(http.StatusOK, res)
	}
}

func (h *handler) UserInv() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		invs, err := h.uc.FindInvitationByUserID(ctx, userID)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var invsRes []invRes

		for _, inv := range invs {
			invsRes = append(invsRes, invRes{
				ID:      int(inv.ID),
				UserID:  inv.UserID,
				Comment: inv.Comment,
				Place:   inv.Place,
			})
		}
		return c.JSON(http.StatusOK, invsRes)
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
