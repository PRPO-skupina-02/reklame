package reklame

import (
	"sync"

	"github.com/PRPO-skupina-02/reklame/clients/spored/models"
)

type MovieWithTimeslots struct {
	Movie     *models.APIMovieResponse      `json:"movie"`
	Timeslots []*models.APITimeSlotResponse `json:"timeslots"`
}

type AdvertisementStore struct {
	mu     sync.RWMutex
	movies map[string][]MovieWithTimeslots
}

func NewAdvertisementStore() *AdvertisementStore {
	return &AdvertisementStore{
		movies: make(map[string][]MovieWithTimeslots),
	}
}

func (s *AdvertisementStore) SetAdvertisements(theaterID string, movies []MovieWithTimeslots) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.movies[theaterID] = movies
}

func (s *AdvertisementStore) GetAdvertisements(theaterID string) []MovieWithTimeslots {
	s.mu.RLock()
	defer s.mu.RUnlock()

	movies, exists := s.movies[theaterID]
	if !exists {
		return []MovieWithTimeslots{}
	}

	result := make([]MovieWithTimeslots, len(movies))
	copy(result, movies)

	return result
}
