package main

type User struct {
	ID        int     `json:"id" db:"id"`
	Key       *string `json:"key" db:"key"`
	Name      *string `json:"nane" db:"name"`
	CreatedAt string  `json:"createdAt" db:"createdAt"`
	UpdatedAt string  `json:"updatedAt" db:"updatedAt"`
	DeletedAt *string `json:"deletedAt" db:"deletedAt"`
}

type Message struct {
	ID           int     `json:"id" db:"id"`
	Text         string  `json:"text" db:"text"`
	DialogId     int     `json:"dialogId" db:"dialogId"`
	FileId       *int    `json:"fileId" db:"fileId"`
	UserId       int     `json:"userId" db:"userId"`
	AttachmentId *int    `json:"attachmentId" db:"attachmentId"`
	CreatedAt    string  `json:"createdAt" db:"createdAt"`
	UpdatedAt    string  `json:"updatedAt" db:"updatedAt"`
	DeletedAt    *string `json:"deletedAt" db:"deletedAt"`
}

type Dialog struct {
	ID             int     `json:"id" db:"id"`
	Key            *string `json:"key" db:"key"`
	FirstMessageId int     `json:"firstMessageId" db:"firstMessageId"`
	LastMessageId  int     `json:"lastMessageId" db:"lastMessageId"`
	CreatedAt      string  `json:"createdAt" db:"createdAt"`
	UpdatedAt      string  `json:"updatedAt" db:"updatedAt"`
	DeletedAt      *string `json:"deletedAt" db:"deletedAt"`
}

type DialogUser struct {
	UnreadedCount       int     `json:"unreadedCount" db:"unreadedCount"`
	DialogId            int     `json:"dialogId" db:"dialogId"`
	LastReadedMessageId *int    `json:"lastReadedMessageId" db:"lastReadedMessageId"`
	UserId              int     `json:"userId" db:"userId"`
	CreatedAt           string  `json:"createdAt" db:"createdAt"`
	UpdatedAt           string  `json:"updatedAt" db:"updatedAt"`
	DeletedAt           *string `json:"deletedAt" db:"deletedAt"`
}
