package repository

import (
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
func (pr *ProductRepository) Get() []Product {
	pr.LoadProducts()
	products := make([]Product, 0, len(pr.Products))
	for _, product := range pr.Products {
		products = append(products, product)
	}
	return products
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
	products := pr.Get()
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
