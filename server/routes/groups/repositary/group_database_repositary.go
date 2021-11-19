package repositary

import (
	"backend/global/utils"
	"backend/server/routes/groups/model"
	"fmt"

	"gorm.io/gorm"
)

type GroupDatabaseRepositary struct {
	db *gorm.DB
}

func NewRepositary(db *gorm.DB) *GroupDatabaseRepositary {
	return &GroupDatabaseRepositary{db: db}
}

func (r *GroupDatabaseRepositary) Create(group model.Group) (model.Group, error) {
	err := r.db.Create(&group).Error
	fmt.Println(err)
	return group, err
}

func (r *GroupDatabaseRepositary) GetByNameAndLocale(locale string, name string) ([]model.Group, error) {
	groups := []model.Group{}

	err := r.db.Where("locale = ? AND name ILIKE ?", locale, utils.QueryFormatter(name)).Find(&groups).Error
	return groups, err
}
