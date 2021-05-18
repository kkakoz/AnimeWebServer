package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Options struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
	Stdout     bool
}

func NewLog(viper *viper.Viper) (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
		o = &Options{}
	)
	err = viper.UnmarshalKey("log", o)
	if err != nil {
		return logger, err
	}

	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize, // megabytes
		MaxBackups: o.MaxBackups,
		MaxAge:     o.MaxAge, // days
	})

	// file core 采用jsonEncoder
	cores := make([]zapcore.Core, 0, 2)
	je := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	cores = append(cores, zapcore.NewCore(je, fw, level))

	if o.Stdout {
		cw := zapcore.Lock(os.Stdout)
		ce := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(ce, cw, level))
	}

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)

	zap.ReplaceGlobals(logger)


	return logger, nil
}

var ProviderSet = wire.NewSet(NewLog)