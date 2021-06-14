package main

import (
	"os"

	"github.com/afistapratama12/projectBackend/auth"
	"github.com/afistapratama12/projectBackend/config"
	"github.com/afistapratama12/projectBackend/handler"
	"github.com/afistapratama12/projectBackend/note"
	"github.com/afistapratama12/projectBackend/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Config()

	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	authService    = auth.NewService()

	noteRepository = note.NewRepository(DB)
	noteService    = note.NewService(noteRepository)

	userHandler = handler.NewUserHandler(userService, authService)
	noteHandler = handler.NewNoteHandler(noteService)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	var port string
	r := gin.Default()

	r.Use(CORSMiddleware())

	if port = os.Getenv("PORT"); port == "" {
		port = ":8080"
	}

	r.POST("/api/register")
	r.POST("/api/login")

	r.GET("/notes")
	r.POST("/notes")
	r.GET("/notes/:note_id")

	r.Run(port)
}
