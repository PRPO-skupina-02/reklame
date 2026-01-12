package api

import (
	"net/http"

	"github.com/PRPO-skupina-02/common/middleware"
	_ "github.com/PRPO-skupina-02/reklame/api/docs"
	"github.com/PRPO-skupina-02/reklame/reklame"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Reklame API
//	@version		1.0
//	@description	API za upravljanje z kinodvoranami in njihovim sporedom

//	@host		localhost:8083
//	@BasePath	/api/v1/reklame

func Register(router *gin.Engine, trans ut.Translator, store *reklame.AdvertisementStore) {
	// Healthcheck
	router.GET("/healthcheck", healthcheck)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// REST API
	v1 := router.Group("/api/v1/reklame")
	v1.Use(middleware.TranslationMiddleware(trans))
	v1.Use(middleware.ErrorMiddleware)

	// Advertisements
	v1.GET("/advertisements/:theaterID", GetAdvertisements(store))
}

func healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
