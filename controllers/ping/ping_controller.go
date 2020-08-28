package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping returns a pong message with a 200 response
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
