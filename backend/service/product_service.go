package service

import (
	"errors"

	"github.com/denizdoganx/product-app/domain"
	"github.com/denizdoganx/product-app/persistence"
	"github.com/denizdoganx/product-app/service/model"
)

type IProductService interface {
	AddProduct(productCreate model.ProductCreate) error
	DeleteProductById(id int64) error
	GetProductById(id int64) (domain.Product, error)
	UpdateProductPrice(id int64, newPrice float32) error
	GetAllProducts() ([]domain.Product, error)
	GetAllProductsByStore(storeName string) ([]domain.Product, error)
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (productService *ProductService) AddProduct(productCreate model.ProductCreate) error {
	isPriceValidated := validateProductPrice(productCreate.Price)

	if !isPriceValidated {
		return errors.New("product price can not be lower or equal to 0")
	}

	return productService.productRepository.AddProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

func (productService *ProductService) DeleteProductById(id int64) error {
	return productService.productRepository.DeleteProductById(id)
}

func (productService *ProductService) GetProductById(id int64) (domain.Product, error) {
	return productService.productRepository.GetProductById(id)
}

func (productService *ProductService) UpdateProductPrice(id int64, price float32) error {
	return productService.productRepository.UpdateProductPrice(id, price)
}

func (productService *ProductService) GetAllProducts() ([]domain.Product, error) {
	return productService.productRepository.GetAllProducts()
}

func (productService *ProductService) GetAllProductsByStore(storeName string) ([]domain.Product, error) {
	return productService.productRepository.GetAllProductsByStore(storeName)
}

func validateProductPrice(price float32) bool {
	if price <= 0 {
		return false
	}

	return true
}
