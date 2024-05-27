package product

import (
	"context"
	"e-commerce-gorm/eksternal/database"
	"e-commerce-gorm/infra/response"
	"e-commerce-gorm/internal/config"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectMysql(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProduct_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "Baju Baru",
		Stock: 10,
		Price: 10_000,
	}
	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProduct_Fail(t *testing.T) {
	t.Run("name is required", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "",
			Stock: 10,
			Price: 10_000,
		}
		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
	t.Run("name is invalid", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "bik",
			Stock: 10,
			Price: 10_000,
		}
		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)
	})
	t.Run("stock is invalid", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "sempak",
			Stock: 0,
			Price: 10_000,
		}
		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)
	})
	t.Run("price is required", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:  "BRAzil",
			Stock: 10,
			Price: 0,
		}
		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})

}

func TestListProduct_Success(t *testing.T) {
	pagination := ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	}

	product, err := svc.ListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, product)
	log.Printf("%+v", product)
}

func TestProductDetail_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "Baju Baru",
		Stock: 10,
		Price: 10_000,
	}
	ctx := context.Background()

	err := svc.CreateProduct(ctx, req)

	products, err := svc.ListProducts(ctx, ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	})
	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)

	product, err := svc.ProductDetail(ctx, products[0].SKU)
	require.Nil(t, err)
	require.NotEmpty(t, product)
	log.Printf("%+v", product)
}
