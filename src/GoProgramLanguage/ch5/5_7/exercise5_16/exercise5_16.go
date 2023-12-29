package main

import (
	"fmt"
	"strings"
)

func join(des []string, sep ...string) string {
	if len(des) == 0 {
		return ""
	}
	defaultSep := ""
	if len(sep) > 0 {
		for _, subsep := range sep {
			defaultSep += subsep
		}
	}
	// 更快
	var builder strings.Builder
	if len(des) > 0 {
		builder.WriteString(des[0])
	}
	for _, substr := range des[1:] {
		builder.WriteString(defaultSep)
		builder.WriteString(substr)
	}
	return builder.String()
	//  var result string
	//	for index, substr := range des {
	//		if index != len(des)-1 {
	//			substr += defaultSep
	//		}
	//		result += substr
	//	}
	//	return result
}

func main() {
	strSlice := []string{"Hello", "World", "Golang"}
	result := join(strSlice, "-", "+", "/*-+")
	fmt.Println(result)
}
