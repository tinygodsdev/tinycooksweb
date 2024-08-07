package cachedstorage

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent"
)

const (
	driver = "sqlite3"
	dsn    = "file:ent?mode=memory&cache=shared&_fk=1"
	chunk  = 100
)

type CachedStorage struct {
	logger  logger.Logger
	storage storage.Storage
	cache   *entstorage.EntStorage
	mx      sync.Mutex
	cfg     Config
}

type Config struct {
	LogQueries bool
}

func Migrate(ctx context.Context, client *ent.Client) error {
	return entstorage.Migrate(ctx, client)
}

func (s *CachedStorage) SaveRecipe(ctx context.Context, recipe *recipe.Recipe) error {
	return s.storage.SaveRecipe(ctx, recipe)
}

func (s *CachedStorage) GetRecipe(ctx context.Context, id uuid.UUID) (*recipe.Recipe, error) {
	return s.cache.GetRecipe(ctx, id)
}

func (s *CachedStorage) GetRecipeBySlug(ctx context.Context, slug string) (*recipe.Recipe, error) {
	return s.cache.GetRecipeBySlug(ctx, slug)
}

func (s *CachedStorage) GetRecipes(ctx context.Context, filter recipe.Filter) ([]*recipe.Recipe, error) {
	return s.cache.GetRecipes(ctx, filter)
}

func (s *CachedStorage) CountRecipes(ctx context.Context, filter recipe.Filter) (int, error) {
	return s.cache.CountRecipes(ctx, filter)
}

func (s *CachedStorage) GetTags(ctx context.Context, locale string) ([]*recipe.Tag, error) {
	return s.cache.GetTags(ctx, locale)
}

func (s *CachedStorage) GetIngredients(ctx context.Context, locale string) ([]*recipe.Ingredient, error) {
	return s.cache.GetIngredients(ctx, locale)
}

func (s *CachedStorage) GetEquipment(ctx context.Context, locale string) ([]*recipe.Equipment, error) {
	return s.cache.GetEquipment(ctx, locale)
}

func NewCachedStorage(
	cfg Config,
	logger logger.Logger,
	store storage.Storage,
) (*CachedStorage, error) {
	cache, err := createCache(cfg, logger)
	if err != nil {
		return nil, err
	}

	return &CachedStorage{
		logger:  logger,
		storage: store,
		cfg:     cfg,
		cache:   cache,
	}, nil
}

func createCache(cfg Config, logger logger.Logger) (*entstorage.EntStorage, error) {
	cache, err := entstorage.NewEntStorage(entstorage.Config{
		StorageDriver: driver,
		StorageDSN:    dsn,
		LogQueries:    cfg.LogQueries,
		Migrate:       true,
	}, logger)
	if err != nil {
		return nil, err
	}

	return cache, nil
}

func (s *CachedStorage) Close() error {
	if err := s.cache.Close(); err != nil {
		return err
	}

	return s.storage.Close()
}

func (s *CachedStorage) UpdateCache(ctx context.Context) error {
	newCache, err := createCache(s.cfg, s.logger)
	if err != nil {
		return fmt.Errorf("failed to create new cache: %w", err)
	}

	// Load data into the new cache from the primary storage
	err = loadDataFromPrimaryStorage(ctx, newCache, s.storage, s.logger)
	if err != nil {
		newCache.Close()
		return fmt.Errorf("failed to load data into new cache: %w", err)
	}

	// Atomically replace the old cache with the new one
	s.mx.Lock()
	oldCache := s.cache
	s.cache = newCache
	s.mx.Unlock()

	// Close the old cache
	oldCache.Close()

	s.logger.Info("Cache updated successfully")
	return nil
}

func loadDataFromPrimaryStorage(
	ctx context.Context,
	cache *entstorage.EntStorage,
	storage storage.Storage,
	log logger.Logger,
) error {
	offset := 0

	for {
		log.Info("Loading recipes from primary storage", "offset", offset)
		recipes, err := storage.GetRecipes(ctx, recipe.Filter{
			Limit:     chunk,
			Offset:    offset,
			WithEdges: true,
		})
		if err != nil {
			return fmt.Errorf("failed to get recipes from storage: %w", err)
		}
		if len(recipes) == 0 {
			break
		}

		for _, r := range recipes {
			err = cache.SaveRecipe(ctx, r)
			if err != nil {
				return fmt.Errorf("failed to save recipe to cache: %w", err)
			}
		}

		log.Info("Loaded recipes", "count", len(recipes))

		offset += chunk
	}

	return nil
}
