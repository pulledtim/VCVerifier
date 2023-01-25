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
	"github.com/fiware/vcbackend/ent/credential"
	"github.com/fiware/vcbackend/ent/predicate"
	"github.com/fiware/vcbackend/ent/user"
)

// CredentialUpdate is the builder for updating Credential entities.
type CredentialUpdate struct {
	config
	hooks    []Hook
	mutation *CredentialMutation
}

// Where appends a list predicates to the CredentialUpdate builder.
func (cu *CredentialUpdate) Where(ps ...predicate.Credential) *CredentialUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetType sets the "type" field.
func (cu *CredentialUpdate) SetType(s string) *CredentialUpdate {
	cu.mutation.SetType(s)
	return cu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cu *CredentialUpdate) SetNillableType(s *string) *CredentialUpdate {
	if s != nil {
		cu.SetType(*s)
	}
	return cu
}

// SetRaw sets the "raw" field.
func (cu *CredentialUpdate) SetRaw(u []uint8) *CredentialUpdate {
	cu.mutation.SetRaw(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CredentialUpdate) SetUpdatedAt(t time.Time) *CredentialUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *CredentialUpdate) SetNillableUpdatedAt(t *time.Time) *CredentialUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// SetAccountID sets the "account" edge to the User entity by ID.
func (cu *CredentialUpdate) SetAccountID(id string) *CredentialUpdate {
	cu.mutation.SetAccountID(id)
	return cu
}

// SetNillableAccountID sets the "account" edge to the User entity by ID if the given value is not nil.
func (cu *CredentialUpdate) SetNillableAccountID(id *string) *CredentialUpdate {
	if id != nil {
		cu = cu.SetAccountID(*id)
	}
	return cu
}

// SetAccount sets the "account" edge to the User entity.
func (cu *CredentialUpdate) SetAccount(u *User) *CredentialUpdate {
	return cu.SetAccountID(u.ID)
}

// Mutation returns the CredentialMutation object of the builder.
func (cu *CredentialUpdate) Mutation() *CredentialMutation {
	return cu.mutation
}

// ClearAccount clears the "account" edge to the User entity.
func (cu *CredentialUpdate) ClearAccount() *CredentialUpdate {
	cu.mutation.ClearAccount()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CredentialUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CredentialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CredentialUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CredentialUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   credential.Table,
			Columns: credential.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: credential.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldType,
		})
	}
	if value, ok := cu.mutation.Raw(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: credential.FieldRaw,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldUpdatedAt,
		})
	}
	if cu.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.AccountTable,
			Columns: []string{credential.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.AccountTable,
			Columns: []string{credential.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CredentialUpdateOne is the builder for updating a single Credential entity.
type CredentialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CredentialMutation
}

// SetType sets the "type" field.
func (cuo *CredentialUpdateOne) SetType(s string) *CredentialUpdateOne {
	cuo.mutation.SetType(s)
	return cuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cuo *CredentialUpdateOne) SetNillableType(s *string) *CredentialUpdateOne {
	if s != nil {
		cuo.SetType(*s)
	}
	return cuo
}

// SetRaw sets the "raw" field.
func (cuo *CredentialUpdateOne) SetRaw(u []uint8) *CredentialUpdateOne {
	cuo.mutation.SetRaw(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CredentialUpdateOne) SetUpdatedAt(t time.Time) *CredentialUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *CredentialUpdateOne) SetNillableUpdatedAt(t *time.Time) *CredentialUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// SetAccountID sets the "account" edge to the User entity by ID.
func (cuo *CredentialUpdateOne) SetAccountID(id string) *CredentialUpdateOne {
	cuo.mutation.SetAccountID(id)
	return cuo
}

// SetNillableAccountID sets the "account" edge to the User entity by ID if the given value is not nil.
func (cuo *CredentialUpdateOne) SetNillableAccountID(id *string) *CredentialUpdateOne {
	if id != nil {
		cuo = cuo.SetAccountID(*id)
	}
	return cuo
}

// SetAccount sets the "account" edge to the User entity.
func (cuo *CredentialUpdateOne) SetAccount(u *User) *CredentialUpdateOne {
	return cuo.SetAccountID(u.ID)
}

// Mutation returns the CredentialMutation object of the builder.
func (cuo *CredentialUpdateOne) Mutation() *CredentialMutation {
	return cuo.mutation
}

// ClearAccount clears the "account" edge to the User entity.
func (cuo *CredentialUpdateOne) ClearAccount() *CredentialUpdateOne {
	cuo.mutation.ClearAccount()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CredentialUpdateOne) Select(field string, fields ...string) *CredentialUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Credential entity.
func (cuo *CredentialUpdateOne) Save(ctx context.Context) (*Credential, error) {
	var (
		err  error
		node *Credential
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CredentialMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Credential)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CredentialMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CredentialUpdateOne) SaveX(ctx context.Context) *Credential {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CredentialUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CredentialUpdateOne) sqlSave(ctx context.Context) (_node *Credential, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   credential.Table,
			Columns: credential.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: credential.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Credential.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, credential.FieldID)
		for _, f := range fields {
			if !credential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != credential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: credential.FieldType,
		})
	}
	if value, ok := cuo.mutation.Raw(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: credential.FieldRaw,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: credential.FieldUpdatedAt,
		})
	}
	if cuo.mutation.AccountCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.AccountTable,
			Columns: []string{credential.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   credential.AccountTable,
			Columns: []string{credential.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Credential{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
