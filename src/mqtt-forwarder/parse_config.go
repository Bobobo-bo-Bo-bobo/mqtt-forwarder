package main

import (
	ini "gopkg.in/ini.v1"
)

func parseConfigurationFile(f string) (Configuration, error) {
	var config Configuration
	var err error

	cfg, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, f)

	if err != nil {
		return config, err
	}

	src, err := cfg.GetSection("source")
	if err != nil {
		return config, err
	}
	err = src.MapTo(&config.Source)
	if err != nil {
		return config, err
	}

	if config.Source.Password != "" {
		config.Source.password = "**redacted**"
	}

	if config.Source.ClientID == "" {
		config.Source.ClientID, err = generateClientID()
		if err != nil {
			return config, err
		}
	}

	if config.Source.Timeout == 0 {
		config.Source.Timeout = 30
	}

	dst, err := cfg.GetSection("destination")
	if err != nil {
		return config, err
	}
	err = dst.MapTo(&config.Destination)
	if err != nil {
		return config, err
	}

	if config.Destination.Password != "" {
		config.Destination.password = "**redacted**"
	}

	if config.Destination.ClientID == "" {
		config.Destination.ClientID, err = generateClientID()
		if err != nil {
			return config, err
		}
	}

	if config.Destination.Timeout == 0 {
		config.Destination.Timeout = 30
	}

	return config, nil
}
