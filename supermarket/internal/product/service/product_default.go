package service

import (
	"strconv"
	internalProduct "supermarket/internal/product"
	"time"
)

type Product = internalProduct.Product
type ProductRepositoryInterface = internalProduct.ProductRepositoryInterface

type ProductService struct {
	ProductRepository ProductRepositoryInterface
}

// NewProductService creates a new ProductService.
func NewProductService(productRepository ProductRepositoryInterface) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

// GetProducts returns the products from the repository.
func (ps *ProductService) GetProducts() ([]Product, error) {
	products, err := ps.ProductRepository.Get()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetProduct returns a product from the repository by id.
func (ps *ProductService) GetProduct(id string) (Product, error) {
	productId, err := strconv.Atoi(id)
	if err != nil {
		return Product{}, internalProduct.ErrInvalidID
	}

	product, err := ps.ProductRepository.GetById(productId)
	if err != nil {
		return Product{}, internalProduct.ErrProductNotFound
	}

	return product, nil

}

// SearchProducts returns the products from the repository that have a price greater than priceGt.
func (ps *ProductService) SearchProductsByPrice(priceGt string) ([]Product, error) {
	price, err := strconv.ParseFloat(priceGt, 64)
	if err != nil {
		return nil, internalProduct.ErrInvalidPriceGt
	}

	products, err := ps.ProductRepository.SearchByPrice(price)
	if err != nil {
		return nil, internalProduct.ErrProductNotFound
	}

	return products, nil
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
		return internalProduct.ErrInvalidID
	}

	err = ps.ProductRepository.Delete(productId)
	if err != nil {
		return internalProduct.ErrProductNotFound
	}

	return nil
}

// GetConsumerPriceProducts receives a list of ids and returns those products and the total price.
func (ps *ProductService) GetConsumerPriceProducts(ids []string) (internalProduct.ConsumerPriceProducts, error) {
	consumerProducts, err := ps.ProductRepository.GetConsumerPriceProducts(ids)
	if err != nil {
		return consumerProducts, err
	}
	return consumerProducts, nil
}

// ValidateProduct validates the product parameters.
func (ps *ProductService) ValidateProduct(product Product, isUpdate bool) error {
	// no value can be empty. Except is_published, where empty means false
	if product.Name == "" || product.Quantity == 0 || product.CodeValue == "" || product.Expiration == "" || product.Price == 0 {
		return internalProduct.ErrInvalidProduct
	}

	// check if CodeValue already exists and ids are different
	products, err := ps.ProductRepository.Get()
	if err != nil {
		return err
	}
	for _, p := range products {
		if p.CodeValue != product.CodeValue {
			continue
		}
		if isUpdate && p.Id == product.Id {
			continue
		}
		return internalProduct.ErrDuplicateCodeValue
	}

	// product.Expiration must be in MM/DD/YYYY format
	_, err = time.Parse("01/02/2006", product.Expiration)
	if err != nil {
		return internalProduct.ErrInvalidProduct
	}

	return nil
}
