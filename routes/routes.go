package routes

import (
	"ci_cd/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", controllers.GetAllPosts)
	e.POST("/posts", controllers.CreatePost)
}
