package routers

import (
	"github.com/echo_CRUD/src/controllers"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/api/v1")
	v1.GET("/users", controllers.GetAllUser)
	v1.POST("/users", controllers.CreateUser)

	return e
}
