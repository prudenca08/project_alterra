package controllers

import (
	"clinic/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRecipeControllers(c echo.Context)error{
	recipes, e := database.GetRecipes()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"status":"success","recipe":recipes,})
}