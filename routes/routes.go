package routes

import (
	"zestream/controllers"

	"github.com/gin-gonic/gin"
)

// Init function will perform all route operations
func Init() *gin.Engine {

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	v1 := r.Group("/api/v1")

	v1.GET("ping", controllers.Ping)

	v1.POST("process-video", controllers.ProcessVideo)

	return r
}
