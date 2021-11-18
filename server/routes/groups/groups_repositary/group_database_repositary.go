package groups_repositary

import (
	"backend/server/routes/groups/groups_model"

	"gorm.io/gorm"
)

type GroupDatabaseRepositary struct {
	db *gorm.DB
}

func NewRepositary(db *gorm.DB) *GroupDatabaseRepositary {
	return &GroupDatabaseRepositary{db: db}
}

func (r *GroupDatabaseRepositary) Create(group groups_model.Group) (groups_model.Group, error) {
	err := r.db.Create(&group).Error

	return group, err
}

func (r *GroupDatabaseRepositary) GetByNameAndLocale(locale string, name string) ([]groups_model.Group, error) {
	groups := []groups_model.Group{}
	err := r.db.Where(groups_model.Group{Locale: locale, Name: name}).Find(&groups).Error
	return groups, err
}
