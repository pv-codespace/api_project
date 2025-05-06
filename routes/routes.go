package routes

import (
	"rest_api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)

	authenticated := server.Group("/")

	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.PUT("/events/:id", updateEvent)

	// server.POST("/events", createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)
	// server.GET("/users", getUsers)
}
