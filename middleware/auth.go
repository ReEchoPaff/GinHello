package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("已经授权")
		context.Next()
	}
}
