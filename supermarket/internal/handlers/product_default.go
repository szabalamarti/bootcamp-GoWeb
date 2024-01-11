package handlers

import (
	"errors"
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
	response.Text(w, http.StatusOK, "pong")
}

// GetProductsHandler returns the products from the repository.
func (h *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := h.ProductService.GetProducts()
	response.JSON(w, http.StatusOK, "products fetched successfully", products)
}

// GetProductHandler returns a product from the repository by id.
func (h *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	product, err := h.ProductService.GetProduct(chi.URLParam(r, "id"))
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrInvalidID):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internal.ErrProductNotFound):
			response.Errorw(w, http.StatusNotFound, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response.JSON(w, http.StatusOK, "product fetched successfully", product)
}

// SearchProductsHandler returns the products from the repository that have a price greater than priceGt.
func (h *ProductHandler) SearchProductsByPriceHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.SearchProductsByPrice(r.URL.Query().Get("priceGt"))
	if err != nil {
		switch err {
		case internal.ErrInvalidPriceGt:
			response.Errorw(w, http.StatusBadRequest, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response.JSON(w, http.StatusOK, "products fetched successfully", products)
}

// CreateProductHandler adds a product to the repository.
func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := request.JSON(r, &product)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	product, err = h.ProductService.CreateProduct(product)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrInvalidProduct):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internal.ErrDuplicateCodeValue):
			response.Errorw(w, http.StatusConflict, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response.JSON(w, http.StatusOK, "product created successfully", product)
}

// UpdateOrCreateProductHandler updates a product in the repository or creates it if it doesn't exist.
func (h *ProductHandler) UpdateOrCreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := request.JSON(r, &product)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	product, err = h.ProductService.UpdateOrCreateProduct(product)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrInvalidProduct):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internal.ErrDuplicateCodeValue):
			response.Errorw(w, http.StatusConflict, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response.JSON(w, http.StatusOK, "product updated or created successfully", product)
}

// UpdateProductHandler updates a product in the repository.
func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var originalProduct Product
	originalProduct, err := h.ProductService.GetProduct(chi.URLParam(r, "id"))
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrInvalidID):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internal.ErrProductNotFound):
			response.Errorw(w, http.StatusNotFound, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	err = request.JSON(r, &originalProduct)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	originalProduct, err = h.ProductService.UpdateProduct(originalProduct)
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrInvalidProduct):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internal.ErrDuplicateCodeValue):
			response.Errorw(w, http.StatusConflict, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response.JSON(w, http.StatusOK, "Product updated successfully", originalProduct)
}

// DeleteProductHandler deletes a product from the repository by id.
func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	err := h.ProductService.DeleteProduct(chi.URLParam(r, "id"))
	if err != nil {
		switch {
		case errors.Is(err, internal.ErrInvalidID):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internal.ErrProductNotFound):
			response.Errorw(w, http.StatusNotFound, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
	}

	response.JSON(w, http.StatusOK, "Product deleted successfully", nil)
}
