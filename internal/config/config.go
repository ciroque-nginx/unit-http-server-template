/*
 * Copyright Steve Wagner. All rights reserved.
 * Use of this source code is governed by the Apache License that can be found in the LICENSE file.
 */

package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

const (
	DefaultHost                 = "0.0.0.0"
	DefaultPort                 = "8888"
	HostEnvironmentVariable     = "BHS_HTTP_HOST"
	LogLevelEnvironmentVariable = "BHS_LOG_LEVEL"
	PortEnvironmentVariable     = "BHS_HTTP_PORT"
)

type Settings struct {
	Host     string
	Port     int
	LogLevel logrus.Level
}

func NewSettings() (*Settings, error) {
	port := os.Getenv(PortEnvironmentVariable)
	if port == "" {
		port = DefaultPort
	}

	portNumber, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("unable to parse SRES_PORT: %v", err)
	}

	host := os.Getenv(HostEnvironmentVariable)
	if host == "" {
		host = DefaultHost
	}

	logLevel := getLogLevel()

	config := &Settings{host, portNumber, logLevel}

	return config, nil
}

func getLogLevel() logrus.Level {
	logLevel := os.Getenv(LogLevelEnvironmentVariable)
	logrus.Debugf("Settings::setLogLevel: %s", logLevel)
	switch logLevel {
	case "panic":
		return logrus.PanicLevel

	case "fatal":
		return logrus.FatalLevel

	case "error":
		return logrus.ErrorLevel

	case "warn":
		return logrus.WarnLevel

	case "info":
		return logrus.InfoLevel

	case "debug":
		return logrus.DebugLevel

	case "trace":
		return logrus.TraceLevel

	default:
		return logrus.InfoLevel
	}
}
