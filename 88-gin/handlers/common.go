package handlers

import "github.com/gin-gonic/gin"

func Ping(ctx *gin.Context) {
	ctx.JSON(200, map[string]any{"Message": "Pong"})
	//ctx.JSON(200, gin.H{"Message": "Pong"})
}
