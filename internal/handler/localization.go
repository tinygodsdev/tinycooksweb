package handler

import "github.com/tinygodsdev/tinycooksweb/pkg/locale"

type (
	UITranslation struct {
		Base *UITransBase
		Home *UITransHome
		// 404
		// about
		// profile
		// result
		// admin
		// privacy
		// terms
	}

	UITransBase struct {
	}

	UITransHome struct {
		Title       string
		Description string
	}
)

func newUITranslation(loc string) *UITranslation {
	switch loc {
	case locale.En:
		return newTranslationEn()
	case locale.Ru:
		return newTranslationRu()
	default:
		return newUITranslation(locale.Default())
	}
}

func newTranslationEn() *UITranslation {
	return &UITranslation{
		Base: &UITransBase{},
		Home: &UITransHome{
			Title:       "TinyCooks",
			Description: "When in doubt - eat",
		},
	}
}

func newTranslationRu() *UITranslation {
	return &UITranslation{
		Base: &UITransBase{},
		Home: &UITransHome{
			Title:       "TinyCooks",
			Description: "Когда вы в сомнениях - ешьте",
		},
	}
}
