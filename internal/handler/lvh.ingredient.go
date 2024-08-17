package handler

import (
	"context"
	"errors"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
)

type (
	IngredientInstance struct {
		*CommonInstance
		Ingredient   *recipe.Ingredient
		Recipes      []*recipe.Recipe
		RecipesCount int
		Type         string
		Title        string
		Filter       recipe.Filter
	}
)

// must be present in all instances
func (ins *IngredientInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) NewIngredientInstance(s live.Socket) *IngredientInstance {
	m, ok := s.Assigns().(*IngredientInstance)
	if !ok {
		return &IngredientInstance{
			CommonInstance: h.NewCommon(s, viewIngredient),
			Filter: recipe.Filter{
				Limit:  0, // query all
				Locale: locale.Default(),
			},
		}
	}

	return m
}

func (ins *IngredientInstance) withError(err error) *IngredientInstance {
	ins.Error = err
	return ins
}

func (h *Handler) Ingredient() live.Handler {
	t := h.template("base.layout.html", "catalog-page")

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewIngredientInstance // NB: make sure constructor is correct
		// SAFE TO COPY

		// update locale logic
		lvh.HandleParams(func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.SetLocale(p.String(paramLocale))
			err := instance.updateForLocale(ctx, s, h)
			if err != nil {
				return nil, err
			}
			return instance, nil
		})

		lvh.HandleError(func(ctx context.Context, err error) {
			h.HandleError(ctx, err)
		})
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (i interface{}, err error) {
		r := live.Request(ctx)
		tagSlug, ok := mux.Vars(r)[ParamIngredientSlug]
		if !ok {
			return nil, errors.New("recipe slug is required")
		}

		instance := h.NewIngredientInstance(s)
		instance.fromContext(ctx)

		instance.Ingredient, err = h.app.GetIngredientBySlug(ctx, tagSlug)
		if err != nil {
			return instance.withError(err), nil
		}

		instance.Title = instance.Ingredient.Product.Name
		instance.Type = instance.UI.Catalog.Ingredient
		instance.Filter = instance.Filter.WithAddIngredient(
			instance.Ingredient.Product.Name,
			recipe.SearchTypeInclude,
		)
		instance.Recipes, err = h.app.GetRecipes(ctx, instance.Filter)
		if err != nil {
			return instance.withError(err), nil
		}

		instance.RecipesCount, err = h.app.CountRecipes(ctx, instance.Filter)
		if err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	return lvh
}
