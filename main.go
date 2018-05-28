package main

import (
	_ "github.com/emipochettino/items-api/entities"
	"github.com/emipochettino/items-api/handlers"
	"github.com/emipochettino/items-api/middlewares"
	"github.com/gin-gonic/gin"
)

//Export GIN_MODE=release for Production
func main() {
	router := gin.Default()
	router.Use(middlewares.RequestTracking)

	router.GET("/ping", handlers.PingHandler)

	router.POST("/category", handlers.CreateCategoryHandler)
	router.PUT("/category", handlers.UpdateCategoryHandler)
	router.GET("/category/:id", handlers.GetCategoryHandler)
	router.GET("/category", handlers.FindCategoryHandler)
	router.DELETE("/category/:id", handlers.DeleteCategoryHandler)

	router.NoMethod(handlers.NoMethodHandler)
	router.NoRoute(handlers.NoRoutHandler)
	router.Run(":8080")
}
