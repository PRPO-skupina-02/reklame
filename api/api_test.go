package api

import (
	"testing"

	"github.com/PRPO-skupina-02/common/validation"
	"github.com/PRPO-skupina-02/reklame/reklame"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestingRouter(t *testing.T) *gin.Engine {
	router := gin.Default()
	trans, err := validation.RegisterValidation()
	require.NoError(t, err)
	store := reklame.NewAdvertisementStore()
	Register(router, trans, store)

	return router
}
