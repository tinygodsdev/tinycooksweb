package recipe

import (
	"github.com/tinygodsdev/tinycooksweb/internal/util"
)

const (
	SearchTypeInclude = "include"
	SearchTypeExclude = "exclude"
)

type Filter struct {
	NameContains  string
	Locale        string
	Equipment     []string
	EquipmentNot  []string
	Tag           []string
	TagNot        []string
	Ingredient    []string
	IngredientNot []string
	Limit         int
	Offset        int
	NewFirst      bool

	WithEdges bool
}

func (f Filter) Clear() Filter {
	f.NameContains = ""
	f.Equipment = []string{}
	f.EquipmentNot = []string{}
	f.Tag = []string{}
	f.TagNot = []string{}
	f.Ingredient = []string{}
	f.IngredientNot = []string{}

	return f
}

func (f Filter) WithAddTag(tag string, searchType string) Filter {
	switch searchType {
	case SearchTypeInclude:
		f.TagNot = util.DeleteString(f.TagNot, tag)
		f.Tag = util.AppendUniqueString(f.Tag, tag)
	case SearchTypeExclude:
		f.Tag = util.DeleteString(f.Tag, tag)
		f.TagNot = util.AppendUniqueString(f.TagNot, tag)
	}

	return f
}

func (f Filter) WithAddIngredient(ingredient string, searchType string) Filter {
	switch searchType {
	case SearchTypeInclude:
		f.IngredientNot = util.DeleteString(f.IngredientNot, ingredient)
		f.Ingredient = util.AppendUniqueString(f.Ingredient, ingredient)
	case SearchTypeExclude:
		f.Ingredient = util.DeleteString(f.Ingredient, ingredient)
		f.IngredientNot = util.AppendUniqueString(f.IngredientNot, ingredient)
	}

	return f
}

func (f Filter) WithAddEquipment(equipment string, searchType string) Filter {
	switch searchType {
	case SearchTypeInclude:
		f.EquipmentNot = util.DeleteString(f.EquipmentNot, equipment)
		f.Equipment = util.AppendUniqueString(f.Equipment, equipment)
	case SearchTypeExclude:
		f.Equipment = util.DeleteString(f.Equipment, equipment)
		f.EquipmentNot = util.AppendUniqueString(f.EquipmentNot, equipment)
	}

	return f
}

func (f Filter) WithRemoveTag(tag string, searchType string) Filter {
	switch searchType {
	case SearchTypeInclude:
		f.Tag = util.DeleteString(f.Tag, tag)
	case SearchTypeExclude:
		f.TagNot = util.DeleteString(f.TagNot, tag)
	}

	return f
}

func (f Filter) WithRemoveIngredient(ingredient string, searchType string) Filter {
	switch searchType {
	case SearchTypeInclude:
		f.Ingredient = util.DeleteString(f.Ingredient, ingredient)
	case SearchTypeExclude:
		f.IngredientNot = util.DeleteString(f.IngredientNot, ingredient)
	}

	return f
}

func (f Filter) WithRemoveEquipment(equipment string, searchType string) Filter {
	switch searchType {
	case SearchTypeInclude:
		f.Equipment = util.DeleteString(f.Equipment, equipment)
	case SearchTypeExclude:
		f.EquipmentNot = util.DeleteString(f.EquipmentNot, equipment)
	}

	return f
}

func (f Filter) WithName(name string) Filter {
	f.NameContains = name
	return f
}

func (f Filter) WithNewFirst(b bool) Filter {
	f.NewFirst = b
	return f
}

func (f Filter) IsEmpty() bool {
	return f.NameContains == "" &&
		len(f.Equipment) == 0 &&
		len(f.EquipmentNot) == 0 &&
		len(f.Tag) == 0 &&
		len(f.TagNot) == 0 &&
		len(f.Ingredient) == 0 &&
		len(f.IngredientNot) == 0
}

func (f Filter) HasTagRules() bool {
	return len(f.Tag) > 0 || len(f.TagNot) > 0
}

func (f Filter) HasIngredientRules() bool {
	return len(f.Ingredient) > 0 || len(f.IngredientNot) > 0
}

func (f Filter) HasEquipmentRules() bool {
	return len(f.Equipment) > 0 || len(f.EquipmentNot) > 0
}
