package config

import (
	"fmt"
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
	"os"
	"strings"
)

// viper 库实例
var ymlConfig *viperlib.Viper

var BasePath string
func init()  {

	ymlConfig = viperlib.New();

	//binary, _ := os.Executable()
	//path := filepath.Dir(filepath.Dir(binary))
	//ymlConfig.SetConfigType("yml")
	//ymlConfig.SetConfigName("config")
	//// 3. 环境变量配置文件查找的路径，相对于 main.go
	//ymlConfig.AddConfigPath(path +"/config")
	//fmt.Println(path+"/config/config.yml")
	BasePath :=getAppPath()
	ymlConfig.SetConfigFile(BasePath + "/config/config.yml")
	ymlConfig.AutomaticEnv()
	err := ymlConfig.ReadInConfig() // 查找并读取配置文件
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
func Get(path string,defaultValue ...interface{}) string {
	return GetString(path,defaultValue)
}
func GetString(path string,defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path,defaultValue))
}
func GetInt(path string,defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path,defaultValue))
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	// config 或者环境变量不存在的情况
	if !ymlConfig.IsSet(path)  {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return ymlConfig.Get(path)
}
func GetEnv(envName string, defaultValue ...interface{}) string {
	if len(defaultValue) > 0 {
		return cast.ToString(internalGet(envName, defaultValue[0]))
	}
	return cast.ToString(internalGet(envName))
}

func getAppPath() string {
	// 1.初始化程序根目录
	if curPath, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(curPath, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = curPath
		}
	}
	return BasePath;
}

