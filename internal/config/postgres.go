package config

import (
	"errors"
	"fmt"
	"os"
	//"gopkg.in/yaml.v3"
)

// PostgresConfig config
type postgresConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func NewPGConfig() (*postgresConfig, error) {
	//var pgConfig PostgresConfig

	host := os.Getenv("PG_HOST")
	if len(host) == 0 {
		return nil, errors.New("pg host not found")
	}

	port := os.Getenv("PG_PORT")
	if len(port) == 0 {
		return nil, errors.New("pg port not found")
	}

	user := os.Getenv("PG_USER")
	if len(user) == 0 {
		return nil, errors.New("pg user not found")
	}

	password := os.Getenv("PG_PASSWORD")
	if len(password) == 0 {
		return nil, errors.New("pg password not found")
	}

	dbname := os.Getenv("PG_DBNAME")
	if len(dbname) == 0 {
		return nil, errors.New("pg dbname not found")
	}

	sslmode := os.Getenv("PG_SSL_MODE")
	if len(sslmode) == 0 {
		return nil, errors.New("pg sslmode not found")
	}

	return &postgresConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
		sslmode:  sslmode,
	}, nil

}

func GetConnectionString() (string, error) {
	cfg, err := NewPGConfig()
	if err != nil {
		return "", fmt.Errorf("unable to load config: %e", err)
	}
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		cfg.host,
		cfg.port,
		cfg.user,
		cfg.password,
		cfg.dbname,
		cfg.sslmode,
	), nil
}
