package product

import (
	"e-commerce-gorm/infra/response"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		product := Product{
			Name:  "Baju Baru",
			Stock: 10,
			Price: 10_000,
		}
		err := product.Validate()
		require.Nil(t, err)
	})
	t.Run("Product Require", func(t *testing.T) {
		product := Product{
			Name:  "",
			Stock: 10,
			Price: 10_000,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
	t.Run("Product Invalid", func(t *testing.T) {
		product := Product{
			Name:  "baj",
			Stock: 10,
			Price: 10_000,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)
	})
	t.Run("Stock Invalid", func(t *testing.T) {
		product := Product{
			Name:  "Baju Baru",
			Stock: 0,
			Price: 10_000,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)
	})
	t.Run("Price Invalid", func(t *testing.T) {
		product := Product{
			Name:  "Baju Baru",
			Stock: 10,
			Price: 0,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})
}
