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
	EquipmentInstance struct {
		*CommonInstance
		Equipment    *recipe.Equipment
		Recipes      []*recipe.Recipe
		RecipesCount int
		Title        string
		Type         string
		Filter       recipe.Filter
	}
)

// must be present in all instances
func (ins *EquipmentInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) NewEquipmentInstance(s live.Socket) *EquipmentInstance {
	m, ok := s.Assigns().(*EquipmentInstance)
	if !ok {
		return &EquipmentInstance{
			CommonInstance: h.NewCommon(s, viewEquipment),
			Filter: recipe.Filter{
				Limit:  0, // query all
				Locale: locale.Default(),
			},
		}
	}

	return m
}

func (ins *EquipmentInstance) withError(err error) *EquipmentInstance {
	ins.Error = err
	return ins
}

func (h *Handler) Equipment() live.Handler {
	t := h.template("base.layout.html", "catalog-page")

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewEquipmentInstance // NB: make sure constructor is correct
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
		equipmentSlug, ok := mux.Vars(r)[ParamEquipmentSlug]
		if !ok {
			return nil, errors.New("recipe slug is required")
		}

		instance := h.NewEquipmentInstance(s)
		instance.fromContext(ctx)

		instance.Equipment, err = h.app.GetEquipmentBySlug(ctx, equipmentSlug)
		if err != nil {
			return instance.withError(err), nil
		}

		instance.Title = instance.Equipment.Name
		instance.Type = instance.UI.Catalog.Equipment
		instance.Filter = instance.Filter.WithAddEquipment(instance.Equipment.Name, recipe.SearchTypeInclude)
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
