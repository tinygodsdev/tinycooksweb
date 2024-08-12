package handler

import (
	"context"
	"fmt"
	"net/url"
	"slices"

	"github.com/jfyne/live"
	"github.com/samber/lo"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
)

const (
	// filter events
	eventHomeNameFilterChange        = "name-filter-form-change"
	eventHomeNameFilterAdd           = "name-filter-form-submit"
	eventHomeTagsFilterChange        = "tags-filter-form-change"
	eventHomeTagsFilterAdd           = "tags-filter-form-submit"
	eventHomeIngredientsFilterChange = "ingredients-filter-form-change"
	eventHomeIngredientsFilterAdd    = "ingredients-filter-form-submit"
	eventHomeEquipmentFilterChange   = "equipment-filter-form-change"
	eventHomeEquipmentFilterAdd      = "equipment-filter-form-submit"
	eventHomeFilterClear             = "filter-clear"
	eventHomeFilterApply             = "filter-apply"

	eventHomeTagsFilterDelete        = "tags-filter-delete"
	eventHomeIngredientsFilterDelete = "ingredients-filter-delete"
	eventHomeEquipmentFilterDelete   = "equipment-filter-delete"

	// params
	paramHomeSearchType   = "searchtype"
	paramHomeTagGroup     = "group"
	paramHomeFilterValue  = "value"
	paramHomeFilterDelete = "delete"
	paramHomePage         = "page"
)

type HomeInstance struct {
	*CommonInstance
	Recipes       []*recipe.Recipe
	RecipesCount  int
	Tags          []*recipe.Tag
	TagGroups     []string
	SelectedGroup string
	FilteredTags  []*recipe.Tag
	Ingredients   []*recipe.Ingredient
	Equipment     []*recipe.Equipment
	Filter        recipe.Filter
	Pagination    Pagination
}

func (h *Handler) NewHomeInstance(s live.Socket) *HomeInstance {
	m, ok := s.Assigns().(*HomeInstance)
	if !ok {
		return &HomeInstance{
			CommonInstance: h.NewCommon(s, viewHome),
			Filter: recipe.Filter{
				Limit:  h.app.Cfg.PageSize,
				Locale: locale.Default(),
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

func (ins *HomeInstance) WithUpdateRecipes(
	ctx context.Context,
	h *Handler,
	s live.Socket,
	resetOffset bool,
) (*HomeInstance, error) {
	if resetOffset && ins.Filter.Offset != 0 {
		ins.Filter.Offset = 0
		v := url.Values{}
		v.Set(paramHomePage, "1")
		s.PatchURL(v)
	}

	ins.Recipes, ins.Error = h.app.GetRecipes(ctx, ins.Filter)
	if ins.Error != nil {
		return ins, nil
	}

	ins.RecipesCount, ins.Error = h.app.CountRecipes(ctx, ins.Filter)
	if ins.Error != nil {
		return ins, nil
	}

	ins.Pagination = calculatePagination(ins.RecipesCount, ins.Filter.Limit, ins.Filter.Offset)
	return ins, nil
}

func (h *Handler) Home() live.Handler {
	t := h.template("base.layout.html", "home")

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

		instance.Tags, err = h.app.GetTags(ctx, instance.Locale())
		if err != nil {
			return instance.withError(err), nil
		}

		instance.TagGroups = recipe.GetTagGroups(instance.Tags)

		instance.Ingredients, err = h.app.GetIngredients(ctx, instance.Locale())
		if err != nil {
			return instance.withError(err), nil
		}

		instance.Equipment, err = h.app.GetEquipment(ctx, instance.Locale())
		if err != nil {
			return instance.withError(err), nil
		}

		instance.updateForLocale(ctx, s, h)
		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleEvent(eventHomeTagsFilterChange, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		fmt.Printf("eventHomeTagSelect: %+v\n", p)
		instance.SelectedGroup = p.String(paramHomeTagGroup)
		instance.FilteredTags = recipe.FilterTagsByGroup(instance.Tags, instance.SelectedGroup)
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

	lvh.HandleEvent(eventHomeNameFilterChange, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.Filter = instance.Filter.WithName(p.String(paramHomeFilterValue))
		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleEvent(eventHomeTagsFilterAdd, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		searchType := p.String(paramHomeSearchType)
		value := p.String(paramHomeFilterValue)
		if slices.Contains[[]string](
			lo.Map(instance.Tags, func(t *recipe.Tag, _ int) string { return t.Name }),
			value,
		) {
			instance.Filter = instance.Filter.WithAddTag(value, searchType)
		}

		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleEvent(eventHomeIngredientsFilterAdd, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		searchType := p.String(paramHomeSearchType)
		value := p.String(paramHomeFilterValue)
		if slices.Contains[[]string](
			lo.Map(instance.Ingredients, func(i *recipe.Ingredient, _ int) string { return i.Product.Name }),
			value,
		) {
			instance.Filter = instance.Filter.WithAddIngredient(value, searchType)
		}

		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleEvent(eventHomeEquipmentFilterAdd, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		searchType := p.String(paramHomeSearchType)
		value := p.String(paramHomeFilterValue)
		if slices.Contains[[]string](
			lo.Map(instance.Equipment, func(e *recipe.Equipment, _ int) string { return e.Name }),
			value,
		) {
			instance.Filter = instance.Filter.WithAddEquipment(value, searchType)
		}

		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleEvent(eventHomeFilterClear, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.Filter = instance.Filter.Clear()
		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleEvent(eventHomeFilterApply, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleEvent(eventHomeTagsFilterDelete, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.Filter = instance.Filter.WithRemoveTag(
			p.String(paramHomeFilterDelete),
			p.String(paramHomeSearchType),
		)
		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleEvent(eventHomeIngredientsFilterDelete, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.Filter = instance.Filter.WithRemoveIngredient(
			p.String(paramHomeFilterDelete),
			p.String(paramHomeSearchType),
		)
		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleEvent(eventHomeEquipmentFilterDelete, func(ctx context.Context, s live.Socket, p live.Params) (i interface{}, err error) {
		instance := h.NewHomeInstance(s)
		instance.Filter = instance.Filter.WithRemoveEquipment(
			p.String(paramHomeFilterDelete),
			p.String(paramHomeSearchType),
		)
		return instance.WithUpdateRecipes(ctx, h, s, true)
	})

	lvh.HandleParams(func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
		instance := h.NewHomeInstance(s)
		page := p.Int(paramHomePage)
		instance.Filter.Offset = onPageClick(page, instance.Filter.Limit)
		return instance.WithUpdateRecipes(ctx, h, s, false)
	})

	return lvh
}
