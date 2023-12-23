package database

import (
	"github.com/RianIhsan/go-ecommerce/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectionPostgres(t *testing.T) {
	t.Run("Success connect to postgres", func(t *testing.T) {
		_, err := ConnectPostgres(config.Cfg.DB)
		require.Nil(t, err)
	})
}
