package model

import "gorm.io/gorm"

type Message struct {
	Id          int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Message     string `json:"message"`
	Sender      int    `json:"sender" gorm:"foreignKey:users"`
	ChatId      int    `json:"chatid" gorm:"index"`
	CreatedAt   int64  `json:"createdat" gorm:"autoCreateTime"`
	MessageType int    `json:"messagetype"` // 0 representing text , 1 representing image
	Image       string `json:"image"`       // will be empty if the image doesn't contain image
}
type GroupMessage Message

func InitMessageModel(db *gorm.DB) {
	db.AutoMigrate(&Message{})
	db.AutoMigrate(&GroupMessage{})

}
