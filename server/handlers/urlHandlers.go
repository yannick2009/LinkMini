package handlers

import (
	"linkmini/service"
	"net/http"

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

func RedirectURL(ctx *gin.Context) {
	URLHash := ctx.Param("hash")
	longURL, err := service.GetLongURL(URLHash)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ctx.Redirect(http.StatusPermanentRedirect, longURL)
}
