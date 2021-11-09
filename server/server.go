package server

import (
	landing "backend/server/routes/landing"
	user "backend/server/routes/user"

	"github.com/gin-gonic/gin"
)

func Run() {
	//routes
	r := gin.Default()
	r.GET("/", landing.MainRouter)
	r.POST("/checkuser", user.CheckUser)
	r.Run(":80")
}
