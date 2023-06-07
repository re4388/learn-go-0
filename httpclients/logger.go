package httpclients

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run_Log() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("httpclients package is loaded")

}
