package library

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLog() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	Log = zerolog.New(output).With().Timestamp().Str("role", filepath.Base(os.Args[0])).Logger()
}
