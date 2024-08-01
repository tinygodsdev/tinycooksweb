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
	}
}
