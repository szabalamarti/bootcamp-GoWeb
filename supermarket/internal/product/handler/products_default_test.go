package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	internalProduct "supermarket/internal/product"
	"supermarket/internal/product/handler"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type ProductServiceInterface struct {
	mock.Mock
}

func (m *ProductServiceInterface) GetProducts() ([]internalProduct.Product, error) {
	args := m.Called()
	return args.Get(0).([]internalProduct.Product), args.Error(1)
}

func (m *ProductServiceInterface) CreateProduct(p internalProduct.Product) (internalProduct.Product, error) {
	args := m.Called(p)
	return args.Get(0).(internalProduct.Product), args.Error(1)
}

func (m *ProductServiceInterface) GetProduct(id string) (internalProduct.Product, error) {
	args := m.Called(id)
	return args.Get(0).(internalProduct.Product), args.Error(1)
}

func (m *ProductServiceInterface) SearchProductsByPrice(priceGt string) ([]internalProduct.Product, error) {
	args := m.Called(priceGt)
	return args.Get(0).([]internalProduct.Product), args.Error(1)
}

func (m *ProductServiceInterface) UpdateOrCreateProduct(p internalProduct.Product) (internalProduct.Product, error) {
	args := m.Called(p)
	return args.Get(0).(internalProduct.Product), args.Error(1)
}

func (m *ProductServiceInterface) UpdateProduct(p internalProduct.Product) (internalProduct.Product, error) {
	args := m.Called(p)
	return args.Get(0).(internalProduct.Product), args.Error(1)
}

func (m *ProductServiceInterface) DeleteProduct(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *ProductServiceInterface) GetConsumerPriceProducts(ids []string) (internalProduct.ConsumerPriceProducts, error) {
	args := m.Called(ids)
	return args.Get(0).(internalProduct.ConsumerPriceProducts), args.Error(1)
}

// TestGetProducts tests the GetProductsHandler method.
func TestGetProducts(t *testing.T) {
	t.Run("success - get products", func(t *testing.T) {
		// arrange
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		products := []internalProduct.Product{
			{
				Id:          1,
				Name:        "product 1",
				Quantity:    10,
				CodeValue:   "code 1",
				IsPublished: true,
				Expiration:  "2021-12-31",
				Price:       100,
			},
			{
				Id:          2,
				Name:        "product 2",
				Quantity:    20,
				CodeValue:   "code 2",
				IsPublished: true,
				Expiration:  "2021-12-31",
				Price:       200,
			},
		}

		// create the expected response
		expectedResponse := `{
            "data": [
                {
                    "id": 1,
                    "name": "product 1",
                    "quantity": 10,
                    "code_value": "code 1",
                    "is_published": true,
                    "expiration": "2021-12-31",
                    "price": 100
                },
                {
                    "id": 2,
                    "name": "product 2",
                    "quantity": 20,
                    "code_value": "code 2",
                    "is_published": true,
                    "expiration": "2021-12-31",
                    "price": 200
                }
            ],
            "message": "products fetched successfully"
        }`

		// create a mock of GetProducts method
		productService.On("GetProducts").Return(products, nil)
		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)
		// create a new request and recorder
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		rr := httptest.NewRecorder()
		// create a new handler func
		handler := http.HandlerFunc(productHandler.GetProductsHandler)

		// act
		// call the handler func
		handler.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusOK, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
		productService.AssertCalled(t, "GetProducts")
	})
}

