package entstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent"
	entEquipment "github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/equipment"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/product"
	entRecipe "github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/recipe"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/tag"
)

func (s *EntStorage) CountRecipes(ctx context.Context, filter recipe.Filter) (int, error) {
	q := s.client.Recipe.Query()

	if filter.NameContains != "" {
		q.Where(entRecipe.NameContains(filter.NameContains))
	}

	if filter.Locale != "" {
		q.Where(entRecipe.LocaleEQ(entRecipe.Locale(filter.Locale)))
	}

	count, err := q.Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("counting recipes: %w", err)
	}

	return count, nil
}

func (s *EntStorage) GetRecipes(ctx context.Context, filter recipe.Filter) ([]*recipe.Recipe, error) {
	q := s.client.Recipe.Query()

	if filter.NameContains != "" {
		q.Where(entRecipe.NameContains(filter.NameContains))
	}

	if filter.Limit > 0 {
		q.Limit(filter.Limit)
	}

	if filter.Offset > 0 {
		q.Offset(filter.Offset)
	}

	if filter.Locale != "" {
		q.Where(entRecipe.LocaleEQ(entRecipe.Locale(filter.Locale)))
	}

	// TODO implement filtering by equipment, tags, ingredients

	recipes, err := q.
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting recipes: %w", err)
	}

	var res []*recipe.Recipe
	for _, r := range recipes {
		res = append(res, entRecipeToRecipe(r))
	}

	return res, nil
}

func (s *EntStorage) GetRecipe(ctx context.Context, id uuid.UUID) (*recipe.Recipe, error) {
	r, err := s.client.Recipe.
		Query().
		Where(entRecipe.IDEQ(id)).
		WithIngredients(func(q *ent.IngredientQuery) {
			q.WithProduct()
		}).
		WithInstructions().
		WithTags().
		WithIdeas().
		// WithSources().
		WithNutrition().
		WithEquipment().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting recipe by ID: %w", err)
	}

	return entRecipeToRecipe(r), nil
}

func (s *EntStorage) GetRecipeBySlug(ctx context.Context, slug string) (*recipe.Recipe, error) {
	r, err := s.client.Recipe.
		Query().
		Where(entRecipe.SlugEQ(slug)).
		WithIngredients(func(q *ent.IngredientQuery) {
			q.WithProduct()
		}).
		WithInstructions().
		WithTags().
		WithIdeas().
		// WithSources().
		WithNutrition().
		WithEquipment().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting recipe by slug: %w", err)
	}

	return entRecipeToRecipe(r), nil
}

