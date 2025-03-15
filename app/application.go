package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rainmore/rest-api/datasources/postgresql/users_db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	users_db.Init()
	mapUrls()
	router.Run(":8080")
}
