package model

import (
	"backend/server/routes/message/model"

	"gorm.io/gorm"
)

type Group struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Image     string `json:"image"`
	Headline  string `json:"headline" gorm:"not null"`
	Title     string `json:"title" gorm:"not null"`
	NOP       int    `json:"nop" gorm:"default:0"` // number of the people in the group
	NOM       int    `json:"nom" gorm:"default:0"` // number of the messages in the group
	CreatedAt int64  `gorm:"autoCreateTime"`
}

type GroupQueryRequestModel struct {
	GroupId string `json:"group_id"`
}

type GroupQueryResponseModel struct {
	DoesExist    bool                 `json:"does_exist"`
	LastMessages []model.GroupMessage `json:"last_messages"`
	Group        Group                `json:"grup"`
}

func AutoMigrateGroupModel(db *gorm.DB) {
	db.AutoMigrate(&Group{})
}
