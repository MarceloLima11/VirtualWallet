package balance

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}
