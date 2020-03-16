package main

import (
	"net/http"

	"github.com/pkg/errors"
)

type Messenger int

type GetUserDialogsArgs struct {
	UserId int
}

type GetUserDialogsResult struct {
	Users       []*User      `json:"users"`
	Messages    []Message    `json:"messages"`
	Dialogs     []Dialog     `json:"dialogs"`
	DialogUsers []DialogUser `json:"dialogsUsers"`
}

func getUserById(userId int) (*User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM \"Users\" WHERE id=$1", userId)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to fetch user")
	}

	return &user, nil
}

func getRest(userId int) ([]Dialog, []Message, []DialogUser, error) {
	var result []struct {
		Dialog     Dialog     `db:"D"`
		DialogUser DialogUser `db:"DU"`
		Message    Message    `db:"M"`
	}
	err := db.Select(&result, `
	SELECT "D"."id" as "D.id", "D"."key" as "D.key", "D"."createdAt" as "D.createdAt", "D"."firstMessageId" as "D.firstMessageId", "D"."lastMessageId" as "D.lastMessageId", "D"."updatedAt" as "D.updatedAt", "D"."deletedAt" as "D.deletedAt", "DU"."dialogId" as "DU.dialogId", "DU"."userId" as "DU.userId", "DU"."unreadedCount" as "DU.unreadedCount", "DU"."lastReadedMessageId" as "DU.lastReadedMessageId", "DU"."createdAt" as "DU.createdAt", "DU"."updatedAt" as "DU.updatedAt", "DU"."deletedAt" as "DU.deletedAt", "M"."id" as "M.id", "M"."text" as "M.text", "M"."dialogId" as "M.dialogId", "M"."fileId" as "M.fileId", "M"."userId" as "M.userId", "M"."createdAt" as "M.createdAt", "M"."updatedAt" as "M.updatedAt", "M"."deletedAt" as "M.deletedAt"
        FROM "Dialogs" "D"
	        LEFT JOIN "DialogUsers" "DU" ON "DU"."dialogId" = "D".id
          LEFT JOIN (
            SELECT DISTINCT ON ("dialogId") *
              FROM "Messages"
              WHERE "deletedAt" IS NULL
              ORDER BY "dialogId", id DESC
          ) "M" ON "M"."dialogId" = "D".id
	      WHERE "DU"."userId" = $1
	        AND "D"."deletedAt" IS NULL 
	        AND "DU"."deletedAt" IS NULL
	      ORDER BY "M".id DESC NULLS LAST, "D".id DESC 
		`, userId)

	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "Unable to fetch data")
	}

	count := len(result)
	dialogs := make([]Dialog, count)
	dialogUsers := make([]DialogUser, count)
	messages := make([]Message, count)

	for i, row := range result {
		dialogs[i] = row.Dialog
		dialogUsers[i] = row.DialogUser
		messages[i] = row.Message
	}

	return dialogs, messages, dialogUsers, nil
}

func (t *Messenger) GetUserDialogs(r *http.Request, args *GetUserDialogsArgs, result *GetUserDialogsResult) error {
	user, err := getUserById(args.UserId)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("User not found")
	}

	dialogs, messages, dialogUsers, err2 := getRest(args.UserId)
	if err2 != nil {
		return err2
	}

	result.Users = []*User{user}
	result.Dialogs = dialogs
	result.Messages = messages
	result.DialogUsers = dialogUsers
	return nil
}
