package api

import (
	"net/http"

	"github.com/PRPO-skupina-02/common/request"
	"github.com/PRPO-skupina-02/reklame/reklame"
	"github.com/gin-gonic/gin"
)

// GetAdvertisements
//
//	@Id				GetAdvertisements
//	@Summary		Get advertisements
//	@Description	Get a list of movies and their showing times for tomorrow for a specific theater
//	@Tags			advertisements
//	@Accept			json
//	@Produce		json
//	@Param			theaterID	path		string	true	"Theater ID"	Format(uuid)
//	@Success		200			{object}	[]reklame.MovieWithTimeslots
//	@Failure		400			{object}	middleware.HttpError
//	@Failure		500			{object}	middleware.HttpError
//	@Router			/advertisements/{theaterID} [get]
func GetAdvertisements(store *reklame.AdvertisementStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		theaterID, err := request.GetUUIDParam(c, "theaterID")
		if err != nil {
			_ = c.Error(err)
			return
		}

		advertisements := store.GetAdvertisements(theaterID.String())
		c.JSON(http.StatusOK, advertisements)
	}
}
