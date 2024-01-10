package handler

import (
	"encoding/json"
	"net/http"
	"supermarket/internal/repository"
	"supermarket/internal/service"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func writeResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{Message: message, Data: data})
}

func writeError(w http.ResponseWriter, statusCode int, err error) {
	writeResponse(w, statusCode, err.Error(), nil)
}

type Handler struct {
	ProductService *service.ProductService
}

// GetPingHandler returns a pong message.
func (h *Handler) GetPingHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Ping successful", "pong")
}

// GetProductsHandler returns the products from the repository.
func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := h.ProductService.GetProducts()
	writeResponse(w, http.StatusOK, "Products fetched successfully", products)
}

// GetProductHandler returns a product from the repository by id.
func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	product, err := h.ProductService.GetProduct(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeResponse(w, http.StatusOK, "Product fetched successfully", product)
}

// SearchProductsHandler returns the products from the repository that have a price greater than priceGt.
func (h *Handler) SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.SearchProducts(r.URL.Query().Get("priceGt"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeResponse(w, http.StatusOK, "Products fetched successfully", products)
}

// AddProductHandler adds a product to the repository.
func (h *Handler) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var product repository.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	product, err = h.ProductService.AddProduct(product)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeResponse(w, http.StatusOK, "Product added successfully", product)
}
