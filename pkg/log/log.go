package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

var (
	level = map[string]zapcore.Level{
		"info":   zapcore.InfoLevel,
		"debug":  zapcore.DebugLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dPanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
)

var log *zap.SugaredLogger

func NewLog(o *Options) error {
	z := zap.NewProductionConfig()
	if v, ok := level[o.Level]; ok {
		z.Level = zap.NewAtomicLevelAt(v)
	} else {
		z.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	z.Development = o.Development
	z.DisableCaller = o.DisableCaller
	z.DisableStacktrace = o.DisableStacktrace
	z.Sampling = o.Sampling
	z.Encoding = o.Encoding
	z.EncoderConfig = o.EncoderConfig
	z.OutputPaths = o.OutputPaths
	z.ErrorOutputPaths = o.ErrorOutputPaths
	z.InitialFields = o.InitialFields

	logger, err := z.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}

	log = logger.Sugar()

	if o.HttpServer.Open {
		go func() {
			http.HandleFunc("/handle/level", z.Level.ServeHTTP)
			if err := http.ListenAndServe(o.HttpServer.Addr, nil); err != nil {
				Info(fmt.Sprintf("log http server not Start to %v", o.HttpServer.Addr))
			}
		}()
	}

	return nil
}

func Info(msg string, keysAndValues ...interface{}) {
	if log == nil {
		return
	}
	log.Infow(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	if log == nil {
		return
	}
	log.Debugw(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	if log == nil {
		return
	}
	log.Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	if log == nil {
		return
	}
	log.Errorw(msg, keysAndValues...)
}

func DPanic(msg string, keysAndValues ...interface{}) {
	if log == nil {
		return
	}
	log.DPanicw(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	if log == nil {
		return
	}
	log.Panicw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	if log == nil {
		return
	}
	log.Fatalw(msg, keysAndValues...)
}
