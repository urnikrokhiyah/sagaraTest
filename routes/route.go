package routes

import (
	"test/constants"
	"test/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	e.GET("/products", controllers.GetListProductController)
	r.POST("/products", controllers.CreateNewProductController)
	e.GET("/products/:id", controllers.GetProductDetailController)
	r.DELETE("/products/:id", controllers.DeleteProductController)
	r.PUT("/products/:id", controllers.UpdateProductController)
	r.POST("/product/:id/images", controllers.UpdateImageProductController)

	e.POST("/register/:role", controllers.RegisterUserController)
	e.POST("/login", controllers.LoginUserController)

	return e
}
