package logger

import "go.uber.org/zap"

var log *zap.SugaredLogger

func Init(serviceName string, isDev bool) {
	var cfg zap.Config
	if isDev {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	cfg.EncoderConfig.TimeKey = "ts"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.NameKey = "logger"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.StacktraceKey = "stacktrace"

	baseLogger, err := cfg.Build(zap.AddCaller(), zap.Fields(zap.String("service", serviceName)))
	if err != nil {
		panic("cannot init logger: " + err.Error())
	}

	log = baseLogger.Sugar()
}

func Info(msg string, fields ...any) {
	log.Infow(msg, fields...)
}

func Error(msg string, err error, fields ...any) {
	log.Errorw(msg, append(fields, "error", err)...)
}

func Warn(msg string, fields ...any) {
	log.Warnw(msg, fields...)
}

func Debug(msg string, fields ...any) {
	log.Debugw(msg, fields...)
}
