package main

import (
	"fmt"
)

func maxProfit(prices []int) (ans int) {
	minPrice := prices[0]
	for _, p := range prices {
		ans = max(ans, p-minPrice)
		minPrice = min(minPrice, p)
	}
	return
}

func main() {
	temp := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(temp)
}
