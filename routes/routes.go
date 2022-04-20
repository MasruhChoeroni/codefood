package routes

import (
	"codefood/controllers"
	"codefood/db"
	"codefood/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	con := db.CreateCon()
	con.Query("CREATE TABLE `cashiers` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT,`name` varchar(255) NOT NULL DEFAULT '',`passcode` varchar(255) NOT NULL DEFAULT '',`created_at` datetime DEFAULT NULL,`updated_at` datetime DEFAULT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB;")
	con.Query("CREATE TABLE `categories` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT,`name` varchar(255) NOT NULL DEFAULT '',`created_at` datetime DEFAULT NULL,`updated_at` datetime DEFAULT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB;")
	con.Query("CREATE TABLE `payment_types` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT,`name` varchar(255) NOT NULL DEFAULT '',`type` varchar(255) NOT NULL DEFAULT '',`logo` varchar(255) NOT NULL DEFAULT '',`created_at` datetime DEFAULT NULL,`updated_at` datetime DEFAULT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB;")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Back-end: Point of Sales API")
	})

	e.GET("/cashiers", controllers.FindAllCashiers)
	e.GET("/cashiers/:id", controllers.FindCashiersById)
	e.POST("/cashiers/:id/login", controllers.LoginCashiers)
	e.POST("/cashiers", controllers.StoreCashiers)
	e.PUT("/cashiers/:id", controllers.UpdateCashiers)
	e.DELETE("/cashiers/:id", controllers.DeleteCashiers)

	e.GET("/cashiers/:id/passcode", controllers.FindCashiersPasscodeById)

	e.GET("/categories", controllers.FindAllCategories, middleware.IsAuthenticated)
	e.GET("/categories/:id", controllers.FindCategoriesById, middleware.IsAuthenticated)
	e.POST("/categories", controllers.StoreCategories, middleware.IsAuthenticated)
	e.PUT("/categories/:id", controllers.UpdateCategories, middleware.IsAuthenticated)
	e.DELETE("/categories/:id", controllers.DeleteCategories, middleware.IsAuthenticated)

	e.GET("/payments", controllers.FindAllPaymentTypes, middleware.IsAuthenticated)
	e.GET("/payments/:id", controllers.FindPaymentTypesById, middleware.IsAuthenticated)
	e.POST("/payments", controllers.StorePaymentTypes, middleware.IsAuthenticated)
	e.PUT("/payments/:id", controllers.UpdatePaymentTypes, middleware.IsAuthenticated)
	e.DELETE("/payments/:id", controllers.DeletePaymentTypes, middleware.IsAuthenticated)

	return e
}
