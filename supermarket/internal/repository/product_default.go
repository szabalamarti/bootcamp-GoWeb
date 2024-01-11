package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"supermarket/internal"
)

type Product = internal.Product

type ProductRepository struct {
	Products map[int]Product
	LastId   int
}

// NewProductRepository creates a new ProductRepository.
func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		Products: make(map[int]Product),
		LastId:   0,
	}
}

// PrintProductsInfo prints the total of products loaded to the repository.
func (ps *ProductRepository) PrintProductsInfo() {
	totalProducts := len(ps.Products)
	fmt.Printf("Loaded a total of %d products to service.\n", totalProducts)
}

// LoadProducts loads the products from a JSON file into the repository.
func (ps *ProductRepository) LoadProducts(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return internal.ErrLoadProducts
	}
	defer file.Close()
	var products []Product
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		return internal.ErrUnmarshal
	}
	ps.Products = make(map[int]Product)
	for _, product := range products {
		ps.Products[product.Id] = product
	}
	ps.LastId = len(ps.Products)
	return nil
}

// Get returns all products from the repository.
func (ps *ProductRepository) Get() []Product {
	products := make([]Product, 0, len(ps.Products))
	for _, product := range ps.Products {
		products = append(products, product)
	}
	return products
}

// GetById returns a product from the repository by id.
func (ps *ProductRepository) GetById(id int) (Product, error) {
	product, ok := ps.Products[id]
	if !ok {
		return Product{}, internal.ErrProductNotFound
	}
	return product, nil
}

// SearchByPrice returns the products from the repository that have a price greater than priceGt.
func (ps *ProductRepository) SearchByPrice(priceGt float64) ([]Product, error) {
	products := ps.Get()
	var filteredProducts []Product
	for _, product := range products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts, nil
}

// Save adds a product to the repository.
func (ps *ProductRepository) Save(product Product) (Product, error) {
	ps.LastId++
	product.Id = ps.LastId
	ps.Products[product.Id] = product
	return product, nil
}

// SaveOrUpdate updates a product in the repository or creates it if it doesn't exist.
func (ps *ProductRepository) SaveOrUpdate(product Product) (Product, error) {
	_, ok := ps.Products[product.Id]
	if !ok {
		return ps.Save(product)
	}
	ps.Products[product.Id] = product
	return product, nil
}

// Update updates a product in the repository.
func (ps *ProductRepository) Update(product Product) (Product, error) {
	_, ok := ps.Products[product.Id]
	if !ok {
		return Product{}, internal.ErrProductNotFound
	}
	ps.Products[product.Id] = product
	return product, nil
}

// Delete deletes a product from the repository by id.
func (ps *ProductRepository) Delete(id int) error {
	_, ok := ps.Products[id]
	if !ok {
		return internal.ErrProductNotFound
	}
	delete(ps.Products, id)
	return nil
}
