package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"supermarket/internal"
)

type Product = internal.Product

type ProductStorage struct {
	filename string
}

func NewProductStorage(filename string) *ProductStorage {
	return &ProductStorage{
		filename: filename,
	}
}

// LoadProducts loads the products from a JSON file into the repository.
func (ps *ProductStorage) LoadProducts() (map[int]Product, error) {
	file, err := os.Open(ps.filename)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", err)
		return nil, internal.ErrFileNotFound
	}
	defer file.Close()

	var productsSlice []Product
	err = json.NewDecoder(file).Decode(&productsSlice)
	if err != nil {
		fmt.Printf("Failed to decode JSON: %s\n", err)
		return nil, internal.ErrInvalidFile
	}

	productsMap := make(map[int]Product)
	for _, product := range productsSlice {
		productsMap[product.Id] = product
	}

	if len(productsMap) == 0 {
		fmt.Println("No products loaded, the file is empty or contains an empty JSON object.")
	} else {
		fmt.Printf("Loaded %d products from %s\n", len(productsMap), ps.filename)
	}

	return productsMap, nil
}

// SaveProducts saves the products from the repository to a JSON file.
func (ps *ProductStorage) SaveProducts(products map[int]Product) error {
	file, err := os.Create(ps.filename)
	if err != nil {
		return internal.ErrSaveProducts
	}
	defer file.Close()

	// Convert map to slice
	var productsSlice []Product
	for _, product := range products {
		productsSlice = append(productsSlice, product)
	}

	err = json.NewEncoder(file).Encode(productsSlice)
	if err != nil {
		return internal.ErrSaveProducts
	}

	fmt.Printf("Saved %d products to %s\n", len(productsSlice), ps.filename)
	return nil
}
