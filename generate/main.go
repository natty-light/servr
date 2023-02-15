package main

import (
	"serveR/generate/controllers"
	"serveR/generate/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	var r *gin.Engine = gin.Default()

	var api *gin.RouterGroup = r.Group("/api")
	{
		api.GET("/ping", controllers.PingHandler)
		api.GET("/login", controllers.HandleLogin)
		api.GET("/refresh", controllers.HandleRefresh)
		var generate *gin.RouterGroup = api.Group("/generate")
		{
			generate.Use(middleware.AuthorizeToken())
			generate.GET("/", controllers.GetGenerate)
		}

		r.Run(":3000")
	}
}
