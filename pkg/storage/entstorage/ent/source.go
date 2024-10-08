// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/recipe"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/source"
)

// Source is the model entity for the Source schema.
type Source struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SourceQuery when eager-loading is set.
	Edges          SourceEdges `json:"edges"`
	recipe_sources *uuid.UUID
	selectValues   sql.SelectValues
}

// SourceEdges holds the relations/edges for other nodes in the graph.
type SourceEdges struct {
	// Recipe holds the value of the recipe edge.
	Recipe *Recipe `json:"recipe,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RecipeOrErr returns the Recipe value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SourceEdges) RecipeOrErr() (*Recipe, error) {
	if e.Recipe != nil {
		return e.Recipe, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: recipe.Label}
	}
	return nil, &NotLoadedError{edge: "recipe"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Source) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case source.FieldName, source.FieldDescription, source.FieldURL:
			values[i] = new(sql.NullString)
		case source.FieldID:
			values[i] = new(uuid.UUID)
		case source.ForeignKeys[0]: // recipe_sources
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Source fields.
func (s *Source) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case source.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case source.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case source.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case source.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				s.URL = value.String
			}
		case source.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field recipe_sources", values[i])
			} else if value.Valid {
				s.recipe_sources = new(uuid.UUID)
				*s.recipe_sources = *value.S.(*uuid.UUID)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Source.
// This includes values selected through modifiers, order, etc.
func (s *Source) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryRecipe queries the "recipe" edge of the Source entity.
func (s *Source) QueryRecipe() *RecipeQuery {
	return NewSourceClient(s.config).QueryRecipe(s)
}

// Update returns a builder for updating this Source.
// Note that you need to call Source.Unwrap() before calling this method if this Source
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Source) Update() *SourceUpdateOne {
	return NewSourceClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Source entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Source) Unwrap() *Source {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Source is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Source) String() string {
	var builder strings.Builder
	builder.WriteString("Source(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(s.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Sources is a parsable slice of Source.
type Sources []*Source
