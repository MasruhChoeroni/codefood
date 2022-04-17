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

	e.GET("/cashiers/:id/passcode", controllers.FindCashiersPasscodeById)

	e.GET("/categories", controllers.FindAllCategories)
	e.GET("/categories/:id", controllers.FindCategoriesById)
	e.POST("/categories", controllers.StoreCategories)
	e.PUT("/categories/:id", controllers.UpdateCategories)
	e.DELETE("/categories/:id", controllers.DeleteCategories)

	e.GET("/payments", controllers.FindAllPaymentTypes)
	e.GET("/payments/:id", controllers.FindPaymentTypesById)
	e.POST("/payments", controllers.StorePaymentTypes)
	e.PUT("/payments/:id", controllers.UpdatePaymentTypes)
	e.DELETE("/payments/:id", controllers.DeletePaymentTypes)

	return e
}
