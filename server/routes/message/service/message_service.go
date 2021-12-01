package service

import (
	"backend/server/routes/message/model"
	"backend/server/routes/message/repositary"
)

type MessageService struct {
	messageRP *repositary.MessageDatabaseRepositary
}

func NewMessageService(messageRP *repositary.MessageDatabaseRepositary) *MessageService {
	return &MessageService{messageRP: messageRP}
}

func (ms *MessageService) CreateMessage(message model.Message) error {
	_, err := ms.messageRP.Create(message)
	return err
}
func (ms *MessageService) CreateGroupMessage(message model.GroupMessage) error {
	_, err := ms.messageRP.Create(message)
	return err
}

func (ms *MessageService) GetMessages(chatId int, chatType int) (interface{}, error) {
	messages, err := ms.messageRP.GetMessages(chatId, chatType)
	if chatType == 0 { // 0 MEANING NORMAL MESSAGE
		return messages.([]model.Message), err
	} // ELSE, 1 MEANING GROUP MESSAGE
	return messages.([]model.GroupMessage), err

}
