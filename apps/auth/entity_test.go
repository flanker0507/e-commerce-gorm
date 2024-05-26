package auth

import (
	"e-commerce-gorm/infra/response"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

func TestValidateAuthEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "yuda@gmail.com",
			Password: "dillacantik",
		}
		err := authEntity.Validate()
		require.Nil(t, err)
	})

	t.Run("email is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "",
			Password: "dillacantik",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})

	t.Run("email is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "yudagmail.com",
			Password: "dillacantik",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})

	t.Run("password is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "yuda@gmail.com",
			Password: "",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})

	t.Run("password is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "yuda@gmail.com",
			Password: "dill",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalid, err)
	})
}

func TestEncrypPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "dilla@gmail.com",
			Password: "yuda",
		}

		err := authEntity.EncryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)

		log.Printf("%+v\n", authEntity)
	})
}
