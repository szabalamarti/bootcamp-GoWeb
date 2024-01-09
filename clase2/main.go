package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
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

var productService ProductService

// LoadProducts loads products from a JSON file
func (ps *ProductService) LoadProducts(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return errors.New("Failed to load products: " + err.Error())
	}

	err = json.Unmarshal(data, &ps.Products)
	if err != nil {
		return errors.New("Failed to unmarshal products: " + err.Error())
	}

	return nil
}

// printProductsInfo prints the total number of products
func (ps *ProductService) printProductsInfo() {
	totalProducts := len(ps.Products)
	fmt.Printf("Loaded a total of %d products to service.\n", totalProducts)
}

// Handlers

// writeResponse writes a JSON response
func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// writeError writes a JSON error response
func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeResponse(w, statusCode, map[string]string{"error": message})
}

// GetPingHandler returns a simple pong response
func GetPingHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "pong")
}

// GetProductsHandler returns all the products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, productService.Products)
}

// GetProductHandler returns a single product by id
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	for _, product := range productService.Products {
		if product.Id == id {
			writeResponse(w, http.StatusOK, product)
			return
		}
	}

	writeError(w, http.StatusNotFound, "Product not found")
}

// SearchProductsHandler returns all the products with a price greater than priceGt
func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	priceGt, err := strconv.ParseFloat(r.URL.Query().Get("priceGt"), 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid priceGt")
		return
	}

	var products []Product
	for _, product := range productService.Products {
		if product.Price > priceGt {
			products = append(products, product)
		}
	}

	writeResponse(w, http.StatusOK, products)
}

func main() {
	err := productService.LoadProducts("clase2/products.json")
	if err != nil {
		panic(err)
	}
	productService.printProductsInfo()

	r := chi.NewRouter()

	r.Get("/ping", GetPingHandler)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", GetProductsHandler)
		r.Get("/{id}", GetProductHandler)
		r.Get("/search", SearchProductsHandler)
	})

	http.ListenAndServe(":8080", r)
}
