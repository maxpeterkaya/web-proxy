package main

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/rs/zerolog/log"
)

func GenerateKey(n int) string {
	s := make([]byte, n)

	_, err := rand.Read(s)
	if err != nil {
		log.Error().Err(err).Msg("Error generating random password")
		return ""
	}

	return base64.StdEncoding.EncodeToString(s)
}
