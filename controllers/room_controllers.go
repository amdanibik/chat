package controllers

import (
	"app/configs"
	"app/models/rooms"
	"net/http"

	"github.com/labstack/echo"
)

func ListRooms(c echo.Context) error {
	var roomCheck []rooms.Room
	userId := c.Param("userid")
	err := configs.DB.Find(&roomCheck, "sender = ? OR receiver = ?", userId, userId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rooms.RoomResponses{
			false, "Failed Get List Room", nil,
		})
	}

	return c.JSON(http.StatusOK, rooms.RoomResponses{
		true, "Success Get List Room", roomCheck,
	})
}
