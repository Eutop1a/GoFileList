package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数kua
	// 要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	k := func(arg string) (string, error) {
		return arg + "帅气逼人", nil
	}
	// 定义模板
	// 合起来写如下
	// t, err := templates.New("f.tmpl").ParseFiles("./f.tmpl")
	t := template.New("f.tmpl") // 创建一个名字是f的模板对象，名字一定要与模板的名字能对应上
	// 告诉模板引擎，我现在多了一个自定义函数，名字是kua(必须在解析前写)
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	// 解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse templates failed, err:%v\n", err)
		return
	}
	name := "wcl"
	// 渲染模板
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 把被包含的写在前面
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	//t, err := templates.ParseFiles("./ul.tmpl", "./t.tmpl")	写反会导致只能解析出来第一个
	if err != nil {
		fmt.Printf("parse templates failed, err:%v\n", err)
		return
	}
	name := "王成林"
	// 渲染模板
	t.Execute(w, name)

}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v\n", err)
		return
	}
}
