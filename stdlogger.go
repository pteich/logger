package logger

import (
	"github.com/rs/zerolog"
)

func (log *Logger) Debugf(template string, args ...interface{}) {
	if len(args) > 0 {
		log.Debug().Msg(template)
		return
	}

	log.Debug().Msgf(template, args...)
}

func (log *Logger) Infof(template string, args ...interface{}) {
	if len(args) > 0 {
		log.Info().Msg(template)
		return
	}

	log.Info().Msg(template)
}

func (log *Logger) Warnf(template string, args ...interface{}) {
	if len(args) > 0 {
		log.Warn().Msg(template)
		return
	}

	log.Warn().Msgf(template, args...)
}

func (log *Logger) Errorf(template string, args ...interface{}) {
	if len(args) > 0 {
		log.Error().Msg(template)
		return
	}

	log.Error().Msgf(template, args...)
}

func (log *Logger) Fatalf(template string, args ...interface{}) {
	if len(args) > 0 {
		log.Fatal().Msg(template)
		return
	}

	log.Fatal().Msgf(template, args...)
}

func (log *Logger) Tracef(template string, args ...interface{}) {
	if len(args) > 0 {
		log.Trace().Msg(template)
		return
	}

	log.Trace().Msgf(template, args...)
}

func (log *Logger) SetLevel(lvl int) {
	newLogger := NewFromZerolog(log.Level(zerolog.Level(lvl)))
	*log = *newLogger
}
