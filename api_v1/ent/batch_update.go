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
	"github.com/shutterbase/shutterbase/ent/batch"
	"github.com/shutterbase/shutterbase/ent/image"
	"github.com/shutterbase/shutterbase/ent/predicate"
	"github.com/shutterbase/shutterbase/ent/project"
	"github.com/shutterbase/shutterbase/ent/user"
)

// BatchUpdate is the builder for updating Batch entities.
type BatchUpdate struct {
	config
	hooks    []Hook
	mutation *BatchMutation
}

// Where appends a list predicates to the BatchUpdate builder.
func (bu *BatchUpdate) Where(ps ...predicate.Batch) *BatchUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BatchUpdate) SetUpdatedAt(t time.Time) *BatchUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetName sets the "name" field.
func (bu *BatchUpdate) SetName(s string) *BatchUpdate {
	bu.mutation.SetName(s)
	return bu
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (bu *BatchUpdate) AddImageIDs(ids ...uuid.UUID) *BatchUpdate {
	bu.mutation.AddImageIDs(ids...)
	return bu
}

// AddImages adds the "images" edges to the Image entity.
func (bu *BatchUpdate) AddImages(i ...*Image) *BatchUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return bu.AddImageIDs(ids...)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (bu *BatchUpdate) SetProjectID(id uuid.UUID) *BatchUpdate {
	bu.mutation.SetProjectID(id)
	return bu
}

// SetProject sets the "project" edge to the Project entity.
func (bu *BatchUpdate) SetProject(p *Project) *BatchUpdate {
	return bu.SetProjectID(p.ID)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (bu *BatchUpdate) SetCreatedByID(id uuid.UUID) *BatchUpdate {
	bu.mutation.SetCreatedByID(id)
	return bu
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (bu *BatchUpdate) SetNillableCreatedByID(id *uuid.UUID) *BatchUpdate {
	if id != nil {
		bu = bu.SetCreatedByID(*id)
	}
	return bu
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (bu *BatchUpdate) SetCreatedBy(u *User) *BatchUpdate {
	return bu.SetCreatedByID(u.ID)
}

// SetUpdatedByID sets the "updated_by" edge to the User entity by ID.
func (bu *BatchUpdate) SetUpdatedByID(id uuid.UUID) *BatchUpdate {
	bu.mutation.SetUpdatedByID(id)
	return bu
}

// SetNillableUpdatedByID sets the "updated_by" edge to the User entity by ID if the given value is not nil.
func (bu *BatchUpdate) SetNillableUpdatedByID(id *uuid.UUID) *BatchUpdate {
	if id != nil {
		bu = bu.SetUpdatedByID(*id)
	}
	return bu
}

// SetUpdatedBy sets the "updated_by" edge to the User entity.
func (bu *BatchUpdate) SetUpdatedBy(u *User) *BatchUpdate {
	return bu.SetUpdatedByID(u.ID)
}

// Mutation returns the BatchMutation object of the builder.
func (bu *BatchUpdate) Mutation() *BatchMutation {
	return bu.mutation
}

// ClearImages clears all "images" edges to the Image entity.
func (bu *BatchUpdate) ClearImages() *BatchUpdate {
	bu.mutation.ClearImages()
	return bu
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (bu *BatchUpdate) RemoveImageIDs(ids ...uuid.UUID) *BatchUpdate {
	bu.mutation.RemoveImageIDs(ids...)
	return bu
}

// RemoveImages removes "images" edges to Image entities.
func (bu *BatchUpdate) RemoveImages(i ...*Image) *BatchUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return bu.RemoveImageIDs(ids...)
}

// ClearProject clears the "project" edge to the Project entity.
func (bu *BatchUpdate) ClearProject() *BatchUpdate {
	bu.mutation.ClearProject()
	return bu
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (bu *BatchUpdate) ClearCreatedBy() *BatchUpdate {
	bu.mutation.ClearCreatedBy()
	return bu
}

// ClearUpdatedBy clears the "updated_by" edge to the User entity.
func (bu *BatchUpdate) ClearUpdatedBy() *BatchUpdate {
	bu.mutation.ClearUpdatedBy()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BatchUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BatchUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BatchUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BatchUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BatchUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := batch.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BatchUpdate) check() error {
	if v, ok := bu.mutation.Name(); ok {
		if err := batch.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Batch.name": %w`, err)}
		}
	}
	if _, ok := bu.mutation.ProjectID(); bu.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Batch.project"`)
	}
	return nil
}

func (bu *BatchUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(batch.Table, batch.Columns, sqlgraph.NewFieldSpec(batch.FieldID, field.TypeUUID))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(batch.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(batch.FieldName, field.TypeString, value)
	}
	if bu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   batch.ImagesTable,
			Columns: []string{batch.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedImagesIDs(); len(nodes) > 0 && !bu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   batch.ImagesTable,
			Columns: []string{batch.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   batch.ImagesTable,
			Columns: []string{batch.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.ProjectTable,
			Columns: []string{batch.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.ProjectTable,
			Columns: []string{batch.ProjectColumn},
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
	if bu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.CreatedByTable,
			Columns: []string{batch.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.CreatedByTable,
			Columns: []string{batch.CreatedByColumn},
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
	if bu.mutation.UpdatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.UpdatedByTable,
			Columns: []string{batch.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.UpdatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.UpdatedByTable,
			Columns: []string{batch.UpdatedByColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{batch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BatchUpdateOne is the builder for updating a single Batch entity.
type BatchUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BatchMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BatchUpdateOne) SetUpdatedAt(t time.Time) *BatchUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetName sets the "name" field.
func (buo *BatchUpdateOne) SetName(s string) *BatchUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (buo *BatchUpdateOne) AddImageIDs(ids ...uuid.UUID) *BatchUpdateOne {
	buo.mutation.AddImageIDs(ids...)
	return buo
}

// AddImages adds the "images" edges to the Image entity.
func (buo *BatchUpdateOne) AddImages(i ...*Image) *BatchUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return buo.AddImageIDs(ids...)
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (buo *BatchUpdateOne) SetProjectID(id uuid.UUID) *BatchUpdateOne {
	buo.mutation.SetProjectID(id)
	return buo
}

// SetProject sets the "project" edge to the Project entity.
func (buo *BatchUpdateOne) SetProject(p *Project) *BatchUpdateOne {
	return buo.SetProjectID(p.ID)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (buo *BatchUpdateOne) SetCreatedByID(id uuid.UUID) *BatchUpdateOne {
	buo.mutation.SetCreatedByID(id)
	return buo
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (buo *BatchUpdateOne) SetNillableCreatedByID(id *uuid.UUID) *BatchUpdateOne {
	if id != nil {
		buo = buo.SetCreatedByID(*id)
	}
	return buo
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (buo *BatchUpdateOne) SetCreatedBy(u *User) *BatchUpdateOne {
	return buo.SetCreatedByID(u.ID)
}

// SetUpdatedByID sets the "updated_by" edge to the User entity by ID.
func (buo *BatchUpdateOne) SetUpdatedByID(id uuid.UUID) *BatchUpdateOne {
	buo.mutation.SetUpdatedByID(id)
	return buo
}

// SetNillableUpdatedByID sets the "updated_by" edge to the User entity by ID if the given value is not nil.
func (buo *BatchUpdateOne) SetNillableUpdatedByID(id *uuid.UUID) *BatchUpdateOne {
	if id != nil {
		buo = buo.SetUpdatedByID(*id)
	}
	return buo
}

// SetUpdatedBy sets the "updated_by" edge to the User entity.
func (buo *BatchUpdateOne) SetUpdatedBy(u *User) *BatchUpdateOne {
	return buo.SetUpdatedByID(u.ID)
}

// Mutation returns the BatchMutation object of the builder.
func (buo *BatchUpdateOne) Mutation() *BatchMutation {
	return buo.mutation
}

// ClearImages clears all "images" edges to the Image entity.
func (buo *BatchUpdateOne) ClearImages() *BatchUpdateOne {
	buo.mutation.ClearImages()
	return buo
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (buo *BatchUpdateOne) RemoveImageIDs(ids ...uuid.UUID) *BatchUpdateOne {
	buo.mutation.RemoveImageIDs(ids...)
	return buo
}

// RemoveImages removes "images" edges to Image entities.
func (buo *BatchUpdateOne) RemoveImages(i ...*Image) *BatchUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return buo.RemoveImageIDs(ids...)
}

// ClearProject clears the "project" edge to the Project entity.
func (buo *BatchUpdateOne) ClearProject() *BatchUpdateOne {
	buo.mutation.ClearProject()
	return buo
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (buo *BatchUpdateOne) ClearCreatedBy() *BatchUpdateOne {
	buo.mutation.ClearCreatedBy()
	return buo
}

// ClearUpdatedBy clears the "updated_by" edge to the User entity.
func (buo *BatchUpdateOne) ClearUpdatedBy() *BatchUpdateOne {
	buo.mutation.ClearUpdatedBy()
	return buo
}

// Where appends a list predicates to the BatchUpdate builder.
func (buo *BatchUpdateOne) Where(ps ...predicate.Batch) *BatchUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BatchUpdateOne) Select(field string, fields ...string) *BatchUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Batch entity.
func (buo *BatchUpdateOne) Save(ctx context.Context) (*Batch, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BatchUpdateOne) SaveX(ctx context.Context) *Batch {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BatchUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BatchUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BatchUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := batch.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BatchUpdateOne) check() error {
	if v, ok := buo.mutation.Name(); ok {
		if err := batch.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Batch.name": %w`, err)}
		}
	}
	if _, ok := buo.mutation.ProjectID(); buo.mutation.ProjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Batch.project"`)
	}
	return nil
}

func (buo *BatchUpdateOne) sqlSave(ctx context.Context) (_node *Batch, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(batch.Table, batch.Columns, sqlgraph.NewFieldSpec(batch.FieldID, field.TypeUUID))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Batch.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, batch.FieldID)
		for _, f := range fields {
			if !batch.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != batch.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(batch.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(batch.FieldName, field.TypeString, value)
	}
	if buo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   batch.ImagesTable,
			Columns: []string{batch.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedImagesIDs(); len(nodes) > 0 && !buo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   batch.ImagesTable,
			Columns: []string{batch.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   batch.ImagesTable,
			Columns: []string{batch.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.ProjectTable,
			Columns: []string{batch.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.ProjectTable,
			Columns: []string{batch.ProjectColumn},
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
	if buo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.CreatedByTable,
			Columns: []string{batch.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.CreatedByTable,
			Columns: []string{batch.CreatedByColumn},
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
	if buo.mutation.UpdatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.UpdatedByTable,
			Columns: []string{batch.UpdatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.UpdatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   batch.UpdatedByTable,
			Columns: []string{batch.UpdatedByColumn},
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
	_node = &Batch{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{batch.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}