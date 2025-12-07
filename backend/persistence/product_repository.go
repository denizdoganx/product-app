package persistence

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"

	"github.com/denizdoganx/product-app/domain"
)

type IProductRepository interface {
	GetAllProducts() ([]domain.Product, error)
	GetAllProductsByStore(storeName string) ([]domain.Product, error)
	AddProduct(product domain.Product) error
	GetProductById(id int64) (domain.Product, error)
	DeleteProductById(id int64) error
	UpdateProductPrice(id int64, newPrice float32) error
}

type ProductRepository struct {
	dbPool *sql.DB
}

func NewProductRepository(dbPool *sql.DB) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

func (productRepository *ProductRepository) GetAllProducts() ([]domain.Product, error) {
	productRows, err := productRepository.dbPool.Query("SELECT * FROM products")

	if err != nil {
		log.Printf("error while getting all products %v", err)

		return []domain.Product{}, err
	}

	products, err := extractProductFromRows(productRows)

	if err != nil {
		return []domain.Product{}, err
	} else {
		return products, nil
	}
}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) ([]domain.Product, error) {
	productRows, err := productRepository.dbPool.Query("SELECT * FROM products WHERE store = ?", storeName)

	if err != nil {
		log.Printf("error while getting all products by store %v\n", err)

		return []domain.Product{}, err
	}

	products, err := extractProductFromRows(productRows)

	if err != nil {
		return []domain.Product{}, err
	} else {
		return products, nil
	}
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	sqlStatement := "INSERT INTO products(name, price, discount, store) VALUES (?, ?, ?, ?)"
	_, err := productRepository.dbPool.Exec(sqlStatement, product.Name, product.Price, product.Discount, product.Store)

	if err != nil {
		log.Printf("error while adding product to products %v\n", err)
		return err
	}

	return nil
}

func (productRepository *ProductRepository) GetProductById(id int64) (domain.Product, error) {
	sqlStatement := "SELECT * FROM products WHERE id = ?"
	productRow, err := productRepository.dbPool.Query(sqlStatement, id)

	if err != nil {
		return domain.Product{}, err
	}

	products, err := extractProductFromRows(productRow)

	if err != nil {
		return domain.Product{}, err
	} else {
		if len(products) > 0 {
			return products[0], nil
		} else {
			return domain.Product{}, fmt.Errorf("there is no product with given id %d", id)
		}
	}
}

func (productRepository *ProductRepository) DeleteProductById(id int64) error {
	sqlStatement := "DELETE FROM products WHERE id = ?"
	_, err := productRepository.dbPool.Exec(sqlStatement, id)

	if err != nil {
		log.Printf("error while deleting product in the products %v\n", err)
		return err
	}

	return nil
}

func (productRepository *ProductRepository) UpdateProductPrice(id int64, newPrice float32) error {
	sqlStatement := "UPDATE products SET price = ? WHERE id = ?"
	_, err := productRepository.dbPool.Exec(sqlStatement, newPrice, id)

	if err != nil {
		log.Printf("error while updating product's price in the products %v\n", err)
		return err
	}

	return nil
}

func extractProductFromRows(productRows *sql.Rows) ([]domain.Product, error) {
	var (
		products = []domain.Product{}
		id       int64
		name     string
		price    float32
		discount float32
		store    string
	)

	for productRows.Next() {
		err := productRows.Scan(&id, &name, &price, &discount, &store)

		if err != nil {
			return []domain.Product{}, errors.New("error occured while scanning products")
		} else {
			products = append(products, domain.Product{
				Id:       id,
				Name:     name,
				Price:    price,
				Discount: discount,
				Store:    store,
			})
		}
	}

	return products, nil
}
