package logger

import (
	configs "github.com/fwchen/jellyfish/config"
	"github.com/juju/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	L             *Logger
	once          sync.Once
	sugaredLogger *zap.SugaredLogger
)

type Logger struct{}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	if sugaredLogger != nil {
		sugaredLogger.Debugw(msg, keysAndValues)
	}
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	if sugaredLogger != nil {
		sugaredLogger.Infow(msg, keysAndValues)
	}
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	if sugaredLogger != nil {
		sugaredLogger.Warnw(msg, keysAndValues)
	}
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	if sugaredLogger != nil {
		sugaredLogger.Errorw(msg, keysAndValues)
	}
}

func init() {
	once.Do(func() {
		initLogger()
	})
}

func initLogger() {
	L = &Logger{}
}

func InitLogger(loggerConfig configs.LoggerConfig) error {
	var level zapcore.Level
	err := level.UnmarshalText([]byte(loggerConfig.Level))
	if err != nil {
		return errors.Trace(err)
	}
	cfg := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Encoding:         "json",
		OutputPaths:      loggerConfig.OutputPaths,
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"service_name": "jellyfish_backend",
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Logger",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		return errors.Trace(err)
	}
	defer func() {
		_ = logger.Sync()
		//if err != nil {
		//	panic(err)
		//}
	}()

	sugaredLogger = logger.Sugar()
	return nil
}
