package reklame

import (
	"log/slog"
	"time"

	"github.com/PRPO-skupina-02/reklame/clients/spored/client"
	"github.com/go-co-op/gocron/v2"
)

func SetupCron(sporedClient *client.Spored, store *AdvertisementStore) error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	// Fetch on startup
	_, err = s.NewJob(
		gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(time.Now().Add(time.Second*10))),
		gocron.NewTask(RefreshAdvertisements, sporedClient, store),
	)
	if err != nil {
		return err
	}

	// Daily schedule at midnight
	j, err := s.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0))),
		gocron.NewTask(RefreshAdvertisements, sporedClient, store),
	)
	if err != nil {
		return err
	}

	s.Start()

	slog.Info("Cron job started", "id", j.ID())
	return nil
}
