package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping/:param1/:param2/:param3/:param4/:param5", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")
	{
		api.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"test": "true",
			})
		})
	}
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
