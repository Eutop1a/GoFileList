package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	fileName := getFileName(url)
	fmt.Println(fileName)
	outPutHandler, err := os.Create(fileName + ".html")
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	nbytes, err := io.Copy(outPutHandler, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs   %7d %s", secs, nbytes, url)
}

func getFileName(url string) string {
	// find //
	index1, index2, n := 0, 0, len(url)
	flag := 0
	for i := 0; i < n; i++ {
		if flag == 1 && url[i] == '/' {
			index1 = i
			break
		}
		if flag == 0 && url[i] == '/' {
			flag++
		}
	}

	// find first appear .
	for i := index1; i < n; i++ {
		if url[i] == '.' {
			index2 = i
		}
	}
	// return
	return url[index1+1 : index2]
}
