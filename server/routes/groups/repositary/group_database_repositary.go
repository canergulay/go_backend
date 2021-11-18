package repositary

import (
	"backend/server/routes/groups/model"

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

	return group, err
}

func (r *GroupDatabaseRepositary) GetByNameAndLocale(locale string, name string) ([]model.Group, error) {
	groups := []model.Group{}
	err := r.db.Where(model.Group{Locale: locale, Name: name}).Find(&groups).Error
	return groups, err
}
