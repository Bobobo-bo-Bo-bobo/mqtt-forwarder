package main

const name = "mqtt-forwarder"
const version = "1.0.0"
const _url = "https://git.ypbind.de/cgit/mqtt-forwarder/"
const pkgname = name

const globalDefaultConfig = "/etc/mqtt-forwarder/config.ini"

const helpText = `Usage: mqtt-forwarder [--help] [--version] [--config=<cfg>] [--quiet]
    --config=<cfg>  Configuration file to use.
                    Default locations:
                      ~/.config/mqtt-forwader.ini
                      /etc/mqtt-forwarder/config.ini

    --help          This text

    --quiet         Only log warning and error messages

    --version       Show version information

`
const (
	_ uint = iota
	_
	_
	// MQTTv3_1 - use MQTT 3.1 protocol
	MQTTv3_1
	// MQTTv3_1_1 - use MQTT 3.1.1 protocol
	MQTTv3_1_1
)
