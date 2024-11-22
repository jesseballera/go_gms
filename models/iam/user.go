package iam

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStatusEnum string

var UserStatus = struct {
	ACTIVE   UserStatusEnum
	INACTIVE UserStatusEnum
	LOCKED   UserStatusEnum
	BLOCKED  UserStatusEnum
	NEW      UserStatusEnum
}{
	ACTIVE:   "Active",
	INACTIVE: "Inactive",
	LOCKED:   "Locked",
	BLOCKED:  "Blocked",
	NEW:      "New",
}

type User struct {
	UserId         uuid.UUID      `json:"id" gorm:"column:user_id;type:uuid;default:uuid_generate_v4()"`
	Username       string         `json:"username" gorm:"column:username;uniqueIndex;not null;type:varchar(50)"`
	Password       string         `json:"password" gorm:"column:password;not null;type:varchar(100)"`
	Email          string         `json:"email" gorm:"column:email;uniqueIndex;not null;type:varchar(100)"`
	UserStatus     UserStatusEnum `json:"userStatus" gorm:"column:user_status;not null;type:enum('ACTIVE','INACTIVE','LOCKED','BLOCKED','NEW')"`
	Organization   Organization   `json:"organization" gorm:"foreignKey:OrganizationId"`
	Role           Role           `json:"role" gorm:"foreignKey:RoleId"`
	UserInfo       UserInfo       `json:"userInfo" gorm:"foreignKey:UserId"`
	TerminalAccess bool           `json:"terminalAccess" gorm:"column:terminal_access;default:false"`
}

// CreateUser creates a new user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAllUsers returns all users
func GetAllUsers(db *gorm.DB, user *[]User) (err error) {

	err = db.Find(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUser returns a user
func GetUser(db *gorm.DB, user *User, id uuid.UUID) (err error) {

	err = db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser updates a user
func UpdateUser(db *gorm.DB, user *User) (err error) {

	err = db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user
func DeleteUser(db *gorm.DB, user *User, id uuid.UUID) (err error) {

	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
