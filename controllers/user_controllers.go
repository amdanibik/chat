package controllers

import (
	"app/configs"
	"app/models/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c echo.Context) error {
	var userCreate users.UserRegister
	c.Bind(&userCreate)
	var userDB users.User
	password, _ := bcrypt.GenerateFromPassword([]byte(userCreate.Password), 14)
	userDB.Name = userCreate.Name
	userDB.Email = userCreate.Email
	userDB.Password = password
	err := configs.DB.Create(&userDB).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed register user database", nil,
		})
	}

	return c.JSON(http.StatusOK, users.UserResponse{
		true, "success register user database", userDB,
	})
}

func LoginUser(c echo.Context) error {
	var userCheck users.User
	c.Bind(&userCheck)
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(userCheck.Password), 14)
	err := configs.DB.Find(&userCheck, "email = ? AND password = ?", userCheck.Email, hashPass).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, users.UserResponses{
			false, "failed login user", nil,
		})
	}

	id := userCheck.Id
	strId := strconv.Itoa(int(id))

	return c.JSON(http.StatusOK, users.UserResponses{
		true, "success login user with id user " + strId, nil,
	})
}
