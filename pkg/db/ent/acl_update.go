// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/acl"
	"github.com/NpoolPlatform/cms-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ACLUpdate is the builder for updating ACL entities.
type ACLUpdate struct {
	config
	hooks     []Hook
	mutation  *ACLMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ACLUpdate builder.
func (au *ACLUpdate) Where(ps ...predicate.ACL) *ACLUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *ACLUpdate) SetCreatedAt(u uint32) *ACLUpdate {
	au.mutation.ResetCreatedAt()
	au.mutation.SetCreatedAt(u)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *ACLUpdate) SetNillableCreatedAt(u *uint32) *ACLUpdate {
	if u != nil {
		au.SetCreatedAt(*u)
	}
	return au
}

// AddCreatedAt adds u to the "created_at" field.
func (au *ACLUpdate) AddCreatedAt(u int32) *ACLUpdate {
	au.mutation.AddCreatedAt(u)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ACLUpdate) SetUpdatedAt(u uint32) *ACLUpdate {
	au.mutation.ResetUpdatedAt()
	au.mutation.SetUpdatedAt(u)
	return au
}

// AddUpdatedAt adds u to the "updated_at" field.
func (au *ACLUpdate) AddUpdatedAt(u int32) *ACLUpdate {
	au.mutation.AddUpdatedAt(u)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *ACLUpdate) SetDeletedAt(u uint32) *ACLUpdate {
	au.mutation.ResetDeletedAt()
	au.mutation.SetDeletedAt(u)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *ACLUpdate) SetNillableDeletedAt(u *uint32) *ACLUpdate {
	if u != nil {
		au.SetDeletedAt(*u)
	}
	return au
}

// AddDeletedAt adds u to the "deleted_at" field.
func (au *ACLUpdate) AddDeletedAt(u int32) *ACLUpdate {
	au.mutation.AddDeletedAt(u)
	return au
}

