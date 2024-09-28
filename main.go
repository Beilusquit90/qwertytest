package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	// r.Use(LoggerMiddleware())
	// клоака и перехват эндпоинтов
	// r.GET("/index.html", cloak.EditedIndex)

	// [GIN] 2024/09/28 - 21:42:40 | 404 |     314.403µs |     185.44.70.0 | GET      "/assets-ucp/EhkO/i9536e69d084378a3147301fc49f47507/_r192x192_png"
	// [GIN] 2024/09/28 - 21:42:40 | 404 |     226.761µs |     185.44.70.0 | GET      "/assets-ucp/EhkO/i9536e69d084378a3147301fc49f47507/_r48x48_png"
	// [GIN] 2024/09/28 - 21:42:40 | 404 |     198.058µs |     185.44.70.0 | GET      "/assets-ucp/EhkO/i9536e69d084378a3147301fc49f47507/_r96x96_png"
	r.StaticFS("/", gin.Dir("./pwa-machine-app-template/dist", false))
	err := r.Run(":80")
	if err != nil {
		panic(err)
	}
}
