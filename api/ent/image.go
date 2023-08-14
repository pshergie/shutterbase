// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shutterbase/shutterbase/ent/batch"
	"github.com/shutterbase/shutterbase/ent/camera"
	"github.com/shutterbase/shutterbase/ent/image"
	"github.com/shutterbase/shutterbase/ent/project"
	"github.com/shutterbase/shutterbase/ent/user"
)

// Image is the model entity for the Image schema.
type Image struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// ThumbnailID holds the value of the "thumbnail_id" field.
	ThumbnailID uuid.UUID `json:"thumbnailId"`
	// FileName holds the value of the "file_name" field.
	FileName string `json:"fileName"`
	// ComputedFileName holds the value of the "computed_file_name" field.
	ComputedFileName string `json:"computedFileName"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ExifData holds the value of the "exif_data" field.
	ExifData map[string]interface{} `json:"exifData"`
	// CapturedAt holds the value of the "captured_at" field.
	CapturedAt time.Time `json:"capturedAt"`
	// CapturedAtCorrected holds the value of the "captured_at_corrected" field.
	CapturedAtCorrected time.Time `json:"capturedAtCorrected"`
	// InferredAt holds the value of the "inferred_at" field.
	InferredAt time.Time `json:"inferredAt"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageQuery when eager-loading is set.
	Edges            ImageEdges `json:"edges"`
	image_user       *uuid.UUID
	image_batch      *uuid.UUID
	image_project    *uuid.UUID
	image_camera     *uuid.UUID
	image_created_by *uuid.UUID
	image_updated_by *uuid.UUID
	selectValues     sql.SelectValues
}

// ImageEdges holds the relations/edges for other nodes in the graph.
type ImageEdges struct {
	// ImageTagAssignments holds the value of the image_tag_assignments edge.
	ImageTagAssignments []*ImageTagAssignment `json:"tagAssignments"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Batch holds the value of the batch edge.
	Batch *Batch `json:"batch,omitempty"`
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// Camera holds the value of the camera edge.
	Camera *Camera `json:"camera,omitempty"`
	// CreatedBy holds the value of the created_by edge.
	CreatedBy *User `json:"createdBy"`
	// UpdatedBy holds the value of the updated_by edge.
	UpdatedBy *User `json:"updatedBy"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// ImageTagAssignmentsOrErr returns the ImageTagAssignments value or an error if the edge
// was not loaded in eager-loading.
func (e ImageEdges) ImageTagAssignmentsOrErr() ([]*ImageTagAssignment, error) {
	if e.loadedTypes[0] {
		return e.ImageTagAssignments, nil
	}
	return nil, &NotLoadedError{edge: "image_tag_assignments"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BatchOrErr returns the Batch value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) BatchOrErr() (*Batch, error) {
	if e.loadedTypes[2] {
		if e.Batch == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: batch.Label}
		}
		return e.Batch, nil
	}
	return nil, &NotLoadedError{edge: "batch"}
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[3] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// CameraOrErr returns the Camera value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) CameraOrErr() (*Camera, error) {
	if e.loadedTypes[4] {
		if e.Camera == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: camera.Label}
		}
		return e.Camera, nil
	}
	return nil, &NotLoadedError{edge: "camera"}
}

// CreatedByOrErr returns the CreatedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) CreatedByOrErr() (*User, error) {
	if e.loadedTypes[5] {
		if e.CreatedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.CreatedBy, nil
	}
	return nil, &NotLoadedError{edge: "created_by"}
}

