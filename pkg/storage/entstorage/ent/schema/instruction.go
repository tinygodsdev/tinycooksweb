package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Instruction holds the schema definition for the Instruction entity.
type Instruction struct {
	ent.Schema
}

// Fields of the Instruction.
func (Instruction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("text"),
		field.Int("order").NonNegative(),
	}
}

// Edges of the Instruction.
func (Instruction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recipe", Recipe.Type).Ref("instructions").Unique(),
	}
}

// Indexes of the Instruction.
func (Instruction) Indexes() []ent.Index {
	return []ent.Index{
		// unique index for text and order
		index.Edges("recipe").Fields("order").Unique(),
	}
}
