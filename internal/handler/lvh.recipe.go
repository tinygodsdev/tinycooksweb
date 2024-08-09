package handler

import (
	"context"
	"errors"

	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"
)

const (
	// events
	// params
	ParamRecipeSlug = "recipeSlug"
)

type (
	RecipeInstance struct {
		*CommonInstance
		*Constants
		RecipeSlug string
		Recipe     *recipe.Recipe
	}
)

// must be present in all instances
func (ins *RecipeInstance) withError(err error) *RecipeInstance {
	ins.Error = err
	return ins
}

// must be present in all instances
func (ins *RecipeInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) NewRecipeInstance(s live.Socket) *RecipeInstance {
	m, ok := s.Assigns().(*RecipeInstance)
	if !ok {
		return &RecipeInstance{
			CommonInstance: h.NewCommon(s, viewRecipe),
			Constants:      h.NewConstants(),
		}
	}

	return m
}

func (h *Handler) Recipe() live.Handler {
	t := h.template("base.layout.html", "page.recipe.html")

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewRecipeInstance // NB: make sure constructor is correct
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
		recipeSlug, ok := mux.Vars(r)[ParamRecipeSlug]
		if !ok {
			return nil, errors.New("recipe slug is required")
		}

		instance := h.NewRecipeInstance(s)
		instance.fromContext(ctx)
		instance.RecipeSlug = recipeSlug

		// get recipe
		recipe, err := h.app.GetRecipe(ctx, recipeSlug)
		if err != nil {
			return instance.withError(err), nil
		}

		instance.Recipe = recipe
		instance.updateForLocale(ctx, s, h)
		return instance, nil
	})

	return lvh
}
