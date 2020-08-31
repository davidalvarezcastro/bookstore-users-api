package app

import (
	"github.com/davidalvarezcastro/bookstore-users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp starts the user api
func StartApp() {
	mapURLs()

	logger.Info("about to start the app")
	router.Run(":8080")
}
