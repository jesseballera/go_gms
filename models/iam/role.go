package iam

import (
	"github.com/google/uuid"
	"github.com/jesseballera/go_gms/models/core"
	"gorm.io/gorm"
)

type RoleLevelEnum string

var RoleLevel = struct {
	SUPERADMIN    RoleLevelEnum
	ADMINISTRATOR RoleLevelEnum
	USER          RoleLevelEnum
}{
	SUPERADMIN:    "SuperAdmin",
	ADMINISTRATOR: "Administrator",
	USER:          "User",
}

func (Role) TableName() string {
	return "roles"
}

type Role struct {
	//gorm.Model
	RoleId      uuid.UUID       `json:"id" gorm:"column:role_id;type:uuid;default:uuid_generate_v4()"`
	Name        string          `json:"name"  gorm:"role_name;unique;not null;size:50"`
	Description string          `json:"description" gorm:"role_description;not null;size:100"`
	RoleLevel   RoleLevelEnum   `json:"roleLevel" gorm:"role_level;not null;type:enum('SUPERADMIN','ADMINISTRATOR','USER')"`
	Status      core.StatusType `json:"status"" gorm:"role_status;not null;type:enum('ACTIVE', 'INACTIVE')"`
	Permissions []Permission    `json:permissions" gorm:"foreignKey:RoleId"`
}

func GetAllRoles(db *gorm.DB, role *[]Role) (err error) {
	err = db.Find(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func GetRoleByID(db *gorm.DB, role *Role, id uuid.UUID) (err error) {
	role.RoleId = id
	err = db.Preload("Permissions").Find(&[]Permission{}).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateRole(db *gorm.DB, role *Role) (err error) {
	err = db.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateRole(db *gorm.DB, role *Role) (err error) {
	db.Save(&role)
	return nil
}

func DeleteRole(db *gorm.DB, role *Role, id uuid.UUID) (err error) {
	role.RoleId = id
	db.Delete(&role)
	return nil
}
