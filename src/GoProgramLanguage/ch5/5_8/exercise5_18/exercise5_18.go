package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	defer func() {
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()
	// 这里要搞明白 return defer 和 返回值被赋值给调用者 的先后执行顺序
	// 这里先说结论 defer的执行顺序在return之后，但是在返回值返回给调用方之前

	// 匿名返回值是在return执行时被声明，有名返回值则是在函数声明的同时被声明，
	// 因此在defer语句中只能访问有名返回值，而不能直接访问匿名返回值

	// return其实应该包含前后两个步骤：
	// 第一步是给返回值赋值(若为有名返回值则直接赋值，若为匿名返回值则先声明再赋值)；
	// 第二步是调用RET返回指令并传入返回值，而RET则会检查defer是否存在，若存在就先逆序插播defer语句，
	// 最后RET携带返回值退出函数；

	// 回到这里这个例子, err是被显示定义的有名返回值变量, 所以defer可以访问err
	// 这个例子中, 执行到defer时如果Close函数发生错误, 那么返回值中的err将会被改变
	return local, n, err
}

func main() {
	url := os.Args[1]
	name, Len, err := fetch(url)
	if err != nil {
		return
	}
	fmt.Println(name, Len)
}

//
//func test() {
//	fmt.Println("134")
//	defer func() {
//		fmt.Println("10")
//	}()
//}
//
//func main() {
//	println("1")
//	test()
//	func() {
//		defer println("2")
//		println("3")
//	}()
//	println("4")
//}
