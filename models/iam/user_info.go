package iam

import "github.com/google/uuid"

type UserInfo struct {
	//gorm.Model
	userInfoId   uuid.UUID `json:"id" gorm:"column:user_info_id;primary_key;auto_increment"`
	FirstName    string    `json:"firstName" gorm:"column:first_name ";not null;type:varchar(100)"`
	LastName     string    `json:"lastName" gorm:"column:last_name ";not null;type:varchar(100)"`
	MobileNumber string    `json:"mobileNumber" gorm:"column:mobile_number;not null;type:varchar(20)"`
	UserId       uuid.UUID `json:"userId" gorm:"column:user_id;foreign_key:UserId"`
	User         *User     `json:"user" gorm:"foreignKey:UserId"`
}
