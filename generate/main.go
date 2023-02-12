package main

import (
	"serveR/generate/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	var r *gin.Engine = gin.Default()
	r.GET("/ping", controllers.PingHandler)

	var api *gin.RouterGroup = r.Group("/api")
	{
		var generate *gin.RouterGroup = api.Group("/generate")
		{
			generate.GET("/", controllers.GetGenerate)
		}

		r.Run(":3000")
	}
}
