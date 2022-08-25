package utils

import (
	"errors"
	"os"
	"witcier/go-api/global"

	"go.uber.org/zap"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}

		return false, errors.New("存在同名文件")
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exists, err := PathExists(v)
		if err != nil {
			return err
		}

		if !exists {
			global.Log.Debug("create direcory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.Log.Error("create direcory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}

	return err
}
