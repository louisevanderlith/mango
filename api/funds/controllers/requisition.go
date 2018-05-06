package controllers

import (
	"encoding/json"

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

}

// @Title GetRequisitionDetail
// @Description Gets all details (including transactions) related to a requisition.
// @Param	requisitionID			path	string 	true		"requisition's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /:requisitionID [get]
func (req *RequisitionController) GetByID() {
	reqID, err := strconv.ParseInt(req.Ctx.Input.Param(":requisitionID"), 10, 64)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		requisition := funds.Context.Requisition.ReadOne()
		req.Data["json"] = map[string]string{"Data": }
	}

	req.ServeJSON()
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

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Requisition has been created."}
	}

	req.ServeJSON()
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

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]string{"Data": "Requisition has been updated."}
	}

	req.ServeJSON()
}
