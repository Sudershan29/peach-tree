// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/preference"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PreferenceCreate is the builder for creating a Preference entity.
type PreferenceCreate struct {
	config
	mutation *PreferenceMutation
	hooks    []Hook
}

// SetFreeWeekends sets the "free_weekends" field.
func (pc *PreferenceCreate) SetFreeWeekends(b bool) *PreferenceCreate {
	pc.mutation.SetFreeWeekends(b)
	return pc
}

// SetNillableFreeWeekends sets the "free_weekends" field if the given value is not nil.
func (pc *PreferenceCreate) SetNillableFreeWeekends(b *bool) *PreferenceCreate {
	if b != nil {
		pc.SetFreeWeekends(*b)
	}
	return pc
}

// SetWeeklyFrequency sets the "weekly_frequency" field.
func (pc *PreferenceCreate) SetWeeklyFrequency(i int) *PreferenceCreate {
	pc.mutation.SetWeeklyFrequency(i)
	return pc
}

// SetNillableWeeklyFrequency sets the "weekly_frequency" field if the given value is not nil.
func (pc *PreferenceCreate) SetNillableWeeklyFrequency(i *int) *PreferenceCreate {
	if i != nil {
		pc.SetWeeklyFrequency(*i)
	}
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PreferenceCreate) SetUserID(id int) *PreferenceCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pc *PreferenceCreate) SetNillableUserID(id *int) *PreferenceCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PreferenceCreate) SetUser(u *User) *PreferenceCreate {
	return pc.SetUserID(u.ID)
}

// Mutation returns the PreferenceMutation object of the builder.
func (pc *PreferenceCreate) Mutation() *PreferenceMutation {
	return pc.mutation
}

// Save creates the Preference in the database.
func (pc *PreferenceCreate) Save(ctx context.Context) (*Preference, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PreferenceCreate) SaveX(ctx context.Context) *Preference {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PreferenceCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PreferenceCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PreferenceCreate) defaults() {
	if _, ok := pc.mutation.FreeWeekends(); !ok {
		v := preference.DefaultFreeWeekends
		pc.mutation.SetFreeWeekends(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PreferenceCreate) check() error {
	if _, ok := pc.mutation.FreeWeekends(); !ok {
		return &ValidationError{Name: "free_weekends", err: errors.New(`ent: missing required field "Preference.free_weekends"`)}
	}
	if v, ok := pc.mutation.WeeklyFrequency(); ok {
		if err := preference.WeeklyFrequencyValidator(v); err != nil {
			return &ValidationError{Name: "weekly_frequency", err: fmt.Errorf(`ent: validator failed for field "Preference.weekly_frequency": %w`, err)}
		}
	}
	return nil
}

func (pc *PreferenceCreate) sqlSave(ctx context.Context) (*Preference, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PreferenceCreate) createSpec() (*Preference, *sqlgraph.CreateSpec) {
	var (
		_node = &Preference{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(preference.Table, sqlgraph.NewFieldSpec(preference.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.FreeWeekends(); ok {
		_spec.SetField(preference.FieldFreeWeekends, field.TypeBool, value)
		_node.FreeWeekends = value
	}
	if value, ok := pc.mutation.WeeklyFrequency(); ok {
		_spec.SetField(preference.FieldWeeklyFrequency, field.TypeInt, value)
		_node.WeeklyFrequency = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   preference.UserTable,
			Columns: []string{preference.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_preference = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PreferenceCreateBulk is the builder for creating many Preference entities in bulk.
type PreferenceCreateBulk struct {
	config
	builders []*PreferenceCreate
}

// Save creates the Preference entities in the database.
func (pcb *PreferenceCreateBulk) Save(ctx context.Context) ([]*Preference, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Preference, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PreferenceMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PreferenceCreateBulk) SaveX(ctx context.Context) []*Preference {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PreferenceCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PreferenceCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
