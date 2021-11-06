package logger

func (log *Logger) Debugf(template string, args ...interface{}) {
	log.Debug().Msgf(template, args)
}

func (log *Logger) Infof(template string, args ...interface{}) {
	log.Info().Msgf(template, args)
}

func (log *Logger) Warnf(template string, args ...interface{}) {
	log.Warn().Msgf(template, args)
}

func (log *Logger) Errorf(template string, args ...interface{}) {
	log.Error().Msgf(template, args)
}
