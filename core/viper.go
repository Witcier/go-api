package core

import (
	"flag"
	"fmt"
	"os"
	"witcier/go-api/core/internal"
	"witcier/go-api/global"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				config = internal.ConfigDefaultFile
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}

	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
