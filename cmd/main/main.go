package main

import (
	env "backend/config/env"
	pg_manager "backend/config/pg_manager"
	"backend/global/authentication"
	"backend/global/logclient"
	server "backend/server"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	env.SetEnvorinmentVariables()
	logclient.InitLogger()
	server := server.App{}
	server.KickASS(pg_manager.InitPostgreSQL(), gin.Default(), &authentication.JwtManager{SecretKey: os.Getenv("JWT_SECRET")})
}
