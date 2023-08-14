// Code generated by ent, DO NOT EDIT.

package image

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldThumbnailID holds the string denoting the thumbnail_id field in the database.
	FieldThumbnailID = "thumbnail_id"
	// FieldFileName holds the string denoting the file_name field in the database.
	FieldFileName = "file_name"
	// FieldComputedFileName holds the string denoting the computed_file_name field in the database.
	FieldComputedFileName = "computed_file_name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldExifData holds the string denoting the exif_data field in the database.
	FieldExifData = "exif_data"
	// FieldCapturedAt holds the string denoting the captured_at field in the database.
	FieldCapturedAt = "captured_at"
	// FieldCapturedAtCorrected holds the string denoting the captured_at_corrected field in the database.
	FieldCapturedAtCorrected = "captured_at_corrected"
	// FieldInferredAt holds the string denoting the inferred_at field in the database.
	FieldInferredAt = "inferred_at"
	// EdgeImageTagAssignments holds the string denoting the image_tag_assignments edge name in mutations.
	EdgeImageTagAssignments = "image_tag_assignments"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeBatch holds the string denoting the batch edge name in mutations.
	EdgeBatch = "batch"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeCamera holds the string denoting the camera edge name in mutations.
	EdgeCamera = "camera"
	// EdgeCreatedBy holds the string denoting the created_by edge name in mutations.
	EdgeCreatedBy = "created_by"
	// EdgeUpdatedBy holds the string denoting the updated_by edge name in mutations.
	EdgeUpdatedBy = "updated_by"
	// Table holds the table name of the image in the database.
	Table = "images"
	// ImageTagAssignmentsTable is the table that holds the image_tag_assignments relation/edge.
	ImageTagAssignmentsTable = "image_tag_assignments"
	// ImageTagAssignmentsInverseTable is the table name for the ImageTagAssignment entity.
	// It exists in this package in order to avoid circular dependency with the "imagetagassignment" package.
	ImageTagAssignmentsInverseTable = "image_tag_assignments"
	// ImageTagAssignmentsColumn is the table column denoting the image_tag_assignments relation/edge.
	ImageTagAssignmentsColumn = "image_tag_assignment_image"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "images"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "image_user"
	// BatchTable is the table that holds the batch relation/edge.
	BatchTable = "images"
	// BatchInverseTable is the table name for the Batch entity.
	// It exists in this package in order to avoid circular dependency with the "batch" package.
	BatchInverseTable = "batches"
	// BatchColumn is the table column denoting the batch relation/edge.
	BatchColumn = "image_batch"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "images"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "image_project"
	// CameraTable is the table that holds the camera relation/edge.
	CameraTable = "images"
	// CameraInverseTable is the table name for the Camera entity.
	// It exists in this package in order to avoid circular dependency with the "camera" package.
	CameraInverseTable = "cameras"
	// CameraColumn is the table column denoting the camera relation/edge.
	CameraColumn = "image_camera"
	// CreatedByTable is the table that holds the created_by relation/edge.
	CreatedByTable = "images"
	// CreatedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatedByInverseTable = "users"
	// CreatedByColumn is the table column denoting the created_by relation/edge.
	CreatedByColumn = "image_created_by"
	// UpdatedByTable is the table that holds the updated_by relation/edge.
	UpdatedByTable = "images"
	// UpdatedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UpdatedByInverseTable = "users"
	// UpdatedByColumn is the table column denoting the updated_by relation/edge.
	UpdatedByColumn = "image_updated_by"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldThumbnailID,
	FieldFileName,
	FieldComputedFileName,
	FieldDescription,
	FieldExifData,
	FieldCapturedAt,
	FieldCapturedAtCorrected,
	FieldInferredAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"image_user",
	"image_batch",
	"image_project",
	"image_camera",
	"image_created_by",
	"image_updated_by",
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
	// FileNameValidator is a validator for the "file_name" field. It is called by the builders before save.
	FileNameValidator func(string) error
	// DefaultComputedFileName holds the default value on creation for the "computed_file_name" field.
	DefaultComputedFileName string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultExifData holds the default value on creation for the "exif_data" field.
	DefaultExifData map[string]interface{}
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Image queries.
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

// ByThumbnailID orders the results by the thumbnail_id field.
func ByThumbnailID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThumbnailID, opts...).ToFunc()
}

// ByFileName orders the results by the file_name field.
func ByFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileName, opts...).ToFunc()
}

// ByComputedFileName orders the results by the computed_file_name field.
func ByComputedFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComputedFileName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCapturedAt orders the results by the captured_at field.
func ByCapturedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCapturedAt, opts...).ToFunc()
}

// ByCapturedAtCorrected orders the results by the captured_at_corrected field.
func ByCapturedAtCorrected(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCapturedAtCorrected, opts...).ToFunc()
}

// ByInferredAt orders the results by the inferred_at field.
func ByInferredAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInferredAt, opts...).ToFunc()
}

// ByImageTagAssignmentsCount orders the results by image_tag_assignments count.
func ByImageTagAssignmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImageTagAssignmentsStep(), opts...)
	}
}

// ByImageTagAssignments orders the results by image_tag_assignments terms.
func ByImageTagAssignments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImageTagAssignmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByBatchField orders the results by batch field.
func ByBatchField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBatchStep(), sql.OrderByField(field, opts...))
	}
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByCameraField orders the results by camera field.
func ByCameraField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCameraStep(), sql.OrderByField(field, opts...))
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
func newImageTagAssignmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImageTagAssignmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ImageTagAssignmentsTable, ImageTagAssignmentsColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
func newBatchStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BatchInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BatchTable, BatchColumn),
	)
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProjectTable, ProjectColumn),
	)
}
func newCameraStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CameraInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CameraTable, CameraColumn),
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
