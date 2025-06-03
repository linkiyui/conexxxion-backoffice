package logger_v2

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/conexxxion/conexxxion-backoffice/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init() {
	if logger != nil {
		return
	}
	var err error
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "lvl",
		TimeKey:      "ts",
		EncodeTime:   zapcore.RFC3339TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		CallerKey:    "func",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	loggerCfg := zap.Config{
		Development:   config.GetConfig().IsDev(),
		DisableCaller: false,
		Encoding:      "json",
		EncoderConfig: encoderConfig,
		OutputPaths:   []string{},
	}
	switch config.GetConfig().LogLevel {
	case "debug":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "fatal":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	case "panic":
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		loggerCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, err = loggerCfg.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	loggerOut := NewLoggerOutput(false, "")
	if config.GetConfig().Verbose {
		loggerOut.SetConsoleOutput()
	}
	if config.GetConfig().LogsFile != "" {
		loggerOut.SetLumberjackLogger(config.GetConfig().LogsFile)
	}
	logger.WithOptions()
	logger = logger.WithOptions(
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderConfig),
				zapcore.AddSync(loggerOut),
				core,
			)
		}),
		zap.AddCallerSkip(2),
	)
}

func Finish() {
	if logger != nil {
		fmt.Println("terminating logger")
		logger.Sync()
	}
}

func Info(msg string, fields map[string]any) {
	_log(context.Background(), "info", msg, fields)
}

func InfoCtx(ctx context.Context, msg string, fields map[string]any) {
	_log(ctx, "info", msg, fields)
}

func Error(msg string, fields map[string]any) {
	_log(context.Background(), "error", msg, fields)
}

func ErrorCtx(ctx context.Context, msg string, fields map[string]any) {
	_log(ctx, "error", msg, fields)
}

func Fatal(msg string, fields map[string]any) {
	_log(context.Background(), "fatal", msg, fields)
}

func FatalCtx(ctx context.Context, msg string, fields map[string]any) {
	_log(ctx, "fatal", msg, fields)
}

func Debug(msg string, fields map[string]any) {
	_log(context.Background(), "debug", msg, fields)
}

func DebugCtx(ctx context.Context, msg string, fields map[string]any) {
	_log(ctx, "debug", msg, fields)
}

func _log(ctx context.Context, level string, msg string, fields map[string]any) {
	var zapFields []zap.Field
	if ctx != nil {
		reqID := ctx.Value("request_id")
		if reqID != nil {
			zapFields = append(zapFields, zap.Any("request_id", reqID))
		}
		clientIP := ctx.Value("client_ip")
		if clientIP != nil {
			zapFields = append(zapFields, zap.Any("client_ip", clientIP))
		}
		path := ctx.Value("path")
		if path != nil {
			zapFields = append(zapFields, zap.Any("path", path))
		}
		method := ctx.Value("method")
		if method != nil {
			zapFields = append(zapFields, zap.Any("method", method))
		}
		userID := ctx.Value("user_id")
		if userID != nil {
			zapFields = append(zapFields, zap.Any("user_id", userID))
		}
		sessID := ctx.Value("session_id")
		if sessID != nil {
			zapFields = append(zapFields, zap.Any("session_id", sessID))
		}
	}
	if fields != nil {
		//zapFields = append(zapFields, zap.Any("context", fields))
		for k, v := range fields {
			zapFields = append(zapFields, zap.Any(k, v))
		}
	}
	switch level {
	case "error":
		logger.Error(msg, zapFields...)
	case "fatal":
		logger.Fatal(msg, zapFields...)
	case "info":
		logger.Info(msg, zapFields...)
	case "debug":
		logger.Debug(msg, zapFields...)
	}
}
