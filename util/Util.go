package util

import (
	"../exceptions"
)

func CheckNeed(args... interface{}) error {
	for _, arg := range args {
		if nil == arg || "" == arg {
			return &exceptions.Error{Msg: "Lack parameter.", Code: 400}
		}
	}
	return nil
}
