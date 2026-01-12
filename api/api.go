package api

import (
	"net/http"

	"github.com/PRPO-skupina-02/common/middleware"
	_ "github.com/PRPO-skupina-02/reklame/api/docs"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Nakup API
//	@version		1.0
//	@description	API za upravljanje z kinodvoranami in njihovim sporedom

//	@host		localhost:8081
//	@BasePath	/api/v1/reklame

func Register(router *gin.Engine, trans ut.Translator) {
	// Healthcheck
	router.GET("/healthcheck", healthcheck)

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// REST API
	v1 := router.Group("/api/v1/reklame")
	v1.Use(middleware.TranslationMiddleware(trans))
	v1.Use(middleware.ErrorMiddleware)
}

func healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
