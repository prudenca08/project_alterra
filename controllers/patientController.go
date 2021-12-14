package controllers

import (
	"clinic/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPatientControllers(c echo.Context)error{
	patients, e:= database.GetPatients()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"status":"success","patient" : patients,})
}