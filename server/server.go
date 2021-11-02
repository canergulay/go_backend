package server

import (
	landing "backend/server/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	//routes
	r := gin.Default()
	r.GET("/", landing.MainRouter)
	r.Run(":80")
}
