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

	e.GET("/cashiers", controllers.FindAllCashiers)
	e.GET("/cashiers/:id", controllers.FindCashiersById)
	e.POST("/cashiers", controllers.StoreCashiers)
	e.PUT("/cashiers/:id", controllers.UpdateCashiers)
	e.DELETE("/cashiers/:id", controllers.DeleteCashiers)
	return e
}
