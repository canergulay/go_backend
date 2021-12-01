package messages_api

import "backend/server/routes/message/model"

type CreateMessageRequest struct {
	Message model.Message `json:"message"`
}

type CreateGroupMessageRequest struct {
	Message model.GroupMessage `json:"message"`
}

type GetMessagesRequest struct {
	ChatId   int `json:"chatid"`
	ChatType int `json:"chattype"`
}
