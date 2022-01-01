package messages_api

import "backend/server/routes/message/model"

type CreateMessageRequest struct {
	Message model.MessageRequest `json:"message"`
}

type CreateGroupMessageRequest struct {
	Message model.MessageRequest `json:"message"`
}

type GetMessagesRequest struct {
	ChatId   int `json:"chatid"`
	ChatType int `json:"chattype"`
}
