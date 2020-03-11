package main

import (
	"crypto/tls"
	"crypto/x509"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func mqttConnect(cfg *MQTTConfiguration, quiet bool) mqtt.Client {
	var tlsCfg = new(tls.Config)

	mqttOptions := mqtt.NewClientOptions()
	mqttOptions.AddBroker(cfg.URL)

	if cfg.User != "" && cfg.Password != "" {
		mqttOptions.SetUsername(cfg.User)
		mqttOptions.SetPassword(cfg.Password)
	} else if cfg.ClientCertificate != "" && cfg.ClientKey != "" {
		cert, err := tls.LoadX509KeyPair(cfg.ClientCertificate, cfg.ClientKey)
		if err != nil {
			log.WithFields(log.Fields{
				"error":                  err.Error(),
				"client_certificate":     cfg.ClientCertificate,
				"client_certificate_key": cfg.ClientKey,
			}).Fatal(formatLogString("Can't read SSL client certificate"))
		}

		tlsCfg.Certificates = make([]tls.Certificate, 1)
		tlsCfg.Certificates[0] = cert

		mqttOptions.SetTLSConfig(tlsCfg)
	} else {
		log.WithFields(log.Fields{
			"client_certificate":     cfg.ClientCertificate,
			"client_certificate_key": cfg.ClientKey,
			"user":                   cfg.User,
			"password":               cfg.password,
		}).Fatal(formatLogString("Unable to determine authentication method"))
	}

	if cfg.InsecureSSL {
		tlsCfg.InsecureSkipVerify = true
	}

	if cfg.CACertificate != "" {
		tlsCfg.RootCAs = x509.NewCertPool()
		cacert, err := ioutil.ReadFile(cfg.CACertificate)
		if err != nil {
			log.WithFields(log.Fields{
				"error":   err.Error(),
				"ca_cert": cfg.CACertificate,
			}).Fatal(formatLogString("Can't read CA certificate"))
		}

		tlsok := tlsCfg.RootCAs.AppendCertsFromPEM(cacert)
		if !tlsok {
			log.WithFields(log.Fields{
				"error":   err.Error(),
				"ca_cert": cfg.CACertificate,
			}).Fatal(formatLogString("Can't add CA certificate to x509.CertPool"))
		}
	}

	mqttOptions.SetClientID(cfg.ClientID)
	mqttOptions.SetAutoReconnect(true)
	mqttOptions.SetConnectRetry(true)
	mqttOptions.SetProtocolVersion(MQTTv3_1_1)

	mqttClient := mqtt.NewClient(mqttOptions)

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
			"insecure_ssl":           cfg.InsecureSSL,
		}).Info(formatLogString("Connecting to MQTT message broker"))
	}

	mqttToken := mqttClient.Connect()
	mqttToken.Wait()

	if mqttToken.Error() != nil {
		log.WithFields(log.Fields{
			"error":                  mqttToken.Error(),
			"server":                 cfg.URL,
			"user":                   cfg.User,
			"password":               cfg.password,
			"client_certificate":     cfg.ClientCertificate,
			"client_certificate_key": cfg.ClientKey,
			"client_id":              cfg.ClientID,
			"ca_cert":                cfg.CACertificate,
			"qos":                    cfg.QoS,
			"topic":                  cfg.Topic,
			"insecure_ssl":           cfg.InsecureSSL,
		}).Fatal(formatLogString("Connection to MQTT message broker failed"))
	}

	return mqttClient
}
