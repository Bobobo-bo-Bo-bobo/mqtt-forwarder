package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

func mqttConnectionLostHandler(c mqtt.Client, err error) {
	r := c.OptionsReader()
	_url := r.Servers()
	url := "?"

	if _url != nil {
		url = _url[0].String()
	}
	log.WithFields(log.Fields{
		"error":  err,
		"server": url,
	}).Warn(formatLogString("MQTT connection lost, reconnecting"))

	mqttToken := c.Connect()
	mqttToken.Wait()

	if mqttToken.Error() != nil {
		log.WithFields(log.Fields{
			"error":  mqttToken.Error(),
			"server": url,
		}).Fatal(formatLogString("MQTT reconnect failed"))
	}
}
