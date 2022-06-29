package fetch

import (
	"go-test/kit"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type request struct {
	// 请求地址
	Url string
	// 请求方法
	Method string
	// 请求体
	Data   string
	Cookie map[string]string
	Header map[string]string
}

type response struct {
	// 响应体
	Data string
	// 响应Cookie
	Cookie map[string]string
}

func Use(router *gin.RouterGroup) {
	router.POST("/fetch", func(ctx *gin.Context) {
		var req request
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			kit.HandleError(err, ctx)
			return
		}
		if req.Method == "GET" {
			newReq, _ := http.NewRequest(req.Method, req.Url, nil)
			for key, value := range req.Header {
				newReq.Header.Add(key, value)
			}
			for key, value := range req.Cookie {
				newReq.AddCookie(&http.Cookie{
					Name:  key,
					Value: value,
				})
			}
			newRes, err := http.DefaultClient.Do(newReq)
			if err != nil {
				kit.HandleError(err, ctx)
				return
			}
			body, _ := ioutil.ReadAll(newRes.Body)
			defer newRes.Body.Close()
			res := response{}
			res.Data = string(body)
			res.Cookie = map[string]string{}
			for _, cookie := range newRes.Cookies() {
				res.Cookie[cookie.Name] = cookie.Value
			}
			ctx.JSON(http.StatusOK, res)
			return
		}

		if req.Method == "POST" {
			newReq, _ := http.NewRequest(req.Method, req.Url, strings.NewReader(req.Data))
			newReq.Header.Add("Content-Type", "application/json;charset=UTF-8")
			for key, value := range req.Header {
				newReq.Header.Add(key, value)
			}
			for key, value := range req.Cookie {
				newReq.AddCookie(&http.Cookie{
					Name:  key,
					Value: value,
				})
			}
			newRes, err := http.DefaultClient.Do(newReq)
			if err != nil {
				kit.HandleError(err, ctx)
				return
			}
			body, _ := ioutil.ReadAll(newRes.Body)
			defer newRes.Body.Close()
			res := response{}
			res.Data = string(body)
			res.Cookie = map[string]string{}
			for _, cookie := range newRes.Cookies() {
				res.Cookie[cookie.Name] = cookie.Value
			}
			ctx.JSON(http.StatusOK, res)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "未指定请求类型",
		})
	})
}
