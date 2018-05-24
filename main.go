package main

import (
	"github.com/gin-gonic/gin"
	"github.com/emipochettino/items-api/handlers"
)

func main() {
	router := gin.Default()
	router.StaticFile("/favicon.ico", "")
	router.GET("/ping", handlers.PingHandler)

	router.NoMethod(handlers.NoMethodHandler)
	router.NoRoute(handlers.NoRoutHandler)
	router.Run(":8080")
}
