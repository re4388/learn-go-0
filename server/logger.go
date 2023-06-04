package server

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Run_Log() {

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// log.Print("learn package is loaded")
	log.Info().Msg("server package is loaded")
	// 	Str("Name", "Tom").
	// 	Send()

	// other APIs:
	// panic(zerolog.PanicLevel, 5)
	// fatal(zerolog.FatalLevel, 4)
	// error(zerolog.ErrorLevel, 3)
	// warn(zerolog.WarnLevel, 2)
	// info(zerolog.InfoLevel, 1)
	// debug(zerolog.DebugLevel, 0)
	// trace(zerolog.TraceLevel, -1)

}
