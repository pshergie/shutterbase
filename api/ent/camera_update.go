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
	"github.com/shutterbase/shutterbase/ent/camera"
	"github.com/shutterbase/shutterbase/ent/image"
	"github.com/shutterbase/shutterbase/ent/predicate"
	"github.com/shutterbase/shutterbase/ent/timeoffset"
	"github.com/shutterbase/shutterbase/ent/user"
)

// CameraUpdate is the builder for updating Camera entities.
type CameraUpdate struct {
	config
	hooks    []Hook
	mutation *CameraMutation
}

// Where appends a list predicates to the CameraUpdate builder.
func (cu *CameraUpdate) Where(ps ...predicate.Camera) *CameraUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CameraUpdate) SetUpdatedAt(t time.Time) *CameraUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetName sets the "name" field.
func (cu *CameraUpdate) SetName(s string) *CameraUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CameraUpdate) SetDescription(s string) *CameraUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// AddTimeOffsetIDs adds the "timeOffsets" edge to the TimeOffset entity by IDs.
func (cu *CameraUpdate) AddTimeOffsetIDs(ids ...uuid.UUID) *CameraUpdate {
	cu.mutation.AddTimeOffsetIDs(ids...)
	return cu
}

// AddTimeOffsets adds the "timeOffsets" edges to the TimeOffset entity.
func (cu *CameraUpdate) AddTimeOffsets(t ...*TimeOffset) *CameraUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.AddTimeOffsetIDs(ids...)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (cu *CameraUpdate) AddImageIDs(ids ...uuid.UUID) *CameraUpdate {
	cu.mutation.AddImageIDs(ids...)
	return cu
}

