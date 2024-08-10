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
	TagInstance struct {
		*CommonInstance
		Tag          *recipe.Tag
		Recipes      []*recipe.Recipe
		RecipesCount int
		Type         string
		Title        string
		Filter       recipe.Filter
	}
)

// must be present in all instances
func (ins *TagInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) NewTagInstance(s live.Socket) *TagInstance {
	m, ok := s.Assigns().(*TagInstance)
	if !ok {
		return &TagInstance{
			CommonInstance: h.NewCommon(s, viewTag),
			Filter: recipe.Filter{
				Limit:  0, // query all
				Locale: locale.Default(),
			},
		}
	}

	return m
}

func (ins *TagInstance) withError(err error) *TagInstance {
	ins.Error = err
	return ins
}

func (h *Handler) Tag() live.Handler {
	t := h.template("base.layout.html", "page.catalog_page.html")

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewTagInstance // NB: make sure constructor is correct
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
		tagSlug, ok := mux.Vars(r)[ParamTagSlug]
		if !ok {
			return nil, errors.New("recipe slug is required")
		}

		instance := h.NewTagInstance(s)
		instance.fromContext(ctx)

		instance.Tag, err = h.app.GetTagBySlug(ctx, tagSlug)
		if err != nil {
			return instance.withError(err), nil
		}

		instance.Title = instance.Tag.Title()
		instance.Type = instance.UI.Catalog.Tag
		instance.Filter = instance.Filter.WithAddTag(instance.Tag.Name, recipe.SearchTypeInclude)
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
