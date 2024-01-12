package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"supermarket/internal/platform/web/request"
	"supermarket/internal/platform/web/response"
	"supermarket/internal/platform/web/serialization"
	"supermarket/internal/platform/web/validator"
	internalProduct "supermarket/internal/product"

	"github.com/go-chi/chi/v5"
)

type Product = internalProduct.Product

type ProductServiceInterface = internalProduct.ProductServiceInterface

type ProductHandler struct {
	ProductService ProductServiceInterface
}

// NewProductHandler returns a new ProductHandler.
func NewProductHandler(productService ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
	}
}

// GetPingHandler returns a pong message.
func (h *ProductHandler) GetPingHandler(w http.ResponseWriter, r *http.Request) {
	response.Text(w, http.StatusOK, "pong")
}

// GetProductsHandler returns the products from the repository.
func (h *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := h.ProductService.GetProducts()
	// serialize products to ProductResponseJSON
	productsResponse := serialization.ProductsToProductsResponse(products)
	response.JSON(w, http.StatusOK, "products fetched successfully", productsResponse)
}

// GetProductHandler returns a product from the repository by id.
func (h *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	product, err := h.ProductService.GetProduct(chi.URLParam(r, "id"))
	if err != nil {
		switch {
		case errors.Is(err, internalProduct.ErrInvalidID):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internalProduct.ErrProductNotFound):
			response.Errorw(w, http.StatusNotFound, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	// serialize product to ProductResponseJSON
	productResponse := serialization.ProductToProductResponse(product)
	response.JSON(w, http.StatusOK, "product fetched successfully", productResponse)
}

// SearchProductsHandler returns the products from the repository that have a price greater than priceGt.
func (h *ProductHandler) SearchProductsByPriceHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.SearchProductsByPrice(r.URL.Query().Get("priceGt"))
	if err != nil {
		switch err {
		case internalProduct.ErrInvalidPriceGt:
			response.Errorw(w, http.StatusBadRequest, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	// serialize products to ProductResponseJSON
	productsResponse := serialization.ProductsToProductsResponse(products)
	response.JSON(w, http.StatusOK, "products fetched successfully", productsResponse)
}

// CreateProductHandler adds a product to the repository.
func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	// read product from request
	var productRequest serialization.ProductRequest
	err := request.JSON(r, &productRequest)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	// deserialize productRequest to Product
	product := serialization.ProductRequestToProduct(productRequest)

	// create product
	product, err = h.ProductService.CreateProduct(product)
	if err != nil {
		switch {
		case errors.Is(err, internalProduct.ErrInvalidProduct):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internalProduct.ErrDuplicateCodeValue):
			response.Errorw(w, http.StatusConflict, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	// serialize product to ProductResponseJSON
	productResponse := serialization.ProductToProductResponse(product)
	response.JSON(w, http.StatusOK, "product created successfully", productResponse)
}

// UpdateOrCreateProductHandler updates a product in the repository or creates it if it doesn't exist.
func (h *ProductHandler) UpdateOrCreateProductHandler(w http.ResponseWriter, r *http.Request) {
	// get id from url
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	// get body to []byte
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	// create a map to validate the body
	var bodyMap map[string]any
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	err = validator.ValidateRequiredKeys(bodyMap, "name", "quantity", "code_value", "expiration", "price")
	if err != nil {
		response.Errorw(w, http.StatusBadRequest, err)
		return
	}

	// read product from bytes
	var productRequest serialization.ProductRequest
	err = json.Unmarshal(body, &productRequest)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	// deserialize productRequest to Product
	product := serialization.ProductRequestToProduct(productRequest)
	product.Id = id

	// update or create product
	product, err = h.ProductService.UpdateOrCreateProduct(product)
	if err != nil {
		switch {
		case errors.Is(err, internalProduct.ErrInvalidProduct):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internalProduct.ErrDuplicateCodeValue):
			response.Errorw(w, http.StatusConflict, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	// serialize product to ProductResponse
	productResponse := serialization.ProductToProductResponse(product)
	response.JSON(w, http.StatusOK, "product updated or created successfully", productResponse)
}

// UpdateProductHandler updates a product in the repository.
func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	// get id from url
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	// find original product to patch
	var originalProduct Product
	originalProduct, err = h.ProductService.GetProduct(chi.URLParam(r, "id"))
	if err != nil {
		switch {
		case errors.Is(err, internalProduct.ErrInvalidID):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internalProduct.ErrProductNotFound):
			response.Errorw(w, http.StatusNotFound, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	// serialize originalProduct to ProductRequest
	updateProductRequest := serialization.ProductToProductRequest(originalProduct)

	// read productRequest from request into updateProductRequest
	err = request.JSON(r, &updateProductRequest)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "bad request")
		return
	}

	// deserialize updateProductRequest to Product
	updateProduct := serialization.ProductRequestToProduct(updateProductRequest)
	updateProduct.Id = id

	// update product
	updateProduct, err = h.ProductService.UpdateProduct(updateProduct)
	if err != nil {
		switch {
		case errors.Is(err, internalProduct.ErrInvalidProduct):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internalProduct.ErrDuplicateCodeValue):
			response.Errorw(w, http.StatusConflict, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	// serialize updateProduct to ProductResponseJSON
	updateProductResponse := serialization.ProductToProductResponse(updateProduct)
	response.JSON(w, http.StatusOK, "Product updated successfully", updateProductResponse)
}

// DeleteProductHandler deletes a product from the repository by id.
func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	// delete product
	err := h.ProductService.DeleteProduct(chi.URLParam(r, "id"))
	if err != nil {
		switch {
		case errors.Is(err, internalProduct.ErrInvalidID):
			response.Errorw(w, http.StatusBadRequest, err)
		case errors.Is(err, internalProduct.ErrProductNotFound):
			response.Errorw(w, http.StatusNotFound, err)
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response.Text(w, http.StatusOK, "Product deleted successfully")
}
