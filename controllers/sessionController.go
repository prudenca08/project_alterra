package controllers

import (
	"clinic/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetSessionControllers(c echo.Context) error{
	sessions, e:= database.GetSessions()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"status":"success","session":sessions,})
}