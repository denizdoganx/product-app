package main

import (
	"context"

	"github.com/denizdoganx/product-app/common/app"
	"github.com/denizdoganx/product-app/common/mysql"
	"github.com/denizdoganx/product-app/controller"
	"github.com/denizdoganx/product-app/persistence"
	"github.com/denizdoganx/product-app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx := context.Background()
	configurationManager := app.NewConfigurationManager()
	databaseInstance := mysql.GetConnectionPool(ctx, configurationManager.MySqlConfig)

	e := echo.New()

	productRepository := persistence.NewProductRepository(databaseInstance)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	productController.RegisterRoutes(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.Start("0.0.0.0:8080")
}
