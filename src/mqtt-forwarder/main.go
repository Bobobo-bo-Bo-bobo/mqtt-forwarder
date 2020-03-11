package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Note: The message handler only accepts mqtt.Client and mqtt.Message as arguments. So we have
//       to declare the configuration as global variable
var configuration Configuration

func main() {
	var help = flag.Bool("help", false, "Show help")
	var version = flag.Bool("version", false, "Show version information")
	var config = flag.String("config", "", "Configuration file")
	var quiet = flag.Bool("quiet", false, "Suppress info messages")
	var configFile string
	var logFmt = new(log.TextFormatter)
	var sigChannel = make(chan os.Signal, 1)
	var err error

	signal.Notify(sigChannel, os.Interrupt, os.Kill, syscall.SIGTERM)

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

	if *quiet {
		configuration.Quiet = true
	}

	configuration, err = parseConfigurationFile(configFile)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal(formatLogString("Configuration parsing failed"))
	}

	err = validateConfiguration(configuration)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal(formatLogString("Invalid configuration"))
	}

	configuration.Source.mqttClient = mqttConnect(&configuration.Source, *quiet)
	configuration.Destination.mqttClient = mqttConnect(&configuration.Destination, *quiet)

	// source connection subscribes to the configured topic and listen for messages
	configuration.Source.mqttClient.Subscribe(configuration.Source.Topic, byte(configuration.Source.QoS), mqttMessageHandler)

	// wait for OS signal to arrive
	<-sigChannel

	mqttDisconnect(&configuration.Source, *quiet)
	mqttDisconnect(&configuration.Destination, *quiet)

	os.Exit(0)
}
