package main

import (
	env "backend/config/env"
	pg_manager "backend/config/pg_manager"
	server "backend/pkg"
)

func main() {
	env.SetEnvorinmentVariables()
	pg_manager.InitPostgreSQL()
	server.Run()
}
