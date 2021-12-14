package controllers

import (
	"clinic/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDoctorControllers(c echo.Context) error {
	doctors, e := database.GetDoctors()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"status" : "success", "doctor" : doctors,})
}