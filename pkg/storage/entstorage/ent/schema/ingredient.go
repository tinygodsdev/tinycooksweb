package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Ingredient holds the schema definition for the Ingredient entity.
type Ingredient struct {
	ent.Schema
}

// Fields of the Ingredient.
func (Ingredient) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("quantity").Optional(),
		field.String("unit").Optional(),
		field.UUID("recipe_id", uuid.UUID{}),
		field.UUID("product_id", uuid.UUID{}),
		field.Bool("optional").Default(false),
	}
}

// Edges of the Ingredient.
func (Ingredient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("recipe", Recipe.Type).Unique().Required().Field("recipe_id"),
		edge.To("product", Product.Type).Unique().Required().Field("product_id"),
	}
}

// Mixin of the Ingredient.
func (Ingredient) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
