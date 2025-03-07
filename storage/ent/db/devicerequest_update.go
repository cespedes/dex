// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/dexidp/dex/storage/ent/db/devicerequest"
	"github.com/dexidp/dex/storage/ent/db/predicate"
)

// DeviceRequestUpdate is the builder for updating DeviceRequest entities.
type DeviceRequestUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceRequestMutation
}

// Where appends a list predicates to the DeviceRequestUpdate builder.
func (dru *DeviceRequestUpdate) Where(ps ...predicate.DeviceRequest) *DeviceRequestUpdate {
	dru.mutation.Where(ps...)
	return dru
}

// SetUserCode sets the "user_code" field.
func (dru *DeviceRequestUpdate) SetUserCode(s string) *DeviceRequestUpdate {
	dru.mutation.SetUserCode(s)
	return dru
}

// SetDeviceCode sets the "device_code" field.
func (dru *DeviceRequestUpdate) SetDeviceCode(s string) *DeviceRequestUpdate {
	dru.mutation.SetDeviceCode(s)
	return dru
}

// SetClientID sets the "client_id" field.
func (dru *DeviceRequestUpdate) SetClientID(s string) *DeviceRequestUpdate {
	dru.mutation.SetClientID(s)
	return dru
}

// SetClientSecret sets the "client_secret" field.
func (dru *DeviceRequestUpdate) SetClientSecret(s string) *DeviceRequestUpdate {
	dru.mutation.SetClientSecret(s)
	return dru
}

// SetScopes sets the "scopes" field.
func (dru *DeviceRequestUpdate) SetScopes(s []string) *DeviceRequestUpdate {
	dru.mutation.SetScopes(s)
	return dru
}

// AppendScopes appends s to the "scopes" field.
func (dru *DeviceRequestUpdate) AppendScopes(s []string) *DeviceRequestUpdate {
	dru.mutation.AppendScopes(s)
	return dru
}

// ClearScopes clears the value of the "scopes" field.
func (dru *DeviceRequestUpdate) ClearScopes() *DeviceRequestUpdate {
	dru.mutation.ClearScopes()
	return dru
}

// SetExpiry sets the "expiry" field.
func (dru *DeviceRequestUpdate) SetExpiry(t time.Time) *DeviceRequestUpdate {
	dru.mutation.SetExpiry(t)
	return dru
}

// Mutation returns the DeviceRequestMutation object of the builder.
func (dru *DeviceRequestUpdate) Mutation() *DeviceRequestMutation {
	return dru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dru *DeviceRequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, DeviceRequestMutation](ctx, dru.sqlSave, dru.mutation, dru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dru *DeviceRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := dru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dru *DeviceRequestUpdate) Exec(ctx context.Context) error {
	_, err := dru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dru *DeviceRequestUpdate) ExecX(ctx context.Context) {
	if err := dru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dru *DeviceRequestUpdate) check() error {
	if v, ok := dru.mutation.UserCode(); ok {
		if err := devicerequest.UserCodeValidator(v); err != nil {
			return &ValidationError{Name: "user_code", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.user_code": %w`, err)}
		}
	}
	if v, ok := dru.mutation.DeviceCode(); ok {
		if err := devicerequest.DeviceCodeValidator(v); err != nil {
			return &ValidationError{Name: "device_code", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.device_code": %w`, err)}
		}
	}
	if v, ok := dru.mutation.ClientID(); ok {
		if err := devicerequest.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.client_id": %w`, err)}
		}
	}
	if v, ok := dru.mutation.ClientSecret(); ok {
		if err := devicerequest.ClientSecretValidator(v); err != nil {
			return &ValidationError{Name: "client_secret", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.client_secret": %w`, err)}
		}
	}
	return nil
}

func (dru *DeviceRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := dru.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   devicerequest.Table,
			Columns: devicerequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: devicerequest.FieldID,
			},
		},
	}
	if ps := dru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dru.mutation.UserCode(); ok {
		_spec.SetField(devicerequest.FieldUserCode, field.TypeString, value)
	}
	if value, ok := dru.mutation.DeviceCode(); ok {
		_spec.SetField(devicerequest.FieldDeviceCode, field.TypeString, value)
	}
	if value, ok := dru.mutation.ClientID(); ok {
		_spec.SetField(devicerequest.FieldClientID, field.TypeString, value)
	}
	if value, ok := dru.mutation.ClientSecret(); ok {
		_spec.SetField(devicerequest.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := dru.mutation.Scopes(); ok {
		_spec.SetField(devicerequest.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := dru.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, devicerequest.FieldScopes, value)
		})
	}
	if dru.mutation.ScopesCleared() {
		_spec.ClearField(devicerequest.FieldScopes, field.TypeJSON)
	}
	if value, ok := dru.mutation.Expiry(); ok {
		_spec.SetField(devicerequest.FieldExpiry, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicerequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dru.mutation.done = true
	return n, nil
}

// DeviceRequestUpdateOne is the builder for updating a single DeviceRequest entity.
type DeviceRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceRequestMutation
}

