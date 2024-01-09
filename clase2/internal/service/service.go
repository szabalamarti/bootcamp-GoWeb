package service

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

type ProductService struct {
	Products []Product
}

var (
	errLoadProducts = errors.New("failed to load products")
	errUnmarshal    = errors.New("failed to unmarshal products")
)

func (ps *ProductService) LoadProducts(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("%w: %v", errLoadProducts, err)
	}

	err = json.Unmarshal(data, &ps.Products)
	if err != nil {
		return fmt.Errorf("%w: %v", errUnmarshal, err)
	}

	return nil
}

func (ps *ProductService) PrintProductsInfo() {
	totalProducts := len(ps.Products)
	fmt.Printf("Loaded a total of %d products to service.\n", totalProducts)
}
