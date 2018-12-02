package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/core/comment"

	"github.com/louisevanderlith/mango/pkg/control"
)

type MessageController struct {
	control.APIController
}

func NewMessageCtrl(ctrlMap *control.ControllerMap) *MessageController {
	result := &MessageController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetMessages
// @Description Gets all comments related to a node.
// @Param	typeID			path 	string 	true		"comment's type"
// @Param	nodeID			path	string 	true		"node's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /:type/:nodeID[get]
func (req *MessageController) Get() {
	commentType := comment.GetCommentType(req.Ctx.Input.Param(":type"))
	nodeKey := husk.ParseKey(req.Ctx.Input.Param(":nodeID"))

	result, err := comment.GetMessage(nodeKey, commentType)

	req.Serve(result, err)
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *MessageController) Post() {
	var entry comment.Message
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &entry)

	if err != nil {
		req.Serve(nil, err)
		return
	}

	rec := comment.SubmitMessage(entry)

	req.Serve(rec, nil)
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *MessageController) Put() {
	entry, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(nil, err)
		return
	}

	err = comment.UpdateMessage(entry.Key, entry.Body.(comment.Message))

	req.Serve(nil, err)
}
