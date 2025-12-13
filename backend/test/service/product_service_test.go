package service

import (
	"os"
	"testing"

	"github.com/denizdoganx/product-app/domain"
	"github.com/denizdoganx/product-app/service"
	"github.com/denizdoganx/product-app/service/model"
	"github.com/stretchr/testify/assert"
)

var productService service.IProductService

func TestMain(m *testing.M) {
	initialProducts := []domain.Product{
		{
			Id:    1,
			Name:  "AirFryer",
			Price: 1000,
			Store: "ABC TECH",
		},
		{
			Id:    2,
			Name:  "Ütü",
			Price: 4000,
			Store: "ABC TECH",
		},
	}

	fakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductRepository)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("ShouldGetAllProducts", func(t *testing.T) {
		products, _ := productService.GetAllProducts()

		assert.Equal(t, 2, len(products))
	})
}

func Test_WhenNoValidationErrorOccurred_ShouldAddProduct(t *testing.T) {
	t.Run("WhenNoValidationErrorOccurred_ShouldAddProduct", func(t *testing.T) {
		productService.AddProduct(model.ProductCreate{
			Name:     "Ütü",
			Price:    2000,
			Discount: 50,
			Store:    "ABC TECH",
		})

		actualProducts, _ := productService.GetAllProducts()
		assert.Equal(t, 3, len(actualProducts))
	})
}

func Test_WhenPriceIsLowerOrEqualToZero_ShouldNotAddProduct(t *testing.T) {
	t.Run("WhenPriceIsLowerOrEqualToZero_ShouldNotAddProduct", func(t *testing.T) {
		err := productService.AddProduct(model.ProductCreate{
			Name:     "Ütü",
			Price:    -5,
			Discount: 50,
			Store:    "ABC TECH",
		})

		actualProducts, _ := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))
		assert.Equal(t, "product price can not be lower or equal to 0", err.Error())
	})
}
