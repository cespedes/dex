// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dexidp/dex/storage/ent/db/connector"
	"github.com/dexidp/dex/storage/ent/db/predicate"
)

// ConnectorDelete is the builder for deleting a Connector entity.
type ConnectorDelete struct {
	config
	hooks    []Hook
	mutation *ConnectorMutation
}

// Where appends a list predicates to the ConnectorDelete builder.
func (cd *ConnectorDelete) Where(ps ...predicate.Connector) *ConnectorDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ConnectorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ConnectorMutation](ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ConnectorDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *ConnectorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: connector.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: connector.FieldID,
			},
		},
	}
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// ConnectorDeleteOne is the builder for deleting a single Connector entity.
type ConnectorDeleteOne struct {
	cd *ConnectorDelete
}

// Exec executes the deletion query.
func (cdo *ConnectorDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{connector.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ConnectorDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
