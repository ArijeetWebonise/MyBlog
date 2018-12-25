package app

import (
	"database/sql"

	"github.com/ArijeetBaruah/MyBlog/app/models"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router      *gin.Engine
	DB          *sql.DB
	UserService models.UserService
}
