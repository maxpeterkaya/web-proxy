package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"

	"github.com/rs/zerolog/log"
)

func SHA256Hash(text string) string {
	h := sha256.New()
	if _, err := io.WriteString(h, text); err != nil {
		log.Error().Err(err).Msg("SHA256Hash")
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))

}
