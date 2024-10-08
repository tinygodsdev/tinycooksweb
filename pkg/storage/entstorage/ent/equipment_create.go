// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/equipment"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/recipe"
)

// EquipmentCreate is the builder for creating a Equipment entity.
type EquipmentCreate struct {
	config
	mutation *EquipmentMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ec *EquipmentCreate) SetName(s string) *EquipmentCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetSlug sets the "slug" field.
func (ec *EquipmentCreate) SetSlug(s string) *EquipmentCreate {
	ec.mutation.SetSlug(s)
	return ec
}

// SetID sets the "id" field.
func (ec *EquipmentCreate) SetID(u uuid.UUID) *EquipmentCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EquipmentCreate) SetNillableID(u *uuid.UUID) *EquipmentCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// AddRecipeIDs adds the "recipes" edge to the Recipe entity by IDs.
func (ec *EquipmentCreate) AddRecipeIDs(ids ...uuid.UUID) *EquipmentCreate {
	ec.mutation.AddRecipeIDs(ids...)
	return ec
}

// AddRecipes adds the "recipes" edges to the Recipe entity.
func (ec *EquipmentCreate) AddRecipes(r ...*Recipe) *EquipmentCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddRecipeIDs(ids...)
}

// Mutation returns the EquipmentMutation object of the builder.
func (ec *EquipmentCreate) Mutation() *EquipmentMutation {
	return ec.mutation
}

// Save creates the Equipment in the database.
func (ec *EquipmentCreate) Save(ctx context.Context) (*Equipment, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EquipmentCreate) SaveX(ctx context.Context) *Equipment {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EquipmentCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EquipmentCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EquipmentCreate) defaults() {
	if _, ok := ec.mutation.ID(); !ok {
		v := equipment.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EquipmentCreate) check() error {
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Equipment.name"`)}
	}
	if v, ok := ec.mutation.Name(); ok {
		if err := equipment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Equipment.name": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Equipment.slug"`)}
	}
	if v, ok := ec.mutation.Slug(); ok {
		if err := equipment.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "Equipment.slug": %w`, err)}
		}
	}
	return nil
}

func (ec *EquipmentCreate) sqlSave(ctx context.Context) (*Equipment, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EquipmentCreate) createSpec() (*Equipment, *sqlgraph.CreateSpec) {
	var (
		_node = &Equipment{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(equipment.Table, sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeUUID))
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.SetField(equipment.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ec.mutation.Slug(); ok {
		_spec.SetField(equipment.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if nodes := ec.mutation.RecipesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipment.RecipesTable,
			Columns: equipment.RecipesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EquipmentCreateBulk is the builder for creating many Equipment entities in bulk.
type EquipmentCreateBulk struct {
	config
	err      error
	builders []*EquipmentCreate
}

// Save creates the Equipment entities in the database.
func (ecb *EquipmentCreateBulk) Save(ctx context.Context) ([]*Equipment, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Equipment, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EquipmentCreateBulk) SaveX(ctx context.Context) []*Equipment {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EquipmentCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EquipmentCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
