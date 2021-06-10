package log

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
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

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func L() *zap.Logger {
	return logger
}

func NewLog(viper *viper.Viper) (*zap.Logger, error) {
	var (
		err    error
		o = &Options{}
	)
	viper.SetDefault("log.filename", "temp/temp.log")
	viper.SetDefault("log.maxSize", 10)
	viper.SetDefault("log.maxBackups", 5)
	viper.SetDefault("log.maxAge", 30)
	viper.SetDefault("log.stdout", true)
	err = viper.UnmarshalKey("log", o)
	if err != nil {
		return logger, errors.Wrap(err, "viper unmarshal失败")
	}

	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize, // 日志文件最大大小(MB)
		MaxBackups: o.MaxBackups,// 保留旧文件最大数量
		MaxAge:     o.MaxAge, // 保留旧文件最长天数
	})

	encoder := getEncoder()

	var core zapcore.Core
	if o.Stdout {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, fw, zapcore.DebugLevel),
			zapcore.NewCore(consoleEncoder, os.Stdout, zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, fw, zapcore.InfoLevel)
	}
	logger = zap.New(core)
	sugarLogger = logger.Sugar()

	zap.ReplaceGlobals(logger)
	return logger, nil
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(config)
}

var LogSet = wire.NewSet(NewLog)