package lib

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gookit/goutil/netutil/httpctype"
)

type HttpRequest struct {
	// 请求地址
	Url string
	// 请求方法
	Method string
	// 请求体
	Data   string
	Cookie map[string]string
	Header map[string]string
}

type HttpResponse struct {
	// 响应体
	Data string
	// 响应Cookie
	Cookie map[string]string
	Header map[string]string
}

func HttpFetch(request *HttpRequest) *HttpResponse {
	var newReq *http.Request
	if request.Method == "GET" {
		newReq, _ = http.NewRequest(request.Method, request.Url, nil)
	} else {
		newReq, _ = http.NewRequest(request.Method, request.Url, strings.NewReader(request.Data))
		newReq.Header.Set(httpctype.Key, httpctype.JSON)
	}
	for key, value := range request.Header {
		newReq.Header.Set(key, value)
	}
	for key, value := range request.Cookie {
		newReq.AddCookie(&http.Cookie{
			Name:  key,
			Value: value,
		})
	}
	newRes, err := http.DefaultClient.Do(newReq)
	if err != nil {
		return &HttpResponse{
			Data:   err.Error(),
			Cookie: map[string]string{},
			Header: map[string]string{},
		}
	}
	body, _ := ioutil.ReadAll(newRes.Body)
	defer newRes.Body.Close()
	response := HttpResponse{}
	response.Data = string(body)
	response.Cookie = map[string]string{}
	for _, cookie := range newRes.Cookies() {
		response.Cookie[cookie.Name] = cookie.Value
	}
	response.Header = map[string]string{}
	for key := range newReq.Header {
		response.Header[key] = newReq.Header.Get(key)
	}
	return &response
}
