package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Source holds the schema definition for the Source entity.
type Source struct {
	ent.Schema
}

// Fields of the Source.
func (Source) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Optional(),
		field.String("description").Optional(),
		field.String("url").Optional(),
	}
}

// Edges of the Source.
func (Source) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipe", Recipe.Type).Ref("sources").Unique(),
	}
}
