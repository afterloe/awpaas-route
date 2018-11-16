package exceptions

import "fmt"

// 自定义 Error
type Error struct {
	Msg string
	Code int
}

func (e *Error) getMsg() string {
	return e.Msg
}

func (e *Error) getCode() int {
	return e.Code
}

// 自定义异常要实现这个方法
func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Msg)
}