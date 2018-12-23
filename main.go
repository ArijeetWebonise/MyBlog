package main

import (
	"net/http"

	"github.com/ArijeetBaruah/MyBlog/app"
	"github.com/ArijeetBaruah/MyBlog/app/config"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	cfg := &config.AppConfig{}
	cfg = cfg.ConstructAppConfig()

	a := app.App{
		Router: router,
	}

	s, _ := a.GetSchema()
	h := handler.New(&handler.Config{
		Schema:     &s,
		Pretty:     true,
		Playground: true,
	})
	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.Any("/", gin.WrapH(h))
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Start and run the server
	router.Run(":" + cfg.Port)
}
