package app

import (
	"github.com/vijay-vikram/bookstore_users-api/controllers"
)


func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
