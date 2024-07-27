package handlers

import (
	"linkmini/service"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Body struct {
	URL string `json:"url"`
}

func CreateShortURLHandler(ctx *gin.Context) {
	request := new(Body)
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := service.CreateShortURLService(request.URL)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, data)
}

var hashParam string = "hash"

func RedirectURL(ctx *gin.Context) {
	URLHash := ctx.Param(hashParam)
	
	URLData, err := service.GetLongURL(URLHash)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	if time.Now().After(URLData.ExpireAt) {
		ctx.Redirect(http.StatusPermanentRedirect, URLData.LongURL)
	} else {
		ctx.Redirect(http.StatusPermanentRedirect, os.Getenv("CLIENT_URL")+"/invalid")
		// TODO: call a service to delete the current shortURL in the database
	}
}
