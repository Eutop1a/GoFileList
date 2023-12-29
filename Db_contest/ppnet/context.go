package ppnet

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Res  http.ResponseWriter
	Req  *http.Request
	Path string

	Params     map[string]string
	Method     string
	StatusCode int

	middlewares []HFunc
	runningIdx  int
}

func newContext(res http.ResponseWriter, req *http.Request) *Context {

	res.Header().Set("Content-Type", "application/json")
	//origin := req.Header.Get("Origin")
	res.Header().Set("Access-Control-Allow-Origin", "*")

	// 必须，设置服务器支持的所有跨域请求的方法
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	//req.Header.Set("Access-Control-Allow-Methods", "POST")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type,Access-Token")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	res.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	res.Header().Set("Access-Control-Allow-Credentials", "true")

	return &Context{
		Res:         res,
		Req:         req,
		Path:        req.URL.Path,
		Params:      make(map[string]string),
		Method:      req.Method,
		StatusCode:  http.StatusOK,
		middlewares: make([]HFunc, 0, 16),
		runningIdx:  -1,
	}

}

func (c *Context) Next() {
	if c.runningIdx == len(c.middlewares)-1 {
		return
	}
	c.runningIdx++
	c.middlewares[c.runningIdx](c)

	c.Next()

}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Get(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetHeader(key, value string) {
	c.Res.Header().Set(key, value)
}

func (c *Context) SetStatusCode(code int) {
	c.StatusCode = code
}

func (c *Context) JSON(code int, data H) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatusCode(code)
	encode := json.NewEncoder(c.Res)
	if err := encode.Encode(data); err != nil {
		http.Error(c.Res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatusCode(code)
	c.Res.Write([]byte(html))
}

func (c *Context) String(code int, format string, vales ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.SetStatusCode(code)
	c.Res.Write([]byte(fmt.Sprintf(format, vales...)))
}

func (c *Context) Data(code int, data []byte) {
	c.SetStatusCode(code)
	c.Res.Write(data)
}
