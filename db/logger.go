package db

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run_dbLog() {

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// log.Print("learn package is loaded")
	log.Info().Msg("db package is loaded")
}
