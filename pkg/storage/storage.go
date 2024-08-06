package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
)

var (
	ErrAlreadyExists = errors.New("resource already exists")
)

type Storage interface {
	SaveRecipe(ctx context.Context, recipe *recipe.Recipe) error
	GetRecipe(ctx context.Context, id uuid.UUID) (*recipe.Recipe, error)
	GetRecipeBySlug(ctx context.Context, slug string) (*recipe.Recipe, error)
	GetRecipes(ctx context.Context, filter recipe.Filter) ([]*recipe.Recipe, error)
	CountRecipes(ctx context.Context, filter recipe.Filter) (int, error)

	GetTags(ctx context.Context, locale string) ([]*recipe.Tag, error)
	GetIngredients(ctx context.Context, locale string) ([]*recipe.Ingredient, error)
	GetEquipment(ctx context.Context, locale string) ([]*recipe.Equipment, error)
}
