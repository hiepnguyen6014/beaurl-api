package controllers

import (
	"net/http"

	"beaurl.vn/api/src/services"
	"github.com/gin-gonic/gin"
)

type PostUrl struct {
	Url string `json:"url"`
}

func ShortenLink(ctx *gin.Context) {

	var url PostUrl
	ctx.BindJSON(&url)

	result, err := services.ShortenLink(url.Url)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"path": result,
	})

}

func GetLink(ctx *gin.Context) {
	// get link from get method
	path := ctx.Query("path")

	result := services.GetLink(path)

	ctx.JSON(http.StatusOK, gin.H{
		"target": result,
	})
}
