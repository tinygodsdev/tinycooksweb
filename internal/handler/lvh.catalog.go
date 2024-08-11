package handler

import (
	"context"

	"github.com/jfyne/live"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
)

type (
	CatalogInstance struct {
		*CommonInstance
		Tags        []*recipe.Tag
		Ingredients []*recipe.Ingredient
		Equipment   []*recipe.Equipment
	}
)

// must be present in all instances
func (ins *CatalogInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) NewCatalogInstance(s live.Socket) *CatalogInstance {
	m, ok := s.Assigns().(*CatalogInstance)
	if !ok {
		return &CatalogInstance{
			CommonInstance: h.NewCommon(s, viewCatalog),
		}
	}

	return m
}

func (ins *CatalogInstance) withError(err error) *CatalogInstance {
	ins.Error = err
	return ins
}

func (h *Handler) Catalog() live.Handler {
	t := h.template("base.layout.html", "catalog")

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewCatalogInstance // NB: make sure constructor is correct
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
		instance := h.NewCatalogInstance(s)
		instance.fromContext(ctx)

		instance.Tags, err = h.app.GetTags(ctx, instance.Locale())
		if err != nil {
			return instance.withError(err), nil
		}

		instance.Ingredients, err = h.app.GetIngredients(ctx, instance.Locale())
		if err != nil {
			return instance.withError(err), nil
		}

		instance.Equipment, err = h.app.GetEquipment(ctx, instance.Locale())
		if err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	return lvh
}
