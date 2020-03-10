package main

import (
	uuid "github.com/nu7hatch/gouuid"
)

func generateClientID() (string, error) {
	_uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return _uuid.String(), nil
}