// TestGetProduct tests the GetProductHandler method.
func TestGetProduct(t *testing.T) {
	t.Run("success - get product", func(t *testing.T) {
		// arrange
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		// create a product
		product := internalProduct.Product{
			Id:          1,
			Name:        "product 1",
			Quantity:    10,
			CodeValue:   "code 1",
			IsPublished: true,
			Expiration:  "2021-12-31",
			Price:       100,
		}
		// create the expected response
		expectedResponse := `{
			"data": {
				"id": 1,
				"name": "product 1",
				"quantity": 10,
				"code_value": "code 1",
				"is_published": true,
				"expiration": "2021-12-31",
				"price": 100
			},
			"message": "product fetched successfully"
		}`

		// create a mock of GetProduct method
		productService.On("GetProduct", "1").Return(product, nil)

		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)

		// create a new request and recorder
		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)

		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("id", "1")

		// set the RouteContext on the request context
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))

		rr := httptest.NewRecorder()
		// create a new handler func
		handler := http.HandlerFunc(productHandler.GetProductHandler)

		// act
		// call the handler func
		handler.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusOK, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
		productService.AssertCalled(t, "GetProduct", "1")
	})
	t.Run("fail - get product invalid id", func(t *testing.T) {
		// arrange
		// expected response
		expectedResponse := `{"message":"invalid id", "status":"Bad Request"}`
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		productService.On("GetProduct", "bad id").Return(internalProduct.Product{}, internalProduct.ErrInvalidID)

		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)

		// create a new request and recorder
		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)

		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("id", "bad id")

		// set the RouteContext on the request context
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
		rr := httptest.NewRecorder()

		// create a new handler func
		handler := http.HandlerFunc(productHandler.GetProductHandler)

		// act
		// call the handler func
		handler.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
	})
	t.Run("fail - get product not found", func(t *testing.T) {
		// arrange
		// expected response
		expectedResponse := `{"message":"product not found", "status":"Not Found"}`
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		productService.On("GetProduct", "1").Return(internalProduct.Product{}, internalProduct.ErrProductNotFound)

		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)

		// create a new request and recorder
		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)

		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("id", "1")

		// set the RouteContext on the request context
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
		rr := httptest.NewRecorder()

		// create a new handler func
		handler := http.HandlerFunc(productHandler.GetProductHandler)

		// act
		// call the handler func
		handler.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusNotFound, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
	})
}

