package controllers

import (
	"github.com/louisevanderlith/mango/util/control"
)

type VoteController struct {
	control.APIController
}

func (req *VoteController) Post() {
	// Up/Down vote a comment
}
