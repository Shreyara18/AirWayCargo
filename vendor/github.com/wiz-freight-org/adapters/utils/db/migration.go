package db

import (
	"github.com/wiz-freight-org/adapters/internal/config"
	"github.com/wiz-freight-org/adapters/internal/config/globals"
	"go.uber.org/zap"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func RunMigration() bool {
	l := globals.Logger

	m, err := migrate.New(config.MigrationFilePath, config.DatabaseURL)
	if err != nil {
		l.Fatal("Unable to get migrate instance", zap.Error(err))
		return false
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		l.Fatal("Unable to migrate the database", zap.Error(err))
		return false
	}

	return true
}
