package controller

import (
	"net/http"

	"github.com/denizdoganx/product-app/service"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products", productController.GetAllProducts)
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.POST("/api/v1/products", productController.AddProduct)
	e.PUT("/api/v1/products/:id", productController.UpdateProductPrice)
	e.DELETE("/api/v1/products/:id", productController.DeleteProduct)
}

func (productController *ProductController) GetAllProducts(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Hello, World!</h1>")
}

func (productController *ProductController) GetProductById(c echo.Context) error {
	return nil
}

func (productController *ProductController) AddProduct(c echo.Context) error {
	return nil
}

func (productController *ProductController) UpdateProductPrice(c echo.Context) error {
	return nil
}

func (productController *ProductController) DeleteProduct(c echo.Context) error {
	return nil
}
