package messages_api

import (
	"backend/server/routes/message/service"

	"github.com/gin-gonic/gin"
)

type MessageApi struct {
	sv *service.MessageService
	g  *gin.Context
}

func NewMessageApi(messageService *service.MessageService, g *gin.Context) *MessageApi {
	return &MessageApi{sv: messageService, g: g}
}

func (a *MessageApi) CreateNormalMessageApi() {
	var requestBody CreateMessageRequest
	a.g.BindJSON(&requestBody)
	err := a.sv.CreateMessage(requestBody.Message)
	if err == nil {
		a.g.JSON(200, requestBody)
	}
	a.g.JSON(500, requestBody)
}

func (a *MessageApi) CreateGroupMessageApi() {
	var requestBody CreateGroupMessageRequest
	a.g.BindJSON(&requestBody)
	err := a.sv.CreateGroupMessage(requestBody.Message)
	if err == nil {
		a.g.JSON(200, requestBody)
	}
	a.g.JSON(500, requestBody)
}

func (a *MessageApi) GetMessagesApi() {
	var requestBody GetMessagesRequest
	a.g.BindJSON(&requestBody)
	messages, err := a.sv.GetMessages(requestBody.ChatId, requestBody.ChatType)
	if err == nil {
		a.g.JSON(200, messages)
	}
	a.g.JSON(500, messages)
}
