package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PRPO-skupina-02/common/xtesting"
	"github.com/PRPO-skupina-02/reklame/clients/spored/models"
	"github.com/PRPO-skupina-02/reklame/reklame"
	"github.com/stretchr/testify/assert"
)

func TestGetAdvertisements(t *testing.T) {
	store := reklame.NewAdvertisementStore()
	store.SetAdvertisements("fb126c8c-d059-11f0-8fa4-b35f33be83b7", []reklame.MovieWithTimeslots{
		{
			Movie: &models.APIMovieResponse{
				ID:            "aa1234567-0123-0123-0123-0123456789ab",
				Name:          "Test Movie",
				Description:   "Test Description",
				ImageURL:      "https://example.com/image.jpg",
				Rating:        8.5,
				LengthMinutes: 120,
				Active:        true,
			},
			Timeslots: []*models.APITimeSlotResponse{
				{
					ID:        "bb1234567-0123-0123-0123-0123456789ab",
					StartTime: "2026-01-13T18:00:00Z",
					EndTime:   "2026-01-13T20:00:00Z",
					RoomID:    "cc1234567-0123-0123-0123-0123456789ab",
					MovieID:   "aa1234567-0123-0123-0123-0123456789ab",
				},
				{
					ID:        "dd1234567-0123-0123-0123-0123456789ab",
					StartTime: "2026-01-13T20:30:00Z",
					EndTime:   "2026-01-13T22:30:00Z",
					RoomID:    "ee1234567-0123-0123-0123-0123456789ab",
					MovieID:   "aa1234567-0123-0123-0123-0123456789ab",
				},
			},
		},
		{
			Movie: &models.APIMovieResponse{
				ID:            "ff1234567-0123-0123-0123-0123456789ab",
				Name:          "Another Movie",
				Description:   "Another Description",
				ImageURL:      "https://example.com/image2.jpg",
				Rating:        7.5,
				LengthMinutes: 100,
				Active:        true,
			},
			Timeslots: []*models.APITimeSlotResponse{
				{
					ID:        "gg1234567-0123-0123-0123-0123456789ab",
					StartTime: "2026-01-13T19:00:00Z",
					EndTime:   "2026-01-13T20:40:00Z",
					RoomID:    "hh1234567-0123-0123-0123-0123456789ab",
					MovieID:   "ff1234567-0123-0123-0123-0123456789ab",
				},
			},
		},
	})
	r := TestingRouter(t, store)

	tests := []struct {
		name      string
		status    int
		theaterID string
	}{
		{
			name:      "ok",
			status:    http.StatusOK,
			theaterID: "fb126c8c-d059-11f0-8fa4-b35f33be83b7",
		},
		{
			name:      "ok-no-advertisements",
			status:    http.StatusOK,
			theaterID: "bae209f6-d059-11f0-b2a4-cbf992c2eb6d",
		},
		{
			name:      "ok-empty-theater",
			status:    http.StatusOK,
			theaterID: "ea0b7f96-ddc9-11f0-9635-23efd36396bd",
		},
		{
			name:      "invalid-theater-id",
			status:    http.StatusOK,
			theaterID: "01234567-0123-0123-0123-0123456789ab",
		},
		{
			name:      "nil-theater-id",
			status:    http.StatusBadRequest,
			theaterID: "00000000-0000-0000-0000-000000000000",
		},
		{
			name:      "malformed-theater-id",
			status:    http.StatusBadRequest,
			theaterID: "000",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {

			targetURL := fmt.Sprintf("/api/v1/reklame/advertisements/%s", testCase.theaterID)

			req := xtesting.NewTestingRequest(t, targetURL, http.MethodGet, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.status, w.Code)
			xtesting.AssertGoldenJSON(t, w)
		})
	}
}
