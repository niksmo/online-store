package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Instance zerolog.Logger

func Init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	zerolog.MessageFieldName = "msg"
	Instance = zerolog.New(os.Stdout).With().Timestamp().Logger()
}
