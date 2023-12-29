package ppnet

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type HFunc func(ctx *Context)

type router struct {
	handlers map[string]*routerPart
}

func newRouter() *router {

	r := router{make(map[string]*routerPart)}

	for _, v := range []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"} {
		r.handlers[v] = newFatherRouter("", false)
	}

	return &r
}

func (r *router) addHandlerFunc(method string, pattern []string, handlerFunc HFunc) error {
	return r.handlers[method].addRouterInTrie(pattern, 1, handlerFunc)
}

// 实现 http.Handler 接口
func (e *engine) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	{
		if req.Method == "OPTIONS" {

			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			//res.Header().Set("Access-Control-Allow-Credentials", "true")
			res.Header().Set("Access-Control-Max-Age", strconv.Itoa(3600))
			res.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Credentials, Content-Type, Authorization")

			res.WriteHeader(200)
			return
		}
	} //flight

	if routerPart, exist := e.router.handlers[req.Method]; exist {
		ctext := newContext(res, req)
		if hFunc, err := routerPart.parseRouterInTrie(ctext, strings.Split(req.URL.Path, "/"), 1); err == nil {
			for _, groups := range e.routers {
				if strings.HasPrefix(req.URL.Path, groups.prefix) {
					ctext.middlewares = append(ctext.middlewares, groups.middlewares...)
				}
			}
			ctext.middlewares = append(ctext.middlewares, hFunc)
			ctext.Next()
			return
		}
	}
	fmt.Fprintf(res, "404 not found！ %v", req.URL)

}
