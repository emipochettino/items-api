package main

import (
	_ "github.com/emipochettino/items-api-go/entities"
	"github.com/emipochettino/items-api-go/handlers"
	"github.com/emipochettino/items-api-go/middlewares"
	"github.com/gin-gonic/gin"
)

//Export GIN_MODE=release for Production
func main() {
	router := gin.Default()
	router.Use(middlewares.RequestTracking)

	router.GET("/ping", handlers.PingHandler)

	//Resource Category
	router.POST("/category", handlers.CreateCategoryHandler)
	router.PUT("/category", handlers.UpdateCategoryHandler)
	router.GET("/category/:id", handlers.GetCategoryHandler)
	router.GET("/category", handlers.FindCategoryHandler)
	router.DELETE("/category/:id", handlers.DeleteCategoryHandler)

	//Resource User
	router.POST("/user", handlers.CreateUserHandler)
	router.PUT("/user", handlers.UpdateUserHandler)
	router.GET("/user/:id", handlers.GetUserHandler)
	router.GET("/user", handlers.FindUsersHandler)
	router.DELETE("/user/:id", handlers.DeleteUserHandler)

	//Resource Item
	router.POST("/item", handlers.CreateItemHandler)
	router.PUT("/item", handlers.UpdateItemHandler)
	router.GET("/item/:id", handlers.GetItemHandler)
	router.GET("/item", handlers.FindItemsHandler)
	router.DELETE("/item/:id", handlers.DeleteItemHandler)

	router.NoMethod(handlers.NoMethodHandler)
	router.NoRoute(handlers.NoRoutHandler)
	router.Run(":8080")
}
