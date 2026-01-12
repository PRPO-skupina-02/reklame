package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/PRPO-skupina-02/common/config"
	"github.com/PRPO-skupina-02/common/logging"
	"github.com/PRPO-skupina-02/common/validation"
	"github.com/PRPO-skupina-02/reklame/api"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/go-swagger/go-swagger/examples/stream-client/client"
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

	router := gin.Default()
	api.Register(router, trans)

	slog.Info("Server startup complete")
	err = router.Run(":8081")
	if err != nil {
		return err
	}

	return nil
}
