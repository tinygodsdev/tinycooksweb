// Code generated by ent, DO NOT EDIT.

package recipe

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the recipe type in the database.
	Label = "recipe"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldServings holds the string denoting the servings field in the database.
	FieldServings = "servings"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// EdgeRequiredProducts holds the string denoting the required_products edge name in mutations.
	EdgeRequiredProducts = "required_products"
	// EdgeInstructions holds the string denoting the instructions edge name in mutations.
	EdgeInstructions = "instructions"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// EdgeIdeas holds the string denoting the ideas edge name in mutations.
	EdgeIdeas = "ideas"
	// EdgeSources holds the string denoting the sources edge name in mutations.
	EdgeSources = "sources"
	// EdgeNutrition holds the string denoting the nutrition edge name in mutations.
	EdgeNutrition = "nutrition"
	// EdgeIngredients holds the string denoting the ingredients edge name in mutations.
	EdgeIngredients = "ingredients"
	// Table holds the table name of the recipe in the database.
	Table = "recipes"
	// RequiredProductsTable is the table that holds the required_products relation/edge. The primary key declared below.
	RequiredProductsTable = "ingredients"
	// RequiredProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	RequiredProductsInverseTable = "products"
	// InstructionsTable is the table that holds the instructions relation/edge.
	InstructionsTable = "instructions"
	// InstructionsInverseTable is the table name for the Instruction entity.
	// It exists in this package in order to avoid circular dependency with the "instruction" package.
	InstructionsInverseTable = "instructions"
	// InstructionsColumn is the table column denoting the instructions relation/edge.
	InstructionsColumn = "recipe_instructions"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "recipe_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// EquipmentTable is the table that holds the equipment relation/edge. The primary key declared below.
	EquipmentTable = "recipe_equipment"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
	// IdeasTable is the table that holds the ideas relation/edge.
	IdeasTable = "ideas"
	// IdeasInverseTable is the table name for the Idea entity.
	// It exists in this package in order to avoid circular dependency with the "idea" package.
	IdeasInverseTable = "ideas"
	// IdeasColumn is the table column denoting the ideas relation/edge.
	IdeasColumn = "recipe_ideas"
	// SourcesTable is the table that holds the sources relation/edge.
	SourcesTable = "sources"
	// SourcesInverseTable is the table name for the Source entity.
	// It exists in this package in order to avoid circular dependency with the "source" package.
	SourcesInverseTable = "sources"
	// SourcesColumn is the table column denoting the sources relation/edge.
	SourcesColumn = "recipe_sources"
	// NutritionTable is the table that holds the nutrition relation/edge.
	NutritionTable = "nutritions"
	// NutritionInverseTable is the table name for the Nutrition entity.
	// It exists in this package in order to avoid circular dependency with the "nutrition" package.
	NutritionInverseTable = "nutritions"
	// NutritionColumn is the table column denoting the nutrition relation/edge.
	NutritionColumn = "recipe_nutrition"
	// IngredientsTable is the table that holds the ingredients relation/edge.
	IngredientsTable = "ingredients"
	// IngredientsInverseTable is the table name for the Ingredient entity.
	// It exists in this package in order to avoid circular dependency with the "ingredient" package.
	IngredientsInverseTable = "ingredients"
	// IngredientsColumn is the table column denoting the ingredients relation/edge.
	IngredientsColumn = "recipe_id"
)

// Columns holds all SQL columns for recipe fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldLocale,
	FieldName,
	FieldSlug,
	FieldDescription,
	FieldText,
	FieldServings,
	FieldTime,
}

var (
	// RequiredProductsPrimaryKey and RequiredProductsColumn2 are the table columns denoting the
	// primary key for the required_products relation (M2M).
	RequiredProductsPrimaryKey = []string{"recipe_id", "product_id"}
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"recipe_id", "tag_id"}
	// EquipmentPrimaryKey and EquipmentColumn2 are the table columns denoting the
	// primary key for the equipment relation (M2M).
	EquipmentPrimaryKey = []string{"recipe_id", "equipment_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// ServingsValidator is a validator for the "servings" field. It is called by the builders before save.
	ServingsValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Locale defines the type for the "locale" enum field.
type Locale string

// LocaleRu is the default value of the Locale enum.
const DefaultLocale = LocaleRu

// Locale values.
const (
	LocaleEn Locale = "en"
	LocaleRu Locale = "ru"
)

func (l Locale) String() string {
	return string(l)
}

// LocaleValidator is a validator for the "locale" field enum values. It is called by the builders before save.
func LocaleValidator(l Locale) error {
	switch l {
	case LocaleEn, LocaleRu:
		return nil
	default:
		return fmt.Errorf("recipe: invalid enum value for locale field: %q", l)
	}
}

// OrderOption defines the ordering options for the Recipe queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByLocale orders the results by the locale field.
func ByLocale(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocale, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySlug orders the results by the slug field.
func BySlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSlug, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByServings orders the results by the servings field.
func ByServings(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldServings, opts...).ToFunc()
}

// ByTime orders the results by the time field.
func ByTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTime, opts...).ToFunc()
}

// ByRequiredProductsCount orders the results by required_products count.
func ByRequiredProductsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRequiredProductsStep(), opts...)
	}
}

// ByRequiredProducts orders the results by required_products terms.
func ByRequiredProducts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRequiredProductsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInstructionsCount orders the results by instructions count.
func ByInstructionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInstructionsStep(), opts...)
	}
}

// ByInstructions orders the results by instructions terms.
func ByInstructions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInstructionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEquipmentCount orders the results by equipment count.
func ByEquipmentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEquipmentStep(), opts...)
	}
}

// ByEquipment orders the results by equipment terms.
func ByEquipment(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEquipmentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByIdeasCount orders the results by ideas count.
func ByIdeasCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIdeasStep(), opts...)
	}
}

// ByIdeas orders the results by ideas terms.
func ByIdeas(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIdeasStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySourcesCount orders the results by sources count.
func BySourcesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSourcesStep(), opts...)
	}
}

// BySources orders the results by sources terms.
func BySources(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSourcesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNutritionField orders the results by nutrition field.
func ByNutritionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNutritionStep(), sql.OrderByField(field, opts...))
	}
}

// ByIngredientsCount orders the results by ingredients count.
func ByIngredientsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIngredientsStep(), opts...)
	}
}

// ByIngredients orders the results by ingredients terms.
func ByIngredients(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIngredientsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRequiredProductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RequiredProductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, RequiredProductsTable, RequiredProductsPrimaryKey...),
	)
}
func newInstructionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InstructionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InstructionsTable, InstructionsColumn),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
	)
}
func newEquipmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EquipmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, EquipmentTable, EquipmentPrimaryKey...),
	)
}
func newIdeasStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IdeasInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IdeasTable, IdeasColumn),
	)
}
func newSourcesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SourcesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SourcesTable, SourcesColumn),
	)
}
func newNutritionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NutritionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, NutritionTable, NutritionColumn),
	)
}
func newIngredientsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IngredientsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, IngredientsTable, IngredientsColumn),
	)
}