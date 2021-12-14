package controllers

import (
	"clinic/lib/database"
	"clinic/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAdminControllers(c echo.Context)error{
	admins, e := database.GetAdmins()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"status" : "success", "admin" : admins,})
}

func RegisterController(c echo.Context) error {
	var admin models.Admin
	c.Bind(&admin)

	// validasi
	// if admin.Name == "" {
	// 	return c.JSON(http.StatusBadRequest, response.BaseResponse{
	// 		Code:    http.StatusBadRequest,
	// 		Message: "Nama masih kosong",
	// 		Data:    nil,
	// 	})
	// }

	// etc

	var userDB models.Admin
	userDB.Name = admin.Name
	userDB.Password = admin.Password
	userDB.Email = admin.Email
	userDB.Username = admin.Username


	result, err := database.CreateAdmin(userDB)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// result := config.DB.Create(&userDB)
	// if result.Error != nil {
	// 	return c.JSON(http.StatusInternalServerError, response.BaseResponse{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: "Error ketika input data user ke DB",
	// 		Data:    nil,
	// 	})
	// }

	return c.JSON(http.StatusOK, result)
}