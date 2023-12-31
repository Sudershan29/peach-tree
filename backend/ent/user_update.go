// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/ent/category"
	"backend/ent/predicate"
	"backend/ent/preference"
	"backend/ent/skill"
	"backend/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmailAddress sets the "email_address" field.
func (uu *UserUpdate) SetEmailAddress(s string) *UserUpdate {
	uu.mutation.SetEmailAddress(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetUUID sets the "uuid" field.
func (uu *UserUpdate) SetUUID(u uuid.UUID) *UserUpdate {
	uu.mutation.SetUUID(u)
	return uu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUUID(u *uuid.UUID) *UserUpdate {
	if u != nil {
		uu.SetUUID(*u)
	}
	return uu
}

// SetPremium sets the "premium" field.
func (uu *UserUpdate) SetPremium(b bool) *UserUpdate {
	uu.mutation.SetPremium(b)
	return uu
}

// SetNillablePremium sets the "premium" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePremium(b *bool) *UserUpdate {
	if b != nil {
		uu.SetPremium(*b)
	}
	return uu
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (uu *UserUpdate) AddSkillIDs(ids ...int) *UserUpdate {
	uu.mutation.AddSkillIDs(ids...)
	return uu
}

// AddSkills adds the "skills" edges to the Skill entity.
func (uu *UserUpdate) AddSkills(s ...*Skill) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddSkillIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (uu *UserUpdate) AddCategoryIDs(ids ...int) *UserUpdate {
	uu.mutation.AddCategoryIDs(ids...)
	return uu
}

// AddCategories adds the "categories" edges to the Category entity.
func (uu *UserUpdate) AddCategories(c ...*Category) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCategoryIDs(ids...)
}

// SetPreferenceID sets the "preference" edge to the Preference entity by ID.
func (uu *UserUpdate) SetPreferenceID(id int) *UserUpdate {
	uu.mutation.SetPreferenceID(id)
	return uu
}

// SetNillablePreferenceID sets the "preference" edge to the Preference entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillablePreferenceID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetPreferenceID(*id)
	}
	return uu
}

// SetPreference sets the "preference" edge to the Preference entity.
func (uu *UserUpdate) SetPreference(p *Preference) *UserUpdate {
	return uu.SetPreferenceID(p.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearSkills clears all "skills" edges to the Skill entity.
func (uu *UserUpdate) ClearSkills() *UserUpdate {
	uu.mutation.ClearSkills()
	return uu
}

// RemoveSkillIDs removes the "skills" edge to Skill entities by IDs.
func (uu *UserUpdate) RemoveSkillIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveSkillIDs(ids...)
	return uu
}

// RemoveSkills removes "skills" edges to Skill entities.
func (uu *UserUpdate) RemoveSkills(s ...*Skill) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveSkillIDs(ids...)
}

// ClearCategories clears all "categories" edges to the Category entity.
func (uu *UserUpdate) ClearCategories() *UserUpdate {
	uu.mutation.ClearCategories()
	return uu
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (uu *UserUpdate) RemoveCategoryIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveCategoryIDs(ids...)
	return uu
}

// RemoveCategories removes "categories" edges to Category entities.
func (uu *UserUpdate) RemoveCategories(c ...*Category) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCategoryIDs(ids...)
}

