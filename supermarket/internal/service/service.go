package service

import (
	"errors"
	"strconv"
	"supermarket/internal"
	"time"
)

type Product = internal.Product

var (
	errInvalidID          = errors.New("invalid id")
	errProductNotFound    = errors.New("product not found")
	errInvalidPriceGt     = errors.New("invalid priceGt")
	errInvalidProduct     = errors.New("invalid product parameters")
	errDuplicateCodeValue = errors.New("duplicated code value")
)

type ProductRepositoryInterface interface {
	GetProducts() []Product
	GetProduct(id int) (Product, error)
	SearchProducts(priceGt float64) ([]Product, error)
	CreateProduct(product Product) (Product, error)
}

type ProductService struct {
	ProductRepository ProductRepositoryInterface
}

// GetProducts returns the products from the repository.
func (ps *ProductService) GetProducts() []Product {
	return ps.ProductRepository.GetProducts()
}

// GetProduct returns a product from the repository by id.
func (ps *ProductService) GetProduct(id string) (Product, error) {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return Product{}, errInvalidID
	}

	product, err := ps.ProductRepository.GetProduct(productId)
	if err != nil {
		return Product{}, errProductNotFound
	}

	return product, nil

}

// SearchProducts returns the products from the repository that have a price greater than priceGt.
func (ps *ProductService) SearchProducts(priceGt string) ([]Product, error) {
	price, err := strconv.ParseFloat(priceGt, 64)
	if err != nil {
		return nil, errInvalidPriceGt
	}

	products, err := ps.ProductRepository.SearchProducts(price)
	if err != nil {
		return nil, errProductNotFound
	}

	return products, nil
}

// CreateProduct adds a product to the repository.
func (ps *ProductService) CreateProduct(product Product) (Product, error) {
	// no value can be empty, except is_published, where empty is false
	if product.Name == "" || product.Quantity == 0 || product.CodeValue == "" || product.Expiration == "" || product.Price == 0 {
		return product, errInvalidProduct
	}

	// check if CodeValue already exists
	products := ps.ProductRepository.GetProducts()
	for _, p := range products {
		if p.CodeValue == product.CodeValue {
			return product, errDuplicateCodeValue
		}
	}

	// product.Expiration must be in MM/DD/YYYY format
	_, err := time.Parse("01/02/2006", product.Expiration)
	if err != nil {
		return product, errInvalidProduct
	}

	// add product to repository
	product, err = ps.ProductRepository.CreateProduct(product)
	if err != nil {
		return product, err
	}

	return product, nil
}
