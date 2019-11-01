package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Maintain() gin.HandlerFunc {
	return func(gin *gin.Context) {
		if os.Getenv("maintain") == "true" {
			//core.Output(gin, 503, make(map[string]string))
			gin.Abort()
		}
		gin.Next()
	}
}
