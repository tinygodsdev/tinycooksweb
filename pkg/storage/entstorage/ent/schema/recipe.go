package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Recipe holds the schema definition for the Recipe entity.
type Recipe struct {
	ent.Schema
}

// Fields of the Recipe.
func (Recipe) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().Unique(),
		field.String("slug").NotEmpty().Unique(),
		field.String("description"),
		field.String("text"),
		field.Float32("rating").Optional().Positive(),
		field.Int("servings").Optional().Nillable().NonNegative(),
		field.Int64("time").Optional().Nillable().GoType(time.Duration(0)),
	}
}

// Edges of the Recipe.
func (Recipe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("required_products", Product.Type).Through("ingredients", Ingredient.Type),
		edge.To("instructions", Instruction.Type),
		edge.To("tags", Tag.Type),
		edge.To("equipment", Equipment.Type),
		edge.To("ideas", Idea.Type),
		edge.To("sources", Source.Type),
		edge.To("nutrition", Nutrition.Type).Unique(),
	}
}

// Mixin of the Recipe.
func (Recipe) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		LocaleMixin{},
	}
}

func (Recipe) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index for create_time
		index.Fields("create_time"),
	}
}
