// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/idea"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent/predicate"
)

// IdeaDelete is the builder for deleting a Idea entity.
type IdeaDelete struct {
	config
	hooks    []Hook
	mutation *IdeaMutation
}

// Where appends a list predicates to the IdeaDelete builder.
func (id *IdeaDelete) Where(ps ...predicate.Idea) *IdeaDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *IdeaDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *IdeaDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *IdeaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(idea.Table, sqlgraph.NewFieldSpec(idea.FieldID, field.TypeUUID))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// IdeaDeleteOne is the builder for deleting a single Idea entity.
type IdeaDeleteOne struct {
	id *IdeaDelete
}

// Where appends a list predicates to the IdeaDelete builder.
func (ido *IdeaDeleteOne) Where(ps ...predicate.Idea) *IdeaDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *IdeaDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{idea.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *IdeaDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}