package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/emipochettino/items-api-go/logger"
	"github.com/emipochettino/items-api-go/errors"
)

func PingHandler(c *gin.Context) {
	logger.Info("PingHandler called")
	c.JSON(http.StatusOK, "pong")
}

func NoMethodHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "method not found"})
}

func NoRoutHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "path not found"})
}

type HandlerFunc func(c *gin.Context) *errors.APIError

func errorWrapper(handlerFunc HandlerFunc, c *gin.Context) {
	err := handlerFunc(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}
