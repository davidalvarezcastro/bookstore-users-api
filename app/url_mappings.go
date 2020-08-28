package app

import (
	"github.com/davidalvarezcastro/bookstore-users-api/controllers/ping"
	"github.com/davidalvarezcastro/bookstore-users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	// router.GET("/users/search", users.Search)
	router.POST("/users", users.Create)
}
