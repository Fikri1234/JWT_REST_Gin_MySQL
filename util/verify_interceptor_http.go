package util

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware ...
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessDetails, err := ExtractFromRedis(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Verify Token failure. Reason: " + err.Error()})
			c.Abort()
			return
		}

		s := strconv.FormatInt(accessDetails.UserID, 10)
		c.Request.Header.Set("userId", s)
		c.Next()
	}
}
