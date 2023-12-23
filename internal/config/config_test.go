package config

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Success get file", func(t *testing.T) {
		filename := "../../cmd/api/config.yaml"
		err := LoadConfig(filename)

		require.Nil(t, err)
		log.Printf("%v\n", Cfg)
	})

	t.Run("Failed get file", func(t *testing.T) {
		filename := "../../cmd/api/config2.yaml"
		err := LoadConfig(filename)

		require.NotNil(t, err)
		log.Printf("%v\n", Cfg)
	})
}
