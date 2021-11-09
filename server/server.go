package server

import (
	landing "backend/server/routes/landing"
	user "backend/server/routes/user"
	"backend/server/routes/user/data/user_db"

	"github.com/gin-gonic/gin"
)

func Run() {
	user_db.AutoMigrate()
	//routes
	r := gin.Default()
	r.GET("/", landing.MainRouter)
	r.POST("/checkuser", user.CheckUser)
	r.Run(":80")
}
