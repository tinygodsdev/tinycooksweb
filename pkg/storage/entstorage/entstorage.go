package entstorage

import (
	"context"
	"fmt"

	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/migrate"

	_ "github.com/lib/pq"
)

type EntStorage struct {
	client *ent.Client
	logger logger.Logger
}

type Config struct {
	StorageDriver string
	StorageDSN    string
	LogQueries    bool
	Migrate       bool
}

func NewEntStorage(
	cfg Config,
	logger logger.Logger,
) (*EntStorage, error) {
	var dbOptions []ent.Option

	if cfg.LogQueries {
		dbOptions = append(dbOptions, ent.Debug())
	}

	client, err := ent.Open(cfg.StorageDriver, cfg.StorageDSN, dbOptions...)
	if err != nil {
		return nil, fmt.Errorf("EntStorage: failed connecting to database: %w", err)
	}
	defer client.Close()

	logger.Info("EntStorage: connected")

	if cfg.Migrate {
		err = Migrate(context.Background(), client) // run db migration
		if err != nil {
			return nil, fmt.Errorf("EntStorage: failed creating schema resources: %w", err)
		}
		logger.Info("EntStorage: migrations done")
	}

	return &EntStorage{
		client: client,
		logger: logger,
	}, nil
}

func Migrate(ctx context.Context, client *ent.Client) error {
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return err
	}
	return nil
}
