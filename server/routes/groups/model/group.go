package model

import "gorm.io/gorm"

type Group struct {
	Id         int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name       string `json:"name" gorm:"not null"`
	Image      string `json:"image"`
	NOP        int    `json:"nop" gorm:"default:0"` // number of the people in the group
	NOM        int    `json:"nom" gorm:"default:0"` // number of the messages in the group
	NOT        int    `json:"not" gorm:"default:0"` // number of the topics created within the group
	Definition string `json:"definiton" gorm:"not null"`
	Locale     string `json:"locale" gorm:"not null"`
	Creator    int    `json:"creator" gorm:"not null"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
}

func AutoMigrateGroupModel(db *gorm.DB) {
	db.AutoMigrate(&Group{})
}
