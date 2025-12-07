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
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, actualProducts, expectedProducts)
	})

	clear(ctx, databaseInstance)
}
