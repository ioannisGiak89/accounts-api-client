package utils

import (
	"github.com/google/uuid"
	"log"
)

// Parses a UUID. Util function to avoid dublication
func ParseUuid(id string) uuid.UUID {
	uID, err := uuid.Parse(id)

	if err != nil {
		log.Fatal(err)
	}

	return uID
}
