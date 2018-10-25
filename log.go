package log

var Log *Logger
var applicationName string

func BuildLogger(app string) *Logger {
	Log = &Logger{}
	applicationName = app
	return Log
}

func Error(msg string, s interface{}) {
	msg = applicationName + ": " + msg
	for _, h := range Log.hooks {
		h.Error(msg, s)
	}
}

func Fatal(msg string, s interface{}) {
	msg = applicationName + ": " + msg
	for _, h := range Log.hooks {
		h.Fatal(msg, s)
	}
}
