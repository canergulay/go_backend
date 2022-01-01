package message

import (
	"backend/server/routes/message/messages_api"
	"backend/server/routes/message/model"
	"backend/server/routes/message/repositary"
	"backend/server/routes/message/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitMessageRouter(r *gin.Engine, db *gorm.DB) {

	model.InitMessageModel(db)
	repo := repositary.NewMessageDBRepositary(db)
	service := service.NewMessageService(repo)
	api := messages_api.NewMessageApi(service)
	authorized := r.Group("/")
	authorized.Use()

	{
		r.POST(messageNormal, api.CreateNormalMessageApi)
		r.POST(messageGroup, api.CreateGroupMessageApi)
		// I WOULD NORMALLY SET THIS ENDPOINT AS A GET METHOD BUT, I GOT A PROBLEM IN CLEINT SIDE WHICH PREVENTS ME SEND BODY USING GET METHOD.
		r.POST(messageGet, api.GetMessagesApi)
	}

}

const messageNormal = "/message/normal"
const messageGroup = "/message/group"
const messageGet = "/messages"
