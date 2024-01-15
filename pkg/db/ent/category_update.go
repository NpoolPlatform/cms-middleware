// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/category"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks     []Hook
	mutation  *CategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CategoryUpdate) SetCreatedAt(u uint32) *CategoryUpdate {
	cu.mutation.ResetCreatedAt()
	cu.mutation.SetCreatedAt(u)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableCreatedAt(u *uint32) *CategoryUpdate {
	if u != nil {
		cu.SetCreatedAt(*u)
	}
	return cu
}

// AddCreatedAt adds u to the "created_at" field.
func (cu *CategoryUpdate) AddCreatedAt(u int32) *CategoryUpdate {
	cu.mutation.AddCreatedAt(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CategoryUpdate) SetUpdatedAt(u uint32) *CategoryUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(u)
	return cu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cu *CategoryUpdate) AddUpdatedAt(u int32) *CategoryUpdate {
	cu.mutation.AddUpdatedAt(u)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CategoryUpdate) SetDeletedAt(u uint32) *CategoryUpdate {
	cu.mutation.ResetDeletedAt()
	cu.mutation.SetDeletedAt(u)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDeletedAt(u *uint32) *CategoryUpdate {
	if u != nil {
		cu.SetDeletedAt(*u)
	}
	return cu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cu *CategoryUpdate) AddDeletedAt(u int32) *CategoryUpdate {
	cu.mutation.AddDeletedAt(u)
	return cu
}

// SetEntID sets the "ent_id" field.
func (cu *CategoryUpdate) SetEntID(u uuid.UUID) *CategoryUpdate {
	cu.mutation.SetEntID(u)
	return cu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableEntID(u *uuid.UUID) *CategoryUpdate {
	if u != nil {
		cu.SetEntID(*u)
	}
	return cu
}

// SetAppID sets the "app_id" field.
func (cu *CategoryUpdate) SetAppID(u uuid.UUID) *CategoryUpdate {
	cu.mutation.SetAppID(u)
	return cu
}

// SetParentID sets the "parent_id" field.
func (cu *CategoryUpdate) SetParentID(u uuid.UUID) *CategoryUpdate {
	cu.mutation.SetParentID(u)
	return cu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableParentID(u *uuid.UUID) *CategoryUpdate {
	if u != nil {
		cu.SetParentID(*u)
	}
	return cu
}

// ClearParentID clears the value of the "parent_id" field.
func (cu *CategoryUpdate) ClearParentID() *CategoryUpdate {
	cu.mutation.ClearParentID()
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableName(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "name" field.
func (cu *CategoryUpdate) ClearName() *CategoryUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetSlug sets the "slug" field.
func (cu *CategoryUpdate) SetSlug(s string) *CategoryUpdate {
	cu.mutation.SetSlug(s)
	return cu
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableSlug(s *string) *CategoryUpdate {
	if s != nil {
		cu.SetSlug(*s)
	}
	return cu
}

// ClearSlug clears the value of the "slug" field.
func (cu *CategoryUpdate) ClearSlug() *CategoryUpdate {
	cu.mutation.ClearSlug()
	return cu
}

// SetEnabled sets the "enabled" field.
func (cu *CategoryUpdate) SetEnabled(b bool) *CategoryUpdate {
	cu.mutation.SetEnabled(b)
	return cu
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableEnabled(b *bool) *CategoryUpdate {
	if b != nil {
		cu.SetEnabled(*b)
	}
	return cu
}

// ClearEnabled clears the value of the "enabled" field.
func (cu *CategoryUpdate) ClearEnabled() *CategoryUpdate {
	cu.mutation.ClearEnabled()
	return cu
}

// SetIndex sets the "index" field.
func (cu *CategoryUpdate) SetIndex(u uint32) *CategoryUpdate {
	cu.mutation.ResetIndex()
	cu.mutation.SetIndex(u)
	return cu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableIndex(u *uint32) *CategoryUpdate {
	if u != nil {
		cu.SetIndex(*u)
	}
	return cu
}

// AddIndex adds u to the "index" field.
func (cu *CategoryUpdate) AddIndex(u int32) *CategoryUpdate {
	cu.mutation.AddIndex(u)
	return cu
}

// ClearIndex clears the value of the "index" field.
func (cu *CategoryUpdate) ClearIndex() *CategoryUpdate {
	cu.mutation.ClearIndex()
	return cu
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
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
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CategoryUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if category.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized category.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := category.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CategoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CategoryUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: category.FieldID,
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
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: category.FieldEntID,
		})
	}
	if value, ok := cu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: category.FieldAppID,
		})
	}
	if value, ok := cu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if cu.mutation.ParentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: category.FieldParentID,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
	}
	if cu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: category.FieldName,
		})
	}
	if value, ok := cu.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldSlug,
		})
	}
	if cu.mutation.SlugCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: category.FieldSlug,
		})
	}
	if value, ok := cu.mutation.Enabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: category.FieldEnabled,
		})
	}
	if cu.mutation.EnabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: category.FieldEnabled,
		})
	}
	if value, ok := cu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldIndex,
		})
	}
	if value, ok := cu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldIndex,
		})
	}
	if cu.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: category.FieldIndex,
		})
	}
	_spec.Modifiers = cu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CategoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CategoryUpdateOne) SetCreatedAt(u uint32) *CategoryUpdateOne {
	cuo.mutation.ResetCreatedAt()
	cuo.mutation.SetCreatedAt(u)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableCreatedAt(u *uint32) *CategoryUpdateOne {
	if u != nil {
		cuo.SetCreatedAt(*u)
	}
	return cuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cuo *CategoryUpdateOne) AddCreatedAt(u int32) *CategoryUpdateOne {
	cuo.mutation.AddCreatedAt(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CategoryUpdateOne) SetUpdatedAt(u uint32) *CategoryUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(u)
	return cuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cuo *CategoryUpdateOne) AddUpdatedAt(u int32) *CategoryUpdateOne {
	cuo.mutation.AddUpdatedAt(u)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CategoryUpdateOne) SetDeletedAt(u uint32) *CategoryUpdateOne {
	cuo.mutation.ResetDeletedAt()
	cuo.mutation.SetDeletedAt(u)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDeletedAt(u *uint32) *CategoryUpdateOne {
	if u != nil {
		cuo.SetDeletedAt(*u)
	}
	return cuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cuo *CategoryUpdateOne) AddDeletedAt(u int32) *CategoryUpdateOne {
	cuo.mutation.AddDeletedAt(u)
	return cuo
}

// SetEntID sets the "ent_id" field.
func (cuo *CategoryUpdateOne) SetEntID(u uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.SetEntID(u)
	return cuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableEntID(u *uuid.UUID) *CategoryUpdateOne {
	if u != nil {
		cuo.SetEntID(*u)
	}
	return cuo
}

// SetAppID sets the "app_id" field.
func (cuo *CategoryUpdateOne) SetAppID(u uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.SetAppID(u)
	return cuo
}

// SetParentID sets the "parent_id" field.
func (cuo *CategoryUpdateOne) SetParentID(u uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.SetParentID(u)
	return cuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableParentID(u *uuid.UUID) *CategoryUpdateOne {
	if u != nil {
		cuo.SetParentID(*u)
	}
	return cuo
}

// ClearParentID clears the value of the "parent_id" field.
func (cuo *CategoryUpdateOne) ClearParentID() *CategoryUpdateOne {
	cuo.mutation.ClearParentID()
	return cuo
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableName(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "name" field.
func (cuo *CategoryUpdateOne) ClearName() *CategoryUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetSlug sets the "slug" field.
func (cuo *CategoryUpdateOne) SetSlug(s string) *CategoryUpdateOne {
	cuo.mutation.SetSlug(s)
	return cuo
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableSlug(s *string) *CategoryUpdateOne {
	if s != nil {
		cuo.SetSlug(*s)
	}
	return cuo
}

// ClearSlug clears the value of the "slug" field.
func (cuo *CategoryUpdateOne) ClearSlug() *CategoryUpdateOne {
	cuo.mutation.ClearSlug()
	return cuo
}

// SetEnabled sets the "enabled" field.
func (cuo *CategoryUpdateOne) SetEnabled(b bool) *CategoryUpdateOne {
	cuo.mutation.SetEnabled(b)
	return cuo
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableEnabled(b *bool) *CategoryUpdateOne {
	if b != nil {
		cuo.SetEnabled(*b)
	}
	return cuo
}

// ClearEnabled clears the value of the "enabled" field.
func (cuo *CategoryUpdateOne) ClearEnabled() *CategoryUpdateOne {
	cuo.mutation.ClearEnabled()
	return cuo
}

// SetIndex sets the "index" field.
func (cuo *CategoryUpdateOne) SetIndex(u uint32) *CategoryUpdateOne {
	cuo.mutation.ResetIndex()
	cuo.mutation.SetIndex(u)
	return cuo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableIndex(u *uint32) *CategoryUpdateOne {
	if u != nil {
		cuo.SetIndex(*u)
	}
	return cuo
}

// AddIndex adds u to the "index" field.
func (cuo *CategoryUpdateOne) AddIndex(u int32) *CategoryUpdateOne {
	cuo.mutation.AddIndex(u)
	return cuo
}

// ClearIndex clears the value of the "index" field.
func (cuo *CategoryUpdateOne) ClearIndex() *CategoryUpdateOne {
	cuo.mutation.ClearIndex()
	return cuo
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	var (
		err  error
		node *Category
	)
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
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
		nv, ok := v.(*Category)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CategoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CategoryUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if category.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized category.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := category.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CategoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CategoryUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   category.Table,
			Columns: category.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: category.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
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
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: category.FieldEntID,
		})
	}
	if value, ok := cuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: category.FieldAppID,
		})
	}
	if value, ok := cuo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: category.FieldParentID,
		})
	}
	if cuo.mutation.ParentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: category.FieldParentID,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
	}
	if cuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: category.FieldName,
		})
	}
	if value, ok := cuo.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldSlug,
		})
	}
	if cuo.mutation.SlugCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: category.FieldSlug,
		})
	}
	if value, ok := cuo.mutation.Enabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: category.FieldEnabled,
		})
	}
	if cuo.mutation.EnabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: category.FieldEnabled,
		})
	}
	if value, ok := cuo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldIndex,
		})
	}
	if value, ok := cuo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: category.FieldIndex,
		})
	}
	if cuo.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: category.FieldIndex,
		})
	}
	_spec.Modifiers = cuo.modifiers
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
