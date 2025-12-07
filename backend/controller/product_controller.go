package controller

import (
	"github.com/denizdoganx/product-app/domain"
	"github.com/denizdoganx/product-app/service"
	"github.com/labstack/echo/v4"
)

type IProductController interface {
	GetAllProducts() []domain.Product
}

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {

}

func (productController *ProductController) GetAllProducts() []domain.Product {
	return nil
}
