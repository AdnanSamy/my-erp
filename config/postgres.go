package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type DB struct{ *sqlx.DB }

var defaultDb = &DB{}

type DbConfig struct {
	Host            string
	Port            int
	SslMode         string
	Name            string
	User            string
	Password        string
	Debug           bool
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifetime time.Duration
}

var port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
var debug, _ = strconv.ParseBool(os.Getenv("DB_DEBUG"))
var maxOpenConn, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
var maxIdleConn, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
var lifeTime, _ = strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

var dbCfg = &DbConfig{
	Host:            os.Getenv("DB_HOST"),
	Port:            port,
	User:            os.Getenv("DB_USER"),
	Password:        os.Getenv("DB_PASSWORD"),
	Name:            os.Getenv("DB_NAME"),
	SslMode:         os.Getenv("DB_SSL_MODE"),
	Debug:           debug,
	MaxOpenConn:     maxOpenConn,
	MaxIdleConn:     maxIdleConn,
	MaxConnLifetime: time.Duration(lifeTime) * time.Second,
}

func (db *DB) connect(cfg *DbConfig) (err error) {
	dbUri := fmt.Sprintf("host=%s port=%d sslmode=%s user=%s password=%s dbname=%s",
		cfg.Host,
		cfg.Port,
		cfg.SslMode,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	sqlx.Connect("pgx", dbUri)

	if err != nil {
		return err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	db.SetConnMaxLifetime(cfg.MaxConnLifetime)

	if err := db.Ping(); err != nil {
		defer db.Close()
		return fmt.Errorf("can't send ping to database, %w", err)
	}

	return nil
}

func GetDb() *DB {
	return defaultDb
}

func ConnectDB() error {
	return defaultDb.connect(dbCfg)
}
