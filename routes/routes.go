package routes

import (
	"codefood/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Back-end: Point of Sales API")
	})

	e.GET("/cashiers", controllers.FindAllChasiers)
	e.POST("/cashiers", controllers.StoreChasiers)
	e.PUT("/cashiers", controllers.UpdateChasiers)
	e.DELETE("/cashiers", controllers.DeleteChasiers)
	return e
}
