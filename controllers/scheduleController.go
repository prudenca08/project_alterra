package controllers

import (
	"clinic/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetScheduleControllers(c echo.Context) error{
	schedules, e := database.GetSchedules()
	if e != nil{
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"status":"success","schedule":schedules,})
}