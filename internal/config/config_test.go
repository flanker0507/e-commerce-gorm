package config

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		filename := "../../cmd/api/config.yaml"
		err := LoadConfig(filename)

		require.Nil(t, err)
		log.Printf("%+v\n", Cfg)
	})
}
