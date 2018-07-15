package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/louisevanderlith/mango/db/funds"
	"github.com/louisevanderlith/mango/util/control"
)

type RequisitionController struct {
	control.APIController
}

// @Title GetUserRequisitions
// @Description Gets the current user's requisitions.
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [get]
func (req *RequisitionController) Get() {
	filter := funds.Requisition{}
	var container []*funds.Requisition
	err := funds.Ctx.Requisition.Read(&filter, &container)

	req.Serve(err, container)
}

// @Title GetRequisitionDetail
// @Description Gets all details (including transactions) related to a requisition.
// @Param	requisitionID			path	string 	true		"requisition's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /:requisitionID [get]
func (req *RequisitionController) GetByID() {
	var result db.IRecord

	reqID, err := strconv.ParseInt(req.Ctx.Input.Param(":requisitionID"), 10, 64)

	if err == nil {
		filter := funds.Requisition{}
		filter.Id = reqID

		result, err = funds.Ctx.Requisition.ReadOne(&filter)
	}

	req.Serve(err, result)
}

// @Title CreateRequisition
// @Description Create requisition on good or services checkout
// @Param	body		body 	funds.Requisition	true		"requisition entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *RequisitionController) Post() {
	var requisition funds.Requisition
	json.Unmarshal(req.Ctx.Input.RequestBody, &requisition)

	_, err := funds.Ctx.Requisition.Create(&requisition)

	req.Serve(err, "Requisition has been created.")
}

// @Title UpdateRequisition
// @Description Update requisition to confirm delivery of goods or services
// @Param	body		body 	funds.Requisition	true		"requisition entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *RequisitionController) Put() {
	var requisition funds.Requisition
	json.Unmarshal(req.Ctx.Input.RequestBody, &requisition)

	err := funds.Ctx.Requisition.Update(&requisition)

	req.Serve(err, "Requisition has been updated.")
}
