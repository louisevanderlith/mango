package enums

import (
	"strings"
)

type RequisitionStatus = int

const (
	Open RequisitionStatus = iota
	Pending
	Delivered
	Paid
	Error
	Cancelled
)

var requisitionStatuses = [...]string{
	"Open",
	"Pending",
	"Delivered",
	"Paid",
	"Error",
	"Cancelled"}

func GetRequisitionStatus(name string) RequisitionStatus {
	var result RequisitionStatus

	for k, v := range requisitionStatuses {
		if strings.ToUpper(name) == strings.ToUpper(v) {
			result = RequisitionStatus(k)
			break
		}
	}

	return result
}
