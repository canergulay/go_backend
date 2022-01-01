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

func (ms *MessageService) CreateMessage(m model.MessageRequest) error {
	msg := model.Message{Message: m.Message, Sender: m.Sender, ChatId: m.ChatId, Image: m.Image}
	_, err := ms.messageRP.CreateMessage(msg)
	return err
}
func (ms *MessageService) CreateGroupMessage(m model.MessageRequest) error {
	msg := model.GroupMessage{Message: m.Message, Sender: m.Sender, ChatId: m.ChatId, Image: m.Image}

	_, err := ms.messageRP.CreateGroupMessage(msg)
	return err
}

func (ms *MessageService) GetMessages(chatId int, chatType int) (interface{}, error) {
	messages, err := ms.messageRP.GetMessages(chatId, chatType)
	if chatType == 0 { // 0 MEANING NORMAL MESSAGE
		return messages.([]model.Message), err
	} // ELSE, 1 MEANING GROUP MESSAGE
	return messages.([]model.GroupMessage), err

}
