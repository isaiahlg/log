package log

import (
	"github.com/muhoro/log/seq"
)

type hook interface {
	Error(msg string, s interface{})
	Fatal(msg string, s interface{})
}

type Logger struct {
	hooks []hook
}

func (logger *Logger) UseFile(filePath string) *Logger {
	// to be implemented
	return logger
}

func (logger *Logger) UseSeq(url, apikey string) *Logger {
	seqHook := &seq.SeqHook{
		BaseUrl: url,
		ApiKey:  apikey,
	}
	logger.hooks = append(logger.hooks, seqHook)
	return logger
}

func (logger *Logger) UseConsole() *Logger {
	// to be implemented
	return logger
}
