package app

import (
	"database/sql"
	"fmt"
	"userservice/internal/config"
)

func mustLoadPostgres(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresConf.Host,
		cfg.PostgresConf.Port,
		cfg.PostgresConf.User,
		cfg.PostgresConf.Password,
		cfg.PostgresConf.DbName,
		cfg.PostgresConf.Sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	return db
}
