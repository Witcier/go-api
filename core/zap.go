package core

import (
	"fmt"
	"os"
	"time"
	"witcier/go-api/core/internal"
	"witcier/go-api/global"
	"witcier/go-api/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Config.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.Config.Zap.Director)
		_ = os.Mkdir(global.Config.Zap.Director, os.ModePerm)
	}

	cores := make([]zapcore.Core, 0, 7)
	debugLevel := getEncoderCore(zap.DebugLevel)
	infoLevel := getEncoderCore(zap.InfoLevel)
	warnLevel := getEncoderCore(zap.WarnLevel)
	errorLevel := getEncoderCore(zap.ErrorLevel)
	dPanicLevel := getEncoderCore(zap.DPanicLevel)
	panicLevel := getEncoderCore(zap.PanicLevel)
	fatalLevel := getEncoderCore(zap.FatalLevel)

	switch global.Config.Zap.Level {
	case "debug", "DEBUG":
		cores = append(cores, debugLevel, infoLevel, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "info", "INFO":
		cores = append(cores, infoLevel, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "warn", "WARN":
		cores = append(cores, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "error", "ERROR":
		cores = append(cores, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	case "dpanic", "DPANIC":
		cores = append(cores, dPanicLevel, panicLevel, fatalLevel)
	case "panic", "PANIC":
		cores = append(cores, panicLevel, fatalLevel)
	case "fatal", "FATAL":
		cores = append(cores, fatalLevel)
	default:
		cores = append(cores, debugLevel, infoLevel, warnLevel, errorLevel, dPanicLevel, panicLevel, fatalLevel)
	}

	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	if global.Config.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.Config.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	switch {
	case global.Config.Zap.EncodeLevel == "LowercaseLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.Config.Zap.EncodeLevel == "LowercaseColorLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.Config.Zap.EncodeLevel == "CapitalLevelEncoder":
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.Config.Zap.EncodeLevel == "CapitalColorLevelEncoder":
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	return config
}

func getEncoder() zapcore.Encoder {
	if global.Config.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}

	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderCore(level zapcore.Level) (core zapcore.Core) {
	writer, err := internal.FileRotatelogs.GetWriteSyncer(level.String())

	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())

		return
	}

	return zapcore.NewCore(getEncoder(), writer, level)
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.Config.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
