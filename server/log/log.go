package log

import (
	// "file"
	"io"
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type Config struct {
	Level string
	Output string
}

func (cfg *Config) Init() Logger {
	var out io.Writer
	switch cfg.Output {
	case `stderr`:
		out = os.Stderr
	default:
		out = os.Stderr
	}
	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		zlog.Info().Msg(`failed to parse log level, defaulting to info`)
		level = zerolog.InfoLevel
	}
	return Logger{zerolog.New(out).Level(level).With().Timestamp().Logger()}
}

type Logger struct { 
	zerolog.Logger
}

func (l *Logger) WithComponent(component string) Logger {
	return Logger{l.With().Str(`component`, component).Logger()}
}