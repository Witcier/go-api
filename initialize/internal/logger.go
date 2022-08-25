package internal

import (
	"fmt"
	"witcier/go-api/global"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	var LogZap bool
	LogZap = global.Config.Mysql.LogZap

	if LogZap {
		global.Log.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
