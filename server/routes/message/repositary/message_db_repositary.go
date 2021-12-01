package repositary

import (
	"backend/server/routes/message/model"

	"gorm.io/gorm"
)

type MessageDatabaseRepositary struct {
	db *gorm.DB
}

func NewMessageDBRepositary(db *gorm.DB) *MessageDatabaseRepositary {
	return &MessageDatabaseRepositary{db: db}
}

func (r MessageDatabaseRepositary) Create(message interface{}) (interface{}, error) {
	err := r.db.Create(&message).Error
	return message, err
}

func (r MessageDatabaseRepositary) GetMessages(chatid int, chatType int) (interface{}, error) {

	if chatType == 0 {
		messages := []model.Message{}
		err := r.db.Order("nop desc").Where("chatid = ?", chatid).Find(&messages).Error
		return messages, err
	}

	groupMessages := []model.GroupMessage{}
	err := r.db.Order("nop desc").Where("chatid = ?", chatid).Find(&groupMessages).Error
	return groupMessages, err
}
