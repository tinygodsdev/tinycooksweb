package locale

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
		Ingredients  string
		Instructions string
		Equipment    string
		Ideas        string
		Optional     string
		Required     string
	}

	UITransShare struct {
		HeaderMessage  string
		ExploreMessage string
		ShareMessage   string
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
			Ingredients:  "Ingredients",
			Instructions: "Instructions",
			Equipment:    "Equipment",
			Ideas:        "Ideas",
			Optional:     "Optional",
			Required:     "Required",
		},
		Share: &UITransShare{
			HeaderMessage:  "recipe 😋",
			ExploreMessage: "Explore",
			ShareMessage:   "Share recipe",
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
			Ingredients:  "Ингредиенты",
			Instructions: "Инструкции",
			Equipment:    "Оборудование",
			Ideas:        "Идеи",
			Optional:     "Опционально",
			Required:     "Обязательно",
		},
		Share: &UITransShare{
			HeaderMessage:  "рецепт 😋",
			ExploreMessage: "Посмотреть",
			ShareMessage:   "Поделиться рецептом",
		},
	}
}
