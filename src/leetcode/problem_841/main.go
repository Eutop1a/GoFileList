package main

import "fmt"

var (
	num   int
	visit []bool
)

// DFS
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num = 0
	visit = make([]bool, n)
	dfs(rooms, 0)
	return num == n
}

func dfs(rooms [][]int, x int) {
	//fmt.Println("x = ", x)
	visit[x] = true
	num++
	for _, it := range rooms[x] {
		if !visit[it] {
			dfs(rooms, it)
		}
	}
}

// BFS
func canVisitAllRooms2(rooms [][]int) bool {
	n := len(rooms)
	num := 0
	vis := make([]bool, n)
	var q []int
	vis[0] = true
	q = append(q, 0)
	for i := 0; i < len(q); i++ {
		x := q[i]
		num++
		for _, it := range rooms[x] {
			if !vis[it] {
				vis[it] = true
				q = append(q, it)
			}
		}
	}
	return num == n
}

func main() {
	var rooms [][]int
	rooms = make([][]int, 5)
	for i := 0; i < 5; i++ {
		rooms[i] = make([]int, 5)
	}
	rooms[0] = []int{1, 3}
	rooms[1] = []int{3, 0, 1, 2}
	rooms[2] = []int{2}
	rooms[3] = []int{0}
	temp := canVisitAllRooms(rooms)
	temp = canVisitAllRooms2(rooms)
	fmt.Println(temp)

}
