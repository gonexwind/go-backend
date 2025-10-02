package routes

import (
	"gonexwind/backend-api/controllers"
	"gonexwind/backend-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//initialize gin
	router := gin.Default()

	// set up CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	public := router.Group("/api")
	{
		// Auth
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)

		// Post
		public.GET("/posts", controllers.ShowPosts)
	}

	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		// Profile
		protected.GET("/profile", controllers.GetProfile)
		//protected.GET("/users", controllers.FindUsers)
		//protected.POST("/users", controllers.CreateUser)
		//protected.GET("/users/:id", controllers.FindUserById)
		//protected.PUT("/users/:id", controllers.UpdateUser)
		//protected.DELETE("/users/:id", controllers.DeleteUser)

		// POST
		protected.POST("/posts", controllers.CreatePost)
		protected.PUT("/posts/:id", controllers.UpdatePost)
		protected.DELETE("/posts/:id", controllers.DeletePost)
	}

	return router
}
