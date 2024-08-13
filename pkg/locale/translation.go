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
		Content     string
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
		Nutrition    *UITransNutrition
		Filter       *UIFilter
	}

	UITransShare struct {
		HeaderMessage  string
		ExploreMessage string
		ShareMessage   string
	}

	UITransNutrition struct {
		Calories              string
		Protein               string
		Fat                   string
		Carbs                 string
		PrecisionAuto         string
		PrecisionProfessional string
		PrecisionExact        string
		PrecisionApprox       string
	}

	UIFilter struct {
		Title        string
		Description  string
		Include      string
		Exclude      string
		NameContains string
		Add          string
		Clear        string
		Apply        string
	}
)

func (nt *UITransNutrition) Precision(prec string) string {
	// TODO: decide something with those magic strings
	switch prec {
	case "auto":
		return nt.PrecisionAuto
	case "professional":
		return nt.PrecisionProfessional
	case "exact":
		return nt.PrecisionExact
	case "approx":
		return nt.PrecisionApprox
	default:
		return nt.PrecisionAuto
	}
}

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
			Content:     "TinyCooks is your simple and convenient kitchen companion. We've gathered recipes from around the world and made them easily accessible, free from distractions. On our site, you'll find a flexible search and clean pages without overwhelming ads or unnecessary information. Just pick a recipe and enjoy cooking. With TinyCooks, cooking becomes an easy and pleasant task.",
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
				Title:        "Search recipes",
				Description:  "Find recipes by name, tags, ingredients, required equipment",
				Include:      "With",
				Exclude:      "Without",
				Add:          "Add",
				Clear:        "Clear",
				Apply:        "Search!",
				NameContains: "Name contains",
			},
			Nutrition: &UITransNutrition{
				Calories:              "Calories",
				Protein:               "Protein",
				Fat:                   "Fat",
				Carbs:                 "Carbs",
				PrecisionAuto:         "Auto",
				PrecisionProfessional: "Professional",
				PrecisionExact:        "Exact",
				PrecisionApprox:       "Approx",
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
			Content:     "TinyCooks — это ваш простой и удобный помощник на кухне. Мы собрали рецепты со всего мира и сделали их доступными для вас без лишних отвлекающих элементов. На нашем сайте вы найдете гибкий поиск и чистые страницы без обилия рекламы и ненужной информации. Просто выбирайте рецепт и готовьте с удовольствием. С TinyCooks готовка становится легкой и приятной задачей.",
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
				Title:        "Поиск рецептов",
				Description:  "Находите рецепты по названию, тегам, ингредиентам, необходимому оборудованию",
				Include:      "C",
				Exclude:      "Без",
				Add:          "Добавить",
				Clear:        "Очистить",
				Apply:        "Искать!",
				NameContains: "Название содержит",
			},
			Nutrition: &UITransNutrition{
				Calories:              "Калории",
				Protein:               "Белки",
				Fat:                   "Жиры",
				Carbs:                 "Углеводы",
				PrecisionAuto:         "Авто",
				PrecisionProfessional: "Профессионально",
				PrecisionExact:        "Точно",
				PrecisionApprox:       "Приблизительно",
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
