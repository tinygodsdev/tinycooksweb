package app

import (
	"context"
	"fmt"
	"slices"

	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/internal/config"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage"
)

type App struct {
	Cfg     *config.Config
	log     logger.Logger
	store   storage.Storage
	locales []string // locale count is not very big so no need to have map
	Errors  []error
}

func New(
	cfg *config.Config,
	logger logger.Logger,
	store storage.Storage,
) (*App, error) {
	app := App{
		Cfg:     cfg,
		log:     logger,
		store:   store,
		locales: locale.List(),
	}

	if cfg.UseMocks {
		if err := app.SaveMocks(); err != nil {
			return nil, err
		}
	}

	return &app, nil
}

func (a *App) IsDev() bool {
	return a.Cfg.Dev()
}

func (a *App) AddError(err error) {
	a.Errors = append(a.Errors, err)
}

func (a *App) SaveMocks() error {
	defer a.Timer("SaveMocks")()

	mocks, _ := recipe.MockRecipes(999, false)

	for _, mock := range mocks {
		if err := a.SaveRecipe(context.Background(), mock); err != nil {
			if err == storage.ErrAlreadyExists {
				a.log.Info("Mock recipe already exists", "slug", mock.Slug)
				continue
			}

			return err
		}
	}

	return nil
}

func (a *App) GetRecipe(ctx context.Context, slug string) (*recipe.Recipe, error) {
	defer a.Timer("GetRecipe", "slug", slug)()

	if slug == "" {
		return nil, fmt.Errorf("empty slug")
	}

	if a.Cfg.MockQueries {
		return recipe.MockRecipe(slug, false)
	}

	return a.store.GetRecipeBySlug(ctx, slug)
}

func (a *App) GetRecipes(ctx context.Context, filter recipe.Filter) ([]*recipe.Recipe, error) {
	defer a.Timer("GetRecipes", "filter", filter)()

	if filter.Locale == "" || !slices.Contains(a.locales, filter.Locale) {
		// filter.Locale = locale.Default()
		return nil, fmt.Errorf("invalid locale: %s", filter.Locale)
	}

	if filter.UseMocks {
		return recipe.MockRecipes(filter.Limit, false)
	}

	return a.store.GetRecipes(ctx, filter)
}

func (a *App) SaveRecipe(ctx context.Context, rec *recipe.Recipe) error {
	if rec == nil {
		return fmt.Errorf("empty recipe")
	}

	defer a.Timer("SaveRecipe", "slug", rec.Slug)()

	return a.store.SaveRecipe(ctx, rec)
}

func (a *App) CountRecipes(ctx context.Context, filter recipe.Filter) (int, error) {
	defer a.Timer("CountRecipes")()

	return a.store.CountRecipes(ctx, filter)
}

func (a *App) GetTags(ctx context.Context, locale string) ([]*recipe.Tag, error) {
	defer a.Timer("GetTags")()

	if locale == "" || !slices.Contains(a.locales, locale) {
		return nil, fmt.Errorf("invalid locale: %s", locale)
	}

	return a.store.GetTags(ctx, locale)
}

func (a *App) GetIngredients(ctx context.Context, locale string) ([]*recipe.Ingredient, error) {
	defer a.Timer("GetIngredients")()

	if locale == "" || !slices.Contains(a.locales, locale) {
		return nil, fmt.Errorf("invalid locale: %s", locale)
	}

	return a.store.GetIngredients(ctx, locale)
}

func (a *App) GetEquipment(ctx context.Context, locale string) ([]*recipe.Equipment, error) {
	defer a.Timer("GetEquipment")()

	if locale == "" || !slices.Contains(a.locales, locale) {
		return nil, fmt.Errorf("invalid locale: %s", locale)
	}

	return a.store.GetEquipment(ctx, locale)
}
