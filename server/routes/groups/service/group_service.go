package service

import (
	"backend/server/routes/groups/model"
	"backend/server/routes/groups/repositary"
	"backend/server/routes/groups/utils"
)

type GroupService struct {
	dbRP *repositary.GroupDatabaseRepositary
}

func NewGroupService(dbRP *repositary.GroupDatabaseRepositary) *GroupService {
	return &GroupService{dbRP: dbRP}
}

func (s *GroupService) CreateGroup(group model.Group) (model.Group, error) {
	// LETS ATTACH A RANDOM IMAGE BEFORE GROUP IS CREATED
	group.Image = utils.GetRandomDefaulGroupPicture()
	return s.dbRP.Create(group)
}

func (s *GroupService) GetGroupsByNameAndLocale(locale string, name string) ([]model.Group, error) {

	return s.dbRP.GetByNameAndLocale(locale, name)
}
