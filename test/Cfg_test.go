package test

import (
	"../config"
	"testing"
	"fmt"
	"reflect"
)

func Test_getWhiteList(t *testing.T) {
	whiteList := config.GetByTarget(config.Get("custom"),"whiteList")

	t.Log(reflect.TypeOf(whiteList))
}

func Test_NumberFormat(t *testing.T) {
	serverCfg := config.Get("server").(map[string]interface{})
	gatewayCfg := config.GetByTarget(serverCfg, "gateway")
	multiCoreCfg := config.GetByTarget(gatewayCfg, "multiCore")
	if nil == multiCoreCfg {
		t.Error("type err! can't find config")
		return
	}
	num := config.GetByTarget(multiCoreCfg, "num")
	var cpuNum int
	if nil != num {
		switch reflect.TypeOf(num).String()  {
		case "float64":
			cpuNum = int(reflect.ValueOf(num).Float())
			break
		case "int":
			cpuNum = int(reflect.ValueOf(num).Int())
			break
		default:
			cpuNum = -1
		}
	} else {
		cpuNum = -1
	}

	t.Log(fmt.Sprintf("cpuNum is %d", cpuNum))
}
