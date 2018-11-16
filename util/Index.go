package util

import (
	"io/ioutil"
	"../exceptions"
	"encoding/json"
	"github.com/satori/go.uuid"
	"strings"
)

/**
	生产UUID

	@return：string - uuid信息
*/
func GeneratorUUID() string {
	code, _ := uuid.NewV4()
	return strings.ToUpper(strings.Replace(code.String(), "-","", -1))
}

/**
	读取文件

	@param：path - 文件路径
	@return：读取文件
	@return：读取异常
*/
func ReadRealFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if nil != err {
		return "", &exceptions.Error{Msg: "no such this file", Code: 500}
	}
	return string(data), nil
}

/**
	文本转换为结构体

	@param：chunk - 文本信息
	@return：结构体
	@return：转化异常
*/
func FormatToStruct(chunk *string) (map[string]interface{}, error){
	rep := make(map[string]interface{})
	err := json.Unmarshal([]byte(*chunk), &rep)
	if nil != err {
		return nil, &exceptions.Error{Msg: "json format error", Code: 500}
	}
	return rep, nil
}

/**
	结构体转化为文本

	@param: vol - 结构体
	@return 结构体文本
*/
func FormatToString(vol interface{}) (string, error){
	buf, err := json.Marshal(vol)
	if nil != err {
		return "", &exceptions.Error{Msg: "format object error", Code: 500}
	}
	return string(buf), nil
}