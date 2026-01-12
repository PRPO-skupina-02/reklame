package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/PRPO-skupina-02/common/config"
	"github.com/PRPO-skupina-02/common/logging"
	"github.com/PRPO-skupina-02/common/validation"
	"github.com/PRPO-skupina-02/reklame/api"
	"github.com/PRPO-skupina-02/reklame/clients/spored/client"
	"github.com/PRPO-skupina-02/reklame/reklame"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
)

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	slog.Info("Starting server")

	logger := logging.GetDefaultLogger()
	slog.SetDefault(logger)

	trans, err := validation.RegisterValidation()
	if err != nil {
		return err
	}

	sporedHost := config.GetEnv("SPORED_HOST")
	transportConfig := client.DefaultTransportConfig().WithHost(sporedHost)
	sporedClient := client.NewHTTPClientWithConfig(strfmt.Default, transportConfig)

	store := reklame.NewAdvertisementStore()

	err = reklame.SetupCron(sporedClient, store)
	if err != nil {
		return err
	}

	router := gin.Default()
	api.Register(router, trans, store)

	slog.Info("Server startup complete")
	err = router.Run(":8083")
	if err != nil {
		return err
	}

	return nil
}