// UpdatedByOrErr returns the UpdatedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageEdges) UpdatedByOrErr() (*User, error) {
	if e.loadedTypes[6] {
		if e.UpdatedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UpdatedBy, nil
	}
	return nil, &NotLoadedError{edge: "updated_by"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Image) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case image.FieldExifData:
			values[i] = new([]byte)
		case image.FieldFileName, image.FieldComputedFileName, image.FieldDescription:
			values[i] = new(sql.NullString)
		case image.FieldCreatedAt, image.FieldUpdatedAt, image.FieldCapturedAt, image.FieldCapturedAtCorrected, image.FieldInferredAt:
			values[i] = new(sql.NullTime)
		case image.FieldID, image.FieldThumbnailID:
			values[i] = new(uuid.UUID)
		case image.ForeignKeys[0]: // image_user
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case image.ForeignKeys[1]: // image_batch
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case image.ForeignKeys[2]: // image_project
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case image.ForeignKeys[3]: // image_camera
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case image.ForeignKeys[4]: // image_created_by
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case image.ForeignKeys[5]: // image_updated_by
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Image fields.
func (i *Image) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case image.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case image.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case image.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case image.FieldThumbnailID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_id", values[j])
			} else if value != nil {
				i.ThumbnailID = *value
			}
		case image.FieldFileName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name", values[j])
			} else if value.Valid {
				i.FileName = value.String
			}
		case image.FieldComputedFileName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field computed_file_name", values[j])
			} else if value.Valid {
				i.ComputedFileName = value.String
			}
		case image.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		case image.FieldExifData:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field exif_data", values[j])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &i.ExifData); err != nil {
					return fmt.Errorf("unmarshal field exif_data: %w", err)
				}
			}
		case image.FieldCapturedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field captured_at", values[j])
			} else if value.Valid {
				i.CapturedAt = value.Time
			}
		case image.FieldCapturedAtCorrected:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field captured_at_corrected", values[j])
			} else if value.Valid {
				i.CapturedAtCorrected = value.Time
			}
		case image.FieldInferredAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field inferred_at", values[j])
			} else if value.Valid {
				i.InferredAt = value.Time
			}
		case image.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field image_user", values[j])
			} else if value.Valid {
				i.image_user = new(uuid.UUID)
				*i.image_user = *value.S.(*uuid.UUID)
			}
		case image.ForeignKeys[1]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field image_batch", values[j])
			} else if value.Valid {
				i.image_batch = new(uuid.UUID)
				*i.image_batch = *value.S.(*uuid.UUID)
			}
		case image.ForeignKeys[2]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field image_project", values[j])
			} else if value.Valid {
				i.image_project = new(uuid.UUID)
				*i.image_project = *value.S.(*uuid.UUID)
			}
		case image.ForeignKeys[3]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field image_camera", values[j])
			} else if value.Valid {
				i.image_camera = new(uuid.UUID)
				*i.image_camera = *value.S.(*uuid.UUID)
			}
		case image.ForeignKeys[4]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field image_created_by", values[j])
			} else if value.Valid {
				i.image_created_by = new(uuid.UUID)
				*i.image_created_by = *value.S.(*uuid.UUID)
			}
		case image.ForeignKeys[5]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field image_updated_by", values[j])
			} else if value.Valid {
				i.image_updated_by = new(uuid.UUID)
				*i.image_updated_by = *value.S.(*uuid.UUID)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Image.
// This includes values selected through modifiers, order, etc.
func (i *Image) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryImageTagAssignments queries the "image_tag_assignments" edge of the Image entity.
func (i *Image) QueryImageTagAssignments() *ImageTagAssignmentQuery {
	return NewImageClient(i.config).QueryImageTagAssignments(i)
}

// QueryUser queries the "user" edge of the Image entity.
func (i *Image) QueryUser() *UserQuery {
	return NewImageClient(i.config).QueryUser(i)
}

// QueryBatch queries the "batch" edge of the Image entity.
func (i *Image) QueryBatch() *BatchQuery {
	return NewImageClient(i.config).QueryBatch(i)
}

// QueryProject queries the "project" edge of the Image entity.
func (i *Image) QueryProject() *ProjectQuery {
	return NewImageClient(i.config).QueryProject(i)
}

// QueryCamera queries the "camera" edge of the Image entity.
func (i *Image) QueryCamera() *CameraQuery {
	return NewImageClient(i.config).QueryCamera(i)
}

// QueryCreatedBy queries the "created_by" edge of the Image entity.
func (i *Image) QueryCreatedBy() *UserQuery {
	return NewImageClient(i.config).QueryCreatedBy(i)
}

// QueryUpdatedBy queries the "updated_by" edge of the Image entity.
func (i *Image) QueryUpdatedBy() *UserQuery {
	return NewImageClient(i.config).QueryUpdatedBy(i)
}

// Update returns a builder for updating this Image.
// Note that you need to call Image.Unwrap() before calling this method if this Image
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Image) Update() *ImageUpdateOne {
	return NewImageClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Image entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Image) Unwrap() *Image {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Image is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Image) String() string {
	var builder strings.Builder
	builder.WriteString("Image(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("thumbnail_id=")
	builder.WriteString(fmt.Sprintf("%v", i.ThumbnailID))
	builder.WriteString(", ")
	builder.WriteString("file_name=")
	builder.WriteString(i.FileName)
	builder.WriteString(", ")
	builder.WriteString("computed_file_name=")
	builder.WriteString(i.ComputedFileName)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(i.Description)
	builder.WriteString(", ")
	builder.WriteString("exif_data=")
	builder.WriteString(fmt.Sprintf("%v", i.ExifData))
	builder.WriteString(", ")
	builder.WriteString("captured_at=")
	builder.WriteString(i.CapturedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("captured_at_corrected=")
	builder.WriteString(i.CapturedAtCorrected.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("inferred_at=")
	builder.WriteString(i.InferredAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Images is a parsable slice of Image.
type Images []*Image
