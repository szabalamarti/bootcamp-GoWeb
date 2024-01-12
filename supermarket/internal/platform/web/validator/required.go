package validator

import (
	"errors"
)

var ErrKeyNotFound = errors.New("key not found")

func ValidateRequiredKeys(m map[string]any, keys ...string) error {
	for _, key := range keys {
		if _, ok := m[key]; !ok {
			return ErrKeyNotFound
		}
	}
	return nil
}
