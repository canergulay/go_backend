package groups_service

import (
	"backend/server/routes/groups/groups_model"
	"backend/server/routes/groups/groups_repositary"
)

type GroupService struct {
	dbRP *groups_repositary.GroupDatabaseRepositary
}

func NewGroupService(dbRP *groups_repositary.GroupDatabaseRepositary) *GroupService {
	return &GroupService{dbRP: dbRP}
}

func (s *GroupService) CreateGroup(group groups_model.Group) (groups_model.Group, error) {
	return s.dbRP.Create(group)
}

func (s *GroupService) GetGroupsByNameAndLocale(locale string, name string) ([]groups_model.Group, error) {
	return s.dbRP.GetByNameAndLocale(locale, name)
}
