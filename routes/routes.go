package routes

import (
	"clinic/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo{
	e := echo.New()
	e.GET("/admins", controllers.GetAdminControllers)
	e.GET("/admins/register", controllers.RegisterController)
	e.GET("/doctors", controllers.GetDoctorControllers)
	e.GET("/patients", controllers.GetPatientControllers)
	e.GET("/recipes", controllers.GetRecipeControllers)
	e.GET("/schedules", controllers.GetScheduleControllers)
	e.GET("/sessions",controllers.GetSessionControllers)
	return e
}