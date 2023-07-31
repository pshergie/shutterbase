// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CamerasColumns holds the columns for the "cameras" table.
	CamerasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "camera_owner", Type: field.TypeUUID, Nullable: true},
		{Name: "camera_created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "camera_updated_by", Type: field.TypeUUID, Nullable: true},
	}
	// CamerasTable holds the schema information for the "cameras" table.
	CamerasTable = &schema.Table{
		Name:       "cameras",
		Columns:    CamerasColumns,
		PrimaryKey: []*schema.Column{CamerasColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cameras_users_owner",
				Columns:    []*schema.Column{CamerasColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "cameras_users_created_by",
				Columns:    []*schema.Column{CamerasColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "cameras_users_updated_by",
				Columns:    []*schema.Column{CamerasColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "fileName", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Default: ""},
		{Name: "exif_data", Type: field.TypeJSON},
		{Name: "image_user", Type: field.TypeUUID},
		{Name: "image_project", Type: field.TypeUUID},
		{Name: "image_camera", Type: field.TypeUUID},
		{Name: "image_created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "image_updated_by", Type: field.TypeUUID, Nullable: true},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "images_users_user",
				Columns:    []*schema.Column{ImagesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "images_projects_project",
				Columns:    []*schema.Column{ImagesColumns[7]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "images_cameras_camera",
				Columns:    []*schema.Column{ImagesColumns[8]},
				RefColumns: []*schema.Column{CamerasColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "images_users_created_by",
				Columns:    []*schema.Column{ImagesColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "images_users_updated_by",
				Columns:    []*schema.Column{ImagesColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImageTagsColumns holds the columns for the "image_tags" table.
	ImageTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "is_album", Type: field.TypeBool, Default: false},
		{Name: "image_tag_project", Type: field.TypeUUID, Nullable: true},
		{Name: "image_tag_created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "image_tag_updated_by", Type: field.TypeUUID, Nullable: true},
	}
	// ImageTagsTable holds the schema information for the "image_tags" table.
	ImageTagsTable = &schema.Table{
		Name:       "image_tags",
		Columns:    ImageTagsColumns,
		PrimaryKey: []*schema.Column{ImageTagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "image_tags_projects_project",
				Columns:    []*schema.Column{ImageTagsColumns[6]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "image_tags_users_created_by",
				Columns:    []*schema.Column{ImageTagsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "image_tags_users_updated_by",
				Columns:    []*schema.Column{ImageTagsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "project_created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "project_updated_by", Type: field.TypeUUID, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "projects_users_created_by",
				Columns:    []*schema.Column{ProjectsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "projects_users_updated_by",
				Columns:    []*schema.Column{ProjectsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProjectAssignmentsColumns holds the columns for the "project_assignments" table.
	ProjectAssignmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "project_assignment_user", Type: field.TypeUUID},
		{Name: "project_assignment_project", Type: field.TypeUUID},
		{Name: "project_assignment_role", Type: field.TypeUUID, Nullable: true},
		{Name: "project_assignment_created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "project_assignment_updated_by", Type: field.TypeUUID, Nullable: true},
	}
	// ProjectAssignmentsTable holds the schema information for the "project_assignments" table.
	ProjectAssignmentsTable = &schema.Table{
		Name:       "project_assignments",
		Columns:    ProjectAssignmentsColumns,
		PrimaryKey: []*schema.Column{ProjectAssignmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_assignments_users_user",
				Columns:    []*schema.Column{ProjectAssignmentsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_assignments_projects_project",
				Columns:    []*schema.Column{ProjectAssignmentsColumns[4]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_assignments_roles_role",
				Columns:    []*schema.Column{ProjectAssignmentsColumns[5]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "project_assignments_users_created_by",
				Columns:    []*schema.Column{ProjectAssignmentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "project_assignments_users_updated_by",
				Columns:    []*schema.Column{ProjectAssignmentsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// TimeOffsetsColumns holds the columns for the "time_offsets" table.
	TimeOffsetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "server_time", Type: field.TypeTime},
		{Name: "camera_time", Type: field.TypeTime},
		{Name: "offset_seconds", Type: field.TypeInt},
		{Name: "time_offset_camera", Type: field.TypeUUID, Nullable: true},
		{Name: "time_offset_created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "time_offset_updated_by", Type: field.TypeUUID, Nullable: true},
	}
	// TimeOffsetsTable holds the schema information for the "time_offsets" table.
	TimeOffsetsTable = &schema.Table{
		Name:       "time_offsets",
		Columns:    TimeOffsetsColumns,
		PrimaryKey: []*schema.Column{TimeOffsetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "time_offsets_cameras_camera",
				Columns:    []*schema.Column{TimeOffsetsColumns[6]},
				RefColumns: []*schema.Column{CamerasColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "time_offsets_users_created_by",
				Columns:    []*schema.Column{TimeOffsetsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "time_offsets_users_updated_by",
				Columns:    []*schema.Column{TimeOffsetsColumns[8]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "copyright_tag", Type: field.TypeString, Unique: true},
		{Name: "email_validated", Type: field.TypeBool, Default: false},
		{Name: "validation_key", Type: field.TypeUUID},
		{Name: "validation_sent_at", Type: field.TypeTime},
		{Name: "password", Type: field.TypeBytes},
		{Name: "password_reset_key", Type: field.TypeUUID},
		{Name: "password_reset_at", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "user_role", Type: field.TypeUUID, Nullable: true},
		{Name: "user_created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "user_updated_by", Type: field.TypeUUID, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_roles_role",
				Columns:    []*schema.Column{UsersColumns[14]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_users_created_by",
				Columns:    []*schema.Column{UsersColumns[15]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_users_updated_by",
				Columns:    []*schema.Column{UsersColumns[16]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImageTagImagesColumns holds the columns for the "image_tag_images" table.
	ImageTagImagesColumns = []*schema.Column{
		{Name: "image_tag_id", Type: field.TypeUUID},
		{Name: "image_id", Type: field.TypeUUID},
	}
	// ImageTagImagesTable holds the schema information for the "image_tag_images" table.
	ImageTagImagesTable = &schema.Table{
		Name:       "image_tag_images",
		Columns:    ImageTagImagesColumns,
		PrimaryKey: []*schema.Column{ImageTagImagesColumns[0], ImageTagImagesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "image_tag_images_image_tag_id",
				Columns:    []*schema.Column{ImageTagImagesColumns[0]},
				RefColumns: []*schema.Column{ImageTagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "image_tag_images_image_id",
				Columns:    []*schema.Column{ImageTagImagesColumns[1]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CamerasTable,
		ImagesTable,
		ImageTagsTable,
		ProjectsTable,
		ProjectAssignmentsTable,
		RolesTable,
		TimeOffsetsTable,
		UsersTable,
		ImageTagImagesTable,
	}
)

func init() {
	CamerasTable.ForeignKeys[0].RefTable = UsersTable
	CamerasTable.ForeignKeys[1].RefTable = UsersTable
	CamerasTable.ForeignKeys[2].RefTable = UsersTable
	ImagesTable.ForeignKeys[0].RefTable = UsersTable
	ImagesTable.ForeignKeys[1].RefTable = ProjectsTable
	ImagesTable.ForeignKeys[2].RefTable = CamerasTable
	ImagesTable.ForeignKeys[3].RefTable = UsersTable
	ImagesTable.ForeignKeys[4].RefTable = UsersTable
	ImageTagsTable.ForeignKeys[0].RefTable = ProjectsTable
	ImageTagsTable.ForeignKeys[1].RefTable = UsersTable
	ImageTagsTable.ForeignKeys[2].RefTable = UsersTable
	ProjectsTable.ForeignKeys[0].RefTable = UsersTable
	ProjectsTable.ForeignKeys[1].RefTable = UsersTable
	ProjectAssignmentsTable.ForeignKeys[0].RefTable = UsersTable
	ProjectAssignmentsTable.ForeignKeys[1].RefTable = ProjectsTable
	ProjectAssignmentsTable.ForeignKeys[2].RefTable = RolesTable
	ProjectAssignmentsTable.ForeignKeys[3].RefTable = UsersTable
	ProjectAssignmentsTable.ForeignKeys[4].RefTable = UsersTable
	TimeOffsetsTable.ForeignKeys[0].RefTable = CamerasTable
	TimeOffsetsTable.ForeignKeys[1].RefTable = UsersTable
	TimeOffsetsTable.ForeignKeys[2].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = RolesTable
	UsersTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[2].RefTable = UsersTable
	ImageTagImagesTable.ForeignKeys[0].RefTable = ImageTagsTable
	ImageTagImagesTable.ForeignKeys[1].RefTable = ImagesTable
}
