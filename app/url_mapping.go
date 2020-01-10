package app

import (
	"github.com/bhawana/bookstore_user-apis/contollers/ping"
	"github.com/bhawana/bookstore_user-apis/contollers/users"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", contollers.SearchUser)
}
