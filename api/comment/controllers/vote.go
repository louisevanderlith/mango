package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/mango/api/comment/logic"
	"github.com/louisevanderlith/mango/util/control"
)

type VoteController struct {
	control.APIController
}

// @Title CreateVote
// @Description Adds a vote to the specified comment.
// @Param	body		body 	logic.Vote 	true		"comment vote entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *VoteController) Post() {
	// Up/Down vote a comment
	var vote logic.Vote
	json.Unmarshal(req.Ctx.Input.RequestBody, &vote)

	err := vote.Submit()

	req.Serve(err, "Vote has been created.")
}
