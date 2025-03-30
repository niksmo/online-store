package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Instance zerolog.Logger

func Init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	Instance = zerolog.New(os.Stderr)
}
