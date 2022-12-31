package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortenLink(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "You're in shorten link",
	})

}
