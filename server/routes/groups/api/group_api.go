package api

import (
	"backend/server/routes/groups/model"
	"backend/server/routes/groups/service"

	"github.com/gin-gonic/gin"
)

type GroupApi struct {
	service *service.GroupService
}

func NewGroupApi(service *service.GroupService) *GroupApi {
	return &GroupApi{service: service}
}

func (a *GroupApi) CreateGroupApi(c *gin.Context) {
	var modelToCreate model.Group
	c.BindJSON(&modelToCreate)
	_, err := a.service.CreateGroup(modelToCreate)
	if err != nil {
		c.JSON(500, "error")
	}
	c.JSON(500, modelToCreate)
}

func (a *GroupApi) GetGroupsByNameAndNameApi(c *gin.Context) {
	var requestBody GetGroupsRequest
	c.BindJSON(&requestBody)

	_, err := a.service.GetGroupsByNameAndLocale(requestBody.Locale, requestBody.Name)
	if err != nil {
		c.JSON(500, "error")
	}
	c.JSON(500, requestBody)
}
