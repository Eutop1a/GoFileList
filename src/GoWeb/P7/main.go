package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index2.tmpl")
	if err != nil {
		fmt.Printf("parse templates failed, err:%v\n", err)
		return
	}
	msg := "王成林"
	// 渲染模板
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./home2.tmpl")
	if err != nil {
		fmt.Printf("parse templates failed, err:%v\n", err)
		return
	}
	msg := "王成林"
	// 渲染模板
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 根模板写前面
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("parse templates failed, err:%v\n", err)
		return
	}
	msg := "王成林"
	// 渲染模板
	t.ExecuteTemplate(w, "index2.tmpl", msg)
}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 根模板写前面
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("parse templates failed, err:%v\n", err)
		return
	}
	msg := "王成林"
	// 渲染模板
	t.ExecuteTemplate(w, "home2.tmpl", msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v\n", err)
		return
	}
}
