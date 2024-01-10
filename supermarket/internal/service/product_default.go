package service

import (
	"strconv"
	"supermarket/internal"
	"time"
)

type Product = internal.Product
type ProductRepositoryInterface = internal.ProductRepositoryInterface

type ProductService struct {
	ProductRepository ProductRepositoryInterface
}

// GetProducts returns the products from the repository.
func (ps *ProductService) GetProducts() []Product {
	return ps.ProductRepository.Get()
}

// GetProduct returns a product from the repository by id.
func (ps *ProductService) GetProduct(id string) (Product, error) {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return Product{}, internal.ErrInvalidID
	}

	product, err := ps.ProductRepository.GetById(productId)
	if err != nil {
		return Product{}, internal.ErrProductNotFound
	}

	return product, nil

}

// SearchProducts returns the products from the repository that have a price greater than priceGt.
func (ps *ProductService) SearchProductsByPrice(priceGt string) ([]Product, error) {
	price, err := strconv.ParseFloat(priceGt, 64)
	if err != nil {
		return nil, internal.ErrInvalidPriceGt
	}

	products, err := ps.ProductRepository.SearchByPrice(price)
	if err != nil {
		return nil, internal.ErrProductNotFound
	}

	return products, nil
}

// ValidateProduct validates the product parameters.
func (ps *ProductService) ValidateProduct(product Product, isUpdate bool) error {
	// no value can be empty. Except is_published, where empty means false
	if product.Name == "" || product.Quantity == 0 || product.CodeValue == "" || product.Expiration == "" || product.Price == 0 {
		return internal.ErrInvalidProduct
	}

	// check if CodeValue already exists and ids are different
	products := ps.ProductRepository.Get()
	for _, p := range products {
		if p.CodeValue == product.CodeValue {
			if isUpdate {
				if p.Id != product.Id {
					return internal.ErrDuplicateCodeValue
				}
			} else {
				return internal.ErrDuplicateCodeValue
			}
		}
	}

	// product.Expiration must be in MM/DD/YYYY format
	_, err := time.Parse("01/02/2006", product.Expiration)
	if err != nil {
		return internal.ErrInvalidProduct
	}

	return nil
}

// CreateProduct adds a product to the repository.
func (ps *ProductService) CreateProduct(product Product) (Product, error) {
	// validate product
	err := ps.ValidateProduct(product, false)
	if err != nil {
		return product, err
	}

	// add product to repository
	product, err = ps.ProductRepository.Save(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

// UpdateOrCreateProduct updates a product in the repository or creates it if it doesn't exist.
func (ps *ProductService) UpdateOrCreateProduct(product Product) (Product, error) {
	// validate product
	err := ps.ValidateProduct(product, true)
	if err != nil {
		return product, err
	}

	// update product in repository
	product, err = ps.ProductRepository.SaveOrUpdate(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

// UpdateProduct updates a product in the repository.
func (ps *ProductService) UpdateProduct(product Product) (Product, error) {
	// validate product
	err := ps.ValidateProduct(product, true)
	if err != nil {
		return product, err
	}

	// update product in repository
	product, err = ps.ProductRepository.Update(product)
	if err != nil {
		return product, err
	}

	return product, nil
}

// DeleteProduct deletes a product from the repository by id.
func (ps *ProductService) DeleteProduct(id string) error {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return internal.ErrInvalidID
	}

	err = ps.ProductRepository.Delete(productId)
	if err != nil {
		return internal.ErrProductNotFound
	}

	return nil
}
