package reklame

import (
	"log/slog"
	"time"

	"github.com/PRPO-skupina-02/reklame/clients/spored/client"
	"github.com/PRPO-skupina-02/reklame/clients/spored/client/movies"
	"github.com/PRPO-skupina-02/reklame/clients/spored/client/rooms"
	"github.com/PRPO-skupina-02/reklame/clients/spored/client/theaters"
	"github.com/PRPO-skupina-02/reklame/clients/spored/client/timeslots"
	"github.com/PRPO-skupina-02/reklame/clients/spored/models"
	"github.com/go-openapi/strfmt"
)

func RefreshAdvertisements(sporedClient *client.Spored, store *AdvertisementStore) {
	slog.Info("Starting advertisement refresh")

	theatersResp, err := sporedClient.Theaters.TheatersList(theaters.NewTheatersListParams())
	if err != nil {
		slog.Error("Failed to fetch theaters", "err", err)
		return
	}

	tomorrow := getTomorrow()

	for _, theater := range theatersResp.Payload.Data {
		fetchTheaterRooms(sporedClient, store, theater.ID, tomorrow)
	}

	slog.Info("Advertisements successfully refreshed")
}

func getTomorrow() time.Time {
	now := time.Now()

	tomorrow := now.Add(24 * time.Hour).Truncate(24 * time.Hour)

	return tomorrow
}

func fetchTheaterRooms(sporedClient *client.Spored, store *AdvertisementStore, theaterID string, tomorrow time.Time) {
	theaterUUID := strfmt.UUID(theaterID)
	params := rooms.NewRoomsListParams().WithTheaterID(theaterUUID)

	roomsResp, err := sporedClient.Rooms.RoomsList(params)
	if err != nil {
		slog.Warn("Failed to fetch rooms", "theater_id", theaterID, "err", err)
		return
	}

	moviesMap := map[string]*MovieWithTimeslots{}

	for _, room := range roomsResp.Payload.Data {
		fetchRoomTimeslots(sporedClient, moviesMap, theaterUUID, room.ID, tomorrow)
	}

	advertisements := []MovieWithTimeslots{}
	for _, movie := range moviesMap {
		advertisements = append(advertisements, *movie)
	}

	slog.Debug("Advertisements collected", "theaterID", theaterID, "advertisements", advertisements)
	store.SetAdvertisements(theaterID, advertisements)
}

func fetchRoomTimeslots(sporedClient *client.Spored, moviesMap map[string]*MovieWithTimeslots, theaterID strfmt.UUID, roomID string, tomorrow time.Time) {
	roomUUID := strfmt.UUID(roomID)

	date := strfmt.Date(tomorrow)

	params := timeslots.NewTimeSlotsListParams().WithTheaterID(theaterID).WithRoomID(roomUUID).WithDate(&date)

	timeslotsResp, err := sporedClient.Timeslots.TimeSlotsList(params)
	if err != nil {
		slog.Warn("Failed to fetch timeslots", "theater_id", theaterID, "room_id", roomID, "err", err)
		return
	}

	for _, timeslot := range timeslotsResp.Payload.Data {
		fetchTimeSlotMovie(sporedClient, moviesMap, timeslot)
	}
}

func fetchTimeSlotMovie(sporedClient *client.Spored, moviesMap map[string]*MovieWithTimeslots, timeslot *models.APITimeSlotResponse) {
	movieID := timeslot.MovieID

	params := movies.NewMoviesShowParams().WithMovieID(strfmt.UUID(movieID))

	if _, ok := moviesMap[movieID]; !ok {
		movieResp, err := sporedClient.Movies.MoviesShow(params)
		if err != nil {
			slog.Warn("Failed to fetch movie details", "movie_id", movieID, "err", err)
			return
		}

		moviesMap[movieID] = &MovieWithTimeslots{
			Movie:     movieResp.Payload,
			Timeslots: []*models.APITimeSlotResponse{},
		}
	}

	moviesMap[movieID].Timeslots = append(moviesMap[movieID].Timeslots, timeslot)
}
