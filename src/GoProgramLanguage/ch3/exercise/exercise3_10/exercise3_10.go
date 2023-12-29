package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer

	n := len(s)
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(string(s[i]))
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("12123456748979"))
}
