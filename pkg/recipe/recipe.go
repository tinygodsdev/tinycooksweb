package recipe

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"golang.org/x/exp/rand"
	"gopkg.in/yaml.v3"
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
	ID               uuid.UUID         `json:"-" yaml:"-"`
	Slug             string            `json:"-" yaml:"-"`
	Name             string            `json:"name"`
	Description      string            `json:"description"` // Short description for catalog
	Text             string            `json:"text"`        // Long description for recipe page
	Tags             []*Tag            `json:"tags"`
	Ingredients      []*Ingredient     `json:"ingredients"`
	Equipment        []*Equipment      `json:"equipment"`
	Instructions     []Instruction     `json:"instructions"` // Steps to prepare the recipe
	Ideas            []*Idea           `json:"ideas"`        // Ideas for variations
	Time             time.Duration     `json:"time"`
	Servings         int               `json:"servings"`
	Nutrition        *Nutrition        `json:"nutrition"`
	Meta             map[string]string `json:"meta"`
	Lang             string            `json:"lang"`
	Rating           float32           `json:"rating"`
	Sources          []*Source         `json:"sources"`
	ModerationStatus string            `json:"moderation_status" yaml:"-"`
	Published        bool              `json:"published"`
}

type Ingredient struct {
	ID       uuid.UUID `json:"-" yaml:"-"`
	Product  *Product  `json:"product"`
	Quantity string    `json:"quantity"`
	Unit     string    `json:"unit"`
	Optional bool      `json:"optional,omitempty" yaml:",omitempty"`
}

type Product struct {
	ID   uuid.UUID `json:"-" yaml:"-"`
	Name string    `json:"name"`
	Slug string    `json:"-" yaml:"-"`
}

type Instruction struct {
	ID   uuid.UUID `json:"-" yaml:"-"`
	Text string    `json:"text"`
}

type Equipment struct {
	ID   uuid.UUID `json:"-" yaml:"-"`
	Name string    `json:"name"`
	Slug string    `json:"-" yaml:"-"`
}

type Idea struct {
	ID   uuid.UUID `json:"-" yaml:"-"`
	Text string    `json:"text"`
}

type Source struct {
	ID          uuid.UUID `json:"-" yaml:"-"`
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

func (r *Recipe) slugifyAll() {
	r.Name = strings.ToLower(r.Name)
	r.Slug = Slugify(r.Name)
	for _, i := range r.Ingredients {
		i.Product.Name = strings.ToLower(i.Product.Name)
		i.Product.Slug = Slugify(i.Product.Name)
	}

	for _, e := range r.Equipment {
		e.Name = strings.ToLower(e.Name)
		e.Slug = Slugify(e.Name)
	}

	for _, t := range r.Tags {
		t.Group = strings.ToLower(t.Group)
		t.Name = strings.ToLower(t.Name)
		t.Slug = Slugify(t.Group + " " + t.Name)
	}
}

func (r *Recipe) addRandomRating() {
	if r.Rating == 0 {
		rand.Seed(uint64(time.Now().UnixNano()))
		ratingOptions := []float32{4.5, 4.6, 4.7, 4.8, 4.9, 5.0}
		r.Rating = ratingOptions[rand.Intn(len(ratingOptions))]
	}
}

func (r *Recipe) addTimeTag() {
	// remove existing time tag
	for i, t := range r.Tags {
		if t.Group == locale.NewUITranslation(r.Lang).Tag.TimeGroup {
			r.Tags = append(r.Tags[:i], r.Tags[i+1:]...)
			break
		}
	}

	trans := locale.NewUITranslation(r.Lang)
	group := trans.Tag.TimeGroup
	var timeTag string

	switch {
	case r.Time.Minutes() < 30:
		timeTag = trans.Tag.TimeFast
	case r.Time.Minutes() < 60:
		timeTag = trans.Tag.TimeMedium
	case r.Time.Minutes() >= 60:
		timeTag = trans.Tag.TimeLong
	}

	if timeTag != "" {
		r.Tags = append(r.Tags, &Tag{Name: timeTag, Group: group})
	}
}

func (r *Recipe) PostProcess() {
	r.addTimeTag()
	r.addRandomRating()
	r.slugifyAll()
}

func (r *Recipe) Link(baseURL string) string {
	var params string
	if r.Lang != locale.Default() {
		params = fmt.Sprintf("?locale=%s", r.Lang)
	}
	path := fmt.Sprintf("/recipe/%s%s", r.Slug, params)
	if baseURL == "" {
		return path
	}
	recipeURL := fmt.Sprintf("%s%s", baseURL, path)
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

func (e *Equipment) Slugify() {
	e.Name = strings.ToLower(e.Name)
	e.Slug = Slugify(e.Name)
}

func (i *Ingredient) Slugify() {
	i.Product.Name = strings.ToLower(i.Product.Name)
	i.Product.Slug = Slugify(i.Product.Name)
}

func LoadRecipeFromString(data string) (*Recipe, error) {
	lines := strings.Split(data, "\n")

	if len(lines) > 1 && strings.HasPrefix(lines[0], "```") && strings.HasPrefix(lines[len(lines)-1], "```") {
		lines = lines[1 : len(lines)-1]
	}

	data = strings.Join(lines, "\n")

	var r Recipe

	err := yaml.Unmarshal([]byte(data), &r)
	if err != nil {
		return nil, err
	}

	r.Time = time.Duration(r.Time)

	r.PostProcess()
	return &r, nil
}

func (n *Nutrition) IsEmpty() bool {
	return n.Calories == 0 && n.Fat == 0 && n.Carbs == 0 && n.Protein == 0 && len(n.Benefits) == 0
}
