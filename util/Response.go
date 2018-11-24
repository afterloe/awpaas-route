package util

import (
	"reflect"
	"fmt"
	"net/http"
	"../exceptions"
)

type responseDTO map[string]interface{}

func (res *responseDTO) String() string {
	val := reflect.ValueOf(res)
	return fmt.Sprintf("Response {data: %s, code: %d, msg: %s}", val.MapIndex(reflect.ValueOf("data")).Interface(),
		val.MapIndex(reflect.ValueOf("code")).Interface(), val.MapIndex(reflect.ValueOf("msg")).Interface())
}

func Success(data interface{}) *responseDTO {
	return Build(data, http.StatusOK, nil)
}

func Error(point interface{}) *responseDTO {
	p := point.(*exceptions.Error)
	return Build(nil, p.Code, p.Msg)
}

func Fail(code int, msg string) *responseDTO {
	return Build(nil, code, msg)
}

func Build(data interface{}, code int, msg interface{}) *responseDTO {
	return &responseDTO {
		"data": data,
		"code": code,
		"msg": msg,
	}
}
