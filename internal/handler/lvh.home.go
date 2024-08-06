package handler

import (
	"context"
	"fmt"
	"html/template"

	"github.com/jfyne/live"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
)

const (
	// filter events
	eventHomeTagsFilterChange        = "tags-filter-form-change"
	eventHomeIngredientsFilterChange = "ingredients-filter-form-change"
	eventHomeEquipmentFilterChange   = "equipment-filter-form-change"
	eventHomeTagsFilterAdd           = "tags-filter-form-submit"
	eventHomeIngredientsFilterAdd    = "ingredients-filter-form-submit"
	eventHomeEquipmentFilterAdd      = "equipment-filter-form-submit"
	eventHomeFilterClear             = "filter-clear"
	eventHomeFilterApply             = "filter-apply"

	// params
	paramHomeSearchType  = "searchType"
	paramHomeFilterValue = "value"
)

type HomeInstance struct {
	*CommonInstance
	Recipes      []*recipe.Recipe
	RecipesCount int
	Tags         []*recipe.Tag
	Ingredients  []*recipe.Ingredient
	Equipment    []*recipe.Equipment
	Filter       recipe.Filter
}

func (h *Handler) NewHomeInstance(s live.Socket) *HomeInstance {
	m, ok := s.Assigns().(*HomeInstance)
	if !ok {
		return &HomeInstance{
			CommonInstance: h.NewCommon(s, viewHome),
			Filter: recipe.Filter{
				Limit:    h.app.Cfg.PageSize,
				Locale:   locale.Default(),
				UseMocks: h.app.Cfg.MockQueries,
			},
		}
	}

	return m
}

func (ins *HomeInstance) withError(err error) *HomeInstance {
	ins.Error = err
	return ins
}

// must be present in all instances
func (ins *HomeInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) Home() live.Handler {
	t := template.Must(template.New("base.layout.html").Funcs(funcMap).ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.home.html",
	))

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewHomeInstance // NB: make sure constructor is correct
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

	// Set the mount function for this handler.
	lvh.HandleMount(func(ctx context.Context, s live.Socket) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.fromContext(ctx)

		instance.Recipes, err = h.app.GetRecipes(ctx, instance.Filter)
		if err != nil {
			return instance.withError(err), nil
		}

		instance.RecipesCount, err = h.app.CountRecipes(ctx, instance.Filter)
		if err != nil {
			return instance.withError(err), nil
		}

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

		instance.updateForLocale(ctx, s, h)
		return instance.withError(err), nil
	})

	lvh.HandleEvent(eventHomeTagsFilterChange, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		fmt.Printf("eventHomeTagSelect: %+v\n", p)
		return instance, nil
	})

	lvh.HandleEvent(eventHomeIngredientsFilterChange, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		fmt.Printf("eventHomeIngredientsFilterChange: %+v\n", p)
		return instance, nil
	})

	lvh.HandleEvent(eventHomeEquipmentFilterChange, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		fmt.Printf("eventHomeEquipmentFilterChange: %+v\n", p)
		return instance, nil
	})

	lvh.HandleEvent(eventHomeTagsFilterAdd, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		searchType := p.String(paramHomeSearchType)
		value := p.String(paramHomeFilterValue)
		instance.Filter = instance.Filter.WithAddTag(value, searchType)
		return instance, nil
	})

	lvh.HandleEvent(eventHomeIngredientsFilterAdd, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		searchType := p.String(paramHomeSearchType)
		value := p.String(paramHomeFilterValue)
		instance.Filter = instance.Filter.WithAddIngredient(value, searchType)
		return instance, nil
	})

	lvh.HandleEvent(eventHomeEquipmentFilterAdd, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		searchType := p.String(paramHomeSearchType)
		value := p.String(paramHomeFilterValue)
		instance.Filter = instance.Filter.WithAddEquipment(value, searchType)
		return instance, nil
	})

	lvh.HandleEvent(eventHomeFilterClear, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.Filter = instance.Filter.Clear()
		return instance, nil
	})

	lvh.HandleEvent(eventHomeFilterApply, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.Recipes, err = h.app.GetRecipes(ctx, instance.Filter)
		if err != nil {
			return instance.withError(err), nil
		}

		return instance, nil
	})

	return lvh
}
