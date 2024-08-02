package recipe

import (
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
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
	Ingredients  []*Ingredient `json:"ingredients"`
	Equipment    []*Equipment  `json:"equipment"`
	Instructions []Instruction `json:"instructions"` // Steps to prepare the recipe
	Tags         []*Tag        `json:"tags"`
	Ideas        []*Idea       `json:"ideas"` // Ideas for variations
	Time         time.Duration `json:"time"`
	Servings     int           `json:"servings"`
	Sources      []*Source     `json:"sources"`
	Nutrition    *Nutrition    `json:"nutrition"`
}

type Ingredient struct {
	ID       uuid.UUID `json:"id"`
	Product  *Product  `json:"product"`
	Quantity string    `json:"quantity"`
	Unit     string    `json:"unit"`
}

type Product struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Instruction struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

type Equipment struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Tag struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Group string    `json:"group"`
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

type Filter struct {
	NameContains  string
	Locale        string
	Equipment     string
	EquipmentNot  string
	Tag           string
	TagNot        string
	Ingredient    string
	IngredientNot string
	Limit         int
	Offset        int

	UseMocks bool
}
