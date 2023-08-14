// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCopyright holds the string denoting the copyright field in the database.
	FieldCopyright = "copyright"
	// FieldCopyrightReference holds the string denoting the copyright_reference field in the database.
	FieldCopyrightReference = "copyright_reference"
	// FieldLocationName holds the string denoting the location_name field in the database.
	FieldLocationName = "location_name"
	// FieldLocationCode holds the string denoting the location_code field in the database.
	FieldLocationCode = "location_code"
	// FieldLocationCity holds the string denoting the location_city field in the database.
	FieldLocationCity = "location_city"
	// EdgeAssignments holds the string denoting the assignments edge name in mutations.
	EdgeAssignments = "assignments"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// EdgeBatches holds the string denoting the batches edge name in mutations.
	EdgeBatches = "batches"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeCreatedBy holds the string denoting the created_by edge name in mutations.
	EdgeCreatedBy = "created_by"
	// EdgeUpdatedBy holds the string denoting the updated_by edge name in mutations.
	EdgeUpdatedBy = "updated_by"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// AssignmentsTable is the table that holds the assignments relation/edge.
	AssignmentsTable = "project_assignments"
	// AssignmentsInverseTable is the table name for the ProjectAssignment entity.
	// It exists in this package in order to avoid circular dependency with the "projectassignment" package.
	AssignmentsInverseTable = "project_assignments"
	// AssignmentsColumn is the table column denoting the assignments relation/edge.
	AssignmentsColumn = "project_assignment_project"
	// ImagesTable is the table that holds the images relation/edge.
	ImagesTable = "images"
	// ImagesInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImagesInverseTable = "images"
	// ImagesColumn is the table column denoting the images relation/edge.
	ImagesColumn = "image_project"
	// BatchesTable is the table that holds the batches relation/edge.
	BatchesTable = "batches"
	// BatchesInverseTable is the table name for the Batch entity.
	// It exists in this package in order to avoid circular dependency with the "batch" package.
	BatchesInverseTable = "batches"
	// BatchesColumn is the table column denoting the batches relation/edge.
	BatchesColumn = "batch_project"
	// TagsTable is the table that holds the tags relation/edge.
	TagsTable = "image_tags"
	// TagsInverseTable is the table name for the ImageTag entity.
	// It exists in this package in order to avoid circular dependency with the "imagetag" package.
	TagsInverseTable = "image_tags"
	// TagsColumn is the table column denoting the tags relation/edge.
	TagsColumn = "image_tag_project"
	// CreatedByTable is the table that holds the created_by relation/edge.
	CreatedByTable = "projects"
	// CreatedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatedByInverseTable = "users"
	// CreatedByColumn is the table column denoting the created_by relation/edge.
	CreatedByColumn = "project_created_by"
	// UpdatedByTable is the table that holds the updated_by relation/edge.
	UpdatedByTable = "projects"
	// UpdatedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UpdatedByInverseTable = "users"
	// UpdatedByColumn is the table column denoting the updated_by relation/edge.
	UpdatedByColumn = "project_updated_by"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
	FieldCopyright,
	FieldCopyrightReference,
	FieldLocationName,
	FieldLocationCode,
	FieldLocationCity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "projects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_created_by",
	"project_updated_by",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultCopyright holds the default value on creation for the "copyright" field.
	DefaultCopyright string
	// DefaultCopyrightReference holds the default value on creation for the "copyright_reference" field.
	DefaultCopyrightReference string
	// DefaultLocationName holds the default value on creation for the "location_name" field.
	DefaultLocationName string
	// DefaultLocationCode holds the default value on creation for the "location_code" field.
	DefaultLocationCode string
	// DefaultLocationCity holds the default value on creation for the "location_city" field.
	DefaultLocationCity string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Project queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCopyright orders the results by the copyright field.
func ByCopyright(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCopyright, opts...).ToFunc()
}

// ByCopyrightReference orders the results by the copyright_reference field.
func ByCopyrightReference(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCopyrightReference, opts...).ToFunc()
}

// ByLocationName orders the results by the location_name field.
func ByLocationName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationName, opts...).ToFunc()
}

// ByLocationCode orders the results by the location_code field.
func ByLocationCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationCode, opts...).ToFunc()
}

// ByLocationCity orders the results by the location_city field.
func ByLocationCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationCity, opts...).ToFunc()
}

// ByAssignmentsCount orders the results by assignments count.
func ByAssignmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAssignmentsStep(), opts...)
	}
}

// ByAssignments orders the results by assignments terms.
func ByAssignments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAssignmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByImagesCount orders the results by images count.
func ByImagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImagesStep(), opts...)
	}
}

// ByImages orders the results by images terms.
func ByImages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBatchesCount orders the results by batches count.
func ByBatchesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBatchesStep(), opts...)
	}
}

// ByBatches orders the results by batches terms.
func ByBatches(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBatchesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedByField orders the results by created_by field.
func ByCreatedByField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedByStep(), sql.OrderByField(field, opts...))
	}
}

// ByUpdatedByField orders the results by updated_by field.
func ByUpdatedByField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUpdatedByStep(), sql.OrderByField(field, opts...))
	}
}
func newAssignmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AssignmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AssignmentsTable, AssignmentsColumn),
	)
}
func newImagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ImagesTable, ImagesColumn),
	)
}
func newBatchesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BatchesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, BatchesTable, BatchesColumn),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, TagsTable, TagsColumn),
	)
}
func newCreatedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CreatedByTable, CreatedByColumn),
	)
}
func newUpdatedByStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UpdatedByInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UpdatedByTable, UpdatedByColumn),
	)
}
