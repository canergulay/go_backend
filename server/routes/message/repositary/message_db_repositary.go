package repositary

import (
	"backend/server/routes/message/model"
	"fmt"

	"gorm.io/gorm"
)

type MessageDatabaseRepositary struct {
	db *gorm.DB
}

func NewMessageDBRepositary(db *gorm.DB) *MessageDatabaseRepositary {
	return &MessageDatabaseRepositary{db: db}
}

func (r MessageDatabaseRepositary) CreateGroupMessage(message model.GroupMessage) (interface{}, error) {
	err := r.db.Create(&message).Error
	return message, err
}

func (r MessageDatabaseRepositary) CreateMessage(message model.Message) (interface{}, error) {
	err := r.db.Create(&message).Error
	return message, err
}

func (r MessageDatabaseRepositary) GetMessages(chatid int, chatType int) (interface{}, error) {
	fmt.Println(chatid)
	fmt.Println(chatType)
	if chatType == 0 {

		messages := []model.Message{}
		err := r.db.Preload("SenderObj").Where("chat_id = ?", chatid).Order("created_at desc").Limit(10).Find(&messages).Error

		return messages, err
	}

	groupMessages := []model.GroupMessage{}
	err := r.db.Preload("SenderObj").Where("chat_id = ?", chatid).Order("created_at desc").Limit(10).Find(&groupMessages).Error

	return groupMessages, err
}
