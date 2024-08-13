package app

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/internal/config"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"github.com/tinygodsdev/tinycooksweb/pkg/moderation"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage"
	"golang.org/x/sync/errgroup"
)

type App struct {
	Cfg             *config.Config
	log             logger.Logger
	store           storage.Storage
	moderationStore moderation.ModerationStore
	scheduler       *gocron.Scheduler
	locales         []string // locale count is not very big so no need to have map
	Errors          []error
}

func New(
	cfg *config.Config,
	logger logger.Logger,
	store storage.Storage,
	moderationStore moderation.ModerationStore,
) (*App, error) {
	s := gocron.NewScheduler(time.UTC)
	s.SetMaxConcurrentJobs(1, gocron.RescheduleMode)

	app := App{
		Cfg:             cfg,
		log:             logger,
		store:           store,
		moderationStore: moderationStore,
		locales:         locale.List(),
		scheduler:       s,
	}

	err := app.addJob("load-approved-recipes", cfg.ModerationCheckSchedule, app.LoadApprovedRecipes)
	if err != nil {
		return nil, err
	}

	if cfg.SaveMocks {
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

	mocks, _ := recipe.MockRecipes(1999, false)

	g, ctx := errgroup.WithContext(context.Background())
	if a.Cfg.StorageDriver == "sqlite3" {
		g.SetLimit(1) // sqlite3 does not support concurrent writes
	} else {
		g.SetLimit(25) // for postgres it's fine
	}

	for _, mock := range mocks {
		mock := mock

		g.Go(func() error {
			if err := a.SaveRecipe(ctx, mock); err != nil {
				if err == storage.ErrAlreadyExists {
					a.log.Info("Mock recipe already exists", "slug", mock.Slug)
					return nil
				}
				return err
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *App) GetRecipe(ctx context.Context, slug string) (*recipe.Recipe, error) {
	defer a.Timer("GetRecipe", "slug", slug)()

	if slug == "" {
		return nil, fmt.Errorf("empty slug")
	}

	return a.store.GetRecipeBySlug(ctx, slug)
}

func (a *App) GetRecipes(ctx context.Context, filter recipe.Filter) ([]*recipe.Recipe, error) {
	defer a.Timer("GetRecipes", "filter", filter)()

	if filter.Locale == "" || !slices.Contains(a.locales, filter.Locale) {
		// filter.Locale = locale.Default()
		return nil, fmt.Errorf("invalid locale: %s", filter.Locale)
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

func (a *App) GetTagBySlug(ctx context.Context, slug string) (*recipe.Tag, error) {
	defer a.Timer("GetTagBySlug", "slug", slug)()

	if slug == "" {
		return nil, fmt.Errorf("empty slug")
	}

	return a.store.GetTagBySlug(ctx, slug)
}

func (a *App) GetIngredientBySlug(ctx context.Context, slug string) (*recipe.Ingredient, error) {
	defer a.Timer("GetIngredientBySlug", "slug", slug)()

	if slug == "" {
		return nil, fmt.Errorf("empty slug")
	}

	return a.store.GetIngredientBySlug(ctx, slug)
}

func (a *App) GetEquipmentBySlug(ctx context.Context, slug string) (*recipe.Equipment, error) {
	defer a.Timer("GetEquipmentBySlug", "slug", slug)()

	if slug == "" {
		return nil, fmt.Errorf("empty slug")
	}

	return a.store.GetEquipmentBySlug(ctx, slug)
}

func (a *App) LoadApprovedRecipes(ctx context.Context) error {
	defer a.Timer("LoadApprovedRecipes")()

	moderationRecords, err := a.moderationStore.GetApproved(ctx)
	if err != nil {
		return err
	}

	if len(moderationRecords) == 0 {
		a.log.Info("No approved recipes to load")
		return nil
	}

	var success, fail int
	for _, mr := range moderationRecords {
		rec := mr.Recipe()
		if err := a.SaveRecipe(ctx, rec); err != nil {
			a.log.Error("Failed to save recipe", "slug", rec.Slug, "error", err)
			mrErr := mr.Errored(ctx, err)
			if mrErr != nil {
				a.log.Error("Failed to update moderation record", "slug", rec.Slug, "error", mrErr)
				fail++
				continue
			}
			fail++
			continue
		}

		if err := mr.Finish(ctx); err != nil {
			a.log.Error("Failed to finish moderation record", "slug", rec.Slug, "error", err)
			fail++
			continue
		}

		success++
	}

	a.log.Info("Processed approved recipes", "count", len(moderationRecords), "successful", success, "failed", fail)
	return nil
}
