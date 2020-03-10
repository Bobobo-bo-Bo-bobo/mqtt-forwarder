package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Configuration - Configuration data
type Configuration struct {
	Source      MQTTConfiguration
	Destination MQTTConfiguration
}

// MQTTConfiguration - MQTT configuration
type MQTTConfiguration struct {
	URL               string `ini:"url"`
	User              string `ini:"user"`
	Password          string `ini:"password"`
	password          string
	ClientCertificate string `ini:"cert"`
	ClientKey         string `ini:"key"`
	CACertificate     string `ini:"ca_cert"`
	Topic             string `ini:"topic"`
	ClientID          string `ini:"client_id"`
	QoS               uint   `ini:"qos"`
	mqttClient        mqtt.Client
}
