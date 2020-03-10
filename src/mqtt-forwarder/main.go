package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	var help = flag.Bool("help", false, "Show help")
	var version = flag.Bool("version", false, "Show version information")
	var config = flag.String("config", "", "Configuration file")
	var configFile string
	var logFmt = new(log.TextFormatter)

	logFmt.FullTimestamp = true
	logFmt.TimestampFormat = time.RFC3339
	log.SetFormatter(logFmt)

	flag.Usage = showUsage
	flag.Parse()
	if *help {
		showUsage()
		os.Exit(0)
	}

	if *version {
		showVersion()
		os.Exit(0)
	}

	if *config != "" {
		configFile = *config
	} else {
		configFile = getDefaultConfig()
	}

	if configFile == "" {
		log.Fatal(formatLogString("No configuration file found"))
	}

	configuration, err := parseConfigurationFile(configFile)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal(formatLogString("Configuration parsing failed"))
	}

	configuration.Source.mqttClient = mqttConnect(&configuration.Source)
	configuration.Destination.mqttClient = mqttConnect(&configuration.Destination)
}
