package main

import (
	"os"

	"github.com/allgor-data/backend/cmd"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("Could not find .env file")
	}
}

func main() {
	err := cmd.Run(os.Args)
	if err != nil {
		log.Err(err).Msg("Failed to run app")
		os.Exit(1)
	}
}
