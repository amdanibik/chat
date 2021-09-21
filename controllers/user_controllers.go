package controllers

import (
	"app/configs"
	"app/models/users"
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	var userCreate users.Users
	c.Bind(&userCreate)
	var userDB users.User
	// password, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 14)
	userDB.Name = userCreate.Name
	userDB.Email = userCreate.Email
	userDB.Password = userCreate.Password
	err := configs.DB.Create(&userCreate).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed register user database", nil,
		})
	}

	return c.JSON(http.StatusOK, users.UserResponse{
		true, "success register user database", userDB,
	})
}