// ClearPreference clears the "preference" edge to the Preference entity.
func (uu *UserUpdate) ClearPreference() *UserUpdate {
	uu.mutation.ClearPreference()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.EmailAddress(); ok {
		_spec.SetField(user.FieldEmailAddress, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UUID(); ok {
		_spec.SetField(user.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := uu.mutation.Premium(); ok {
		_spec.SetField(user.FieldPremium, field.TypeBool, value)
	}
	if uu.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SkillsTable,
			Columns: []string{user.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSkillsIDs(); len(nodes) > 0 && !uu.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SkillsTable,
			Columns: []string{user.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SkillsTable,
			Columns: []string{user.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CategoriesTable,
			Columns: []string{user.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !uu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CategoriesTable,
			Columns: []string{user.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CategoriesTable,
			Columns: []string{user.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PreferenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PreferenceTable,
			Columns: []string{user.PreferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preference.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PreferenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PreferenceTable,
			Columns: []string{user.PreferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preference.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmailAddress sets the "email_address" field.
func (uuo *UserUpdateOne) SetEmailAddress(s string) *UserUpdateOne {
	uuo.mutation.SetEmailAddress(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetUUID sets the "uuid" field.
func (uuo *UserUpdateOne) SetUUID(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetUUID(u)
	return uuo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUUID(u *uuid.UUID) *UserUpdateOne {
	if u != nil {
		uuo.SetUUID(*u)
	}
	return uuo
}

// SetPremium sets the "premium" field.
func (uuo *UserUpdateOne) SetPremium(b bool) *UserUpdateOne {
	uuo.mutation.SetPremium(b)
	return uuo
}

// SetNillablePremium sets the "premium" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePremium(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetPremium(*b)
	}
	return uuo
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (uuo *UserUpdateOne) AddSkillIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddSkillIDs(ids...)
	return uuo
}

// AddSkills adds the "skills" edges to the Skill entity.
func (uuo *UserUpdateOne) AddSkills(s ...*Skill) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddSkillIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (uuo *UserUpdateOne) AddCategoryIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddCategoryIDs(ids...)
	return uuo
}

// AddCategories adds the "categories" edges to the Category entity.
func (uuo *UserUpdateOne) AddCategories(c ...*Category) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCategoryIDs(ids...)
}

// SetPreferenceID sets the "preference" edge to the Preference entity by ID.
func (uuo *UserUpdateOne) SetPreferenceID(id int) *UserUpdateOne {
	uuo.mutation.SetPreferenceID(id)
	return uuo
}

// SetNillablePreferenceID sets the "preference" edge to the Preference entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePreferenceID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetPreferenceID(*id)
	}
	return uuo
}

// SetPreference sets the "preference" edge to the Preference entity.
func (uuo *UserUpdateOne) SetPreference(p *Preference) *UserUpdateOne {
	return uuo.SetPreferenceID(p.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearSkills clears all "skills" edges to the Skill entity.
func (uuo *UserUpdateOne) ClearSkills() *UserUpdateOne {
	uuo.mutation.ClearSkills()
	return uuo
}

// RemoveSkillIDs removes the "skills" edge to Skill entities by IDs.
func (uuo *UserUpdateOne) RemoveSkillIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveSkillIDs(ids...)
	return uuo
}

// RemoveSkills removes "skills" edges to Skill entities.
func (uuo *UserUpdateOne) RemoveSkills(s ...*Skill) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveSkillIDs(ids...)
}

// ClearCategories clears all "categories" edges to the Category entity.
func (uuo *UserUpdateOne) ClearCategories() *UserUpdateOne {
	uuo.mutation.ClearCategories()
	return uuo
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (uuo *UserUpdateOne) RemoveCategoryIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveCategoryIDs(ids...)
	return uuo
}

// RemoveCategories removes "categories" edges to Category entities.
func (uuo *UserUpdateOne) RemoveCategories(c ...*Category) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCategoryIDs(ids...)
}

// ClearPreference clears the "preference" edge to the Preference entity.
func (uuo *UserUpdateOne) ClearPreference() *UserUpdateOne {
	uuo.mutation.ClearPreference()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.EmailAddress(); ok {
		_spec.SetField(user.FieldEmailAddress, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UUID(); ok {
		_spec.SetField(user.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := uuo.mutation.Premium(); ok {
		_spec.SetField(user.FieldPremium, field.TypeBool, value)
	}
	if uuo.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SkillsTable,
			Columns: []string{user.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSkillsIDs(); len(nodes) > 0 && !uuo.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SkillsTable,
			Columns: []string{user.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SkillsTable,
			Columns: []string{user.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CategoriesTable,
			Columns: []string{user.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !uuo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CategoriesTable,
			Columns: []string{user.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CategoriesTable,
			Columns: []string{user.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PreferenceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PreferenceTable,
			Columns: []string{user.PreferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preference.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PreferenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.PreferenceTable,
			Columns: []string{user.PreferenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preference.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
