package controllers

import (
	"app/configs"
	"app/models/user"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {
	var userRegister user.UserRegister
	c.Bind(&userRegister)
	var userDB user.User
	password, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 14)
	userDB.Name = userRegister.Name
	userDB.Email = userRegister.Email
	userDB.Password = password
	err := configs.DB.Create(&userRegister).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, user.UserResponses{
			false, "failed register user database", nil,
		})
	}

	return c.JSON(http.StatusOK, user.UserResponse{
		true, "success register user database", userDB,
	})
}
