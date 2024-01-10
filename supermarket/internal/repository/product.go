package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"supermarket/internal"
)

type Product = internal.Product

type ProductRepository struct {
	Products map[int]Product
	LastId   int
}

var (
	errLoadProducts    = errors.New("failed to load products")
	errUnmarshal       = errors.New("failed to unmarshal products")
	errProductNotFound = errors.New("product not found")
)

// LoadProducts loads the products from a JSON file into the repository.
func (ps *ProductRepository) LoadProducts(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return errLoadProducts
	}
	defer file.Close()
	var products []Product
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		return errUnmarshal
	}
	ps.Products = make(map[int]Product)
	for _, product := range products {
		ps.Products[product.Id] = product
	}
	ps.LastId = len(ps.Products)
	return nil
}

// GetProducts returns all products from the repository.
func (ps *ProductRepository) GetProducts() []Product {
	products := make([]Product, 0, len(ps.Products))
	for _, product := range ps.Products {
		products = append(products, product)
	}
	return products
}

// GetProduct returns a product from the repository by id.
func (ps *ProductRepository) GetProduct(id int) (Product, error) {
	product, ok := ps.Products[id]
	if !ok {
		return Product{}, errProductNotFound
	}
	return product, nil
}

// SearchProducts returns the products from the repository that have a price greater than priceGt.
func (ps *ProductRepository) SearchProducts(priceGt float64) ([]Product, error) {
	products := ps.GetProducts()
	var filteredProducts []Product
	for _, product := range products {
		if product.Price > priceGt {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts, nil
}

// CreateProduct adds a product to the repository.
func (ps *ProductRepository) CreateProduct(product Product) (Product, error) {
	ps.LastId++
	product.Id = ps.LastId
	ps.Products[product.Id] = product
	return product, nil
}

// PrintProductsInfo prints the total of products loaded to the repository.
func (ps *ProductRepository) PrintProductsInfo() {
	totalProducts := len(ps.Products)
	fmt.Printf("Loaded a total of %d products to service.\n", totalProducts)
}
