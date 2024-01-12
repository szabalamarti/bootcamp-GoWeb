package storage

import (
	"encoding/json"
	"os"
	internalProduct "supermarket/internal/product"
)

type Product = internalProduct.Product

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
		return nil, internalProduct.ErrFileNotFound
	}
	defer file.Close()

	var productsSlice []Product
	err = json.NewDecoder(file).Decode(&productsSlice)
	if err != nil {
		return nil, internalProduct.ErrInvalidFile
	}

	productsMap := make(map[int]Product)
	for _, product := range productsSlice {
		productsMap[product.Id] = product
	}

	return productsMap, nil
}

// SaveProducts saves the products from the repository to a JSON file.
func (ps *ProductStorage) SaveProducts(products map[int]Product) error {
	file, err := os.Create(ps.filename)
	if err != nil {
		return internalProduct.ErrSaveProducts
	}
	defer file.Close()

	// Convert map to slice
	var productsSlice []Product
	for _, product := range products {
		productsSlice = append(productsSlice, product)
	}

	err = json.NewEncoder(file).Encode(productsSlice)
	if err != nil {
		return internalProduct.ErrSaveProducts
	}

	return nil
}
