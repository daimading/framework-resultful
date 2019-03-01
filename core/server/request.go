package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"net/http"
)

func HandleRequest(c *gin.Context, request interface{}) (err error) {
	switch c.Request.Method {
	case http.MethodGet:
		m := make(map[string]interface{})
		q := c.Request.URL.Query()
		for k, v := range q {
			m[k] = v[0]
		}
		//序列化GET请求数据
		data, _ := jsoniter.Marshal(m)
		err = jsoniter.Unmarshal(data, request)
		break
	default:
		//body为空时无法转换，异常错误无法拦截
		if c.Request.Body == nil {
			return
		}
		err = c.BindJSON(request)
		break
	}
	if err != nil {
		fmt.Println("request error:")
	}
	return err
}
