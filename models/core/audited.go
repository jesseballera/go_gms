package core

import "fmt"

type AuditTable struct {
	CreatedBy string
	UpdatedBy string
}

func (model *AuditTable) SetCreatedBy(createdBy string) {
	model.CreatedBy = fmt.Sprint("%v", createdBy)
}

func (model *AuditTable) GetCreatedBy() string {
	return model.CreatedBy
}

func (model *AuditTable) SetUpdatedBy(updatedBy string) {
	model.UpdatedBy = fmt.Sprint("%v", updatedBy)
}

func (model *AuditTable) GetUpdatedBy() string {
	return model.UpdatedBy
}
