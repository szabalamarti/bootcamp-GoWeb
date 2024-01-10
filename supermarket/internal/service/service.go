package service

import (
	"errors"
	"strconv"
	"supermarket/internal/repository"
	"time"
)

var (
	errInvalidID          = errors.New("invalid id")
	errProductNotFound    = errors.New("product not found")
	errInvalidPriceGt     = errors.New("invalid priceGt")
	errInvalidProduct     = errors.New("invalid product parameters")
	errDuplicateCodeValue = errors.New("duplicated code value")
)

type ProductService struct {
	ProductRepository *repository.ProductRepository
}

// GetProducts returns the products from the repository.
func (ps *ProductService) GetProducts() []repository.Product {
	return ps.ProductRepository.Products
}

// GetProduct returns a product from the repository by id.
func (ps *ProductService) GetProduct(id string) (repository.Product, error) {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return repository.Product{}, errInvalidID
	}

	for _, product := range ps.ProductRepository.Products {
		if product.Id == productId {
			return product, nil
		}
	}

	return repository.Product{}, errProductNotFound
}

// SearchProducts returns the products from the repository that have a price greater than priceGt.
func (ps *ProductService) SearchProducts(priceGt string) ([]repository.Product, error) {
	price, err := strconv.ParseFloat(priceGt, 64)
	if err != nil {
		return nil, errInvalidPriceGt
	}

	var products []repository.Product
	for _, product := range ps.ProductRepository.Products {
		if product.Price > price {
			products = append(products, product)
		}
	}

	return products, nil
}

// AddProduct adds a product to the repository.
func (ps *ProductService) AddProduct(product repository.Product) (repository.Product, error) {
	// no value can be empty, except is_published, where empty is false
	if product.Name == "" || product.Quantity == 0 || product.CodeValue == "" || product.Expiration == "" || product.Price == 0 {
		return product, errInvalidProduct
	}

	// update id to latest id + 1
	product.Id = ps.ProductRepository.LastId + 1

	// check if CodeValue already exists
	for _, p := range ps.ProductRepository.Products {
		if p.CodeValue == product.CodeValue {
			return product, errDuplicateCodeValue
		}
	}

	// product.Expiration mut be in MM/DD/YYYY format
	_, err := time.Parse("01/02/2006", product.Expiration)
	if err != nil {
		return product, errInvalidProduct
	}

	// add product to repository
	ps.ProductRepository.AddProduct(product)
	return product, nil

}
