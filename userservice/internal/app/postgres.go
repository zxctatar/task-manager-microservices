package app

import (
	"database/sql"
	"fmt"
	"userservice/internal/config"
	"userservice/internal/infrastructure/postgres"
)

func mustLoadPostgres(cfg *config.Config) *postgres.Postgres {
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
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	pos := postgres.NewPostgres(db)

	return pos
}
