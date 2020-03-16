package main

import (
	"github.com/pkg/errors"
)

func getUserById(userId int) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM \"Users\" WHERE id=$1", userId)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to fetch user")
	}

	return &user, nil
}

func getDialogById(dialogId int) (*Dialog, error) {
	dialog := Dialog{}
	err := db.Get(&dialog, "SELECT * FROM \"Dialogs\" WHERE id=$1", dialogId)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to fetch dialog")
	}

	return &dialog, nil
}

func insertMessage(userId int, dialogId int, text string) (*Message, error) {
	message := Message{}
	err := db.Get(&message, "INSERT INTO \"Messages\" (\"dialogId\", \"userId\", \"text\")VALUES ($1, $2, $3) RETURNING *", dialogId, userId, text)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to create message")
	}

	return &message, nil
}
