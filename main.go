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
	port string
	DB   *gorm.DB = config.Config()

	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	authService    = auth.NewService()

	noteRepository = note.NewRepository(DB)
	noteService    = note.NewService(noteRepository)

	userHandler = handler.NewUserHandler(userService, authService)
	noteHandler = handler.NewNoteHandler(noteService)

	middleware = handler.Middleware(userService, authService)
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

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	if port = os.Getenv("PORT"); port == "" {
		port = ":8080"
	}

	r.Static("/images", "./images")

	routeAPI := r.Group("/api")
	{
		routeAPI.POST("/register", userHandler.RegisterUser)
		routeAPI.POST("/login", userHandler.LoginUser)
		routeAPI.POST("/register/admin", userHandler.RegisterAdmin)
		routeAPI.POST("/email_confirmation/:confirmation_key", userHandler.VerificationEmailUser)

		routeAPI.GET("/notes", noteHandler.GetAllNote)
		routeAPI.GET("/users/notes", middleware, noteHandler.GetAllNoteByUser)
		routeAPI.POST("/notes", middleware, noteHandler.SaveNewNote)
		routeAPI.GET("/notes/:note_id", middleware, noteHandler.GetByIDNote)
		routeAPI.PUT("/notes/:note_id", middleware, noteHandler.UpdateNote)
		routeAPI.PATCH("/notes/:note_id", middleware, noteHandler.UnDeleteNote)
		routeAPI.DELETE("/notes/:note_id", middleware, noteHandler.DeleteNote)
	}

	return r
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	router := SetupRouter()

	// user, err := userService.GetByUsername("afistapratama")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(user)

	router.Run(port)
}
