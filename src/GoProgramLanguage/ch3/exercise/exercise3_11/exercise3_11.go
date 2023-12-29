package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	if s[0] == '+' || s[0] == '-' { //去除符号
		buf.WriteByte(s[0])
		s = s[1:]
	}
	index := strings.LastIndex(s, ".")
	var tmp string //小数部分
	tmp = s[index+1:]
	s = s[:index]

	n := len(s)
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(string(s[i]))
	}

	buf.WriteByte('.')

	for i := 0; i < len(tmp); i++ {
		buf.WriteByte(tmp[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("-12312345566.166623"))
}
