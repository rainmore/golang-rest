package log

import (
	"go.uber.org/zap"
	"os"
)

var (
	initialized                    = false
	logger      *zap.SugaredLogger = nil
)

func InitLog() {
	if !initialized {
		logEnv := zap.Must(zap.NewDevelopment())
		if os.Getenv("APP_ENV") == "prod" {
			logEnv = zap.Must(zap.NewProduction())
		}
		zap.ReplaceGlobals(logEnv)
		logger = zap.S()
		initialized = true
	}
}

func PanicHandler(logger *zap.Logger) {
	if r := recover(); r != nil {
		logger.Panic("caught panic", zap.Any("panicDetails", r))
		// Don't swallow the panic
		panic(r)
	}
}

func Logger() *zap.SugaredLogger {
	InitLog()
	return logger
}
