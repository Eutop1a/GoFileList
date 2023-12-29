package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	v := make([]int, 1002)
	w := make([]int, 1002)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i], &w[i])
	}

	f := make([][]int, 1002)
	for i := 0; i < 1002; i++ {
		f[i] = make([]int, 1002)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j < v[i] {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = max(f[i-1][j], f[i-1][j-v[i]]+w[i])
			}
		}
	}
	fmt.Println(f[n][m])
}
