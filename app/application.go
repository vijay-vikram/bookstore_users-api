package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vijay-vikram/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("Starting Application ...")
	router.Run()
}
