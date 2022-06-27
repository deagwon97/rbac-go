package models

// Subject: Service users.
// RBAC system is dependent on User service that is Subject to this system.

type SubjectAssignment struct {
	ID        int `gorm:"primaryKey;column:ID" json:"ID"`
	SubjectID int `gorm:"column:SubjectID"    json:"SubjectID"`
	RoleID    int `gorm:"column:RoleID"       json:"RoleID"`
}

func (SubjectAssignment) TableName() string {
	return "SubjectAssignment"
}

type Role struct {
	ID          int    `gorm:"primaryKey;column:ID"  json:"ID"`
	Name        string `gorm:"column:Name"           json:"Name"`
	Description string `gorm:"column:Description"    json:"Description"`
}

func (Role) TableName() string {
	return "Role"
}

type PermissionAssignment struct {
	ID           int `gorm:"primaryKey;column:ID"   json:"ID"`
	RoleID       int `gorm:"column:RoleID"         json:"RoleID"`
	PermissionID int `gorm:"column:PermissionID"   json:"PermissionID"`
}

func (PermissionAssignment) TableName() string {
	return "PermissionAssignment"
}

type Permission struct {
	ID          int    `gorm:"primaryKey;column:ID" json:"ID"`
	ServiceName string `gorm:"column:ServiceName"  json:"ServiceName"`
	Name        string `gorm:"column:Name"          json:"Name"`
	Action      string `gorm:"column:Action"        json:"Action"`
	Object      string `gorm:"column:Object"        json:"Object"`
}

func (Permission) TableName() string {
	return "Permission"
}
