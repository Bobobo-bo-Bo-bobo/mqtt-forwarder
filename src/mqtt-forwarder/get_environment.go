package main

import (
	"os"
	"strings"
)

func getEnvironment() map[string]string {
	var env = make(map[string]string)

	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		key := kv[0]
		value := kv[1]
		env[key] = value
	}

	return env
}