// SetEntID sets the "ent_id" field.
func (au *ACLUpdate) SetEntID(u uuid.UUID) *ACLUpdate {
	au.mutation.SetEntID(u)
	return au
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (au *ACLUpdate) SetNillableEntID(u *uuid.UUID) *ACLUpdate {
	if u != nil {
		au.SetEntID(*u)
	}
	return au
}

// SetAppID sets the "app_id" field.
func (au *ACLUpdate) SetAppID(u uuid.UUID) *ACLUpdate {
	au.mutation.SetAppID(u)
	return au
}

// SetRoleID sets the "role_id" field.
func (au *ACLUpdate) SetRoleID(u uuid.UUID) *ACLUpdate {
	au.mutation.SetRoleID(u)
	return au
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (au *ACLUpdate) SetNillableRoleID(u *uuid.UUID) *ACLUpdate {
	if u != nil {
		au.SetRoleID(*u)
	}
	return au
}

// ClearRoleID clears the value of the "role_id" field.
func (au *ACLUpdate) ClearRoleID() *ACLUpdate {
	au.mutation.ClearRoleID()
	return au
}

// SetArticleKey sets the "article_key" field.
func (au *ACLUpdate) SetArticleKey(u uuid.UUID) *ACLUpdate {
	au.mutation.SetArticleKey(u)
	return au
}

// SetNillableArticleKey sets the "article_key" field if the given value is not nil.
func (au *ACLUpdate) SetNillableArticleKey(u *uuid.UUID) *ACLUpdate {
	if u != nil {
		au.SetArticleKey(*u)
	}
	return au
}

// ClearArticleKey clears the value of the "article_key" field.
func (au *ACLUpdate) ClearArticleKey() *ACLUpdate {
	au.mutation.ClearArticleKey()
	return au
}

// Mutation returns the ACLMutation object of the builder.
func (au *ACLUpdate) Mutation() *ACLMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ACLUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := au.defaults(); err != nil {
		return 0, err
	}
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ACLMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ACLUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ACLUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ACLUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *ACLUpdate) defaults() error {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		if acl.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized acl.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := acl.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *ACLUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ACLUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *ACLUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   acl.Table,
			Columns: acl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: acl.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldDeletedAt,
		})
	}
	if value, ok := au.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldDeletedAt,
		})
	}
	if value, ok := au.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldEntID,
		})
	}
	if value, ok := au.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldAppID,
		})
	}
	if value, ok := au.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldRoleID,
		})
	}
	if au.mutation.RoleIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: acl.FieldRoleID,
		})
	}
	if value, ok := au.mutation.ArticleKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldArticleKey,
		})
	}
	if au.mutation.ArticleKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: acl.FieldArticleKey,
		})
	}
	_spec.Modifiers = au.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{acl.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ACLUpdateOne is the builder for updating a single ACL entity.
type ACLUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ACLMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (auo *ACLUpdateOne) SetCreatedAt(u uint32) *ACLUpdateOne {
	auo.mutation.ResetCreatedAt()
	auo.mutation.SetCreatedAt(u)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *ACLUpdateOne) SetNillableCreatedAt(u *uint32) *ACLUpdateOne {
	if u != nil {
		auo.SetCreatedAt(*u)
	}
	return auo
}

// AddCreatedAt adds u to the "created_at" field.
func (auo *ACLUpdateOne) AddCreatedAt(u int32) *ACLUpdateOne {
	auo.mutation.AddCreatedAt(u)
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ACLUpdateOne) SetUpdatedAt(u uint32) *ACLUpdateOne {
	auo.mutation.ResetUpdatedAt()
	auo.mutation.SetUpdatedAt(u)
	return auo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (auo *ACLUpdateOne) AddUpdatedAt(u int32) *ACLUpdateOne {
	auo.mutation.AddUpdatedAt(u)
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *ACLUpdateOne) SetDeletedAt(u uint32) *ACLUpdateOne {
	auo.mutation.ResetDeletedAt()
	auo.mutation.SetDeletedAt(u)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *ACLUpdateOne) SetNillableDeletedAt(u *uint32) *ACLUpdateOne {
	if u != nil {
		auo.SetDeletedAt(*u)
	}
	return auo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (auo *ACLUpdateOne) AddDeletedAt(u int32) *ACLUpdateOne {
	auo.mutation.AddDeletedAt(u)
	return auo
}

// SetEntID sets the "ent_id" field.
func (auo *ACLUpdateOne) SetEntID(u uuid.UUID) *ACLUpdateOne {
	auo.mutation.SetEntID(u)
	return auo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (auo *ACLUpdateOne) SetNillableEntID(u *uuid.UUID) *ACLUpdateOne {
	if u != nil {
		auo.SetEntID(*u)
	}
	return auo
}

// SetAppID sets the "app_id" field.
func (auo *ACLUpdateOne) SetAppID(u uuid.UUID) *ACLUpdateOne {
	auo.mutation.SetAppID(u)
	return auo
}

// SetRoleID sets the "role_id" field.
func (auo *ACLUpdateOne) SetRoleID(u uuid.UUID) *ACLUpdateOne {
	auo.mutation.SetRoleID(u)
	return auo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (auo *ACLUpdateOne) SetNillableRoleID(u *uuid.UUID) *ACLUpdateOne {
	if u != nil {
		auo.SetRoleID(*u)
	}
	return auo
}

// ClearRoleID clears the value of the "role_id" field.
func (auo *ACLUpdateOne) ClearRoleID() *ACLUpdateOne {
	auo.mutation.ClearRoleID()
	return auo
}

// SetArticleKey sets the "article_key" field.
func (auo *ACLUpdateOne) SetArticleKey(u uuid.UUID) *ACLUpdateOne {
	auo.mutation.SetArticleKey(u)
	return auo
}

// SetNillableArticleKey sets the "article_key" field if the given value is not nil.
func (auo *ACLUpdateOne) SetNillableArticleKey(u *uuid.UUID) *ACLUpdateOne {
	if u != nil {
		auo.SetArticleKey(*u)
	}
	return auo
}

// ClearArticleKey clears the value of the "article_key" field.
func (auo *ACLUpdateOne) ClearArticleKey() *ACLUpdateOne {
	auo.mutation.ClearArticleKey()
	return auo
}

// Mutation returns the ACLMutation object of the builder.
func (auo *ACLUpdateOne) Mutation() *ACLMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ACLUpdateOne) Select(field string, fields ...string) *ACLUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated ACL entity.
func (auo *ACLUpdateOne) Save(ctx context.Context) (*ACL, error) {
	var (
		err  error
		node *ACL
	)
	if err := auo.defaults(); err != nil {
		return nil, err
	}
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ACLMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ACL)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ACLMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ACLUpdateOne) SaveX(ctx context.Context) *ACL {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ACLUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ACLUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *ACLUpdateOne) defaults() error {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		if acl.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized acl.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := acl.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *ACLUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ACLUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *ACLUpdateOne) sqlSave(ctx context.Context) (_node *ACL, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   acl.Table,
			Columns: acl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: acl.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ACL.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, acl.FieldID)
		for _, f := range fields {
			if !acl.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != acl.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldDeletedAt,
		})
	}
	if value, ok := auo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldDeletedAt,
		})
	}
	if value, ok := auo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldEntID,
		})
	}
	if value, ok := auo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldAppID,
		})
	}
	if value, ok := auo.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldRoleID,
		})
	}
	if auo.mutation.RoleIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: acl.FieldRoleID,
		})
	}
	if value, ok := auo.mutation.ArticleKey(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldArticleKey,
		})
	}
	if auo.mutation.ArticleKeyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: acl.FieldArticleKey,
		})
	}
	_spec.Modifiers = auo.modifiers
	_node = &ACL{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{acl.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
