package iam

import (
	"github.com/google/uuid"
	"github.com/jesseballera/go_gms/models/core"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type OrganizationAgent struct {
	//gorm.Model
	OrganizationAgentID uuid.UUID       `json:"id" gorm:"column:organization_agent_id;type:uuid;default:uuid_generate_v4()"`
	Name                string          `json:"name" gorm:"type:varchar(50);not null"`
	AgentLevel          int16           `json:"agentLeve" gorm:"type:smallint;not null;default:0"`
	Percentage          decimal.Decimal `json:"percentage" gorm:"type:decimal(10,2)"`
	Description         string          `json:"description, omitempty" gorm:"type:varchar(100);not null"`
	Status              core.StatusType `json:"status" gorm:"type:enum('ACTIVE','INACTIVE') not null"`
	OrganizationId      uuid.UUID
	Organization        Organization

	//core.AuditTable
}

func CreateOrganizationAgent(db *gorm.DB, OrganizationAgent *OrganizationAgent) (err error) {

	err = db.Create(OrganizationAgent).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrganizationAgent(db *gorm.DB, OrganizationAgent *OrganizationAgent, id string) (err error) {
	err = db.Where("id = ?", id).First(OrganizationAgent).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrganizationAgents(db *gorm.DB, OrganizationAgents *[]OrganizationAgent) (err error) {
	err = db.Find(OrganizationAgents).Error
	if err != nil {
		return err
	}
	return nil
}
