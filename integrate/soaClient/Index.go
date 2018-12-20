package soaClient

import (
	"../logger"
	"net/http"
	"net"
	"time"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
)

var (
	maxIdleConn,
	maxIdleConnPerHost,
	idleConnTimeout int
)

func init() {
	maxIdleConn = 100
	maxIdleConnPerHost = 100
	idleConnTimeout = 90
}

func JsonToObject(chunk string) (map[string]interface{}, error){
	rep := make(map[string]interface{})
	err := json.Unmarshal([]byte(chunk), &rep)
	if nil != err {
		return nil, err
	}
	return rep, nil
}

func Call(method, serviceName, url string, body io.ReadCloser, header map[string]string) (map[string]interface{}, error) {
	remote, err := http.NewRequest(method, fmt.Sprintf("http://%s%s", serviceName, url), body)
	for key, value := range header {
		remote.Header.Add(key, value)
	}
	if nil != err {
	}
	response, err := GeneratorClient().Do(remote)
	if err != nil && response == nil {
		logger.Error("soa-client", fmt.Sprintf("forward %+v", err))
		return nil, err
	} else {
		defer response.Body.Close()
		logger.Logger("soa-client", fmt.Sprintf("%3d | %-7s | %s", response.StatusCode, method, url))
		reply, err := ioutil.ReadAll(response.Body)
		if nil != err {
			return nil, err
		}
		return JsonToObject(string(reply))
	}
}

func GeneratorClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{ Timeout: 30 * time.Second,}).DialContext,
			MaxIdleConns:        maxIdleConn,
			MaxIdleConnsPerHost: maxIdleConnPerHost,
			IdleConnTimeout:	 time.Duration(idleConnTimeout)* time.Second,
		},
		Timeout: 30 * time.Second,
	}
	return client
}