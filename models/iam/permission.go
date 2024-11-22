package iam

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	//gorm.Model
	PermissionId uuid.UUID   `json:"id" gorm:"column:organization_type_id;type:uuid;default:uuid_generate_v4()"`
	Name         string      `json:"name" gorm:"column:permission_name; not null; unique"`
	//Description  string `json:"description" gorm:"column:permission_description;not null"`
	RoleId       uuid.UUID
	Role         Role
}

// CreatePermission creates a new Permission
func CreatePermission(db *gorm.DB, Permission *Permission) (err error) {
	// Create the Permission
	err = db.Create(Permission).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAllPermissions returns all Permissions
func GetAllPermissions(db *gorm.DB, Permissions *[]Permission) (err error) {
	err = db.Find(Permissions).Error
	if err != nil {
		return err
	}
	return nil
}

// GetPermission returns a Permission
func GetPermission(db *gorm.DB, Permission *Permission, id uuid.UUID) (err error) {
	err = db.Where("id = ?", id).First(Permission).Error
	if err != nil {
		return err
	}
	return nil
}
