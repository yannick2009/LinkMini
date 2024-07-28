package handlers

import (
	"linkmini/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Body struct {
	URL string `json:"url"`
}

func CreateShortURLHandler(ctx *gin.Context) {
	body := new(Body)
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := service.CreateShortURLService(body.URL)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, data)
}

var hashParam string = "hash"

func RedirectURL(ctx *gin.Context) {
	URLHash := ctx.Param(hashParam)

	URLData, err := service.GetLongURLAndUpdate(URLHash)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	if URLData != nil && time.Now().Before(URLData.ExpireAt) {
		ctx.Redirect(http.StatusPermanentRedirect, URLData.LongURL)
	} else {
		ctx.Redirect(http.StatusPermanentRedirect, "https://www.mongodb.com/products/platform/atlas-database")
		_ = service.DeleteURLExpired(URLHash)
	}
}

func GetURLStats(ctx *gin.Context) {
	URLHash := ctx.Param(hashParam)

	URLData, err := service.GetLongURL(URLHash)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, URLData)
}
