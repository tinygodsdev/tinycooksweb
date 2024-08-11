package recipe

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
)

const (
	NutritionPrecisionExact        = "exact"
	NutritionPrecisionApprox       = "approx"
	NutritionPrecisionProfessional = "professional" // calculated by a professional
	NutritionPrecisionAuto         = "auto"         // calculated by the system

	LangEn = "en"
	LangRu = "ru"
)

type Recipe struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Lang         string        `json:"lang"`
	Slug         string        `json:"slug"`
	Description  string        `json:"description"` // Short description for catalog
	Text         string        `json:"text"`        // Long description for recipe page
	Tags         []*Tag        `json:"tags"`
	Ingredients  []*Ingredient `json:"ingredients"`
	Equipment    []*Equipment  `json:"equipment"`
	Instructions []Instruction `json:"instructions"` // Steps to prepare the recipe
	Ideas        []*Idea       `json:"ideas"`        // Ideas for variations
	Time         time.Duration `json:"time"`
	Servings     int           `json:"servings"`
	Sources      []*Source     `json:"sources"`
	Rating       float32       `json:"rating"`
	Nutrition    *Nutrition    `json:"nutrition"`
}

type Ingredient struct {
	ID       uuid.UUID `json:"id"`
	Product  *Product  `json:"product"`
	Quantity string    `json:"quantity"`
	Unit     string    `json:"unit"`
	Optional bool      `json:"optional"`
}

type Product struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
}

type Instruction struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

type Equipment struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
}

type Idea struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

type Source struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
}

type Nutrition struct {
	Calories  int      `json:"calories"`
	Fat       int      `json:"fat"`
	Carbs     int      `json:"carbs"`
	Protein   int      `json:"protein"`
	Precision string   `json:"precision"`
	Benefits  []string `json:"benefits"`
}

func Slugify(name string) string {
	return slug.Make(name)
}

func (r *Recipe) SlugifyAll() {
	r.Slug = Slugify(r.Name)
	for _, i := range r.Ingredients {
		i.Product.Slug = Slugify(i.Product.Name)
	}

	for _, e := range r.Equipment {
		e.Slug = Slugify(e.Name)
	}

	for _, t := range r.Tags {
		t.Slug = Slugify(t.Group + " " + t.Name)
	}
}

func (r *Recipe) Link(domain string) string {
	var params string
	if r.Lang != locale.Default() {
		params = fmt.Sprintf("?locale=%s", r.Lang)
	}
	path := fmt.Sprintf("/recipe/%s%s", r.Slug, params)
	if domain == "" {
		return path
	}
	recipeURL := fmt.Sprintf("https://%s%s", domain, path)
	return recipeURL
}

func (r *Recipe) ShareText() string {
	trans := locale.NewUITranslation(r.Lang)
	res := fmt.Sprintf("%s - %s\n\n", r.Name, trans.Share.HeaderMessage)
	if len(r.Description) > 215 {
		res += r.Description[:215]
	} else {
		res += r.Description
	}

	res += "\n\n" + trans.Share.ExploreMessage
	return res
}