// AddImages adds the "images" edges to the Image entity.
func (cu *CameraUpdate) AddImages(i ...*Image) *CameraUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.AddImageIDs(ids...)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (cu *CameraUpdate) SetCreatedByID(id uuid.UUID) *CameraUpdate {
	cu.mutation.SetCreatedByID(id)
	return cu
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (cu *CameraUpdate) SetNillableCreatedByID(id *uuid.UUID) *CameraUpdate {
	if id != nil {
		cu = cu.SetCreatedByID(*id)
	}
	return cu
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (cu *CameraUpdate) SetCreatedBy(u *User) *CameraUpdate {
	return cu.SetCreatedByID(u.ID)
}

// SetModifiedByID sets the "modified_by" edge to the User entity by ID.
func (cu *CameraUpdate) SetModifiedByID(id uuid.UUID) *CameraUpdate {
	cu.mutation.SetModifiedByID(id)
	return cu
}

// SetNillableModifiedByID sets the "modified_by" edge to the User entity by ID if the given value is not nil.
func (cu *CameraUpdate) SetNillableModifiedByID(id *uuid.UUID) *CameraUpdate {
	if id != nil {
		cu = cu.SetModifiedByID(*id)
	}
	return cu
}

// SetModifiedBy sets the "modified_by" edge to the User entity.
func (cu *CameraUpdate) SetModifiedBy(u *User) *CameraUpdate {
	return cu.SetModifiedByID(u.ID)
}

// Mutation returns the CameraMutation object of the builder.
func (cu *CameraUpdate) Mutation() *CameraMutation {
	return cu.mutation
}

// ClearTimeOffsets clears all "timeOffsets" edges to the TimeOffset entity.
func (cu *CameraUpdate) ClearTimeOffsets() *CameraUpdate {
	cu.mutation.ClearTimeOffsets()
	return cu
}

// RemoveTimeOffsetIDs removes the "timeOffsets" edge to TimeOffset entities by IDs.
func (cu *CameraUpdate) RemoveTimeOffsetIDs(ids ...uuid.UUID) *CameraUpdate {
	cu.mutation.RemoveTimeOffsetIDs(ids...)
	return cu
}

// RemoveTimeOffsets removes "timeOffsets" edges to TimeOffset entities.
func (cu *CameraUpdate) RemoveTimeOffsets(t ...*TimeOffset) *CameraUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cu.RemoveTimeOffsetIDs(ids...)
}

// ClearImages clears all "images" edges to the Image entity.
func (cu *CameraUpdate) ClearImages() *CameraUpdate {
	cu.mutation.ClearImages()
	return cu
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (cu *CameraUpdate) RemoveImageIDs(ids ...uuid.UUID) *CameraUpdate {
	cu.mutation.RemoveImageIDs(ids...)
	return cu
}

// RemoveImages removes "images" edges to Image entities.
func (cu *CameraUpdate) RemoveImages(i ...*Image) *CameraUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.RemoveImageIDs(ids...)
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (cu *CameraUpdate) ClearCreatedBy() *CameraUpdate {
	cu.mutation.ClearCreatedBy()
	return cu
}

// ClearModifiedBy clears the "modified_by" edge to the User entity.
func (cu *CameraUpdate) ClearModifiedBy() *CameraUpdate {
	cu.mutation.ClearModifiedBy()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CameraUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CameraUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CameraUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CameraUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CameraUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := camera.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CameraUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := camera.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Camera.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Description(); ok {
		if err := camera.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Camera.description": %w`, err)}
		}
	}
	return nil
}

func (cu *CameraUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(camera.Table, camera.Columns, sqlgraph.NewFieldSpec(camera.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(camera.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(camera.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(camera.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.TimeOffsetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.TimeOffsetsTable,
			Columns: []string{camera.TimeOffsetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timeoffset.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedTimeOffsetsIDs(); len(nodes) > 0 && !cu.mutation.TimeOffsetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.TimeOffsetsTable,
			Columns: []string{camera.TimeOffsetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timeoffset.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TimeOffsetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.TimeOffsetsTable,
			Columns: []string{camera.TimeOffsetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timeoffset.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.ImagesTable,
			Columns: []string{camera.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedImagesIDs(); len(nodes) > 0 && !cu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.ImagesTable,
			Columns: []string{camera.ImagesColumn},
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
	if nodes := cu.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.ImagesTable,
			Columns: []string{camera.ImagesColumn},
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
	if cu.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   camera.CreatedByTable,
			Columns: []string{camera.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   camera.CreatedByTable,
			Columns: []string{camera.CreatedByColumn},
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
	if cu.mutation.ModifiedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   camera.ModifiedByTable,
			Columns: []string{camera.ModifiedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ModifiedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   camera.ModifiedByTable,
			Columns: []string{camera.ModifiedByColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{camera.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CameraUpdateOne is the builder for updating a single Camera entity.
type CameraUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CameraMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CameraUpdateOne) SetUpdatedAt(t time.Time) *CameraUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CameraUpdateOne) SetName(s string) *CameraUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CameraUpdateOne) SetDescription(s string) *CameraUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// AddTimeOffsetIDs adds the "timeOffsets" edge to the TimeOffset entity by IDs.
func (cuo *CameraUpdateOne) AddTimeOffsetIDs(ids ...uuid.UUID) *CameraUpdateOne {
	cuo.mutation.AddTimeOffsetIDs(ids...)
	return cuo
}

// AddTimeOffsets adds the "timeOffsets" edges to the TimeOffset entity.
func (cuo *CameraUpdateOne) AddTimeOffsets(t ...*TimeOffset) *CameraUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.AddTimeOffsetIDs(ids...)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (cuo *CameraUpdateOne) AddImageIDs(ids ...uuid.UUID) *CameraUpdateOne {
	cuo.mutation.AddImageIDs(ids...)
	return cuo
}

// AddImages adds the "images" edges to the Image entity.
func (cuo *CameraUpdateOne) AddImages(i ...*Image) *CameraUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.AddImageIDs(ids...)
}

// SetCreatedByID sets the "created_by" edge to the User entity by ID.
func (cuo *CameraUpdateOne) SetCreatedByID(id uuid.UUID) *CameraUpdateOne {
	cuo.mutation.SetCreatedByID(id)
	return cuo
}

// SetNillableCreatedByID sets the "created_by" edge to the User entity by ID if the given value is not nil.
func (cuo *CameraUpdateOne) SetNillableCreatedByID(id *uuid.UUID) *CameraUpdateOne {
	if id != nil {
		cuo = cuo.SetCreatedByID(*id)
	}
	return cuo
}

// SetCreatedBy sets the "created_by" edge to the User entity.
func (cuo *CameraUpdateOne) SetCreatedBy(u *User) *CameraUpdateOne {
	return cuo.SetCreatedByID(u.ID)
}

// SetModifiedByID sets the "modified_by" edge to the User entity by ID.
func (cuo *CameraUpdateOne) SetModifiedByID(id uuid.UUID) *CameraUpdateOne {
	cuo.mutation.SetModifiedByID(id)
	return cuo
}

// SetNillableModifiedByID sets the "modified_by" edge to the User entity by ID if the given value is not nil.
func (cuo *CameraUpdateOne) SetNillableModifiedByID(id *uuid.UUID) *CameraUpdateOne {
	if id != nil {
		cuo = cuo.SetModifiedByID(*id)
	}
	return cuo
}

// SetModifiedBy sets the "modified_by" edge to the User entity.
func (cuo *CameraUpdateOne) SetModifiedBy(u *User) *CameraUpdateOne {
	return cuo.SetModifiedByID(u.ID)
}

// Mutation returns the CameraMutation object of the builder.
func (cuo *CameraUpdateOne) Mutation() *CameraMutation {
	return cuo.mutation
}

// ClearTimeOffsets clears all "timeOffsets" edges to the TimeOffset entity.
func (cuo *CameraUpdateOne) ClearTimeOffsets() *CameraUpdateOne {
	cuo.mutation.ClearTimeOffsets()
	return cuo
}

// RemoveTimeOffsetIDs removes the "timeOffsets" edge to TimeOffset entities by IDs.
func (cuo *CameraUpdateOne) RemoveTimeOffsetIDs(ids ...uuid.UUID) *CameraUpdateOne {
	cuo.mutation.RemoveTimeOffsetIDs(ids...)
	return cuo
}

// RemoveTimeOffsets removes "timeOffsets" edges to TimeOffset entities.
func (cuo *CameraUpdateOne) RemoveTimeOffsets(t ...*TimeOffset) *CameraUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cuo.RemoveTimeOffsetIDs(ids...)
}

// ClearImages clears all "images" edges to the Image entity.
func (cuo *CameraUpdateOne) ClearImages() *CameraUpdateOne {
	cuo.mutation.ClearImages()
	return cuo
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (cuo *CameraUpdateOne) RemoveImageIDs(ids ...uuid.UUID) *CameraUpdateOne {
	cuo.mutation.RemoveImageIDs(ids...)
	return cuo
}

// RemoveImages removes "images" edges to Image entities.
func (cuo *CameraUpdateOne) RemoveImages(i ...*Image) *CameraUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.RemoveImageIDs(ids...)
}

// ClearCreatedBy clears the "created_by" edge to the User entity.
func (cuo *CameraUpdateOne) ClearCreatedBy() *CameraUpdateOne {
	cuo.mutation.ClearCreatedBy()
	return cuo
}

// ClearModifiedBy clears the "modified_by" edge to the User entity.
func (cuo *CameraUpdateOne) ClearModifiedBy() *CameraUpdateOne {
	cuo.mutation.ClearModifiedBy()
	return cuo
}

// Where appends a list predicates to the CameraUpdate builder.
func (cuo *CameraUpdateOne) Where(ps ...predicate.Camera) *CameraUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CameraUpdateOne) Select(field string, fields ...string) *CameraUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Camera entity.
func (cuo *CameraUpdateOne) Save(ctx context.Context) (*Camera, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CameraUpdateOne) SaveX(ctx context.Context) *Camera {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CameraUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CameraUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CameraUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := camera.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CameraUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := camera.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Camera.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Description(); ok {
		if err := camera.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Camera.description": %w`, err)}
		}
	}
	return nil
}

func (cuo *CameraUpdateOne) sqlSave(ctx context.Context) (_node *Camera, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(camera.Table, camera.Columns, sqlgraph.NewFieldSpec(camera.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Camera.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, camera.FieldID)
		for _, f := range fields {
			if !camera.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != camera.FieldID {
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
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(camera.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(camera.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(camera.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.TimeOffsetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.TimeOffsetsTable,
			Columns: []string{camera.TimeOffsetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timeoffset.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedTimeOffsetsIDs(); len(nodes) > 0 && !cuo.mutation.TimeOffsetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.TimeOffsetsTable,
			Columns: []string{camera.TimeOffsetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timeoffset.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TimeOffsetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.TimeOffsetsTable,
			Columns: []string{camera.TimeOffsetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timeoffset.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.ImagesTable,
			Columns: []string{camera.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedImagesIDs(); len(nodes) > 0 && !cuo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.ImagesTable,
			Columns: []string{camera.ImagesColumn},
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
	if nodes := cuo.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   camera.ImagesTable,
			Columns: []string{camera.ImagesColumn},
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
	if cuo.mutation.CreatedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   camera.CreatedByTable,
			Columns: []string{camera.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   camera.CreatedByTable,
			Columns: []string{camera.CreatedByColumn},
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
	if cuo.mutation.ModifiedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   camera.ModifiedByTable,
			Columns: []string{camera.ModifiedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ModifiedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   camera.ModifiedByTable,
			Columns: []string{camera.ModifiedByColumn},
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
	_node = &Camera{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{camera.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}