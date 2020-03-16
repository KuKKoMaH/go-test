package main

import (
	"net/http"

	"github.com/pkg/errors"
)

type CreateMessageArgs struct {
	DialogId int
	UserId   int
	Text     string
}

type CreateMessageResult struct {
	Message *Message `json:"message"`
}

func (t *Messenger) CreateMessage(r *http.Request, args *CreateMessageArgs, result *CreateMessageResult) error {
	var (
		user    *User
		dialog  *Dialog
		message *Message
		err     error
	)
	user, err = getUserById(args.UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("User not found")
	}

	dialog, err = getDialogById(args.DialogId)
	if err != nil {
		return err
	}
	if dialog == nil {
		return errors.New("Dialog not found")
	}

	message, err = insertMessage(args.UserId, args.DialogId, args.Text)
	if err != nil {
		return err
	}
	if message == nil {
		return errors.New("Message create error")
	}

	result.Message = message

	return nil
}
