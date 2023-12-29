package main

import (
	"fmt"
	"math/rand"
	"time"
)

//type Address struct {
//	Province   string
//	City       string
//	UpdateTime string
//}
//
//type Email struct {
//	Addr       string
//	UpdateTime string
//}
//
//type Person struct {
//	Name   string
//	Gender string
//	Age    int
//	Address
//	Email
//}
//
//func main() {
//	p1 := Person{
//		Name:   "wcl",
//		Gender: "man",
//		Age:    18,
//		Address: Address{
//			Province:   "gansu",
//			City:       "Jiuquan",
//			UpdateTime: "2020-2-2",
//		},
//		Email: Email{
//			Addr:       "w10909i80@gmail.com",
//			UpdateTime: "2018-1-8",
//		},
//	}
//
//	fmt.Printf("%#v\n", p1)
//
//	fmt.Println(p1.Address.UpdateTime)
//	fmt.Println(p1.Email.UpdateTime)
//}
//
//
//
//type Animal struct {
//	name string
//}
//
//func (a *Animal) move() {
//	fmt.Printf("%s动了\n", a.name)
//}
//
//type Dog struct {
//	Feet    uint8
//	*Animal // 匿名嵌套
//}
//
//func (d *Dog) speak() {
//	fmt.Printf("%s狗叫\n", d.name)
//}
//
//func main() {
//	d1 := &Dog{
//		Feet: 4,
//		Animal: &Animal{
//			name: "dog",
//		},
//	}
//	d1.move()
//	d1.speak()
//}
//
//type student struct {
//	ID   int
//	Name string
//}
//
//func newStudent(id int, name string) student {
//	return student{
//		ID:   id,
//		Name: name,
//	}
//}
//
//type class struct {
//	Title    string    `json:"title"`
//	Students []student `json:"student_list" db:"student" xml:"ss"`
//}
//
//func main() {
//	c1 := class{
//		Title:    "火箭101",
//		Students: make([]student, 0, 10),
//	}
//	for i := 0; i < 3; i++ {
//		tmp := newStudent(i, fmt.Sprintf("Stu%02d", i))
//		c1.Students = append(c1.Students, tmp)
//	}
//	// fmt.Printf("%#v\n", c1)
//
//	// JSON序列化: go语言中的数据 -> JSON格式的字符串
//	data, err := json.Marshal(c1)
//	if err != nil {
//		fmt.Println("Json marshal failed, err", err)
//		return
//	}
//	fmt.Printf("%T\n", data)
//	fmt.Printf("%s\n", data)
//
//	// JSON反序列化: JSON格式的字符串 -> go语言中的数据
//	str := `{"Title":"火箭101","Students":[{"ID":0,"Name":"Stu00"},{"ID":1,"Name":"Stu01"},{"ID":2,"Name":"Stu02"}]}
//`
//	var c2 class
//	err = json.Unmarshal([]byte(str), &c2)
//	if err != nil {
//		fmt.Println("Json Unmarshal failed, err", err)
//		return
//	}
//	fmt.Printf("%#v\n", c2)
//}

const LEN = 20

func GenerateArray() []int {
	rand.Seed(time.Now().UnixNano())
	array := make([]int, LEN)
	for i := 0; i < LEN; i++ {
		array[i] = rand.Intn(100)
	}
	return array
}

func SelectSort(arr []int) {
	for i := 0; i < LEN-1; i++ {
		min := i
		for j := i + 1; j < LEN; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}

func BubbleSort(arr []int) {
	for i := 0; i < LEN; i++ {
		for j := 0; j < LEN-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func SequentialSearch(arr []int, num int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}

func BruteForceStringMatch(T, P string) int {
	for i := 0; i < len(T)-len(P); i++ {
		j := 0
		for ; j < len(P) && P[j] == T[i+j]; j++ {
		}
		if i == len(P) {
			return i
		}
	}
	return -1
}

func InsertionSort(arr []int) {
	for i := 1; i < LEN-1; i++ {
		key := arr[i]
		j := i - 1
		// 遍历前面的
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}

func BinarySearch(arr []int, key int) int {
	left := 0
	right := len(arr)
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > key {
			right = mid - 1
		} else if arr[mid] == key {
			return mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)
	quickSortHelper(arr, low, pivotIndex-1)
	quickSortHelper(arr, pivotIndex+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func MergeSort(arr []int) []int { // 递归把整个数组分段
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	result := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...) // 用于将一个切片（slice）的所有元素展开为独立的参数
	result = append(result, right...)
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 迭代方式实现归并排序
	for step := 1; step < len(arr); step *= 2 {
		for i := 0; i < len(arr)-step; i += step * 2 {
			left := i
			mid := i + step
			right := i + step*2
			// 防止越界
			if mid > len(arr) {
				mid = len(arr)
			}
			if right > len(arr) {
				right = len(arr)
			}
			merge(arr, left, mid, right)
		}
	}
	return arr
}

func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left)
	i := left
	j := mid
	k := 0
	for i < mid && j < right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i < mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j < right {
		temp[k] = arr[j]
		j++
		k++
	}
	copy(arr[left:right], temp)
}

type Graph struct {
	vertices int           // 顶点数
	adjList  map[int][]int // 邻接表
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	// 通过顶点u索引的邻接列表。我们将顶点v添加到g.adjList[u]的末尾，使用append函数实现。
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func DFS(g *Graph, start int, visited map[int]bool) {
	visited[start] = true
	fmt.Printf("%d ", start)

	for _, v := range g.adjList[start] {
		if !visited[v] {
			DFS(g, v, visited)
		}
	}
}

func BFS(g *Graph, start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		fmt.Printf("%d ", node)
		queue = queue[1:]

		for _, v := range g.adjList[node] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
}

func main() {
	//arr1 := GenerateArray()
	//begin := 0
	//end := len(arr1) - 1
	//
	//QuickSort(arr1, begin, end)
	//
	//fmt.Println(arr1) // 输出排序后的结果
	//arr := []int{9, 3, 7, 5, 1, 8, 2, 6, 4}
	//sortedArr := MergeSort(arr)
	//fmt.Println(sortedArr)
	//
	//arr1 := []int{9, 3, 7, 5, 1, 8, 2, 6, 4}
	//sortedArr1 := mergeSort(arr1)
	//fmt.Println(sortedArr1)
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)
	g.AddEdge(4, 0)

	fmt.Println("Depth First Traversal (DFS):")
	visited := make(map[int]bool)
	DFS(g, 0, visited)

	fmt.Println("\n\nBreadth First Traversal (BFS):")
	BFS(g, 0)
}
