package routes

import (
	"apps/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Alhamdulillah belajar Echo Framework!")
	})

	e.GET("/users", controllers.FetchAllUsers)
	e.POST("/users", controllers.StoreUsers)
	e.PUT("/users", controllers.UpdateUsers)
	e.DELETE("/users", controllers.DeleteUsers)
	return e
}
