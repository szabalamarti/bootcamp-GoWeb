package service

import (
	internalProduct "supermarket/internal/product"

	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (m *ProductServiceMock) GetProducts() ([]internalProduct.Product, error) {
	args := m.Called()
	return args.Get(0).([]internalProduct.Product), args.Error(1)
}

func (m *ProductServiceMock) CreateProduct(p internalProduct.Product) (internalProduct.Product, error) {
	args := m.Called(p)
	return args.Get(0).(internalProduct.Product), args.Error(1)
}

func (m *ProductServiceMock) GetProduct(id string) (internalProduct.Product, error) {
	args := m.Called(id)
	return args.Get(0).(internalProduct.Product), args.Error(1)
}

func (m *ProductServiceMock) SearchProductsByPrice(priceGt string) ([]internalProduct.Product, error) {
	args := m.Called(priceGt)
	return args.Get(0).([]internalProduct.Product), args.Error(1)
}

func (m *ProductServiceMock) UpdateOrCreateProduct(p internalProduct.Product) (internalProduct.Product, error) {
	args := m.Called(p)
	return args.Get(0).(internalProduct.Product), args.Error(1)
}

func (m *ProductServiceMock) UpdateProduct(p internalProduct.Product) (internalProduct.Product, error) {
	args := m.Called(p)
	return args.Get(0).(internalProduct.Product), args.Error(1)
}

func (m *ProductServiceMock) DeleteProduct(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *ProductServiceMock) GetConsumerPriceProducts(ids []string) (internalProduct.ConsumerPriceProducts, error) {
	args := m.Called(ids)
	return args.Get(0).(internalProduct.ConsumerPriceProducts), args.Error(1)
}
