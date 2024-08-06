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
		Base   *UITransBase
		Home   *UITransHome
		Menu   *UITransMenu
		Footer *UITransFooter
		Recipe *UITransRecipe
		Share  *UITransShare
		// 404
		// about
		// profile
		// result
		// admin
		// privacy
		// terms
	}

	UITransBase struct {
		Title         string
		Description   string
		TwitterHandle string
	}

	UITransHome struct {
		Title       string
		Description string
	}

	UITransMenu struct {
		Home  string
		About string
		Back  string
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
		Include string
		Exclude string
		Add     string
		Clear   string
		Apply   string
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
			Title:       siteTitle,
			Description: "When in doubt - eat",
		},
		Menu: &UITransMenu{
			Home:  "Home",
			About: "About",
			Back:  "Back",
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
				Include: "Include",
				Exclude: "Exclude",
				Add:     "Add",
				Clear:   "Clear",
				Apply:   "Search!",
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
		Home: &UITransHome{
			Title:       siteTitle,
			Description: "Когда вы в сомнениях - ешьте",
		},
		Menu: &UITransMenu{
			Home:  "Главная",
			About: "О сайте",
			Back:  "Назад",
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
				Include: "Включить",
				Exclude: "Исключить",
				Add:     "Добавить",
				Clear:   "Очистить",
				Apply:   "Искать!",
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
