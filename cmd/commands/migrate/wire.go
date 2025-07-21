//go:build wireinject
// +build wireinject

package migrate

import (
	"context"
	"github.com/compico/em-task/cmd/di"
	"github.com/compico/em-task/internal/pkg/config"
	"github.com/compico/em-task/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/wire"
)

type (
	Migrator struct {
		logger  logger.Logger
		migrate *migrate.Migrate
	}
)

func InitializeMigrator(
	ctx context.Context,
	filepath string,
) (*Migrator, error) {
	panic(wire.Build(
		di.BaseSet,
		di.DatabaseConfigProvider,
		di.LoggerSet,
		MigrateProvider,

		wire.Struct(new(Migrator), "*"),
	))
}

func MigrateProvider(dbConfig config.Database) (*migrate.Migrate, error) {
	return migrate.New(
		dbConfig.GetMigrationDir(),
		dbConfig.GetDsn(),
	)
}