// SetUserCode sets the "user_code" field.
func (druo *DeviceRequestUpdateOne) SetUserCode(s string) *DeviceRequestUpdateOne {
	druo.mutation.SetUserCode(s)
	return druo
}

// SetDeviceCode sets the "device_code" field.
func (druo *DeviceRequestUpdateOne) SetDeviceCode(s string) *DeviceRequestUpdateOne {
	druo.mutation.SetDeviceCode(s)
	return druo
}

// SetClientID sets the "client_id" field.
func (druo *DeviceRequestUpdateOne) SetClientID(s string) *DeviceRequestUpdateOne {
	druo.mutation.SetClientID(s)
	return druo
}

// SetClientSecret sets the "client_secret" field.
func (druo *DeviceRequestUpdateOne) SetClientSecret(s string) *DeviceRequestUpdateOne {
	druo.mutation.SetClientSecret(s)
	return druo
}

// SetScopes sets the "scopes" field.
func (druo *DeviceRequestUpdateOne) SetScopes(s []string) *DeviceRequestUpdateOne {
	druo.mutation.SetScopes(s)
	return druo
}

// AppendScopes appends s to the "scopes" field.
func (druo *DeviceRequestUpdateOne) AppendScopes(s []string) *DeviceRequestUpdateOne {
	druo.mutation.AppendScopes(s)
	return druo
}

// ClearScopes clears the value of the "scopes" field.
func (druo *DeviceRequestUpdateOne) ClearScopes() *DeviceRequestUpdateOne {
	druo.mutation.ClearScopes()
	return druo
}

// SetExpiry sets the "expiry" field.
func (druo *DeviceRequestUpdateOne) SetExpiry(t time.Time) *DeviceRequestUpdateOne {
	druo.mutation.SetExpiry(t)
	return druo
}

// Mutation returns the DeviceRequestMutation object of the builder.
func (druo *DeviceRequestUpdateOne) Mutation() *DeviceRequestMutation {
	return druo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (druo *DeviceRequestUpdateOne) Select(field string, fields ...string) *DeviceRequestUpdateOne {
	druo.fields = append([]string{field}, fields...)
	return druo
}

// Save executes the query and returns the updated DeviceRequest entity.
func (druo *DeviceRequestUpdateOne) Save(ctx context.Context) (*DeviceRequest, error) {
	return withHooks[*DeviceRequest, DeviceRequestMutation](ctx, druo.sqlSave, druo.mutation, druo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (druo *DeviceRequestUpdateOne) SaveX(ctx context.Context) *DeviceRequest {
	node, err := druo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (druo *DeviceRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := druo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (druo *DeviceRequestUpdateOne) ExecX(ctx context.Context) {
	if err := druo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (druo *DeviceRequestUpdateOne) check() error {
	if v, ok := druo.mutation.UserCode(); ok {
		if err := devicerequest.UserCodeValidator(v); err != nil {
			return &ValidationError{Name: "user_code", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.user_code": %w`, err)}
		}
	}
	if v, ok := druo.mutation.DeviceCode(); ok {
		if err := devicerequest.DeviceCodeValidator(v); err != nil {
			return &ValidationError{Name: "device_code", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.device_code": %w`, err)}
		}
	}
	if v, ok := druo.mutation.ClientID(); ok {
		if err := devicerequest.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.client_id": %w`, err)}
		}
	}
	if v, ok := druo.mutation.ClientSecret(); ok {
		if err := devicerequest.ClientSecretValidator(v); err != nil {
			return &ValidationError{Name: "client_secret", err: fmt.Errorf(`db: validator failed for field "DeviceRequest.client_secret": %w`, err)}
		}
	}
	return nil
}

func (druo *DeviceRequestUpdateOne) sqlSave(ctx context.Context) (_node *DeviceRequest, err error) {
	if err := druo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   devicerequest.Table,
			Columns: devicerequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: devicerequest.FieldID,
			},
		},
	}
	id, ok := druo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "DeviceRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := druo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicerequest.FieldID)
		for _, f := range fields {
			if !devicerequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != devicerequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := druo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := druo.mutation.UserCode(); ok {
		_spec.SetField(devicerequest.FieldUserCode, field.TypeString, value)
	}
	if value, ok := druo.mutation.DeviceCode(); ok {
		_spec.SetField(devicerequest.FieldDeviceCode, field.TypeString, value)
	}
	if value, ok := druo.mutation.ClientID(); ok {
		_spec.SetField(devicerequest.FieldClientID, field.TypeString, value)
	}
	if value, ok := druo.mutation.ClientSecret(); ok {
		_spec.SetField(devicerequest.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := druo.mutation.Scopes(); ok {
		_spec.SetField(devicerequest.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := druo.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, devicerequest.FieldScopes, value)
		})
	}
	if druo.mutation.ScopesCleared() {
		_spec.ClearField(devicerequest.FieldScopes, field.TypeJSON)
	}
	if value, ok := druo.mutation.Expiry(); ok {
		_spec.SetField(devicerequest.FieldExpiry, field.TypeTime, value)
	}
	_node = &DeviceRequest{config: druo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, druo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicerequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	druo.mutation.done = true
	return _node, nil
}
