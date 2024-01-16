package repository

import (
	"strconv"
	internalProduct "supermarket/internal/product"
)

type Product = internalProduct.Product

type ProductRepository struct {
	Storage  internalProduct.ProductStorageInterface
	Products map[int]Product
	LastId   int
}

// NewProductRepository creates a new ProductRepository.
func NewProductRepository(storage internalProduct.ProductStorageInterface) *ProductRepository {
	return &ProductRepository{
		Storage: storage,
	}
}

// LoadProducts loads products from storage to the repository.
func (pr *ProductRepository) LoadProducts() error {
	products, err := pr.Storage.LoadProducts()
	if err != nil {
		return err
	}
	pr.Products = products

	// Set LastId to the highest product ID
	pr.LastId = 0
	for id := range pr.Products {
		if id > pr.LastId {
			pr.LastId = id
		}
	}

	return nil
}

// SaveProducts saves products from the repository to storage.
func (pr *ProductRepository) SaveProducts() error {
	return pr.Storage.SaveProducts(pr.Products)
}

// Get returns all products from the repository.
func (pr *ProductRepository) Get() ([]Product, error) {
	err := pr.LoadProducts()
	if err != nil {
		return nil, err
	}
	products := make([]Product, 0, len(pr.Products))
	for _, product := range pr.Products {
		products = append(products, product)
	}
	return products, nil
}

// GetById returns a product from the repository by id.
func (pr *ProductRepository) GetById(id int) (Product, error) {
	pr.LoadProducts()
	product, ok := pr.Products[id]
	if !ok {
		return Product{}, internalProduct.ErrProductNotFound
	}
	return product, nil
}

// SearchByPrice returns the products from the repository that have a price greater than priceGt.
func (pr *ProductRepository) SearchByPrice(priceGt float64) ([]Product, error) {
	pr.LoadProducts()
	products, err := pr.Get()
	if err != nil {
		return nil, err
	}
	var filteredProducts []Product
	for _, product := range products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts, nil
}

// Save adds a product to the repository.
func (pr *ProductRepository) Save(product Product) (Product, error) {
	pr.LoadProducts()
	pr.LastId++
	product.Id = pr.LastId
	pr.Products[product.Id] = product
	pr.SaveProducts()
	return product, nil
}

// SaveOrUpdate updates a product in the repository or creates it if it doesn't exist.
func (pr *ProductRepository) SaveOrUpdate(product Product) (Product, error) {
	pr.LoadProducts()
	_, ok := pr.Products[product.Id]
	if !ok {
		return pr.Save(product)
	}
	pr.Products[product.Id] = product
	pr.SaveProducts()
	return product, nil
}

// Update updates a product in the repository.
func (pr *ProductRepository) Update(product Product) (Product, error) {
	pr.LoadProducts()
	_, ok := pr.Products[product.Id]
	if !ok {
		return Product{}, internalProduct.ErrProductNotFound
	}
	pr.Products[product.Id] = product
	pr.SaveProducts()
	return product, nil
}

// Delete deletes a product from the repository by id.
func (pr *ProductRepository) Delete(id int) error {
	pr.LoadProducts()
	_, ok := pr.Products[id]
	if !ok {
		return internalProduct.ErrProductNotFound
	}
	delete(pr.Products, id)
	pr.SaveProducts()
	return nil
}

// GetConsumerPriceProducts receives a list of ids and returns those products and the total price.
func (pr *ProductRepository) GetConsumerPriceProducts(ids []string) (internalProduct.ConsumerPriceProducts, error) {
	consumerProducts := internalProduct.ConsumerPriceProducts{
		Products:   []Product{},
		TotalPrice: 0,
	}

	if ids[0] == "" {
		products, err := pr.Get()
		if err != nil {
			return consumerProducts, err
		}
		for _, product := range products {
			consumerProducts.TotalPrice += product.Price
			consumerProducts.Products = append(consumerProducts.Products, product)
		}
	} else {
		quantityMap := make(map[string]int)
		for _, id := range ids {
			productId, err := strconv.Atoi(id)
			if err != nil {
				return consumerProducts, internalProduct.ErrInvalidID
			}

			product, err := pr.GetById(productId)
			if err != nil {
				return consumerProducts, internalProduct.ErrProductNotFound
			}

			quantityMap[id]++
			if quantityMap[id] > product.Quantity {
				return consumerProducts, internalProduct.ErrInsufficientQuantity
			}
			consumerProducts.TotalPrice += product.Price
			consumerProducts.Products = append(consumerProducts.Products, product)
		}
	}

	totalProducts := len(consumerProducts.Products)
	switch {
	case totalProducts < 10:
		consumerProducts.TotalPrice *= 1.21
	case totalProducts <= 20:
		consumerProducts.TotalPrice *= 1.17
	default:
		consumerProducts.TotalPrice *= 1.15
	}

	return consumerProducts, nil
}
