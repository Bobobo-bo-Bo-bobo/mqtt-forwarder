package main

import (
	"fmt"
)

func formatLogString(s string) string {
	return fmt.Sprintf("%-50s", s)
}
