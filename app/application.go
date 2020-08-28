package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApp starts the user api
func StartApp() {
	mapURLs()
	router.Run(":8080")
}
