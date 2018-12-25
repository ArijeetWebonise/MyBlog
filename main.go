package main

import (
	"fmt"
	"net/http"

	"github.com/ArijeetBaruah/MyBlog/app"
	"github.com/ArijeetBaruah/MyBlog/app/config"
	"github.com/ArijeetBaruah/MyBlog/app/models"
	"github.com/ArijeetBaruah/MyBlog/pkg/database"
	"github.com/ArijeetBaruah/MyBlog/pkg/logger"
	"github.com/ArijeetBaruah/MyBlog/pkg/redis"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"github.com/graphql-go/handler"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	logger := &logger.RealLogger{}
	logger.Initialise()

	// Set the router as the default one shipped with Gin
	router := gin.Default()
	cfg := &config.AppConfig{}
	cfg = cfg.ConstructAppConfig()

	db := &database.DatabaseWrapper{}
	dbConn, dbErr := db.Initialise(&cfg.DB)
	if dbErr != nil {
		logger.Panic(dbErr)
		return
	}

	migrations := migrate.PackrMigrationSource{
		Box: packr.NewBox("./migrations"),
	}

	n, err := migrate.Exec(dbConn, cfg.DB.DbDriverName, migrations, migrate.Up)
	if err != nil {
		logger.Panic(err)
		return
	}
	logger.Infof("%d migration executed", n)

	redisService := redis.RedisServiceImpl{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	}
	fmt.Println(redisService.Init())

	a := app.App{
		Router:      router,
		DB:          dbConn,
		UserService: &models.UserServiceImpl{DB: dbConn},
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
	api := router.Group("/graphql")
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
