package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Nutrition holds the schema definition for the Nutrition entity.
type Nutrition struct {
	ent.Schema
}

// Fields of the Nutrition.
func (Nutrition) Fields() []ent.Field {
	return []ent.Field{
		field.Int("calories").Optional(),
		field.Int("fat").Optional(),
		field.Int("carbs").Optional(),
		field.Int("protein").Optional(),
		field.String("precision").Optional(),
		field.Strings("benefits").Optional(),
	}
}

// Edges of the Nutrition.
func (Nutrition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipe", Recipe.Type).Ref("nutrition").Unique(),
	}
}
