package service

import (
	"github.com/denizdoganx/product-app/domain"
	"github.com/denizdoganx/product-app/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeProductRepository *FakeProductRepository) GetAllProducts() ([]domain.Product, error) {
	return fakeProductRepository.products, nil
}

func (fakeProductRepository *FakeProductRepository) GetAllProductsByStore(storeName string) ([]domain.Product, error) {
	return nil, nil
}

func (fakeProductRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeProductRepository.products = append(fakeProductRepository.products, domain.Product{
		Id:       int64(len(fakeProductRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})

	return nil
}

func (fakeProductRepository *FakeProductRepository) GetProductById(id int64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (fakeProductRepository *FakeProductRepository) DeleteProductById(id int64) error {
	return nil
}

func (fakeProductRepository *FakeProductRepository) UpdateProductPrice(id int64, newPrice float32) error {
	return nil
}
