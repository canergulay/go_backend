package messages_api

import (
	"backend/server/routes/message/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MessageApi struct {
	sv *service.MessageService
}

func NewMessageApi(messageService *service.MessageService) *MessageApi {
	return &MessageApi{sv: messageService}
}

func (a *MessageApi) CreateNormalMessageApi(c *gin.Context) {
	var requestBody CreateMessageRequest
	c.BindJSON(&requestBody)
	err := a.sv.CreateMessage(requestBody.Message)
	if err == nil {
		c.JSON(200, requestBody)
		return
	}
	c.JSON(500, requestBody)
}

func (a *MessageApi) CreateGroupMessageApi(c *gin.Context) {
	var requestBody CreateGroupMessageRequest
	c.BindJSON(&requestBody)
	err := a.sv.CreateGroupMessage(requestBody.Message)
	if err == nil {
		c.JSON(200, requestBody)
		return

	}
	c.JSON(500, requestBody)
}

func (a *MessageApi) GetMessagesApi(c *gin.Context) {
	var requestBody GetMessagesRequest
	c.BindJSON(&requestBody)
	messages, err := a.sv.GetMessages(requestBody.ChatId, requestBody.ChatType)
	fmt.Println("msjes")
	fmt.Println(messages)
	if err == nil {
		c.JSON(200, messages)
		return
	}
	c.JSON(500, messages)
}
