package handler

type ProductRequestJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductResponseJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func deserializeProductJSONToProduct(productRequest ProductRequestJSON) Product {
	return Product{
		Name:        productRequest.Name,
		Quantity:    productRequest.Quantity,
		CodeValue:   productRequest.CodeValue,
		IsPublished: productRequest.IsPublished,
		Expiration:  productRequest.Expiration,
		Price:       productRequest.Price,
	}
}

func serializeProductToProductRequestJSON(product Product) ProductRequestJSON {
	return ProductRequestJSON{
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
}

func serializeProductToProductResponseJSON(product Product) ProductResponseJSON {
	return ProductResponseJSON{
		Id:          product.Id,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
}

func serializeProductsToProductsResponseJSON(products []Product) []ProductResponseJSON {
	productsResponse := make([]ProductResponseJSON, len(products))
	for i, product := range products {
		productsResponse[i] = serializeProductToProductResponseJSON(product)
	}
	return productsResponse
}
