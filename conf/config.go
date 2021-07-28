package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

func init() {
	viper.SetConfigName("config")
	_, fileStr, _, _ := runtime.Caller(0)
	viper.AddConfigPath(filepath.Join(fileStr, "../../"))
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
