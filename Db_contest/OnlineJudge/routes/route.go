package routes

import (
	"OnlineJudge/controller"
	_ "OnlineJudge/docs"
	"OnlineJudge/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title OnlineJudge API
// @version 1.0
// @description This is a sample API documentation for your Go application using Gin.
// @termsOfService http://swagger.io/terms/

// @contact.name Eutop1a
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:65533
// @BasePath /
// @schemes http

// @Produce json
// @Consumes json

// SetUp initializes and configures a Gin Engine with routes for your application.
// It sets up various middleware, CORS headers, and defines API routes.
//
// Parameters:
//   - mode: The execution mode of the application (e.g., gin.DebugMode, gin.ReleaseMode).
//
// Returns:
//   - *gin.Engine: The configured Gin Engine.
func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		// Set Gin to release mode
		gin.SetMode(gin.ReleaseMode)
	}
	// Create a new Gin Engine with default middleware
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// Set up CORS headers middleware
	r.Use(middlewares.Cors())

	// Define public routes
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.POST("/check", controller.CheckUserName)
	// Group for authenticated routes
	p := r.Group("/api")
	p.Use(middlewares.AuthorizationMiddleWare)
	{
		p.POST("/sort", controller.SortByScore)
		p.POST("/changeScore", controller.AddScore)

		p.POST("/admin", controller.RegAdmin)
		p.POST("/open", controller.CheckToken)

		p.POST("/list", controller.GetList)
		p.POST("/problem/add", controller.AddProblem)
		p.POST("/problem/:id", controller.GetProblem)
	} // CRUD

	// Group for file-related routes
	{
		p.POST("/problem/file/:id", controller.GetFiles)
		p.POST("/problem/file/add/:id", controller.AddFiles)
	} // Files

	// Define a route for code judging
	p.POST("/judge", controller.JudgeCode)

	// Ping route for testing
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 404 route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})

	return r
}
