package main

import (
	"fmt"
	"strings"
)

func validateConfiguration(cfg Configuration) error {
	if cfg.Source.Topic == "" {
		return fmt.Errorf("Topic of source MQTT broker is not set")
	}

	if strings.Index(cfg.Destination.Topic, "+") != -1 || strings.Index(cfg.Destination.Topic, "#") != -1 {
		return fmt.Errorf("Topic of destination MQTT broker can't contain MQTT wildcards")
	}

	if cfg.Source.QoS > 2 {
		return fmt.Errorf("QoS of source MQTT broker must be 0, 1 or 2")
	}
	if cfg.Destination.QoS > 2 {
		return fmt.Errorf("QoS of destination MQTT broker must be 0, 1 or 2")
	}

	return nil
}
