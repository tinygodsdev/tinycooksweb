package recipe

import (
	"time"

	"github.com/google/uuid"
)

const (
	NutritionPrecisionExact  = "exact"
	NutritionPrecisionApprox = "approx"

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
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"` // Steps to prepare the recipe
	Tags         []Tag         `json:"tags"`
	Ideas        []Idea        `json:"ideas"` // Ideas for variations
	Time         time.Duration `json:"time"`
	Servings     int           `json:"servings"`
	Sources      []Source      `json:"sources"`
	Nutrition    Nutrition     `json:"nutrition"`
}

type Ingredient struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Quantity string    `json:"quantity"`
	Unit     string    `json:"unit"`
}

type Instruction struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
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
