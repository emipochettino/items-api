package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func NoMethodHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "method not found"})
}

func NoRoutHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "path not found"})
}
