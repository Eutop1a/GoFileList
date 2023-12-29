package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

//// 解析html
//func countWordsAndImages(url string) (int, int, error) {
//	// 获取html
//	resp, err := http.Get(url)
//	if err != nil {
//		return 0, 0, err
//	}
//	// 解析html
//	doc, err := html.Parse(resp.Body)
//	if err != nil {
//		return 0, 0, fmt.Errorf("html Parse:%s", err)
//	}
//	words, images := countWordsAndImages(doc)
//	return words, images, nil
//}

// 获取单词以及图片的个数
func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	// 文本节点
	if n.Type == html.TextNode {
		// 单词以空格分隔
		sArr := strings.Split(n.Data, " ")
		words += len(sArr)
	} else if n.Data == "img" {
		images = 1
	}
	// 不是style和script才递归它们的子节点
	var wc, ic int
	if n.Data != "style" && n.Data != "script" {
		// 进入子节点
		wc, ic = countWordsAndImages(n.FirstChild)
	}
	// 进入兄弟节点
	wb, ib := countWordsAndImages(n.NextSibling)
	// 合并总数
	words = words + wc + wb
	images = images + ib + ic
	return
}
func main() {

}
