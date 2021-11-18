package groups

import (
	"backend/server/routes/groups/groups_api"
	"backend/server/routes/groups/groups_model"
	"backend/server/routes/groups/groups_repositary"
	"backend/server/routes/groups/groups_service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitGroupRouter(r *gin.Engine, db *gorm.DB) {
	groups_model.AutoMigrateGroupModel(db)
	repositary := groups_repositary.NewRepositary(db)
	service := groups_service.NewGroupService(repositary)
	api := groups_api.NewGroupApi(service)

	r.POST(create, api.CreateGroupApi)
	r.GET(groups, api.GetGroupsByNameAndNameApi)

}

const groups = "/groups"
const create = "/groups/create/"
