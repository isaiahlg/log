package log

var Log *Logger

func BuildLogger() *Logger {
	Log = &Logger{}
	return Log
}

func Error(msg string, s interface{}) {
	for _, h := range Log.hooks {
		h.Error(msg, s)
	}
}

func Fatal(msg string, s interface{}) {
	for _, h := range Log.hooks {
		h.Fatal(msg, s)
	}
}
