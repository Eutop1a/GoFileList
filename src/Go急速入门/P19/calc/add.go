package calc

import (
	"Package/snow"
	"fmt"
)

var Name = "nihao"

func Add(x, y int) int {
	snow.Snow()
	return x + y
}

func init() {
	fmt.Println("calc init()")
}
