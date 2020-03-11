package main

import (
	log "github.com/sirupsen/logrus"
)

func mqttDisconnect(cfg *MQTTConfiguration, quiet bool) {
	if !quiet {
		log.WithFields(log.Fields{
			"server":                 cfg.URL,
			"user":                   cfg.User,
			"password":               cfg.password,
			"client_certificate":     cfg.ClientCertificate,
			"client_certificate_key": cfg.ClientKey,
			"client_id":              cfg.ClientID,
			"ca_cert":                cfg.CACertificate,
			"qos":                    cfg.QoS,
			"topic":                  cfg.Topic,
		}).Info(formatLogString("Disconnecting from MQTT message broker"))
	}

	if cfg.mqttClient.IsConnected() {
		cfg.mqttClient.Disconnect(100)
	} else {
		log.WithFields(log.Fields{
			"server":                 cfg.URL,
			"user":                   cfg.User,
			"password":               cfg.password,
			"client_certificate":     cfg.ClientCertificate,
			"client_certificate_key": cfg.ClientKey,
			"client_id":              cfg.ClientID,
			"ca_cert":                cfg.CACertificate,
			"qos":                    cfg.QoS,
			"topic":                  cfg.Topic,
		}).Warn(formatLogString("Client is not connected to MQTT message broker"))
	}
}
