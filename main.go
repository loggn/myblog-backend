package main

import (
	"myblog-backend/config"
	"myblog-backend/router"
)

func main() {
	config.InitDB()
	r := router.SetupRouter()
	r.Run(":8080")
}
