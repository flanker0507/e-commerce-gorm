package database

import (
	"e-commerce-gorm/internal/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func ConnectMysql(cfg config.DBConfig) (db *sqlx.DB, err error) {
	// Definisikan string koneksi: user:password@tcp(host:port)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	// Membuka koneksi ke database
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	// Menguji koneksi
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdletimeConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpenConnection))
	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIdleConnection))

	return db, nil
}
