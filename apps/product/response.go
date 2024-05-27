package product

type ProductListResponse struct {
	Id    int    `json:"id"`
	SKU   string `json:"SKU"`
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

func NewProductListResponseFromEntity(product []Product) []ProductListResponse {
	var productList = []ProductListResponse{}

	for _, product := range product {
		productList = append(productList, product.ToProductListResponse())
	}
	return productList
}

type ProductDetailResponse struct {
	Id        int    `json:"id"`
	SKU       string `json:"SKU"`
	Name      string `json:"name"`
	Stock     int16  `json:"stock"`
	Price     int    `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
