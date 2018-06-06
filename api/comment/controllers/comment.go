package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/louisevanderlith/mango/api/comment/logic"
	"github.com/louisevanderlith/mango/db/comment"

	"github.com/louisevanderlith/mango/util/control"
)

type CommentController struct {
	control.APIController
}

// @Title GetComments
// @Description Gets all comments related to a node.
// @Param	typeID			path 	string 	true		"comment's type"
// @Param	nodeID			path	string 	true		"node's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /:type/:nodeID[get]
func (req *CommentController) Get() {
	var result logic.CommentChain

	commentType := comment.GetCommentType(req.Ctx.Input.Param(":type"))
	nodeID, err := strconv.ParseInt(req.Ctx.Input.Param(":nodeID"), 10, 64)

	if err == nil {
		result, err = logic.GetCommentChain(nodeID, commentType)
	}

	req.Serve(err, result)
}

// @Title CreateComment
// @Description Creates a comment under a node
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *CommentController) Post() {
	var comment logic.MessageEntry
	json.Unmarshal(req.Ctx.Input.RequestBody, &comment)

	sessionID := req.Ctx.GetCookie("avosession")
	userID, err := control.GetUserID(sessionID)

	if err == nil {
		err = logic.SubmitComment(userID, comment)
	}

	req.Serve(err, "Comment has been created.")
}
