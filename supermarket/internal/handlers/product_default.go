package handlers

import (
	"net/http"
	"supermarket/internal"
	"supermarket/internal/platform/web/request"
	"supermarket/internal/platform/web/response"

	"github.com/go-chi/chi/v5"
)

type Product = internal.Product
type ProductServiceInterface = internal.ProductServiceInterface

type ProductHandler struct {
	ProductService ProductServiceInterface
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
		switch err {
		case internal.ErrInvalidID:
			response.WriteError(w, http.StatusBadRequest, err)
		case internal.ErrProductNotFound:
			response.WriteError(w, http.StatusNotFound, err)
		default:
			response.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	response.WriteResponseJSON(w, http.StatusOK, "Product fetched successfully", product)
}

// SearchProductsHandler returns the products from the repository that have a price greater than priceGt.
func (h *ProductHandler) SearchProductsByPriceHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.SearchProductsByPrice(r.URL.Query().Get("priceGt"))
	if err != nil {
		switch err {
		case internal.ErrInvalidPriceGt:
			response.WriteError(w, http.StatusBadRequest, err)
		default:
			response.WriteError(w, http.StatusInternalServerError, err)
		}
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
		switch err {
		case internal.ErrInvalidProduct:
			response.WriteError(w, http.StatusBadRequest, err)
		case internal.ErrDuplicateCodeValue:
			response.WriteError(w, http.StatusConflict, err)
		default:
			response.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	response.WriteResponseJSON(w, http.StatusOK, "Product created successfully", product)
}

// UpdateOrCreateProductHandler updates a product in the repository or creates it if it doesn't exist.
func (h *ProductHandler) UpdateOrCreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := request.ReadRequestJSON(r, &product)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product, err = h.ProductService.UpdateOrCreateProduct(product)
	if err != nil {
		switch err {
		case internal.ErrInvalidProduct:
			response.WriteError(w, http.StatusBadRequest, err)
		case internal.ErrDuplicateCodeValue:
			response.WriteError(w, http.StatusConflict, err)
		default:
			response.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	response.WriteResponseJSON(w, http.StatusOK, "Product updated or created successfully", product)
}

// UpdateProductHandler updates a product in the repository.
func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var originalProduct Product
	originalProduct, err := h.ProductService.GetProduct(chi.URLParam(r, "id"))
	if err != nil {
		switch err {
		case internal.ErrInvalidID:
			response.WriteError(w, http.StatusBadRequest, err)
		case internal.ErrProductNotFound:
			response.WriteError(w, http.StatusNotFound, err)
		default:
			response.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	err = request.ReadRequestJSON(r, &originalProduct)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}

	originalProduct, err = h.ProductService.UpdateProduct(originalProduct)
	if err != nil {
		switch err {
		case internal.ErrInvalidProduct:
			response.WriteError(w, http.StatusBadRequest, err)
		case internal.ErrDuplicateCodeValue:
			response.WriteError(w, http.StatusConflict, err)
		default:
			response.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	response.WriteResponseJSON(w, http.StatusOK, "Product updated successfully", originalProduct)
}

// DeleteProductHandler deletes a product from the repository by id.
func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	err := h.ProductService.DeleteProduct(chi.URLParam(r, "id"))
	if err != nil {
		switch err {
		case internal.ErrInvalidID:
			response.WriteError(w, http.StatusBadRequest, err)
		case internal.ErrProductNotFound:
			response.WriteError(w, http.StatusNotFound, err)
		default:
			response.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	response.WriteResponseJSON(w, http.StatusOK, "Product deleted successfully", nil)
}
