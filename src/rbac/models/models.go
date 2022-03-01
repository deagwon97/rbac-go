package models

// Subject: Service users.
// RBAC system is dependent on User service that is Subject to this system.

type SubjectAssignment struct {
	ID        int `gorm:"primaryKey;column:id" json:"id"`
	SubjectID int `gorm:"column:subject_id"    json:"subject_id"`
	RoleID    int `gorm:"column:role_id"       json:"role_id"`
}

func (SubjectAssignment) TableName() string {
	return "subject_assignment"
}

type Role struct {
	ID          int    `gorm:"primaryKey;column:id"  json:"id"`
	Name        string `gorm:"column:name"           json:"name"`
	Description string `gorm:"column:description"    json:"description"`
}

func (Role) TableName() string {
	return "role"
}

type PermissionAssignment struct {
	ID           int `gorm:"primaryKey;column:id"   json:"id"`
	RoleID       int `gorm:"column:role_id"         json:"role_id"`
	PermissionID int `gorm:"column:permission_id"   json:"permission_id"`
}

func (PermissionAssignment) TableName() string {
	return "permission_assignment"
}

type Permission struct {
	ID          int    `gorm:"primaryKey;column:id" json:"id"`
	ServiceName string `gorm:"column:service_name"  json:"service_name"`
	Name        string `gorm:"column:name"          json:"name"`
	Action      string `gorm:"column:action"        json:"action"`
	Object      string `gorm:"column:object"        json:"object"`
}

func (Permission) TableName() string {
	return "permission"
}
