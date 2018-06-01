package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/emipochettino/items-api-go/context"
)

const request = "request"

func RequestTracking(c *gin.Context) {
	requestID := c.Request.Header.Get(request)
	if len(requestID) == 0 {
		requestID = uuid.NewV4().String()
	}

	context.Context.ContextMap[request] = requestID
	c.Next()
}
