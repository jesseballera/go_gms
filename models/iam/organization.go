package iam

import (
	"github.com/google/uuid"
	"github.com/jesseballera/go_gms/models/core"
	"gorm.io/gorm"
)

type Organization struct {
	//gorm.Model
	OrganizationId    uuid.UUID           `json:"id" gorm:"column:organization_id;type:uuid;default:uuid_generate_v4()"`
	Name              string              `json:"name" gorm:"uniqueIndex;column:organization_name;not null;type:varchar(50)"`
	Description       string              `json:"description" gorm:"column:organization_description;type:varchar(100) not null"`
	Status            core.StatusType     `json:"status" gorm:"type:enum('ACTIVE','INACTIVE') not null"`
	OrganizationType  OrganizationType    `json:"organizationType" gorm:"foreignKey:OrganizationTypeId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	OrganizationAgent []OrganizationAgent `json:"organizationAgent" gorm:"foreignKey:OrganizationId"`
}

func GetOrganizations(db *gorm.DB, Organizations *[]Organization) (err error) {
	err = db.Find(Organizations).Error
	if err != nil {
		return err
	}
	return nil
}
func GetOrganization(db *gorm.DB, Organization *Organization, id uuid.UUID) (err error) {
	err = db.Where("id = ?", id).First(Organization).Error
	if err != nil {
		return err
	}
	return nil
}
func CreateOrganization(db *gorm.DB, Organization *Organization) (err error) {
	err = db.Create(Organization).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrganization(db *gorm.DB, Organization *Organization) (err error) {
	db.Save(Organization)
	return nil
}
func DeleteOrganization(db *gorm.DB, Organization *Organization, id uuid.UUID) (err error) {
	db.Where("id =?", id).Delete(Organization)
	return nil
}
