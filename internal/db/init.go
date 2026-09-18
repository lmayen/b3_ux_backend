package db

import (
	"b3_ux_backend/internal/config"
	"b3_ux_backend/internal/entities"
	"b3_ux_backend/internal/fsutils"
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/olekukonko/errors"
	_ "modernc.org/sqlite"
)

var Client *entities.Client

func InitDb(ctx context.Context, cfg *config.Config) (*entities.Client, error) {
	if !fsutils.Exists(cfg.App.SqlitePath) {
		CreateAppEntities()
	}

	dsn := fmt.Sprintf("file:%s?_fk=1", filepath.ToSlash(cfg.App.SqlitePath))
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.SQLite, db)
	Client = entities.NewClient(entities.Driver(drv))

	if err := Client.Schema.Create(ctx); err != nil {
		return nil, errors.Wrapf(err, "InitDb ERR")
	}

	return Client, nil
}
