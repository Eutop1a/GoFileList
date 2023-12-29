package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse templates failed, err:%v\n", err)
		return
	}
	msg := "王成林"
	// 渲染模板
	t.Execute(w, msg)
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 解析之前定义一个自定义的函数safe
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")

	if err != nil {
		fmt.Printf("parse templates failed, err:%v\n", err)
		return
	}
	// 渲染模板
	str1 := "<script>alter(123);</script>"
	str2 := "<a href='https://liwenzhou.com'>李文周的博客</a>"

	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v\n", err)
		return
	}
}
