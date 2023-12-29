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
	"github.com/google/uuid"
)

// ACLCreate is the builder for creating a ACL entity.
type ACLCreate struct {
	config
	mutation *ACLMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ac *ACLCreate) SetCreatedAt(u uint32) *ACLCreate {
	ac.mutation.SetCreatedAt(u)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ACLCreate) SetNillableCreatedAt(u *uint32) *ACLCreate {
	if u != nil {
		ac.SetCreatedAt(*u)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ACLCreate) SetUpdatedAt(u uint32) *ACLCreate {
	ac.mutation.SetUpdatedAt(u)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ACLCreate) SetNillableUpdatedAt(u *uint32) *ACLCreate {
	if u != nil {
		ac.SetUpdatedAt(*u)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *ACLCreate) SetDeletedAt(u uint32) *ACLCreate {
	ac.mutation.SetDeletedAt(u)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *ACLCreate) SetNillableDeletedAt(u *uint32) *ACLCreate {
	if u != nil {
		ac.SetDeletedAt(*u)
	}
	return ac
}

// SetEntID sets the "ent_id" field.
func (ac *ACLCreate) SetEntID(u uuid.UUID) *ACLCreate {
	ac.mutation.SetEntID(u)
	return ac
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ac *ACLCreate) SetNillableEntID(u *uuid.UUID) *ACLCreate {
	if u != nil {
		ac.SetEntID(*u)
	}
	return ac
}

// SetAppID sets the "app_id" field.
func (ac *ACLCreate) SetAppID(u uuid.UUID) *ACLCreate {
	ac.mutation.SetAppID(u)
	return ac
}

// SetRoleID sets the "role_id" field.
func (ac *ACLCreate) SetRoleID(u uuid.UUID) *ACLCreate {
	ac.mutation.SetRoleID(u)
	return ac
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (ac *ACLCreate) SetNillableRoleID(u *uuid.UUID) *ACLCreate {
	if u != nil {
		ac.SetRoleID(*u)
	}
	return ac
}

// SetArticleKey sets the "article_key" field.
func (ac *ACLCreate) SetArticleKey(u uuid.UUID) *ACLCreate {
	ac.mutation.SetArticleKey(u)
	return ac
}

// SetNillableArticleKey sets the "article_key" field if the given value is not nil.
func (ac *ACLCreate) SetNillableArticleKey(u *uuid.UUID) *ACLCreate {
	if u != nil {
		ac.SetArticleKey(*u)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ACLCreate) SetID(u uint32) *ACLCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the ACLMutation object of the builder.
func (ac *ACLCreate) Mutation() *ACLMutation {
	return ac.mutation
}

// Save creates the ACL in the database.
func (ac *ACLCreate) Save(ctx context.Context) (*ACL, error) {
	var (
		err  error
		node *ACL
	)
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ACLMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (ac *ACLCreate) SaveX(ctx context.Context) *ACL {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ACLCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ACLCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ACLCreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		if acl.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized acl.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := acl.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		if acl.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized acl.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := acl.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		if acl.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized acl.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := acl.DefaultDeletedAt()
		ac.mutation.SetDeletedAt(v)
	}
	if _, ok := ac.mutation.EntID(); !ok {
		if acl.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized acl.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := acl.DefaultEntID()
		ac.mutation.SetEntID(v)
	}
	if _, ok := ac.mutation.RoleID(); !ok {
		if acl.DefaultRoleID == nil {
			return fmt.Errorf("ent: uninitialized acl.DefaultRoleID (forgotten import ent/runtime?)")
		}
		v := acl.DefaultRoleID()
		ac.mutation.SetRoleID(v)
	}
	if _, ok := ac.mutation.ArticleKey(); !ok {
		if acl.DefaultArticleKey == nil {
			return fmt.Errorf("ent: uninitialized acl.DefaultArticleKey (forgotten import ent/runtime?)")
		}
		v := acl.DefaultArticleKey()
		ac.mutation.SetArticleKey(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *ACLCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ACL.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ACL.updated_at"`)}
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "ACL.deleted_at"`)}
	}
	if _, ok := ac.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "ACL.ent_id"`)}
	}
	if _, ok := ac.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "ACL.app_id"`)}
	}
	return nil
}

func (ac *ACLCreate) sqlSave(ctx context.Context) (*ACL, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (ac *ACLCreate) createSpec() (*ACL, *sqlgraph.CreateSpec) {
	var (
		_node = &ACL{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: acl.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: acl.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: acl.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := ac.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := ac.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldRoleID,
		})
		_node.RoleID = value
	}
	if value, ok := ac.mutation.ArticleKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: acl.FieldArticleKey,
		})
		_node.ArticleKey = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ACL.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ACLUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ac *ACLCreate) OnConflict(opts ...sql.ConflictOption) *ACLUpsertOne {
	ac.conflict = opts
	return &ACLUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ACL.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ac *ACLCreate) OnConflictColumns(columns ...string) *ACLUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &ACLUpsertOne{
		create: ac,
	}
}

type (
	// ACLUpsertOne is the builder for "upsert"-ing
	//  one ACL node.
	ACLUpsertOne struct {
		create *ACLCreate
	}

	// ACLUpsert is the "OnConflict" setter.
	ACLUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ACLUpsert) SetCreatedAt(v uint32) *ACLUpsert {
	u.Set(acl.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ACLUpsert) UpdateCreatedAt() *ACLUpsert {
	u.SetExcluded(acl.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ACLUpsert) AddCreatedAt(v uint32) *ACLUpsert {
	u.Add(acl.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ACLUpsert) SetUpdatedAt(v uint32) *ACLUpsert {
	u.Set(acl.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ACLUpsert) UpdateUpdatedAt() *ACLUpsert {
	u.SetExcluded(acl.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ACLUpsert) AddUpdatedAt(v uint32) *ACLUpsert {
	u.Add(acl.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ACLUpsert) SetDeletedAt(v uint32) *ACLUpsert {
	u.Set(acl.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ACLUpsert) UpdateDeletedAt() *ACLUpsert {
	u.SetExcluded(acl.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ACLUpsert) AddDeletedAt(v uint32) *ACLUpsert {
	u.Add(acl.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *ACLUpsert) SetEntID(v uuid.UUID) *ACLUpsert {
	u.Set(acl.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *ACLUpsert) UpdateEntID() *ACLUpsert {
	u.SetExcluded(acl.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *ACLUpsert) SetAppID(v uuid.UUID) *ACLUpsert {
	u.Set(acl.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ACLUpsert) UpdateAppID() *ACLUpsert {
	u.SetExcluded(acl.FieldAppID)
	return u
}

// SetRoleID sets the "role_id" field.
func (u *ACLUpsert) SetRoleID(v uuid.UUID) *ACLUpsert {
	u.Set(acl.FieldRoleID, v)
	return u
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *ACLUpsert) UpdateRoleID() *ACLUpsert {
	u.SetExcluded(acl.FieldRoleID)
	return u
}

// ClearRoleID clears the value of the "role_id" field.
func (u *ACLUpsert) ClearRoleID() *ACLUpsert {
	u.SetNull(acl.FieldRoleID)
	return u
}

// SetArticleKey sets the "article_key" field.
func (u *ACLUpsert) SetArticleKey(v uuid.UUID) *ACLUpsert {
	u.Set(acl.FieldArticleKey, v)
	return u
}

// UpdateArticleKey sets the "article_key" field to the value that was provided on create.
func (u *ACLUpsert) UpdateArticleKey() *ACLUpsert {
	u.SetExcluded(acl.FieldArticleKey)
	return u
}

// ClearArticleKey clears the value of the "article_key" field.
func (u *ACLUpsert) ClearArticleKey() *ACLUpsert {
	u.SetNull(acl.FieldArticleKey)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ACL.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(acl.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ACLUpsertOne) UpdateNewValues() *ACLUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(acl.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ACL.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ACLUpsertOne) Ignore() *ACLUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ACLUpsertOne) DoNothing() *ACLUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ACLCreate.OnConflict
// documentation for more info.
func (u *ACLUpsertOne) Update(set func(*ACLUpsert)) *ACLUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ACLUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ACLUpsertOne) SetCreatedAt(v uint32) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ACLUpsertOne) AddCreatedAt(v uint32) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ACLUpsertOne) UpdateCreatedAt() *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ACLUpsertOne) SetUpdatedAt(v uint32) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ACLUpsertOne) AddUpdatedAt(v uint32) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ACLUpsertOne) UpdateUpdatedAt() *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ACLUpsertOne) SetDeletedAt(v uint32) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ACLUpsertOne) AddDeletedAt(v uint32) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ACLUpsertOne) UpdateDeletedAt() *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *ACLUpsertOne) SetEntID(v uuid.UUID) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *ACLUpsertOne) UpdateEntID() *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *ACLUpsertOne) SetAppID(v uuid.UUID) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ACLUpsertOne) UpdateAppID() *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateAppID()
	})
}

// SetRoleID sets the "role_id" field.
func (u *ACLUpsertOne) SetRoleID(v uuid.UUID) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *ACLUpsertOne) UpdateRoleID() *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateRoleID()
	})
}

// ClearRoleID clears the value of the "role_id" field.
func (u *ACLUpsertOne) ClearRoleID() *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.ClearRoleID()
	})
}

// SetArticleKey sets the "article_key" field.
func (u *ACLUpsertOne) SetArticleKey(v uuid.UUID) *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.SetArticleKey(v)
	})
}

// UpdateArticleKey sets the "article_key" field to the value that was provided on create.
func (u *ACLUpsertOne) UpdateArticleKey() *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateArticleKey()
	})
}

// ClearArticleKey clears the value of the "article_key" field.
func (u *ACLUpsertOne) ClearArticleKey() *ACLUpsertOne {
	return u.Update(func(s *ACLUpsert) {
		s.ClearArticleKey()
	})
}

// Exec executes the query.
func (u *ACLUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ACLCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ACLUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ACLUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ACLUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ACLCreateBulk is the builder for creating many ACL entities in bulk.
type ACLCreateBulk struct {
	config
	builders []*ACLCreate
	conflict []sql.ConflictOption
}

// Save creates the ACL entities in the database.
func (acb *ACLCreateBulk) Save(ctx context.Context) ([]*ACL, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*ACL, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ACLMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ACLCreateBulk) SaveX(ctx context.Context) []*ACL {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ACLCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ACLCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ACL.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ACLUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (acb *ACLCreateBulk) OnConflict(opts ...sql.ConflictOption) *ACLUpsertBulk {
	acb.conflict = opts
	return &ACLUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ACL.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acb *ACLCreateBulk) OnConflictColumns(columns ...string) *ACLUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &ACLUpsertBulk{
		create: acb,
	}
}

// ACLUpsertBulk is the builder for "upsert"-ing
// a bulk of ACL nodes.
type ACLUpsertBulk struct {
	create *ACLCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ACL.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(acl.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ACLUpsertBulk) UpdateNewValues() *ACLUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(acl.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ACL.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ACLUpsertBulk) Ignore() *ACLUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ACLUpsertBulk) DoNothing() *ACLUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ACLCreateBulk.OnConflict
// documentation for more info.
func (u *ACLUpsertBulk) Update(set func(*ACLUpsert)) *ACLUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ACLUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ACLUpsertBulk) SetCreatedAt(v uint32) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ACLUpsertBulk) AddCreatedAt(v uint32) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ACLUpsertBulk) UpdateCreatedAt() *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ACLUpsertBulk) SetUpdatedAt(v uint32) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ACLUpsertBulk) AddUpdatedAt(v uint32) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ACLUpsertBulk) UpdateUpdatedAt() *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ACLUpsertBulk) SetDeletedAt(v uint32) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ACLUpsertBulk) AddDeletedAt(v uint32) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ACLUpsertBulk) UpdateDeletedAt() *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *ACLUpsertBulk) SetEntID(v uuid.UUID) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *ACLUpsertBulk) UpdateEntID() *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *ACLUpsertBulk) SetAppID(v uuid.UUID) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *ACLUpsertBulk) UpdateAppID() *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateAppID()
	})
}

// SetRoleID sets the "role_id" field.
func (u *ACLUpsertBulk) SetRoleID(v uuid.UUID) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *ACLUpsertBulk) UpdateRoleID() *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateRoleID()
	})
}

// ClearRoleID clears the value of the "role_id" field.
func (u *ACLUpsertBulk) ClearRoleID() *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.ClearRoleID()
	})
}

// SetArticleKey sets the "article_key" field.
func (u *ACLUpsertBulk) SetArticleKey(v uuid.UUID) *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.SetArticleKey(v)
	})
}

// UpdateArticleKey sets the "article_key" field to the value that was provided on create.
func (u *ACLUpsertBulk) UpdateArticleKey() *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.UpdateArticleKey()
	})
}

// ClearArticleKey clears the value of the "article_key" field.
func (u *ACLUpsertBulk) ClearArticleKey() *ACLUpsertBulk {
	return u.Update(func(s *ACLUpsert) {
		s.ClearArticleKey()
	})
}

// Exec executes the query.
func (u *ACLUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ACLCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ACLCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ACLUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
