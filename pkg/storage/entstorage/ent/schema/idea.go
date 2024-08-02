package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Idea holds the schema definition for the Idea entity.
type Idea struct {
	ent.Schema
}

// Fields of the Idea.
func (Idea) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("text"),
	}
}

// Edges of the Idea.
func (Idea) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipe", Recipe.Type).Ref("ideas").Unique(),
	}
}
