package main

import (
	"gonexwind/backend-api/config"
	"gonexwind/backend-api/database"
	"gonexwind/backend-api/routes"
)

func main() {

	//load config .env
	config.LoadEnv()

	//inisialisasi database
	database.InitDB()

	//setup router
	r := routes.SetupRouter()

	//mulai server dengan port 3000
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
