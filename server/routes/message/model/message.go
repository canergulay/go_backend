package model

import (
	"backend/server/routes/user/model"

	"gorm.io/gorm"
)

type Message struct {
	Id          int        `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Message     string     `json:"message"`
	Sender      int        `json:"senderid" gorm:"foreignKey:users"`
	SenderObj   model.User `json:"sender" gorm:"foreignKey:ID;references:Sender"`
	ChatId      int        `json:"chatid" gorm:"index"`
	CreatedAt   int64      `json:"createdat" gorm:"autoCreateTime"`
	MessageType int        `json:"messagetype"`
	Image       string     `json:"image"`
}

type MessageRequest struct {
	Message     string `json:"message"`
	Sender      int    `json:"senderid" gorm:"foreignKey:users"`
	ChatId      int    `json:"chatid" gorm:"index"`
	MessageType int    `json:"messagetype"`
	Image       string `json:"image"`
}

type GroupMessage Message

func InitMessageModel(db *gorm.DB) {
	db.AutoMigrate(&Message{})
	db.AutoMigrate(&GroupMessage{})
}
