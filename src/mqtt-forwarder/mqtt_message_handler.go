package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

func mqttMessageHandler(c mqtt.Client, msg mqtt.Message) {
	var destTopic string
	var token mqtt.Token

	if !configuration.Quiet {
		log.WithFields(log.Fields{
			"topic":      msg.Topic(),
			"duplicate":  msg.Duplicate(),
			"qos":        uint(msg.Qos()),
			"retained":   msg.Retained(),
			"message_id": msg.MessageID(),
			"broker_url": configuration.Source.URL,
			"msg_length": len(msg.Payload()),
		}).Info(formatLogString("Received MQTT message from source broker"))
	}

	destTopic = filepath.Join(configuration.Destination.Topic, msg.Topic()) + "/"
	token = configuration.Destination.mqttClient.Publish(destTopic, msg.Qos(), msg.Retained(), msg.Payload())

	if token.Error() != nil {
		log.WithFields(log.Fields{
			"topic":      destTopic,
			"duplicate":  msg.Duplicate(),
			"qos":        uint(msg.Qos()),
			"retained":   msg.Retained(),
			"message_id": msg.MessageID(),
			"broker_url": configuration.Destination.URL,
			"msg_length": len(msg.Payload()),
			"error":      token.Error(),
		}).Error(formatLogString("MQTT message forwarding to destination MQTT broker failed"))
	} else {
		if !configuration.Quiet {
			log.WithFields(log.Fields{
				"topic":      destTopic,
				"duplicate":  msg.Duplicate(),
				"qos":        uint(msg.Qos()),
				"retained":   msg.Retained(),
				"message_id": msg.MessageID(),
				"broker_url": configuration.Destination.URL,
				"msg_length": len(msg.Payload()),
			}).Info(formatLogString("MQTT message forwarded to destination MQTT broker"))
		}
	}
}
