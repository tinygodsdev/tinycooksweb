// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/ingredient"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/predicate"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/product"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/recipe"
)

// IngredientUpdate is the builder for updating Ingredient entities.
type IngredientUpdate struct {
	config
	hooks    []Hook
	mutation *IngredientMutation
}

// Where appends a list predicates to the IngredientUpdate builder.
func (iu *IngredientUpdate) Where(ps ...predicate.Ingredient) *IngredientUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdateTime sets the "update_time" field.
func (iu *IngredientUpdate) SetUpdateTime(t time.Time) *IngredientUpdate {
	iu.mutation.SetUpdateTime(t)
	return iu
}

// SetQuantity sets the "quantity" field.
func (iu *IngredientUpdate) SetQuantity(s string) *IngredientUpdate {
	iu.mutation.SetQuantity(s)
	return iu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (iu *IngredientUpdate) SetNillableQuantity(s *string) *IngredientUpdate {
	if s != nil {
		iu.SetQuantity(*s)
	}
	return iu
}

// ClearQuantity clears the value of the "quantity" field.
func (iu *IngredientUpdate) ClearQuantity() *IngredientUpdate {
	iu.mutation.ClearQuantity()
	return iu
}

// SetUnit sets the "unit" field.
func (iu *IngredientUpdate) SetUnit(s string) *IngredientUpdate {
	iu.mutation.SetUnit(s)
	return iu
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (iu *IngredientUpdate) SetNillableUnit(s *string) *IngredientUpdate {
	if s != nil {
		iu.SetUnit(*s)
	}
	return iu
}

// ClearUnit clears the value of the "unit" field.
func (iu *IngredientUpdate) ClearUnit() *IngredientUpdate {
	iu.mutation.ClearUnit()
	return iu
}

// SetRecipeID sets the "recipe_id" field.
func (iu *IngredientUpdate) SetRecipeID(u uuid.UUID) *IngredientUpdate {
	iu.mutation.SetRecipeID(u)
	return iu
}

// SetNillableRecipeID sets the "recipe_id" field if the given value is not nil.
func (iu *IngredientUpdate) SetNillableRecipeID(u *uuid.UUID) *IngredientUpdate {
	if u != nil {
		iu.SetRecipeID(*u)
	}
	return iu
}

// SetProductID sets the "product_id" field.
func (iu *IngredientUpdate) SetProductID(u uuid.UUID) *IngredientUpdate {
	iu.mutation.SetProductID(u)
	return iu
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (iu *IngredientUpdate) SetNillableProductID(u *uuid.UUID) *IngredientUpdate {
	if u != nil {
		iu.SetProductID(*u)
	}
	return iu
}

// SetRecipe sets the "recipe" edge to the Recipe entity.
func (iu *IngredientUpdate) SetRecipe(r *Recipe) *IngredientUpdate {
	return iu.SetRecipeID(r.ID)
}

// SetProduct sets the "product" edge to the Product entity.
func (iu *IngredientUpdate) SetProduct(p *Product) *IngredientUpdate {
	return iu.SetProductID(p.ID)
}

// Mutation returns the IngredientMutation object of the builder.
func (iu *IngredientUpdate) Mutation() *IngredientMutation {
	return iu.mutation
}

// ClearRecipe clears the "recipe" edge to the Recipe entity.
func (iu *IngredientUpdate) ClearRecipe() *IngredientUpdate {
	iu.mutation.ClearRecipe()
	return iu
}

// ClearProduct clears the "product" edge to the Product entity.
func (iu *IngredientUpdate) ClearProduct() *IngredientUpdate {
	iu.mutation.ClearProduct()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IngredientUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IngredientUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IngredientUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IngredientUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *IngredientUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := ingredient.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IngredientUpdate) check() error {
	if iu.mutation.RecipeCleared() && len(iu.mutation.RecipeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Ingredient.recipe"`)
	}
	if iu.mutation.ProductCleared() && len(iu.mutation.ProductIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Ingredient.product"`)
	}
	return nil
}

func (iu *IngredientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(ingredient.Table, ingredient.Columns, sqlgraph.NewFieldSpec(ingredient.FieldID, field.TypeUUID))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.SetField(ingredient.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := iu.mutation.Quantity(); ok {
		_spec.SetField(ingredient.FieldQuantity, field.TypeString, value)
	}
	if iu.mutation.QuantityCleared() {
		_spec.ClearField(ingredient.FieldQuantity, field.TypeString)
	}
	if value, ok := iu.mutation.Unit(); ok {
		_spec.SetField(ingredient.FieldUnit, field.TypeString, value)
	}
	if iu.mutation.UnitCleared() {
		_spec.ClearField(ingredient.FieldUnit, field.TypeString)
	}
	if iu.mutation.RecipeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ingredient.RecipeTable,
			Columns: []string{ingredient.RecipeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RecipeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ingredient.RecipeTable,
			Columns: []string{ingredient.RecipeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ingredient.ProductTable,
			Columns: []string{ingredient.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ingredient.ProductTable,
			Columns: []string{ingredient.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ingredient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// IngredientUpdateOne is the builder for updating a single Ingredient entity.
type IngredientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IngredientMutation
}

// SetUpdateTime sets the "update_time" field.
func (iuo *IngredientUpdateOne) SetUpdateTime(t time.Time) *IngredientUpdateOne {
	iuo.mutation.SetUpdateTime(t)
	return iuo
}

// SetQuantity sets the "quantity" field.
func (iuo *IngredientUpdateOne) SetQuantity(s string) *IngredientUpdateOne {
	iuo.mutation.SetQuantity(s)
	return iuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (iuo *IngredientUpdateOne) SetNillableQuantity(s *string) *IngredientUpdateOne {
	if s != nil {
		iuo.SetQuantity(*s)
	}
	return iuo
}

// ClearQuantity clears the value of the "quantity" field.
func (iuo *IngredientUpdateOne) ClearQuantity() *IngredientUpdateOne {
	iuo.mutation.ClearQuantity()
	return iuo
}

// SetUnit sets the "unit" field.
func (iuo *IngredientUpdateOne) SetUnit(s string) *IngredientUpdateOne {
	iuo.mutation.SetUnit(s)
	return iuo
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (iuo *IngredientUpdateOne) SetNillableUnit(s *string) *IngredientUpdateOne {
	if s != nil {
		iuo.SetUnit(*s)
	}
	return iuo
}

// ClearUnit clears the value of the "unit" field.
func (iuo *IngredientUpdateOne) ClearUnit() *IngredientUpdateOne {
	iuo.mutation.ClearUnit()
	return iuo
}

// SetRecipeID sets the "recipe_id" field.
func (iuo *IngredientUpdateOne) SetRecipeID(u uuid.UUID) *IngredientUpdateOne {
	iuo.mutation.SetRecipeID(u)
	return iuo
}

// SetNillableRecipeID sets the "recipe_id" field if the given value is not nil.
func (iuo *IngredientUpdateOne) SetNillableRecipeID(u *uuid.UUID) *IngredientUpdateOne {
	if u != nil {
		iuo.SetRecipeID(*u)
	}
	return iuo
}

// SetProductID sets the "product_id" field.
func (iuo *IngredientUpdateOne) SetProductID(u uuid.UUID) *IngredientUpdateOne {
	iuo.mutation.SetProductID(u)
	return iuo
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (iuo *IngredientUpdateOne) SetNillableProductID(u *uuid.UUID) *IngredientUpdateOne {
	if u != nil {
		iuo.SetProductID(*u)
	}
	return iuo
}

// SetRecipe sets the "recipe" edge to the Recipe entity.
func (iuo *IngredientUpdateOne) SetRecipe(r *Recipe) *IngredientUpdateOne {
	return iuo.SetRecipeID(r.ID)
}

// SetProduct sets the "product" edge to the Product entity.
func (iuo *IngredientUpdateOne) SetProduct(p *Product) *IngredientUpdateOne {
	return iuo.SetProductID(p.ID)
}

// Mutation returns the IngredientMutation object of the builder.
func (iuo *IngredientUpdateOne) Mutation() *IngredientMutation {
	return iuo.mutation
}

// ClearRecipe clears the "recipe" edge to the Recipe entity.
func (iuo *IngredientUpdateOne) ClearRecipe() *IngredientUpdateOne {
	iuo.mutation.ClearRecipe()
	return iuo
}

// ClearProduct clears the "product" edge to the Product entity.
func (iuo *IngredientUpdateOne) ClearProduct() *IngredientUpdateOne {
	iuo.mutation.ClearProduct()
	return iuo
}

// Where appends a list predicates to the IngredientUpdate builder.
func (iuo *IngredientUpdateOne) Where(ps ...predicate.Ingredient) *IngredientUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IngredientUpdateOne) Select(field string, fields ...string) *IngredientUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Ingredient entity.
func (iuo *IngredientUpdateOne) Save(ctx context.Context) (*Ingredient, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IngredientUpdateOne) SaveX(ctx context.Context) *Ingredient {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IngredientUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IngredientUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *IngredientUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := ingredient.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IngredientUpdateOne) check() error {
	if iuo.mutation.RecipeCleared() && len(iuo.mutation.RecipeIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Ingredient.recipe"`)
	}
	if iuo.mutation.ProductCleared() && len(iuo.mutation.ProductIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Ingredient.product"`)
	}
	return nil
}

func (iuo *IngredientUpdateOne) sqlSave(ctx context.Context) (_node *Ingredient, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(ingredient.Table, ingredient.Columns, sqlgraph.NewFieldSpec(ingredient.FieldID, field.TypeUUID))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ingredient.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ingredient.FieldID)
		for _, f := range fields {
			if !ingredient.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ingredient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.SetField(ingredient.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.Quantity(); ok {
		_spec.SetField(ingredient.FieldQuantity, field.TypeString, value)
	}
	if iuo.mutation.QuantityCleared() {
		_spec.ClearField(ingredient.FieldQuantity, field.TypeString)
	}
	if value, ok := iuo.mutation.Unit(); ok {
		_spec.SetField(ingredient.FieldUnit, field.TypeString, value)
	}
	if iuo.mutation.UnitCleared() {
		_spec.ClearField(ingredient.FieldUnit, field.TypeString)
	}
	if iuo.mutation.RecipeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ingredient.RecipeTable,
			Columns: []string{ingredient.RecipeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RecipeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ingredient.RecipeTable,
			Columns: []string{ingredient.RecipeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ingredient.ProductTable,
			Columns: []string{ingredient.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ingredient.ProductTable,
			Columns: []string{ingredient.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Ingredient{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ingredient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}