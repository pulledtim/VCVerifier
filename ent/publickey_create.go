// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/fiware/vcbackend/ent/publickey"
)

// PublicKeyCreate is the builder for creating a PublicKey entity.
type PublicKeyCreate struct {
	config
	mutation *PublicKeyMutation
	hooks    []Hook
}

// SetKty sets the "kty" field.
func (pkc *PublicKeyCreate) SetKty(s string) *PublicKeyCreate {
	pkc.mutation.SetKty(s)
	return pkc
}

// SetAlg sets the "alg" field.
func (pkc *PublicKeyCreate) SetAlg(s string) *PublicKeyCreate {
	pkc.mutation.SetAlg(s)
	return pkc
}

// SetNillableAlg sets the "alg" field if the given value is not nil.
func (pkc *PublicKeyCreate) SetNillableAlg(s *string) *PublicKeyCreate {
	if s != nil {
		pkc.SetAlg(*s)
	}
	return pkc
}

// SetJwk sets the "jwk" field.
func (pkc *PublicKeyCreate) SetJwk(u []uint8) *PublicKeyCreate {
	pkc.mutation.SetJwk(u)
	return pkc
}

// SetCreatedAt sets the "created_at" field.
func (pkc *PublicKeyCreate) SetCreatedAt(t time.Time) *PublicKeyCreate {
	pkc.mutation.SetCreatedAt(t)
	return pkc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pkc *PublicKeyCreate) SetNillableCreatedAt(t *time.Time) *PublicKeyCreate {
	if t != nil {
		pkc.SetCreatedAt(*t)
	}
	return pkc
}

// SetUpdatedAt sets the "updated_at" field.
func (pkc *PublicKeyCreate) SetUpdatedAt(t time.Time) *PublicKeyCreate {
	pkc.mutation.SetUpdatedAt(t)
	return pkc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pkc *PublicKeyCreate) SetNillableUpdatedAt(t *time.Time) *PublicKeyCreate {
	if t != nil {
		pkc.SetUpdatedAt(*t)
	}
	return pkc
}

// SetID sets the "id" field.
func (pkc *PublicKeyCreate) SetID(s string) *PublicKeyCreate {
	pkc.mutation.SetID(s)
	return pkc
}

// Mutation returns the PublicKeyMutation object of the builder.
func (pkc *PublicKeyCreate) Mutation() *PublicKeyMutation {
	return pkc.mutation
}

// Save creates the PublicKey in the database.
func (pkc *PublicKeyCreate) Save(ctx context.Context) (*PublicKey, error) {
	var (
		err  error
		node *PublicKey
	)
	pkc.defaults()
	if len(pkc.hooks) == 0 {
		if err = pkc.check(); err != nil {
			return nil, err
		}
		node, err = pkc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PublicKeyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pkc.check(); err != nil {
				return nil, err
			}
			pkc.mutation = mutation
			if node, err = pkc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pkc.hooks) - 1; i >= 0; i-- {
			if pkc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pkc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pkc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PublicKey)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PublicKeyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pkc *PublicKeyCreate) SaveX(ctx context.Context) *PublicKey {
	v, err := pkc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pkc *PublicKeyCreate) Exec(ctx context.Context) error {
	_, err := pkc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pkc *PublicKeyCreate) ExecX(ctx context.Context) {
	if err := pkc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pkc *PublicKeyCreate) defaults() {
	if _, ok := pkc.mutation.CreatedAt(); !ok {
		v := publickey.DefaultCreatedAt()
		pkc.mutation.SetCreatedAt(v)
	}
	if _, ok := pkc.mutation.UpdatedAt(); !ok {
		v := publickey.DefaultUpdatedAt()
		pkc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pkc *PublicKeyCreate) check() error {
	if _, ok := pkc.mutation.Kty(); !ok {
		return &ValidationError{Name: "kty", err: errors.New(`ent: missing required field "PublicKey.kty"`)}
	}
	if _, ok := pkc.mutation.Jwk(); !ok {
		return &ValidationError{Name: "jwk", err: errors.New(`ent: missing required field "PublicKey.jwk"`)}
	}
	if _, ok := pkc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PublicKey.created_at"`)}
	}
	if _, ok := pkc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PublicKey.updated_at"`)}
	}
	return nil
}

func (pkc *PublicKeyCreate) sqlSave(ctx context.Context) (*PublicKey, error) {
	_node, _spec := pkc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pkc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected PublicKey.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (pkc *PublicKeyCreate) createSpec() (*PublicKey, *sqlgraph.CreateSpec) {
	var (
		_node = &PublicKey{config: pkc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: publickey.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: publickey.FieldID,
			},
		}
	)
	if id, ok := pkc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pkc.mutation.Kty(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: publickey.FieldKty,
		})
		_node.Kty = value
	}
	if value, ok := pkc.mutation.Alg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: publickey.FieldAlg,
		})
		_node.Alg = value
	}
	if value, ok := pkc.mutation.Jwk(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: publickey.FieldJwk,
		})
		_node.Jwk = value
	}
	if value, ok := pkc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: publickey.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pkc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: publickey.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// PublicKeyCreateBulk is the builder for creating many PublicKey entities in bulk.
type PublicKeyCreateBulk struct {
	config
	builders []*PublicKeyCreate
}

// Save creates the PublicKey entities in the database.
func (pkcb *PublicKeyCreateBulk) Save(ctx context.Context) ([]*PublicKey, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pkcb.builders))
	nodes := make([]*PublicKey, len(pkcb.builders))
	mutators := make([]Mutator, len(pkcb.builders))
	for i := range pkcb.builders {
		func(i int, root context.Context) {
			builder := pkcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublicKeyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pkcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pkcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pkcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pkcb *PublicKeyCreateBulk) SaveX(ctx context.Context) []*PublicKey {
	v, err := pkcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pkcb *PublicKeyCreateBulk) Exec(ctx context.Context) error {
	_, err := pkcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pkcb *PublicKeyCreateBulk) ExecX(ctx context.Context) {
	if err := pkcb.Exec(ctx); err != nil {
		panic(err)
	}
}