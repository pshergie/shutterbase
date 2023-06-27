// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/shutterbase/shutterbase/ent/project"
	"github.com/shutterbase/shutterbase/ent/user"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges               ProjectEdges `json:"edges"`
	project_created_by  *uuid.UUID
	project_modified_by *uuid.UUID
	selectValues        sql.SelectValues
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// Assignments holds the value of the assignments edge.
	Assignments []*ProjectAssignment `json:"assignments,omitempty"`
	// Images holds the value of the images edge.
	Images []*Image `json:"images,omitempty"`
	// Tags holds the value of the tags edge.
	Tags []*ImageTag `json:"tags,omitempty"`
	// CreatedBy holds the value of the created_by edge.
	CreatedBy *User `json:"createdBy"`
	// ModifiedBy holds the value of the modified_by edge.
	ModifiedBy *User `json:"modifiedBy"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// AssignmentsOrErr returns the Assignments value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) AssignmentsOrErr() ([]*ProjectAssignment, error) {
	if e.loadedTypes[0] {
		return e.Assignments, nil
	}
	return nil, &NotLoadedError{edge: "assignments"}
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ImagesOrErr() ([]*Image, error) {
	if e.loadedTypes[1] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) TagsOrErr() ([]*ImageTag, error) {
	if e.loadedTypes[2] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// CreatedByOrErr returns the CreatedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) CreatedByOrErr() (*User, error) {
	if e.loadedTypes[3] {
		if e.CreatedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.CreatedBy, nil
	}
	return nil, &NotLoadedError{edge: "created_by"}
}

// ModifiedByOrErr returns the ModifiedBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProjectEdges) ModifiedByOrErr() (*User, error) {
	if e.loadedTypes[4] {
		if e.ModifiedBy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.ModifiedBy, nil
	}
	return nil, &NotLoadedError{edge: "modified_by"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldName, project.FieldDescription:
			values[i] = new(sql.NullString)
		case project.FieldCreatedAt, project.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case project.FieldID:
			values[i] = new(uuid.UUID)
		case project.ForeignKeys[0]: // project_created_by
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case project.ForeignKeys[1]: // project_modified_by
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case project.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case project.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case project.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field project_created_by", values[i])
			} else if value.Valid {
				pr.project_created_by = new(uuid.UUID)
				*pr.project_created_by = *value.S.(*uuid.UUID)
			}
		case project.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field project_modified_by", values[i])
			} else if value.Valid {
				pr.project_modified_by = new(uuid.UUID)
				*pr.project_modified_by = *value.S.(*uuid.UUID)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Project.
// This includes values selected through modifiers, order, etc.
func (pr *Project) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryAssignments queries the "assignments" edge of the Project entity.
func (pr *Project) QueryAssignments() *ProjectAssignmentQuery {
	return NewProjectClient(pr.config).QueryAssignments(pr)
}

// QueryImages queries the "images" edge of the Project entity.
func (pr *Project) QueryImages() *ImageQuery {
	return NewProjectClient(pr.config).QueryImages(pr)
}

// QueryTags queries the "tags" edge of the Project entity.
func (pr *Project) QueryTags() *ImageTagQuery {
	return NewProjectClient(pr.config).QueryTags(pr)
}

// QueryCreatedBy queries the "created_by" edge of the Project entity.
func (pr *Project) QueryCreatedBy() *UserQuery {
	return NewProjectClient(pr.config).QueryCreatedBy(pr)
}

// QueryModifiedBy queries the "modified_by" edge of the Project entity.
func (pr *Project) QueryModifiedBy() *UserQuery {
	return NewProjectClient(pr.config).QueryModifiedBy(pr)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return NewProjectClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project
