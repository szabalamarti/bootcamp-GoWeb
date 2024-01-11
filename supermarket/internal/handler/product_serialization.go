package handler

type ProductRequest struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func ProductRequestToProduct(productRequest ProductRequest) Product {
	return Product{
		Name:        productRequest.Name,
		Quantity:    productRequest.Quantity,
		CodeValue:   productRequest.CodeValue,
		IsPublished: productRequest.IsPublished,
		Expiration:  productRequest.Expiration,
		Price:       productRequest.Price,
	}
}

func ProductToProductRequest(product Product) ProductRequest {
	return ProductRequest{
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
}

func ProductToProductResponse(product Product) ProductResponse {
	return ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}
}

func ProductsToProductsResponse(products []Product) []ProductResponse {
	productsResponse := make([]ProductResponse, len(products))
	for i, product := range products {
		productsResponse[i] = ProductToProductResponse(product)
	}
	return productsResponse
}
