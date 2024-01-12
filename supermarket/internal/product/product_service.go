package product

import "errors"

var (
	ErrInvalidID            = errors.New("invalid id")
	ErrProductNotFound      = errors.New("product not found")
	ErrInvalidPriceGt       = errors.New("invalid priceGt")
	ErrInvalidProduct       = errors.New("invalid product parameters")
	ErrDuplicateCodeValue   = errors.New("duplicated code value")
	ErrInsufficientQuantity = errors.New("insufficient quantity of product")
)

type ConsumerPriceProducts struct {
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"total_price"`
}

type ProductServiceInterface interface {
	GetProducts() ([]Product, error)
	GetProduct(id string) (Product, error)
	SearchProductsByPrice(priceGt string) ([]Product, error)
	CreateProduct(product Product) (Product, error)
	UpdateOrCreateProduct(product Product) (Product, error)
	UpdateProduct(product Product) (Product, error)
	DeleteProduct(id string) error
	GetConsumerPriceProducts(ids []string) (ConsumerPriceProducts, error)
}
