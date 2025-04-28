package storage

import (
	"context"
	"fmt"

	"github.com/GeekchanskiY/migratigo"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"finworker/internal/config"
)

func NewConn(lc fx.Lifecycle, config *config.Config, logger *zap.Logger) *sqlx.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	connector, err := migratigo.New(db.DB, Migrations, "migrations", logger)
	if err != nil {
		panic(err)
	}

	err = connector.RunMigrations(false)
	if err != nil {
		panic(err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.Ping(); err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	logger.Info("connected to database", zap.String("host", config.DbHost))

	return db
}

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(
			NewConn,
		),
		fx.Invoke(
			func(db *sqlx.DB, logger *zap.Logger) {
				// This ensures the database connection is actually used
				logger.Info("Database module initialized")
			},
		),
	)
}
