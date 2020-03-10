package main

import (
	ini "gopkg.in/ini.v1"
)

func parseConfigurationFile(f string) (*Configuration, error) {
	var config Configuration
	var err error

	cfg, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, f)

	if err != nil {
		return nil, err
	}

	src, err := cfg.GetSection("source")
	if err != nil {
		return nil, err
	}
	err = src.MapTo(&config.Source)
	if err != nil {
		return nil, err
	}

	if config.Source.Password != "" {
		config.Source.password = "**redacted**"
	}

	if config.Source.ClientID == "" {
		config.Source.ClientID, err = generateClientID()
		if err != nil {
			return nil, err
		}
	}

	dst, err := cfg.GetSection("destination")
	if err != nil {
		return nil, err
	}
	err = dst.MapTo(&config.Destination)
	if err != nil {
		return nil, err
	}

	if config.Destination.Password != "" {
		config.Destination.password = "**redacted**"
	}

	if config.Destination.ClientID == "" {
		config.Destination.ClientID, err = generateClientID()
		if err != nil {
			return nil, err
		}
	}

	return &config, nil
}
