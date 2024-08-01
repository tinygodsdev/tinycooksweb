package recipe

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

func MockRecipes(recipeCount int, withError bool) ([]*Recipe, error) {
	if withError {
		return nil, errors.New("mock error")
	}

	rand.Seed(123)
	var recipes []*Recipe
	for i := 0; i < recipeCount; i++ {
		recipe := &Recipe{
			ID:           uuid.New(),
			Name:         randomString(10),
			Lang:         randomLang(),
			Slug:         randomString(10),
			Description:  randomString(50),
			Text:         randomString(200),
			Ingredients:  randomIngredients(),
			Instructions: randomInstructions(),
			Tags:         randomTags(),
			Ideas:        randomIdeas(),
			Time:         time.Duration(rand.Intn(120)) * time.Minute,
			Servings:     rand.Intn(10) + 1,
			Sources:      randomSources(),
			Nutrition:    randomNutrition(),
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomLang() string {
	langs := []string{LangEn, LangRu}
	return langs[rand.Intn(len(langs))]
}

func randomNutritionPrecision() string {
	precisions := []string{NutritionPrecisionExact, NutritionPrecisionApprox}
	return precisions[rand.Intn(len(precisions))]
}

func randomTags() []Tag {
	var tags []Tag
	tagCount := rand.Intn(5) + 1
	for i := 0; i < tagCount; i++ {
		tags = append(tags, Tag{
			ID:    uuid.New(),
			Name:  randomString(6),
			Group: randomString(4),
		})
	}
	return tags
}

func randomIngredients() []Ingredient {
	var ingredients []Ingredient
	ingredientCount := rand.Intn(10) + 1
	for i := 0; i < ingredientCount; i++ {
		ingredients = append(ingredients, Ingredient{
			ID:       uuid.New(),
			Name:     randomString(10),
			Quantity: randomString(3),
			Unit:     randomString(2),
		})
	}
	return ingredients
}

func randomInstructions() []Instruction {
	var instructions []Instruction
	instructionCount := rand.Intn(10) + 1
	for i := 0; i < instructionCount; i++ {
		instructions = append(instructions, Instruction{
			ID:   uuid.New(),
			Text: randomString(20),
		})
	}
	return instructions
}

func randomIdeas() []Idea {
	var ideas []Idea
	ideaCount := rand.Intn(5) + 1
	for i := 0; i < ideaCount; i++ {
		ideas = append(ideas, Idea{
			ID:   uuid.New(),
			Text: randomString(20),
		})
	}
	return ideas
}

func randomSources() []Source {
	var sources []Source
	sourceCount := rand.Intn(3) + 1
	for i := 0; i < sourceCount; i++ {
		sources = append(sources, Source{
			ID:          uuid.New(),
			Name:        randomString(10),
			Description: randomString(30),
			URL:         "http://" + randomString(10) + ".com",
		})
	}
	return sources
}

func randomNutrition() Nutrition {
	return Nutrition{
		Calories:  rand.Intn(500),
		Fat:       rand.Intn(50),
		Carbs:     rand.Intn(100),
		Protein:   rand.Intn(50),
		Precision: randomNutritionPrecision(),
		Benefits:  []string{randomString(10), randomString(10)},
	}
}
