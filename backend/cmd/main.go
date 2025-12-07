package main

import (
	"context"
	"net/http"

	"github.com/denizdoganx/product-app/common/app"
	"github.com/denizdoganx/product-app/common/mysql"
	"github.com/denizdoganx/product-app/domain"
	"github.com/denizdoganx/product-app/persistence"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ProductResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func main() {
	ctx := context.Background()
	configurationManager := app.NewConfigurationManager()
	databaseInstance := mysql.GetConnectionPool(ctx, configurationManager.MySqlConfig)

	productRepository := persistence.NewProductRepository(databaseInstance)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.GET("/get-all-products", func(c echo.Context) error {

		products := productRepository.GetAllProducts()

		return c.JSON(http.StatusOK, ToResponseList(products))
	})

	e.Start("0.0.0.0:8080")
}

func ToResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}
func ToResponseList(products []domain.Product) []ProductResponse {
	var productResponseList = []ProductResponse{}
	for _, product := range products {
		productResponseList = append(productResponseList, ToResponse(product))
	}
	return productResponseList
}
