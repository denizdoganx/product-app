package infrastructure

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/denizdoganx/product-app/common/mysql"
	"github.com/denizdoganx/product-app/domain"
	"github.com/denizdoganx/product-app/persistence"
	"github.com/stretchr/testify/assert"
)

var ctx context.Context = context.Background()
var databaseInstance *sql.DB = mysql.GetConnectionPool(ctx, mysql.Config{
	Host:                  "localhost",
	Port:                  "3307",
	Username:              "root",
	Password:              "denizalper..2023",
	DbName:                "productapp",
	MaxConnections:        10,
	MaxIdleConnections:    5,
	MaxConnectionIdleTime: 5 * time.Minute,
	Timeout:               "5s",
	ReadTimeout:           "5s",
	WriteTimeout:          "5s",
})
var productRepository persistence.IProductRepository = persistence.NewProductRepository(databaseInstance)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *sql.DB) {
	TestDataInitialize(ctx, dbPool)
}

func clear(ctx context.Context, dbPool *sql.DB) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	setup(ctx, databaseInstance)

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000,
			Discount: 22,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500,
			Discount: 10,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000,
			Discount: 15,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000,
			Discount: 0,
			Store:    "Dekorasyon Sarayı",
		},
	}

	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts, _ := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, actualProducts, expectedProducts)
	})

	clear(ctx, databaseInstance)
}

func TestGetAllProductsByStore(t *testing.T) {
	setup(ctx, databaseInstance)

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000,
			Discount: 22,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500,
			Discount: 10,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000,
			Discount: 15,
			Store:    "ABC TECH",
		},
	}

	t.Run("GetAllProductByStore", func(t *testing.T) {
		actualProducts, _ := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, actualProducts, expectedProducts)
	})

	clear(ctx, databaseInstance)
}

func TestAddProduct(t *testing.T) {
	t.Run("AddProduct", func(t *testing.T) {
		err := productRepository.AddProduct(domain.Product{
			Name:     "Beko Buzdolabı",
			Price:    43999,
			Discount: 6001,
			Store:    "Beko Mağazası",
		})

		assert.Equal(t, nil, err)
	})

	clear(ctx, databaseInstance)
}

func TestGetProductById(t *testing.T) {
	setup(ctx, databaseInstance)

	expectedProduct := domain.Product{
		Id:       1,
		Name:     "AirFryer",
		Price:    3000,
		Discount: 22,
		Store:    "ABC TECH",
	}

	t.Run("GetProductById", func(t *testing.T) {
		actualProduct, _ := productRepository.GetProductById(1)

		assert.Equal(t, expectedProduct, actualProduct)
	})

	clear(ctx, databaseInstance)
}

func TestDeleteProductById(t *testing.T) {
	setup(ctx, databaseInstance)

	t.Run("DeleteProductById", func(t *testing.T) {
		err := productRepository.DeleteProductById(4)
		products, _ := productRepository.GetAllProducts()

		assert.Equal(t, nil, err)
		assert.Equal(t, 3, len(products))
	})

	clear(ctx, databaseInstance)
}

func TestUpdateProductPrice(t *testing.T) {
	setup(ctx, databaseInstance)

	t.Run("UpdateProductPrice", func(t *testing.T) {
		err := productRepository.UpdateProductPrice(1, 5000)
		product, err2 := productRepository.GetProductById(1)

		assert.Equal(t, nil, err)
		assert.Equal(t, nil, err2)
		assert.Equal(t, float32(5000), product.Price)
	})

	clear(ctx, databaseInstance)
}
