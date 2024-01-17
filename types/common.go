package types

import (
	"log"
)

func handleNotSupported(solidityType string, actualType string) string {
	log.Printf("Non-supported type %s not handled. Actual type: %s", solidityType, actualType)
	return ""
}