func (s *EntStorage) SaveRecipe(ctx context.Context, rec *recipe.Recipe) error {
	// Check if recipe with the same slug already exists
	existingRecipe, err := s.GetRecipeBySlug(ctx, rec.Slug)
	if err == nil && existingRecipe != nil {
		return storage.ErrAlreadyExists
	}
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("checking existing recipe: %w", err)
	}

	tx, err := s.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	var servings *int
	if rec.Servings > 0 {
		servings = &rec.Servings
	}

	var recipeTime *time.Duration
	if rec.Time > 0 {
		recipeTime = &rec.Time
	}

	// Create recipe
	r, err := tx.Recipe.
		Create().
		SetName(rec.Name).
		SetSlug(rec.Slug).
		SetDescription(rec.Description).
		SetText(rec.Text).
		SetNillableServings(servings).
		SetNillableTime(recipeTime).
		Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("creating recipe: %w", err))
	}

	// Create instructions
	for i, inst := range rec.Instructions {
		_, err := tx.Instruction.
			Create().
			SetText(inst.Text).
			SetOrder(i).
			SetRecipe(r).
			Save(ctx)
		if err != nil {
			return rollback(tx, fmt.Errorf("creating instruction: %w", err))
		}
	}

	// Create or attach products and ingredients
	for _, ing := range rec.Ingredients {
		product, err := getOrCreateProduct(ctx, tx, ing.Product)
		if err != nil {
			return rollback(tx, fmt.Errorf("getting or creating product: %w", err))
		}

		_, err = tx.Ingredient.
			Create().
			SetQuantity(ing.Quantity).
			SetUnit(ing.Unit).
			SetOptional(ing.Optional).
			SetRecipe(r).
			SetProduct(product).
			Save(ctx)
		if err != nil {
			return rollback(tx, fmt.Errorf("creating ingredient: %w", err))
		}
	}

	// Create or attach tags
	for _, tag := range rec.Tags {
		existingTag, err := getOrCreateTag(ctx, tx, tag)
		if err != nil {
			return rollback(tx, fmt.Errorf("getting or creating tag: %w", err))
		}

		err = tx.Recipe.Update().
			Where(entRecipe.IDEQ(r.ID)).
			AddTagIDs(existingTag.ID).
			Exec(ctx)
		if err != nil {
			return rollback(tx, fmt.Errorf("attaching tag: %w", err))
		}
	}

	// Create or attach equipment
	for _, equip := range rec.Equipment {
		existingEquip, err := getOrCreateEquipment(ctx, tx, equip)
		if err != nil {
			return rollback(tx, fmt.Errorf("getting or creating equipment: %w", err))
		}

		err = tx.Recipe.Update().
			Where(entRecipe.IDEQ(r.ID)).
			AddEquipmentIDs(existingEquip.ID).
			Exec(ctx)
		if err != nil {
			return rollback(tx, fmt.Errorf("attaching equipment: %w", err))
		}
	}

	// Create ideas
	for _, idea := range rec.Ideas {
		_, err := tx.Idea.
			Create().
			SetText(idea.Text).
			SetRecipe(r).
			Save(ctx)
		if err != nil {
			return rollback(tx, fmt.Errorf("creating idea: %w", err))
		}
	}

	// Create sources
	for _, source := range rec.Sources {
		_, err := tx.Source.
			Create().
			SetName(source.Name).
			SetDescription(source.Description).
			SetURL(source.URL).
			SetRecipe(r).
			Save(ctx)
		if err != nil {
			return rollback(tx, fmt.Errorf("creating source: %w", err))
		}
	}

	// Create nutrition
	if rec.Nutrition != nil {
		_, err := tx.Nutrition.
			Create().
			SetCalories(rec.Nutrition.Calories).
			SetFat(rec.Nutrition.Fat).
			SetCarbs(rec.Nutrition.Carbs).
			SetProtein(rec.Nutrition.Protein).
			SetPrecision(rec.Nutrition.Precision).
			SetRecipe(r).
			Save(ctx)
		if err != nil {
			return rollback(tx, fmt.Errorf("creating nutrition: %w", err))
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

func getOrCreateProduct(ctx context.Context, tx *ent.Tx, prod *recipe.Product) (*ent.Product, error) {
	p, err := tx.Product.
		Query().
		Where(product.NameEQ(prod.Name)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("querying product: %w", err)
	}
	if p != nil {
		return p, nil
	}

	p, err = tx.Product.
		Create().
		SetName(prod.Name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating product: %w", err)
	}
	return p, nil
}

func getOrCreateTag(ctx context.Context, tx *ent.Tx, tg *recipe.Tag) (*ent.Tag, error) {
	t, err := tx.Tag.
		Query().
		Where(tag.NameEQ(tg.Name), tag.GroupEQ(tg.Group)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("querying tag: %w", err)
	}
	if t != nil {
		return t, nil
	}

	t, err = tx.Tag.
		Create().
		SetName(tg.Name).
		SetGroup(tg.Group).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating tag: %w", err)
	}
	return t, nil
}

func getOrCreateEquipment(ctx context.Context, tx *ent.Tx, equip *recipe.Equipment) (*ent.Equipment, error) {
	e, err := tx.Equipment.
		Query().
		Where(entEquipment.NameEQ(equip.Name)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("querying equipment: %w", err)
	}
	if e != nil {
		return e, nil
	}

	e, err = tx.Equipment.
		Create().
		SetName(equip.Name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating equipment: %w", err)
	}
	return e, nil
}
