package auth

import (
	"github.com/RianIhsan/go-ecommerce/infra/response"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

func TestAuthEntity(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "yon@codenexus.me",
			Password: "mypassword",
		}

		err := authEntity.Validate()
		require.Nil(t, err)
	})

	t.Run("Failed: Email required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "",
			Password: "mypassword",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, err, response.ErrEmailRequired)
	})

	t.Run("Failed: Email invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "yon",
			Password: "mypassword",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, err, response.ErrEmailInvalid)
	})

	t.Run("Failed: Password required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "yon@codenexus.me",
			Password: "",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, err, response.ErrPasswordRequired)
	})

	t.Run("Failed: Password invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "yon@codenexus.me",
			Password: "fail",
		}
		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, err, response.ErrPasswordInvalid)
	})
}

func TestAuthEntity_EncryptPassword(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email:    "yon@codenexus.me",
			Password: "mypassword",
		}

		err := authEntity.EncryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)
		log.Printf("%v\n", authEntity)
	})
}
