package comment_test

import (
	"testing"

	"github.com/louisevanderlith/mango/core/comment"
)

func TestSubmitMessage_AllEmpty_Invalid(t *testing.T) {
	msg := comment.Message{}
	_, err := comment.SubmitMessage(msg)

	if err == nil {
		t.Error("Expecting validation errors.")
	}
}

func TestSubmitMessage_TextEmpty_Invalid(t *testing.T) {
	msg := comment.Message{}
	msg.CommentType = comment.Advert

	_, err := comment.SubmitMessage(msg)

	if err == nil {
		t.Error("Expecting 'Text cant be empty'")
	}
}

func TestSubmitMessage_CommentTypeEmpty_Invalid(t *testing.T) {
	msg := comment.Message{}
	msg.Text = "Testing some message"

	_, err := comment.SubmitMessage(msg)

	if err == nil {
		t.Error("Expecting 'CommentType cant be empty'")
	}
}

func TestSubmitMessage_RequiredOnly_Valid(t *testing.T) {
	msg := comment.Message{}
	msg.CommentType = comment.Advert
	msg.Text = "Testing some message"

	_, err := comment.SubmitMessage(msg)

	if err != nil {
		t.Error(err)
	}
}
