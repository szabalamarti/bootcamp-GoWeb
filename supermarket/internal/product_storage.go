package internal

import (
	"errors"
)

var (
	ErrFileNotFound = errors.New("file not found")
	ErrInvalidFile  = errors.New("invalid file")
	ErrSaveProducts = errors.New("error saving products")
)

type ProductStorageInterface interface {
	LoadProducts() (map[int]Product, error)
	SaveProducts(products map[int]Product) error
}
