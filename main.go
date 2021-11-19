package main

import (
	env "backend/config/env"
	pg_manager "backend/config/pg_manager"
	server "backend/server"

	"github.com/gin-gonic/gin"
)

func main() {
	env.SetEnvorinmentVariables()
	server := server.App{}
	server.KickASS(pg_manager.InitPostgreSQL(), gin.Default())
}
