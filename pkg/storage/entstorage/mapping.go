package entstorage

import (
	"time"

	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent"
)

func entRecipeToRecipe(r *ent.Recipe) *recipe.Recipe {
	var servings int
	if r.Servings != nil {
		servings = *r.Servings
	}

	var time time.Duration
	if r.Time != nil {
		time = *r.Time
	}

	var nutrition *recipe.Nutrition
	if r.Edges.Nutrition != nil {
		nutrition = entNutritionToNutrition(r.Edges.Nutrition)
	}

	return &recipe.Recipe{
		ID:           r.ID,
		Name:         r.Name,
		Lang:         r.Locale.String(),
		Slug:         r.Slug,
		Description:  r.Description,
		Text:         r.Text,
		Servings:     servings,
		Time:         time,
		Rating:       r.Rating,
		Ingredients:  mapEntIngredientsToIngredients(r.Edges.Ingredients),
		Instructions: mapEntInstructionsToInstructions(r.Edges.Instructions),
		Tags:         mapEntTagsToTags(r.Edges.Tags),
		Ideas:        mapEntIdeasToIdeas(r.Edges.Ideas),
		Sources:      mapEntSourcesToSources(r.Edges.Sources),
		Nutrition:    nutrition,
		Equipment:    mapEntEquipmentToEquipment(r.Edges.Equipment),
	}
}

func entIngredientToIngredient(i *ent.Ingredient) *recipe.Ingredient {
	return &recipe.Ingredient{
		ID:       i.ID,
		Product:  entProductToProduct(i.Edges.Product),
		Quantity: i.Quantity,
		Unit:     i.Unit,
		Optional: i.Optional,
	}
}

func entProductToProduct(p *ent.Product) *recipe.Product {
	return &recipe.Product{
		ID:   p.ID,
		Name: p.Name,
		Slug: p.Slug,
	}
}

func entInstructionToInstruction(i *ent.Instruction) recipe.Instruction {
	return recipe.Instruction{
		ID:   i.ID,
		Text: i.Text,
	}
}

func entTagToTag(t *ent.Tag) *recipe.Tag {
	return &recipe.Tag{
		ID:    t.ID,
		Name:  t.Name,
		Group: t.Group,
		Slug:  t.Slug,
	}
}

func entIdeaToIdea(i *ent.Idea) *recipe.Idea {
	return &recipe.Idea{
		ID:   i.ID,
		Text: i.Text,
	}
}

func entSourceToSource(s *ent.Source) *recipe.Source {
	return &recipe.Source{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		URL:         s.URL,
	}
}

func entNutritionToNutrition(n *ent.Nutrition) *recipe.Nutrition {
	return &recipe.Nutrition{
		Calories:  n.Calories,
		Fat:       n.Fat,
		Carbs:     n.Carbs,
		Protein:   n.Protein,
		Precision: n.Precision,
		Benefits:  n.Benefits,
	}
}

func entEquipmentToEquipment(e *ent.Equipment) *recipe.Equipment {
	return &recipe.Equipment{
		ID:   e.ID,
		Name: e.Name,
		Slug: e.Slug,
	}
}

func entProductToIngredient(p *ent.Product) *recipe.Ingredient {
	return &recipe.Ingredient{
		Product: entProductToProduct(p),
	}
}

func mapEntIngredientsToIngredients(ings []*ent.Ingredient) []*recipe.Ingredient {
	var res []*recipe.Ingredient
	for _, i := range ings {
		res = append(res, entIngredientToIngredient(i))
	}
	return res
}

func mapEntInstructionsToInstructions(insts []*ent.Instruction) []recipe.Instruction {
	var res []recipe.Instruction
	for _, i := range insts {
		res = append(res, entInstructionToInstruction(i))
	}
	return res
}

func mapEntTagsToTags(tags []*ent.Tag) []*recipe.Tag {
	var res []*recipe.Tag
	for _, t := range tags {
		res = append(res, entTagToTag(t))
	}
	return res
}

func mapEntIdeasToIdeas(ideas []*ent.Idea) []*recipe.Idea {
	var res []*recipe.Idea
	for _, i := range ideas {
		res = append(res, entIdeaToIdea(i))
	}
	return res
}

func mapEntSourcesToSources(sources []*ent.Source) []*recipe.Source {
	var res []*recipe.Source
	for _, s := range sources {
		res = append(res, entSourceToSource(s))
	}
	return res
}

func mapEntEquipmentToEquipment(equipment []*ent.Equipment) []*recipe.Equipment {
	var res []*recipe.Equipment
	for _, e := range equipment {
		res = append(res, entEquipmentToEquipment(e))
	}
	return res
}

func mapEntProductsToIngredients(products []*ent.Product) []*recipe.Ingredient {
	var res []*recipe.Ingredient
	for _, p := range products {
		res = append(res, entProductToIngredient(p))
	}
	return res
}
