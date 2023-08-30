package main

import (
	"AvitoProject/controllers"
	"AvitoProject/models"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	models.ConnectDataBase()

	router := gin.Default()
	router.Use(cors.AllowAll())
	router.Static("/static", "./static")

	users := router.Group("/users")
	segments := router.Group("segments")

	segments.POST("/create", controllers.CreateSegment_HTTP)
	segments.POST("/delete", controllers.DeleteSegment_HTTP)
	users.POST("/addUserToSegment", controllers.AddUserToSegment_HTTP)
	users.GET("/activeSegments", controllers.ActiveSegments_HTTP)

	router.Run(":8080")
}
