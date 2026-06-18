package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	var server *gin.Engine = gin.Default()
	routes.RegisterRoutes(server)
	server.Run("127.0.0.1:8080") // localhost:8080
}
