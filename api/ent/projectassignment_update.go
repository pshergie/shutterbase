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
	"github.com/google/uuid"
	"github.com/shutterbase/shutterbase/ent/predicate"
	"github.com/shutterbase/shutterbase/ent/project"
	"github.com/shutterbase/shutterbase/ent/projectassignment"
	"github.com/shutterbase/shutterbase/ent/role"
	"github.com/shutterbase/shutterbase/ent/user"
)

// ProjectAssignmentUpdate is the builder for updating ProjectAssignment entities.
type ProjectAssignmentUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectAssignmentMutation
}

// Where appends a list predicates to the ProjectAssignmentUpdate builder.
func (pau *ProjectAssignmentUpdate) Where(ps ...predicate.ProjectAssignment) *ProjectAssignmentUpdate {
	pau.mutation.Where(ps...)
	return pau
}

// SetUpdatedAt sets the "updated_at" field.
func (pau *ProjectAssignmentUpdate) SetUpdatedAt(t time.Time) *ProjectAssignmentUpdate {
	pau.mutation.SetUpdatedAt(t)
	return pau
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pau *ProjectAssignmentUpdate) SetUserID(id uuid.UUID) *ProjectAssignmentUpdate {
	pau.mutation.SetUserID(id)
	return pau
}

// SetUser sets the "user" edge to the User entity.
func (pau *ProjectAssignmentUpdate) SetUser(u *User) *ProjectAssignmentUpdate {
	return pau.SetUserID(u.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (pau *ProjectAssignmentUpdate) SetProjectID(id uuid.UUID) *ProjectAssignmentUpdate {
	pau.mutation.SetProjectID(id)
	return pau
}

// SetProject sets the "project" edge to the Project entity.
func (pau *ProjectAssignmentUpdate) SetProject(p *Project) *ProjectAssignmentUpdate {
	return pau.SetProjectID(p.ID)
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (pau *ProjectAssignmentUpdate) SetRoleID(id uuid.UUID) *ProjectAssignmentUpdate {
	pau.mutation.SetRoleID(id)
	return pau
}

// SetNillableRoleID sets the "role" edge to the Role entity by ID if the given value is not nil.
func (pau *ProjectAssignmentUpdate) SetNillableRoleID(id *uuid.UUID) *ProjectAssignmentUpdate {
	if id != nil {
		pau = pau.SetRoleID(*id)
	}
	return pau
}

// SetRole sets the "role" edge to the Role entity.
func (pau *ProjectAssignmentUpdate) SetRole(r *Role) *ProjectAssignmentUpdate {
	return pau.SetRoleID(r.ID)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (pau *ProjectAssignmentUpdate) SetCreatedByID(id uuid.UUID) *ProjectAssignmentUpdate {
	pau.mutation.SetCreatedByID(id)
	return pau
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (pau *ProjectAssignmentUpdate) SetNillableCreatedByID(id *uuid.UUID) *ProjectAssignmentUpdate {
	if id != nil {
		pau = pau.SetCreatedByID(*id)
	}
	return pau
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (pau *ProjectAssignmentUpdate) SetCreatedBy(u *User) *ProjectAssignmentUpdate {
	return pau.SetCreatedByID(u.ID)
}

// SetUpdatedByID sets the "updated_by" edge to the User entity by ID.
func (pau *ProjectAssignmentUpdate) SetUpdatedByID(id uuid.UUID) *ProjectAssignmentUpdate {
	pau.mutation.SetUpdatedByID(id)
	return pau
}

// SetNillableUpdatedByID sets the "updated_by" edge to the User entity by ID if the given value is not nil.
func (pau *ProjectAssignmentUpdate) SetNillableUpdatedByID(id *uuid.UUID) *ProjectAssignmentUpdate {
	if id != nil {
		pau = pau.SetUpdatedByID(*id)
	}
	return pau
}

// SetUpdatedBy sets the "updated_by" edge to the User entity.
func (pau *ProjectAssignmentUpdate) SetUpdatedBy(u *User) *ProjectAssignmentUpdate {
	return pau.SetUpdatedByID(u.ID)
}

// Mutation returns the ProjectAssignmentMutation object of the builder.
func (pau *ProjectAssignmentUpdate) Mutation() *ProjectAssignmentMutation {
	return pau.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pau *ProjectAssignmentUpdate) ClearUser() *ProjectAssignmentUpdate {
	pau.mutation.ClearUser()
	return pau
}

// ClearProject clears the "project" edge to the Project entity.
func (pau *ProjectAssignmentUpdate) ClearProject() *ProjectAssignmentUpdate {
	pau.mutation.ClearProject()
	return pau
}

// ClearRole clears the "role" edge to the Role entity.
func (pau *ProjectAssignmentUpdate) ClearRole() *ProjectAssignmentUpdate {
	pau.mutation.ClearRole()
	return pau
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (pau *ProjectAssignmentUpdate) ClearCreatedBy() *ProjectAssignmentUpdate {
	pau.mutation.ClearCreatedBy()
	return pau
}

// ClearUpdatedBy clears the "updated_by" edge to the User entity.
func (pau *ProjectAssignmentUpdate) ClearUpdatedBy() *ProjectAssignmentUpdate {
	pau.mutation.ClearUpdatedBy()
	return pau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *ProjectAssignmentUpdate) Save(ctx context.Context) (int, error) {
	pau.defaults()
	return withHooks(ctx, pau.sqlSave, pau.mutation, pau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pau *ProjectAssignmentUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *ProjectAssignmentUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *ProjectAssignmentUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pau *ProjectAssignmentUpdate) defaults() {
	if _, ok := pau.mutation.UpdatedAt(); !ok {
		v := projectassignment.UpdateDefaultUpdatedAt()
		pau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pau *ProjectAssignmentUpdate) check() error {
	if _, ok := pau.mutation.UserID(); pau.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectAssignment.user"`)
	}
	if _, ok := pau.mutation.ProjectID(); pau.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectAssignment.project"`)
	}
	return nil
}

func (pau *ProjectAssignmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(projectassignment.Table, projectassignment.Columns, sqlgraph.NewFieldSpec(projectassignment.FieldID, field.TypeUUID))
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.UpdatedAt(); ok {
		_spec.SetField(projectassignment.FieldUpdatedAt, field.TypeTime, value)
	}
	if pau.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.UserTable,
			Columns: []string{projectassignment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.UserTable,
			Columns: []string{projectassignment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pau.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.ProjectTable,
			Columns: []string{projectassignment.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.ProjectTable,
			Columns: []string{projectassignment.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pau.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.RoleTable,
			Columns: []string{projectassignment.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.RoleTable,
			Columns: []string{projectassignment.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pau.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.CreatedByTable,
			Columns: []string{projectassignment.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.CreatedByTable,
			Columns: []string{projectassignment.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pau.mutation.UpdatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.UpdatedByTable,
			Columns: []string{projectassignment.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.UpdatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.UpdatedByTable,
			Columns: []string{projectassignment.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectassignment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pau.mutation.done = true
	return n, nil
}

// ProjectAssignmentUpdateOne is the builder for updating a single ProjectAssignment entity.
type ProjectAssignmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectAssignmentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pauo *ProjectAssignmentUpdateOne) SetUpdatedAt(t time.Time) *ProjectAssignmentUpdateOne {
	pauo.mutation.SetUpdatedAt(t)
	return pauo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pauo *ProjectAssignmentUpdateOne) SetUserID(id uuid.UUID) *ProjectAssignmentUpdateOne {
	pauo.mutation.SetUserID(id)
	return pauo
}

// SetUser sets the "user" edge to the User entity.
func (pauo *ProjectAssignmentUpdateOne) SetUser(u *User) *ProjectAssignmentUpdateOne {
	return pauo.SetUserID(u.ID)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (pauo *ProjectAssignmentUpdateOne) SetProjectID(id uuid.UUID) *ProjectAssignmentUpdateOne {
	pauo.mutation.SetProjectID(id)
	return pauo
}

// SetProject sets the "project" edge to the Project entity.
func (pauo *ProjectAssignmentUpdateOne) SetProject(p *Project) *ProjectAssignmentUpdateOne {
	return pauo.SetProjectID(p.ID)
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (pauo *ProjectAssignmentUpdateOne) SetRoleID(id uuid.UUID) *ProjectAssignmentUpdateOne {
	pauo.mutation.SetRoleID(id)
	return pauo
}

// SetNillableRoleID sets the "role" edge to the Role entity by ID if the given value is not nil.
func (pauo *ProjectAssignmentUpdateOne) SetNillableRoleID(id *uuid.UUID) *ProjectAssignmentUpdateOne {
	if id != nil {
		pauo = pauo.SetRoleID(*id)
	}
	return pauo
}

// SetRole sets the "role" edge to the Role entity.
func (pauo *ProjectAssignmentUpdateOne) SetRole(r *Role) *ProjectAssignmentUpdateOne {
	return pauo.SetRoleID(r.ID)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (pauo *ProjectAssignmentUpdateOne) SetCreatedByID(id uuid.UUID) *ProjectAssignmentUpdateOne {
	pauo.mutation.SetCreatedByID(id)
	return pauo
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (pauo *ProjectAssignmentUpdateOne) SetNillableCreatedByID(id *uuid.UUID) *ProjectAssignmentUpdateOne {
	if id != nil {
		pauo = pauo.SetCreatedByID(*id)
	}
	return pauo
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (pauo *ProjectAssignmentUpdateOne) SetCreatedBy(u *User) *ProjectAssignmentUpdateOne {
	return pauo.SetCreatedByID(u.ID)
}

// SetUpdatedByID sets the "updated_by" edge to the User entity by ID.
func (pauo *ProjectAssignmentUpdateOne) SetUpdatedByID(id uuid.UUID) *ProjectAssignmentUpdateOne {
	pauo.mutation.SetUpdatedByID(id)
	return pauo
}

// SetNillableUpdatedByID sets the "updated_by" edge to the User entity by ID if the given value is not nil.
func (pauo *ProjectAssignmentUpdateOne) SetNillableUpdatedByID(id *uuid.UUID) *ProjectAssignmentUpdateOne {
	if id != nil {
		pauo = pauo.SetUpdatedByID(*id)
	}
	return pauo
}

// SetUpdatedBy sets the "updated_by" edge to the User entity.
func (pauo *ProjectAssignmentUpdateOne) SetUpdatedBy(u *User) *ProjectAssignmentUpdateOne {
	return pauo.SetUpdatedByID(u.ID)
}

// Mutation returns the ProjectAssignmentMutation object of the builder.
func (pauo *ProjectAssignmentUpdateOne) Mutation() *ProjectAssignmentMutation {
	return pauo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pauo *ProjectAssignmentUpdateOne) ClearUser() *ProjectAssignmentUpdateOne {
	pauo.mutation.ClearUser()
	return pauo
}

// ClearProject clears the "project" edge to the Project entity.
func (pauo *ProjectAssignmentUpdateOne) ClearProject() *ProjectAssignmentUpdateOne {
	pauo.mutation.ClearProject()
	return pauo
}

// ClearRole clears the "role" edge to the Role entity.
func (pauo *ProjectAssignmentUpdateOne) ClearRole() *ProjectAssignmentUpdateOne {
	pauo.mutation.ClearRole()
	return pauo
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (pauo *ProjectAssignmentUpdateOne) ClearCreatedBy() *ProjectAssignmentUpdateOne {
	pauo.mutation.ClearCreatedBy()
	return pauo
}

// ClearUpdatedBy clears the "updated_by" edge to the User entity.
func (pauo *ProjectAssignmentUpdateOne) ClearUpdatedBy() *ProjectAssignmentUpdateOne {
	pauo.mutation.ClearUpdatedBy()
	return pauo
}

// Where appends a list predicates to the ProjectAssignmentUpdate builder.
func (pauo *ProjectAssignmentUpdateOne) Where(ps ...predicate.ProjectAssignment) *ProjectAssignmentUpdateOne {
	pauo.mutation.Where(ps...)
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *ProjectAssignmentUpdateOne) Select(field string, fields ...string) *ProjectAssignmentUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated ProjectAssignment entity.
func (pauo *ProjectAssignmentUpdateOne) Save(ctx context.Context) (*ProjectAssignment, error) {
	pauo.defaults()
	return withHooks(ctx, pauo.sqlSave, pauo.mutation, pauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *ProjectAssignmentUpdateOne) SaveX(ctx context.Context) *ProjectAssignment {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *ProjectAssignmentUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *ProjectAssignmentUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pauo *ProjectAssignmentUpdateOne) defaults() {
	if _, ok := pauo.mutation.UpdatedAt(); !ok {
		v := projectassignment.UpdateDefaultUpdatedAt()
		pauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pauo *ProjectAssignmentUpdateOne) check() error {
	if _, ok := pauo.mutation.UserID(); pauo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectAssignment.user"`)
	}
	if _, ok := pauo.mutation.ProjectID(); pauo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProjectAssignment.project"`)
	}
	return nil
}

func (pauo *ProjectAssignmentUpdateOne) sqlSave(ctx context.Context) (_node *ProjectAssignment, err error) {
	if err := pauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(projectassignment.Table, projectassignment.Columns, sqlgraph.NewFieldSpec(projectassignment.FieldID, field.TypeUUID))
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProjectAssignment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectassignment.FieldID)
		for _, f := range fields {
			if !projectassignment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != projectassignment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.UpdatedAt(); ok {
		_spec.SetField(projectassignment.FieldUpdatedAt, field.TypeTime, value)
	}
	if pauo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.UserTable,
			Columns: []string{projectassignment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.UserTable,
			Columns: []string{projectassignment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pauo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.ProjectTable,
			Columns: []string{projectassignment.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.ProjectTable,
			Columns: []string{projectassignment.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pauo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.RoleTable,
			Columns: []string{projectassignment.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.RoleTable,
			Columns: []string{projectassignment.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pauo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.CreatedByTable,
			Columns: []string{projectassignment.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.CreatedByTable,
			Columns: []string{projectassignment.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pauo.mutation.UpdatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.UpdatedByTable,
			Columns: []string{projectassignment.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.UpdatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   projectassignment.UpdatedByTable,
			Columns: []string{projectassignment.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProjectAssignment{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{projectassignment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pauo.mutation.done = true
	return _node, nil
}
