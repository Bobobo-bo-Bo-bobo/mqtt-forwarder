package main

import (
	"os"
	"os/user"
	"path/filepath"
)

func getUserDefaultConfig() (string, error) {
	var loc string

	env := getEnvironment()

	/* XDG Base Directory Specification:
	 *  XDG_CONFIG_HOME defines the base directory relative to which user specific
	 *  configuration files should be stored. If $XDG_CONFIG_HOME is either not set
	 *  or empty, a default equal to $HOME/.config should be used.
	 *
	 * see: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
	 */
	loc, found := env["XDG_CONFIG_HOME"]
	if found {
		return filepath.Join(loc, pkgname, "config.ini"), nil
	}

	home, found := env["HOME"]
	if found {
		return filepath.Join(home, ".config", pkgname, "config.ini"), nil
	}

	_user, err := user.Current()
	if err != nil {
		return "", err
	}
	if _user.HomeDir != "" {
		return filepath.Join(_user.HomeDir, ".config", pkgname, "config.ini"), nil
	}
	return filepath.Join("/home", _user.Username, ".config", pkgname, "config.ini"), nil
}

func getDefaultConfig() string {
	var err error
	var lcfg string

	// if a user configuration exists read it
	lcfg, err = getUserDefaultConfig()
	if err == nil {
		_, err = os.Open(lcfg)
		if err == nil {
			return lcfg
		}
	}

	// otherwise look for the global configuration
	_, err = os.Open(globalDefaultConfig)
	if err == nil {
		return globalDefaultConfig
	}

	// no configuration
	return ""
}