// TestCreateProduct tests the CreateProductHandler method.
func TestCreateProduct(t *testing.T) {
	t.Run("success - create product", func(t *testing.T) {
		// arrange
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		// create a product
		product := internalProduct.Product{
			Id:          0,
			Name:        "product 1",
			Quantity:    10,
			CodeValue:   "code 1",
			IsPublished: false,
			Expiration:  "12/31/2022",
			Price:       100,
		}
		// body
		body := `{
			"name": "product 1",
			"quantity": 10,
			"code_value": "code 1",
			"expiration": "12/31/2022",
			"price": 100.0
		}`

		// create the expected response
		expectedResponse := `{
			"data": {
				"id": 0,
				"name": "product 1",
				"quantity": 10,
				"code_value": "code 1",
				"is_published": false,
				"expiration": "12/31/2022",
				"price": 100
			},
			"message": "product created successfully"
		}`

		// create a mock of CreateProduct method
		productService.On("CreateProduct", product).Return(product, nil)

		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)

		// create a new request and recorder
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		// create a new handler func
		handler := http.HandlerFunc(productHandler.CreateProductHandler)

		// act
		// call the handler func
		handler.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusCreated, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
		productService.AssertCalled(t, "CreateProduct", product)
	})
	t.Run("fail - create product bad request", func(t *testing.T) {
		// arrange
		// expected response
		expectedResponse := `{"message":"bad request", "status":"Bad Request"}`
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		// body
		body := `a really bad json body`
		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)
		// create a new request and recorder
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		// create a new handler func
		handler := http.HandlerFunc(productHandler.CreateProductHandler)
		// act
		// call the handler func
		handler.ServeHTTP(rr, req)
		// assert
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
	})
	t.Run("fail - create product invalid product", func(t *testing.T) {
		// arrange
		// expected response
		expectedResponse := `{"message":"invalid product parameters", "status":"Bad Request"}`
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		// create a product
		product := internalProduct.Product{
			Id:          0,
			Name:        "product 1",
			Quantity:    10,
			CodeValue:   "code 1",
			IsPublished: false,
			Expiration:  "12/31/2022",
			Price:       100,
		}
		// body
		body := `{
			"name": "product 1",
			"quantity": 10,
			"code_value": "code 1",
			"expiration": "12/31/2022",
			"price": 100.0
		}`
		// create a mock of CreateProduct method
		productService.On("CreateProduct", product).Return(product, internalProduct.ErrInvalidProduct)
		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)
		// create a new request and recorder
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		// create a new handler func
		handler := http.HandlerFunc(productHandler.CreateProductHandler)
		// act
		// call the handler func
		handler.ServeHTTP(rr, req)
		// assert
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
	})
	t.Run("fail - create product internal duplicated code value", func(t *testing.T) {
		// arrange
		// expected response
		expectedResponse := `{"message":"duplicated code value", "status":"Conflict"}`
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		// create a product
		product := internalProduct.Product{
			Id:          0,
			Name:        "product 1",
			Quantity:    10,
			CodeValue:   "code 1",
			IsPublished: false,
			Expiration:  "12/31/2022",
			Price:       100,
		}
		// body
		body := `{
			"name": "product 1",
			"quantity": 10,
			"code_value": "code 1",
			"expiration": "12/31/2022",
			"price": 100.0
		}`
		// create a mock of CreateProduct method
		productService.On("CreateProduct", product).Return(product, internalProduct.ErrDuplicateCodeValue)
		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)
		// create a new request and recorder
		req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		// create a new handler func
		handler := http.HandlerFunc(productHandler.CreateProductHandler)
		// act
		// call the handler func
		handler.ServeHTTP(rr, req)
		// assert
		require.Equal(t, http.StatusConflict, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("success - delete product", func(t *testing.T) {
		// arrange
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		// create the expected response
		expectedResponse := "product deleted successfully"

		// create a mock of DeleteProduct method
		productService.On("DeleteProduct", "1").Return(nil)

		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)

		// create a new request and recorder
		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)

		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("id", "1")

		// set the RouteContext on the request context
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))

		rr := httptest.NewRecorder()
		// create a new handler func
		handler := http.HandlerFunc(productHandler.DeleteProductHandler)

		// act
		// call the handler func
		handler.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, expectedResponse, rr.Body.String())
		productService.AssertCalled(t, "DeleteProduct", "1")
	})
	t.Run("fail - delete product not found", func(t *testing.T) {
		// arrange
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)

		// create the expected response
		expectedResponse := `{"message":"product not found", "status":"Not Found"}`

		// create a mock of DeleteProduct method
		productService.On("DeleteProduct", "1").Return(internalProduct.ErrProductNotFound)

		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)

		// create a new request and recorder
		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)

		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("id", "1")

		// set the RouteContext on the request context
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
		rr := httptest.NewRecorder()

		// create a new handler func
		handler := http.HandlerFunc(productHandler.DeleteProductHandler)

		// act
		// call the handler func
		handler.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusNotFound, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
	})
	t.Run("fail - delete product invalid id", func(t *testing.T) {
		// arrange
		// create a mock of ProductServiceInterface
		productService := new(ProductServiceInterface)
		productService.On("DeleteProduct", "1").Return(internalProduct.ErrInvalidID)

		// create the expected response
		expectedResponse := `{"message":"invalid id", "status":"Bad Request"}`

		// create a new ProductHandler
		productHandler := handler.NewProductHandler(productService)

		// create a new request and recorder
		req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)

		// create a new RouteContext and set the "id" URL parameter
		routeContext := chi.NewRouteContext()
		routeContext.URLParams.Add("id", "1")

		// set the RouteContext on the request context
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeContext))
		rr := httptest.NewRecorder()

		// create a new handler func
		handler := http.HandlerFunc(productHandler.DeleteProductHandler)

		// act
		// call the handler func
		handler.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.JSONEq(t, expectedResponse, rr.Body.String())
	})
}
