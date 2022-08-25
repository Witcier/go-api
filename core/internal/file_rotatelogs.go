package internal

import (
	"os"
	"path"
	"time"
	"witcier/go-api/global"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

func (r *fileRotatelogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(global.Config.Zap.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.ForceNewFile(),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.Config.Zap.MaxAge)*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}

	return zapcore.AddSync(fileWriter), err
}
