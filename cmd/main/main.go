/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package main

import (
	"ciroque/go-http-server/internal/config"
	httpserver "ciroque/go-http-server/internal/http-server"
	"ciroque/go-http-server/internal/metrics"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	settings, err := config.NewSettings()
	if err != nil {
		logrus.Fatalf("Error creating configuration settings: %v", err)
	}

	logger := logrus.New()
	logger.SetLevel(settings.LogLevel)

	metricClient, err := metrics.NewMetrics()
	if err != nil {
		logger.Fatalf("Error creating metrics client: %v", err)
	}
	defer metricClient.Shutdown()

	abortChannel := make(chan string)
	defer close(abortChannel)

	httpServer := httpserver.Server{
		AbortChannel: abortChannel,
		Logger:       logrus.NewEntry(logger),
		Settings:     settings,
		Metrics:      &metricClient,
	}

	go httpServer.Run()

	sigTerm := make(chan os.Signal, 1)

	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	signal.Notify(sigTerm, syscall.SIGKILL)

	select {
	case <-sigTerm:
		{
			logger.Info("Exiting per SIGTERM")
		}
	case err := <-abortChannel:
		{
			logger.Error(err)
		}
	}
}
