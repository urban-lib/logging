package logging

// Debug formatting string
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof formatting string
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf formatting string
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf formatting string
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf formatting string
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Panicf formatting string
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// WithFields ...
func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}
