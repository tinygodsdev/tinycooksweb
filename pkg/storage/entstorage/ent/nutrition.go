// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/nutrition"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/recipe"
)

// Nutrition is the model entity for the Nutrition schema.
type Nutrition struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Calories holds the value of the "calories" field.
	Calories int `json:"calories,omitempty"`
	// Fat holds the value of the "fat" field.
	Fat int `json:"fat,omitempty"`
	// Carbs holds the value of the "carbs" field.
	Carbs int `json:"carbs,omitempty"`
	// Protein holds the value of the "protein" field.
	Protein int `json:"protein,omitempty"`
	// Precision holds the value of the "precision" field.
	Precision string `json:"precision,omitempty"`
	// Benefits holds the value of the "benefits" field.
	Benefits []string `json:"benefits,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NutritionQuery when eager-loading is set.
	Edges            NutritionEdges `json:"edges"`
	recipe_nutrition *uuid.UUID
	selectValues     sql.SelectValues
}

// NutritionEdges holds the relations/edges for other nodes in the graph.
type NutritionEdges struct {
	// Recipe holds the value of the recipe edge.
	Recipe *Recipe `json:"recipe,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RecipeOrErr returns the Recipe value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NutritionEdges) RecipeOrErr() (*Recipe, error) {
	if e.Recipe != nil {
		return e.Recipe, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: recipe.Label}
	}
	return nil, &NotLoadedError{edge: "recipe"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Nutrition) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case nutrition.FieldBenefits:
			values[i] = new([]byte)
		case nutrition.FieldID, nutrition.FieldCalories, nutrition.FieldFat, nutrition.FieldCarbs, nutrition.FieldProtein:
			values[i] = new(sql.NullInt64)
		case nutrition.FieldPrecision:
			values[i] = new(sql.NullString)
		case nutrition.ForeignKeys[0]: // recipe_nutrition
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Nutrition fields.
func (n *Nutrition) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case nutrition.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case nutrition.FieldCalories:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field calories", values[i])
			} else if value.Valid {
				n.Calories = int(value.Int64)
			}
		case nutrition.FieldFat:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fat", values[i])
			} else if value.Valid {
				n.Fat = int(value.Int64)
			}
		case nutrition.FieldCarbs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field carbs", values[i])
			} else if value.Valid {
				n.Carbs = int(value.Int64)
			}
		case nutrition.FieldProtein:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field protein", values[i])
			} else if value.Valid {
				n.Protein = int(value.Int64)
			}
		case nutrition.FieldPrecision:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field precision", values[i])
			} else if value.Valid {
				n.Precision = value.String
			}
		case nutrition.FieldBenefits:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field benefits", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.Benefits); err != nil {
					return fmt.Errorf("unmarshal field benefits: %w", err)
				}
			}
		case nutrition.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field recipe_nutrition", values[i])
			} else if value.Valid {
				n.recipe_nutrition = new(uuid.UUID)
				*n.recipe_nutrition = *value.S.(*uuid.UUID)
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Nutrition.
// This includes values selected through modifiers, order, etc.
func (n *Nutrition) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryRecipe queries the "recipe" edge of the Nutrition entity.
func (n *Nutrition) QueryRecipe() *RecipeQuery {
	return NewNutritionClient(n.config).QueryRecipe(n)
}

// Update returns a builder for updating this Nutrition.
// Note that you need to call Nutrition.Unwrap() before calling this method if this Nutrition
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Nutrition) Update() *NutritionUpdateOne {
	return NewNutritionClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Nutrition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Nutrition) Unwrap() *Nutrition {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Nutrition is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Nutrition) String() string {
	var builder strings.Builder
	builder.WriteString("Nutrition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("calories=")
	builder.WriteString(fmt.Sprintf("%v", n.Calories))
	builder.WriteString(", ")
	builder.WriteString("fat=")
	builder.WriteString(fmt.Sprintf("%v", n.Fat))
	builder.WriteString(", ")
	builder.WriteString("carbs=")
	builder.WriteString(fmt.Sprintf("%v", n.Carbs))
	builder.WriteString(", ")
	builder.WriteString("protein=")
	builder.WriteString(fmt.Sprintf("%v", n.Protein))
	builder.WriteString(", ")
	builder.WriteString("precision=")
	builder.WriteString(n.Precision)
	builder.WriteString(", ")
	builder.WriteString("benefits=")
	builder.WriteString(fmt.Sprintf("%v", n.Benefits))
	builder.WriteByte(')')
	return builder.String()
}

// Nutritions is a parsable slice of Nutrition.
type Nutritions []*Nutrition