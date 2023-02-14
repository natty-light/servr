package main

import (
	"serveR/generate/controllers"

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
			generate.GET("/", controllers.GetGenerate)
		}

		r.Run(":3000")
	}
}
