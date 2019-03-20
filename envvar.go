package main

import (
	"log"
	"os"
)

var (
	trueStrings = []string{
		"true", "1", "yes", "tak",
	}
)

// mustGetEnv gets sets value or sets it to default when not set
func mustGetEnv(key, fallbackValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	if fallbackValue == "" {
		log.Fatalf("Required env var (%s) not set", key)
	}

	return fallbackValue
}
