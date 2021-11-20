package groups

import (
	"backend/server/routes/groups/api"
	"backend/server/routes/groups/model"
	"backend/server/routes/groups/repositary"
	"backend/server/routes/groups/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitGroupRouter(r *gin.Engine, db *gorm.DB) {
	// yeah that's better //
	// IT'S LIKE A BUILT IN DEPENDENCY INJECTION THAT'S SUPPORTED BY THE LANGUAGE ITSELF
	model.AutoMigrateGroupModel(db)
	repositary := repositary.NewRepositary(db)
	serv := service.NewGroupService(repositary)
	api := api.NewGroupApi(serv)

	// I WONDER IF IT COULD BE MORE CLEAR ?
	r.POST(groups, api.GetGroupsByNameAndLocaleApi)
	r.POST(create, api.CreateGroupApi)

}

const groups = "/groups"
const create = "/groups/create/"
