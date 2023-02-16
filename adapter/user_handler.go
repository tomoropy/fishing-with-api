package adapter

import (
	"net/http"
	// "strconv"

	"github.com/labstack/echo/v4"
	"github.com/tomoropy/clean-arc-go/usecase"
)

type userHandler struct {
	us usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *userHandler {
	return &userHandler{
		us: uu,
	}
}

type userResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Text     string `json:"text"`
	Avater   string `json:"iamge"`
	Header   string `json:"header"`
}

func (uh *userHandler) FindAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		users, err := uh.us.FindAllUser(ctx)
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
				Avater:   user.Image,
				Header:   user.Header,
			})
		}
		return c.JSON(http.StatusOK, userRes)
	}
}

// func (uh *userHandler) FindUserByID() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := c.Request().Context()
// 		userID, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err.Error())
// 		}
// 		user, err := uh.usecase.FindUserByID(ctx, userID)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err.Error())
// 		}

// 		var userRes userResponse

// 		userRes.ID = user.ID
// 		userRes.Username = user.Username
// 		userRes.Email = user.Email
// 		userRes.Age = user.Age

// 		return c.JSON(http.StatusOK, userRes)
// 	}
// }

// type userPostRequest struct {
// 	Username string `json:"username" validate:"required"`
// 	Password string `json:"password" validate:"required"`
// 	Email    string `json:"email" validate:"required"`
// 	Age      int    `age:"age"`
// }

// func (uh *userHandler) CreateUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := c.Request().Context()

// 		if c.Request().Method != "POST" {
// 			return c.JSON(http.StatusBadRequest, "Post method only allow")
// 		}

// 		params := &userPostRequest{}
// 		if err := c.Bind(params); err != nil {
// 			c.JSON(http.StatusInternalServerError, err.Error())
// 		}

// 		createdUser, err := uh.usecase.CreateUser(ctx, params.Username, params.Email, params.Password, params.Age)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, err.Error())
// 		}

// 		var userRes userResponse

// 		userRes.ID = createdUser.ID
// 		userRes.Username = createdUser.Username
// 		userRes.Email = createdUser.Email
// 		userRes.Age = createdUser.Age

// 		return c.JSON(http.StatusOK, userRes)
// 	}
// }

// func (uh *userHandler) UpdateUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := c.Request().Context()

// 		if c.Request().Method != "POST" {
// 			return c.JSON(http.StatusBadRequest, "Post method only allow")
// 		}

// 		params := &userPostRequest{}
// 		if err := c.Bind(params); err != nil {
// 			c.JSON(http.StatusInternalServerError, err.Error())
// 		}
// 		userID, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err.Error())
// 		}

// 		updateUser, err := uh.usecase.UpdateUser(ctx, userID, params.Username, params.Email, params.Password, params.Age)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, err.Error())
// 		}

// 		var userRes userResponse

// 		userRes.ID = updateUser.ID
// 		userRes.Username = updateUser.Username
// 		userRes.Email = updateUser.Email
// 		userRes.Age = updateUser.Age

// 		return c.JSON(http.StatusOK, userRes)

// 	}
// }

// func (uh *userHandler) DeleteUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := c.Request().Context()

// 		if c.Request().Method != "DELETE" {
// 			return c.JSON(http.StatusBadRequest, "DELETE method only allow")
// 		}

// 		userID, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err.Error())
// 		}

// 		_, err = uh.usecase.DeleteUser(ctx, userID)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, err.Error())
// 		}

// 		return c.JSON(http.StatusOK, "succsess!")
// 	}
// }

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
