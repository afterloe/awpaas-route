package config

import (
	"../util"
	"../integrate/logger"
	"../exceptions"
	"os"
	"path/filepath"
	"log"
	"strings"
	"reflect"
)

var packageJson map[string]interface{}

func checkError(err error) {
	if nil != err {
		logger.Error("service", err.Error())
		os.Exit(101)
		return
	}
}

/**
 *  获取代码运行目录
 */
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/**
 * 初始化包函数
 */
func init() {
	dir := getCurrentDirectory()
	configInfo, err := util.ReadRealFile(filepath.Join(dir, "./package.json"))
	checkError(err)
	pkg, err := util.FormatToStruct(&configInfo)
	checkError(err)
	if 0 == len(pkg) {
		checkError(&exceptions.Error{Msg: "read json fail", Code: 500})
	}
	packageJson = pkg
	readEnv() // 读取env中的信息进行覆盖package.json中的信息
}

/**
	读取env中的数据进行覆盖package.json中的内容
 */
func readEnv() {
	redis_addr := os.Getenv("REDIS_ADDR")
	redis_port := os.Getenv("REDIS_PORT")
	server := packageJson["server"]
	if "" != redis_addr {
		setByTarget(GetByTarget(server, "cache"), "addr", redis_addr)
	}
	if "" != redis_port {
		setByTarget(GetByTarget(server, "cache"), "port", redis_port)
	}
}

/**
	获取配置项

 	@param key string 配置项key
	@return interface{} 配置内容
 */
func Get(key string) interface{} {
	return packageJson[key]
}

/**
	反射设置map
 */
func setByTarget(target, key, value interface{}) {
	reflect.ValueOf(target).SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
}

/**
	反射获取map配置
 */
func GetByTarget(target interface{}, key interface{}) interface{} {
	v := reflect.ValueOf(target)
	value := v.MapIndex(reflect.ValueOf(key))
	if !value.IsValid() {
		return nil
	}
	return value.Interface()
}
