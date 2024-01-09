package handler

import (
	"clase2/internal/service"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var (
	errInvalidID       = errors.New("invalid id")
	errProductNotFound = errors.New("product not found")
	errInvalidPriceGt  = errors.New("invalid priceGt")
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

func (h *Handler) GetPingHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Ping successful", "pong")
}

func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "Products fetched successfully", h.ProductService.Products)
}

func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, errInvalidID)
		return
	}

	for _, product := range h.ProductService.Products {
		if product.Id == id {
			writeResponse(w, http.StatusOK, "Product fetched successfully", product)
			return
		}
	}

	writeError(w, http.StatusNotFound, errProductNotFound)
}

func (h *Handler) SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	priceGt, err := strconv.ParseFloat(r.URL.Query().Get("priceGt"), 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, errInvalidPriceGt)
		return
	}

	var products []service.Product
	for _, product := range h.ProductService.Products {
		if product.Price > priceGt {
			products = append(products, product)
		}
	}

	writeResponse(w, http.StatusOK, "Products fetched successfully", products)
}
