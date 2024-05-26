package database

import (
	"e-commerce-gorm/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func init() {
	// Muat konfigurasi saat package diinisialisasi
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectMysql(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Panggil fungsi ConnectMysql dengan konfigurasi yang dimuat
		db, err := ConnectMysql(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)

		// Uji koneksi
		err = db.Ping()
		require.Nil(t, err)
	})
}
