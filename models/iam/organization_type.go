package iam

import (
	"github.com/google/uuid"
	"github.com/jesseballera/go_gms/models/core"
	"gorm.io/gorm"
)

func (OrganizationType) TableName() string {
	return "organization_types"
}

type OrganizationType struct {
	//*gorm.Model
	OrganizationTypeId uuid.UUID       `json:"id" gorm:"column:organization_type_id;type:uuid;default:uuid_generate_v4()"`
	Name               string          `json:"name" gorm:"uniqueIndex;column:name;not null;type:varchar(50)"`
	Description        string          `json:"description, omitempty" "gorm:"column:organization_type_description;type:varchar(100);not null"`
	Status             core.StatusType `json:"status" gorm:"column:status;type:enum('ACTIVE','INACTIVE');not null"`
}

func CreateOrganizationType(db *gorm.DB, OrganizationType *OrganizationType) (err error) {
	err = db.Create(OrganizationType).Error
	if err != nil {
		return err
	}
	return nil
}
func GetOrganizationTypes(db *gorm.DB, OrganizationTypes *[]OrganizationType) (err error) {
	err = db.Find(OrganizationTypes).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrganizationTypeById(db *gorm.DB, OrganizationType *OrganizationType, id uuid.UUID) (err error) {
	err = db.Where("id = ?", id).First(OrganizationType).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateOrganizationType(db *gorm.DB, OrganizationType *OrganizationType) (err error) {
	db.Save(OrganizationType)
	return nil
}
func DeleteOrganizationType(db *gorm.DB, OrganizationType *OrganizationType, id uuid.UUID) (err error) {
	db.Where("id = ?", id).Delete(OrganizationType)
	return nil
}

func GetOrganizationTypeByName(db *gorm.DB, OrganizationType *OrganizationType, name string) (err error) {
	err = db.Where("name = ?", name).First(OrganizationType).Error
	if err != nil {
		return err
	}
	return nil
}
