package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductRepository struct {
	Products []Product
	LastId   int
}

var (
	errLoadProducts = errors.New("failed to load products")
	errUnmarshal    = errors.New("failed to unmarshal products")
)

// LoadProducts loads the products from a JSON file into the repository.
func (ps *ProductRepository) LoadProducts(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return errLoadProducts
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&ps.Products)
	if err != nil {
		return errUnmarshal
	}
	// TODO: Unsorted products
	ps.LastId = ps.Products[len(ps.Products)-1].Id

	return nil
}

// AddProduct adds a product to the repository.
func (ps *ProductRepository) AddProduct(product Product) {
	ps.Products = append(ps.Products, product)
}

// PrintProductsInfo prints the total of products loaded to the repository.
func (ps *ProductRepository) PrintProductsInfo() {
	totalProducts := len(ps.Products)
	fmt.Printf("Loaded a total of %d products to service.\n", totalProducts)
}
