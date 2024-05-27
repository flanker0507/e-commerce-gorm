package product

import (
	"e-commerce-gorm/infra/response"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id        int    `db:"id"`
	SKU       string `db:"sku"`
	Name      string `db:"name"`
	Stock     int16  `db:"stock"`
	Price     int    `db:"price"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type ProductPagniation struct {
	Cursor int `json:"cursor"`
	Size   int `json:"size"`
}

func NewProductPaginationFromListProductRequest(req ListProductRequestPayload) ProductPagniation {
	req = req.GenerateDefaultValue()
	return ProductPagniation{
		Cursor: req.Cursor,
		Size:   req.Size,
	}
}

func NewProductFromCreateProductRequest(req CreateProductRequestPayload) Product {
	return Product{
		SKU:       uuid.NewString(),
		Name:      req.Name,
		Stock:     req.Stock,
		Price:     req.Price,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (p Product) Validate() (err error) {
	if err = p.ValidateName(); err != nil {
		return
	}
	if err = p.ValidateStock(); err != nil {
		return
	}
	if err = p.ValidatePrice(); err != nil {
		return
	}
	return
}

func (p Product) ValidateName() (err error) {
	if p.Name == "" {
		return response.ErrProductRequired
	}
	if len(p.Name) < 4 {
		return response.ErrProductInvalid
	}
	return
}

func (p Product) ValidateStock() (err error) {
	if p.Stock <= 0 {
		return response.ErrStockInvalid
	}
	return
}

func (p Product) ValidatePrice() (err error) {
	if p.Price <= 0 {
		return response.ErrPriceInvalid
	}
	return
}

func (p Product) ToProductListResponse() ProductListResponse {
	return ProductListResponse{
		Id:    p.Id,
		SKU:   p.SKU,
		Name:  p.Name,
		Stock: p.Stock,
		Price: p.Price,
	}
}

func (p Product) ToProductDetailResponse() ProductDetailResponse {
	return ProductDetailResponse{
		Id:        p.Id,
		SKU:       p.SKU,
		Name:      p.Name,
		Stock:     p.Stock,
		Price:     p.Price,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
