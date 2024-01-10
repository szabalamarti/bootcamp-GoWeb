package internal

import "errors"

var (
	ErrLoadProducts = errors.New("failed to load products")
	ErrUnmarshal    = errors.New("failed to unmarshal products")
)

type ProductRepositoryInterface interface {
	Get() []Product
	GetById(id int) (Product, error)
	SearchByPrice(priceGt float64) ([]Product, error)
	Save(product Product) (Product, error)
	SaveOrUpdate(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(id int) error
}
