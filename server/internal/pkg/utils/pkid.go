package utils

import (
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
)

func noDash(u uuid.UUID) string {
	return hex.EncodeToString(u[:])
}

func PKID() string {
	v7, err := uuid.NewV7()
	if err != nil {
		return strings.ReplaceAll(uuid.NewString(), "-", "")
	} else {
		return hex.EncodeToString(v7[:])
	}
}
