package core

import "database/sql/driver"

type StatusType string

const (
	ACTIVE   StatusType = "Active"
	INACTIVE StatusType = "Inactive"
)

func (st *StatusType) Scan(value string) error {
	*st = StatusType(value)
	return nil
}

func (st StatusType) Value() (driver.Value, error) {
	return string(st), nil
}

//var Status = struct {
//	ACTIVE   StatusEnum
//	INACTIVE StatusEnum
//}{
//	ACTIVE:   "Active",
//	INACTIVE: "Inactive",
//}
