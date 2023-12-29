package main

import (
	"fmt"
	"os"
)

var tempDir []string

func main() {
	for i := 0; i < 10; i++ {
		var str = fmt.Sprintf("./node%d", i)
		tempDir = append(tempDir, str)
	}

	fmt.Println(tempDir)
	var rmdirs []func()
	for _, dir := range tempDir {
		// crucial:
		// if this statement is being omitted, this program will delete only the last directory
		//dir := dir
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			fmt.Println(dir)
			os.RemoveAll(dir)
		})
	}
	//fmt.Println("Wait for a moment")
	//time.Sleep(time.Second * 10)
	for _, rmdir := range rmdirs {
		rmdir()
	}
}
