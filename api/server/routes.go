package server

import (
	"github.com/gin-gonic/gin"
)

func loadRoutes(router *gin.Engine, env *Env) {

	api := router.Group("/api")
	{
		api.GET("/health", env.HealthCheck)

		api.GET("/auth/token", env.UserCreate)
		api.GET("/auth/token/:token", env.UserValidate)
		api.Use(env.UserVerify)
		router.NoRoute(noRoute)

		api.GET("/post/:id", env.PostGet)
		api.PATCH("/post/:id", env.PostPatch)
		api.DELETE("/post/:id", env.PostDelete)
		api.POST("/post/:postID/post", env.PostPost)
		api.POST("/post", env.PostPost)

		api.POST("/post/:postID/vote", env.VotePost)

		api.GET("/posts", env.PostsGet)
	}
}
