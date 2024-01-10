package handlers

import (
	"net/http"
	"supermarket/internal"
	"supermarket/internal/platform/web/request"
	"supermarket/internal/platform/web/response"

	"github.com/go-chi/chi/v5"
)

type Product = internal.Product

type ServiceInterface interface {
	GetProducts() []Product
	GetProduct(id string) (Product, error)
	SearchProducts(priceGt string) ([]Product, error)
	CreateProduct(product Product) (Product, error)
}

type ProductHandler struct {
	ProductService ServiceInterface
}

// GetPingHandler returns a pong message.
func (h *ProductHandler) GetPingHandler(w http.ResponseWriter, r *http.Request) {
	response.WriteResponseText(w, http.StatusOK, "pong")
}

// GetProductsHandler returns the products from the repository.
func (h *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := h.ProductService.GetProducts()
	response.WriteResponseJSON(w, http.StatusOK, "Products fetched successfully", products)
}

// GetProductHandler returns a product from the repository by id.
func (h *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	product, err := h.ProductService.GetProduct(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response.WriteResponseJSON(w, http.StatusOK, "Product fetched successfully", product)
}

// SearchProductsHandler returns the products from the repository that have a price greater than priceGt.
func (h *ProductHandler) SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.SearchProducts(r.URL.Query().Get("priceGt"))
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response.WriteResponseJSON(w, http.StatusOK, "Products fetched successfully", products)
}

// CreateProductHandler adds a product to the repository.
func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := request.ReadRequestJSON(r, &product)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product, err = h.ProductService.CreateProduct(product)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response.WriteResponseJSON(w, http.StatusOK, "Product created successfully", product)
}
