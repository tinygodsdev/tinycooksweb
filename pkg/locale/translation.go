package locale

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	siteTitle    = "TinyCooks"
	developer    = "TinyGods"
	developerURL = "https://tinygods.dev"
)

type (
	UITranslation struct {
		Base    *UITransBase
		Home    *UITransHome
		About   *UITransAbout
		Catalog *UITransCatalog
		Menu    *UITransMenu
		Footer  *UITransFooter
		Recipe  *UITransRecipe
		Share   *UITransShare
		// 404
		// privacy
		// terms
	}

	UITransBase struct {
		Title         string
		Description   string
		TwitterHandle string
	}

	UITransAbout struct {
		Title       string
		Description string
	}

	UITransCatalog struct {
		Title            string
		Description      string
		RecipesTotal     string
		TagsTotal        string
		IngredientsTotal string
		EquipmentTotal   string
		ByTags           string
		ByIngredients    string
		ByEquipment      string
		Tag              string
		Ingredient       string
		Equipment        string
	}

	UITransHome struct {
		Title        string
		Description  string
		RecipesTotal string
		RecipesFound string
		Next         string
		Previous     string
	}

	UITransMenu struct {
		Home    string
		About   string
		Back    string
		Catalog string
	}

	UITransFooter struct {
		DevelopedBy  string
		Developer    string
		DeveloperURL string
	}

	UITransRecipe struct {
		Tags         string
		Ingredients  string
		Instructions string
		Equipment    string
		Ideas        string
		Optional     string
		Required     string
		Filter       *UIFilter
	}

	UITransShare struct {
		HeaderMessage  string
		ExploreMessage string
		ShareMessage   string
	}

	UIFilter struct {
		Include      string
		Exclude      string
		NameContains string
		Add          string
		Clear        string
		Apply        string
	}
)

func NewUITranslation(loc string) *UITranslation {
	switch loc {
	case En:
		return newTranslationEn()
	case Ru:
		return newTranslationRu()
	default:
		return NewUITranslation(Default())
	}
}

func newTranslationEn() *UITranslation {
	return &UITranslation{
		Base: &UITransBase{
			Title:         siteTitle,
			Description:   "When in doubt - eat",
			TwitterHandle: "danipolani",
		},
		Home: &UITransHome{
			Title:        siteTitle,
			Description:  "When in doubt - eat",
			RecipesTotal: "Total recipes",
			RecipesFound: "Recipes found",
			Next:         "Next",
			Previous:     "Previous",
		},
		About: &UITransAbout{
			Title:       "About",
			Description: "TinyCooks is a collection of recipes from around the world",
		},
		Catalog: &UITransCatalog{
			Title:            "Catalog",
			Description:      "Browse all recipes",
			RecipesTotal:     "Total recipes",
			TagsTotal:        "Total tags",
			IngredientsTotal: "Total ingredients",
			EquipmentTotal:   "Total equipment",
			ByTags:           "By Tags",
			ByIngredients:    "By Ingredients",
			ByEquipment:      "By Equipment",
			Tag:              "Tag",
			Ingredient:       "Ingredient",
			Equipment:        "Equipment",
		},
		Menu: &UITransMenu{
			Home:    "Home",
			About:   "About",
			Back:    "Back",
			Catalog: "Catalog",
		},
		Footer: &UITransFooter{
			DevelopedBy:  "Developed by",
			Developer:    developer,
			DeveloperURL: developerURL,
		},
		Recipe: &UITransRecipe{
			Tags:         "Tags",
			Ingredients:  "Ingredients",
			Instructions: "Instructions",
			Equipment:    "Equipment",
			Ideas:        "Ideas",
			Optional:     "Optional",
			Required:     "Required",
			Filter: &UIFilter{
				Include:      "Include",
				Exclude:      "Exclude",
				Add:          "Add",
				Clear:        "Clear",
				Apply:        "Search!",
				NameContains: "Name contains...",
			},
		},
		Share: &UITransShare{
			HeaderMessage:  "recipe 😋",
			ExploreMessage: "Explore",
			ShareMessage:   "Share on",
		},
	}
}

func newTranslationRu() *UITranslation {
	return &UITranslation{
		Base: &UITransBase{
			Title:         siteTitle,
			Description:   "Когда вы в сомнениях - ешьте",
			TwitterHandle: "danipolani",
		},
		About: &UITransAbout{
			Title:       "О сайте",
			Description: "TinyCooks - это коллекция рецептов со всего мира",
		},
		Catalog: &UITransCatalog{
			Title:            "Каталог",
			Description:      "Просмотреть все рецепты",
			RecipesTotal:     "Всего рецептов",
			TagsTotal:        "Всего тегов",
			IngredientsTotal: "Всего ингредиентов",
			EquipmentTotal:   "Всего оборудования",
			ByTags:           "По тегам",
			ByIngredients:    "По ингредиентам",
			ByEquipment:      "По оборудованию",
			Tag:              "Тег",
			Ingredient:       "Ингредиент",
			Equipment:        "Оборудование",
		},
		Home: &UITransHome{
			Title:        siteTitle,
			Description:  "Когда вы в сомнениях - ешьте",
			RecipesTotal: "Всего рецептов",
			RecipesFound: "Найдено рецептов",
			Next:         "Дальше",
			Previous:     "Назад",
		},
		Menu: &UITransMenu{
			Home:    "Главная",
			About:   "О сайте",
			Back:    "Назад",
			Catalog: "Каталог",
		},
		Footer: &UITransFooter{
			DevelopedBy:  "Разработано",
			Developer:    developer,
			DeveloperURL: developerURL,
		},
		Recipe: &UITransRecipe{
			Tags:         "Теги",
			Ingredients:  "Ингредиенты",
			Instructions: "Инструкции",
			Equipment:    "Оборудование",
			Ideas:        "Идеи",
			Optional:     "Опционально",
			Required:     "Обязательно",
			Filter: &UIFilter{
				Include:      "Включить",
				Exclude:      "Исключить",
				Add:          "Добавить",
				Clear:        "Очистить",
				Apply:        "Искать!",
				NameContains: "Название содержит...",
			},
		},
		Share: &UITransShare{
			HeaderMessage:  "рецепт 😋",
			ExploreMessage: "Посмотреть",
			ShareMessage:   "Поделиться в",
		},
	}
}

// Validate checks if all fields are set
func (t *UITranslation) Validate() error {
	return validateStruct(t)
}

func validateStruct(s interface{}) error {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return errors.New("expected a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return fmt.Errorf("field %s is empty", fieldType.Name)
			}
		case reflect.Ptr:
			if field.IsNil() {
				return fmt.Errorf("field %s is nil", fieldType.Name)
			}
			// Recursively validate nested structs
			if err := validateStruct(field.Interface()); err != nil {
				return err
			}
		case reflect.Struct:
			// Recursively validate nested structs
			if err := validateStruct(field.Interface()); err != nil {
				return err
			}
		}
	}

	return nil
}
