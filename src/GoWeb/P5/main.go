package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse templates failed, err:%v\n", err)
		return
	}
	// 渲染模板
	u1 := User{
		Name:   "wcl",
		Gender: "man",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"name":   "wcl123",
		"gender": "woman",
		"age":    20,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v\n", err)
		return
	}
}
