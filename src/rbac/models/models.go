package models

type Role struct {
	ID          int    `gorm:"primaryKey;column:id"  json:"id"`
	Name        string `gorm:"column:name"           json:"name"`
	Description string `gorm:"column:description"    json:"description"`
}

type SubjectAssignment struct {
	ID        int `gorm:"primaryKey;column:id" json:"id"`
	SubjectID int `gorm:"column:subject_id"    json:"subject_id"`
	RoleID    int `gorm:"column:role_id"       json:"role_id"`
}

type PermissionAssignment struct {
	ID               int    `gorm:"primaryKey;column:id"     json:"id"`
	RoleID           int    `gorm:"column:role_id"           json:"role_id"`
	PermissionName   string `gorm:"column:permission_name"   json:"subject_name"`
	PermissionAction string `gorm:"column:permission_action" json:"permission_action"`
	PermissionObject string `gorm:"column:permission_object" json:"permission_object"`
}

type PermissionObject struct {
	PermissionObject string `gorm:"column:permission_object" json:"permission_object"`
}
