package auth

import (
	"context"
	"fmt"
	"github.com/RianIhsan/go-ecommerce/external/database"
	"github.com/RianIhsan/go-ecommerce/infra/response"
	"github.com/RianIhsan/go-ecommerce/internal/config"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@codenexus.me", uuid.NewString()),
		Password: "mypassword",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)
}

func TestRegister_Failed(t *testing.T) {
	t.Run("Failed: Email already used", func(t *testing.T) {
		// preparation for duplicate email
		email := fmt.Sprintf("%v@codenexus.me", uuid.NewString())
		req := RegisterRequestPayload{
			Email:    email,
			Password: "mypassword",
		}
		err := svc.register(context.Background(), req)
		require.Nil(t, err)

		// end preparation

		err = svc.register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailAlreadyUsed, err)
	})
}
