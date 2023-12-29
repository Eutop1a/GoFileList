package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers0(l1 *ListNode, l2 *ListNode) *ListNode {
	num1, num2, dig := 0, 0, 0
	digit := 0
	for l1 != nil {
		num1 += l1.Val * int(math.Pow(10, float64(dig)))
		dig++
		l1 = l1.Next
	}
	dig = 0
	for l2 != nil {
		num2 += l2.Val * int(math.Pow(10, float64(dig)))
		dig++
		l2 = l2.Next
	}
	res := num1 + num2
	fmt.Println(res)
	var head *ListNode
	var tail *ListNode
	if res == 0 {
		head = &ListNode{Val: res}
		return head
	}
	for res > 0 {
		res, digit = res/10, res%10
		newNode := &ListNode{Val: digit}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum, carry := 0, 0
	var head *ListNode
	var tail *ListNode
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum = n1 + n2 + carry
		sum, carry = sum%10, sum/10
		newNode := &ListNode{Val: sum}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

// 辅助函数：将链表转换为整数，用于测试
func convertToList(val int) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead

	for val > 0 {
		digit := val % 10
		current.Next = &ListNode{Val: digit}
		current = current.Next
		val /= 10
	}

	return dummyHead.Next
}

// 辅助函数：打印链表，用于测试
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func subtractProductAndSum(n int) int {
	mul, sum := 1, 0
	for n > 0 {
		temp := n % 10
		mul *= temp
		fmt.Printf("%d *= %d\n", mul, temp)
		sum += temp
		fmt.Printf("%d += %d\n", sum, temp)
		n /= 10
	}
	return mul - sum
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	for i < len(nums1) {
		res = append(res, nums1[i])
		i++
	}
	for j < len(nums2) {
		res = append(res, nums2[j])
		j++
	}
	return median(res)
}
func median(s []int) float64 {
	if len(s)%2 == 0 {
		return float64(s[len(s)/2]+s[(len(s)/2)-1]) / 2.0
	} else {
		return float64(s[(len(s) / 2)])
	}
}

//
//func findMedianSortedArrays0(nums1 []int, nums2 []int) float64 {
//	totalLength := len(nums1) + len(nums2)
//	if totalLength%2 == 1 { // 总长度为奇数
//		midIndex := totalLength / 2
//		return float64(getKthElement(nums1, nums2, midIndex+1))
//	} else { // 总长度为偶数
//		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
//		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
//	}
//	return 0
//}
//
//func getKthElement(nums1, nums2 []int, k int) int {
//	index1, index2 := 0, 0
//	for {
//		if index1 == len(nums1) {
//			return nums2[index2+k-1]
//		}
//		if index2 == len(nums2) {
//			return nums1[index1+k-1]
//		}
//		if k == 1 {
//			return min(nums1[index1], nums2[index2])
//		}
//		half := k / 2
//		newIndex1 := min(index1+half, len(nums1)) - 1
//		newIndex2 := min(index2+half, len(nums2)) - 1
//		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
//		if pivot1 <= pivot2 {
//			k -= newIndex1 - index1 + 1
//			index1 = newIndex1 + 1
//		} else {
//			k -= newIndex2 - index2 + 1
//			index2 = newIndex2 + 1
//		}
//	}
//	return 0
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}

func findMedianSortedArrays0(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	totalLen := m + n
	if totalLen%2 == 0 { // 偶数
		midIndex1, midIndex2 := totalLen/2-1, totalLen/2
		return float64(getKthElement(nums1, nums2, midIndex1+1) + +getKthElement(nums1, nums2, midIndex2+1)) / 2.0

	} else { // 奇数
		midIndex := totalLen / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1

		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	size := m + n
	id0, id1 := (size+1)/2, (size+2)/2
	val0 := findKth(nums1, 0, nums2, 0, id0)
	val1 := findKth(nums1, 0, nums2, 0, id1)
	fmt.Println(val0, val1)
	return float64(val0+val1) / 2.0
}

// 从a[sta, m-1] 和 b[stb, n-1]里面找kth个元素
func findKth(a []int, sta int, b []int, stb, kth int) int {
	if sta >= len(a) {
		return b[stb+kth-1]
	}
	if stb >= len(b) {
		return a[sta+kth-1]
	}
	if kth == 1 {
		return min(a[sta], b[stb])
	}

	valA := 0
	h := kth / 2
	if len(a)-sta >= h {
		valA = a[sta+h-1]
	} else {
		valA = a[len(a)-1]
	}
	countA := 0
	if len(a)-stb >= h {
		countA = h
	} else {
		countA = len(a) - sta
	}

	valB := 0
	if len(b)-stb >= h {
		valB = b[stb+h-1]
	} else {
		valB = b[len(b)-1]
	}
	countB := 0
	if len(b)-stb >= h {
		countB = h
	} else {
		countB = len(b) - stb
	}

	if valA <= valB {
		return findKth(a, sta+countA, b, stb, kth-countA)
	}
	return findKth(a, sta, b, stb+countB, kth-countB)
}

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// dp[i][j] 表示子串 s[i:j+1] 是否为回文子串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true // 单个字符是回文
	}

	start, maxLen := 0, 1 // 初始化最长回文子串的起始位置和长度

	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i <= 2 || dp[i+1][j-1] {
					dp[i][j] = true
					if j-i+1 > maxLen {
						start = i
						maxLen = j - i + 1
					}
				}
			}
		}
	}

	return s[start : start+maxLen]
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func longestPalindrome1(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	maxLen := 0
	result := ""

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isPalindrome(s[i:j+1]) && j-i+1 > maxLen {
				maxLen = j - i + 1
				result = s[i : j+1]
			}
		}
	}

	return result
}

func longestPalindrome2(s string) string {
	n := len(s)

	if n <= 1 {
		return s
	}

	dp := make([][]bool, n)
	// 记录初始状态, st是初始下标
	maxLen, st := 0, -1

	maxLen = 1
	st = 0
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for Len := 2; Len <= n; Len++ { // Len表示子串长度
		for i := 0; i+Len <= n; i++ { // i 是左边界
			j := i + Len - 1 //j 是又边界
			if s[i] != s[j] {
				continue
			}
			if Len > 2 && !dp[i+1][j-1] {
				continue
			}
			dp[i][j] = true
			maxLen = Len
			st = i
		}
	}
	return s[st : st+maxLen]
}

//func convert(s string, numRows int) string {
//	if numRows == 1 {
//		return s
//	}
//	rows := make([]string, numRows)
//	n := 2*numRows - 2
//	for i, char := range s {
//		x := i % n
//		rows[min(x, n-x)] += string(char)
//	}
//	return strings.Join(rows, "")
//}

func convert(s string, numRows int) (ans string) {
	n := len(s)
	if numRows >= n || numRows == 1 {
		ans = s
		return
	}
	rows := make([]string, numRows)
	row, down := 0, false
	for _, v := range s {
		rows[row] += string(v)
		if row == 0 || row == numRows-1 {
			down = !down
		}
		if down {
			row++
		} else {
			row--
		}
	}

	for i := 0; i < numRows; i++ {
		ans += rows[i]
	}
	return
}

func convert0(s string, numRows int) (ans string) {
	n := len(s)
	if numRows == 1 || n < numRows {
		ans = s
		return
	}

	row, dir := 0, false
	rows := make([]string, numRows)

	for _, v := range s {
		rows[row] += string(v)
		if row == 0 || row == numRows-1 {
			dir = !dir
		}
		if dir {
			row++
		} else {
			row--
		}
	}

	for _, v := range rows {
		ans += v
	}
	return
}

func rev(str string) int {
	bytes := []byte(str)
	n := len(bytes)
	for i := 0; i < n/2; i++ {
		tmp := bytes[n-i-1]
		bytes[n-i-1] = bytes[i]
		bytes[i] = tmp
	}
	str = string(bytes)
	num, _ := strconv.Atoi(str)
	return num
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	if x >= 0 {
		num := rev(str)
		if num > math.MaxInt32 {
			return 0
		}
		return num
	} else {
		// 去掉符号
		str = str[1:]
		num := rev(str)
		if num > math.MaxInt32 {
			return 0
		}
		return -num
	}
}

func reverse0(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}

func myAtoi(s string) (ans int) {
	ans = 0
	n := len(s)
	i, sign := 0, 1
	// 去除前导空格
	for ; i < n && s[i] == ' '; i++ {
	}
	if i == n {
		return 0
	}
	// 读入字符
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}
	// 读入数字
	for j := i; j < n && s[j] >= '0' && s[j] <= '9'; j++ {
		ans = ans*10 + sign*int(s[j]-'0')
		// 溢出判断
		if ans < 0 && ans <= math.MinInt32 {
			ans = math.MinInt32
		}

		if ans > 0 && ans >= math.MaxInt32 {
			ans = math.MaxInt32
		}
	}
	return
}

func isPalindrome0(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}
	return x == rev || x == rev/10
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	count, n := 0, len(str)
	if n%2 != 0 {
		j := n/2 + 1
		for i := n/2 - 1; i >= 0; i-- {
			if j >= n {
				break
			}
			if str[i] == str[j] {
				count++
			}
			j++
		}
		if count*2 == n-1 {
			return true
		} else {
			return false
		}

	}

	j := n / 2
	for i := n/2 - 1; i >= 0; i-- {
		if j >= n {
			break
		}
		if str[i] == str[j] {
			count++
		}
		j++
	}
	if count*2 == n {
		return true
	}
	return false
}

func ext(s string) string {
	res := make([]rune, len(s)+1)
	res[0] = '0'
	copy(res[1:], []rune(s))
	return string(res)
}

func isMatch0(s string, p string) bool {
	n, m := len(s), len(p)

	s = ext(s)
	p = ext(p)

	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	// 初始化
	f[0][0] = true
	for j := 0; j <= m; j++ {
		if p[j] == '*' {
			f[0][j] = f[0][j-1]
		} else if j+1 > m || p[j+1] != '*' {
			break
		} else {
			f[0][j] = true
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j] == '*' {
				f[i][j] = f[i][j-1]
				continue
			}
			if j+1 <= m && p[j+1] == '*' {
				f[i][j] = f[i][j-1] ||
					(f[i-1][j-1] && (p[j] == '.' || s[i] == p[j])) ||
					(f[i-1][j] && (p[j] == '.' || s[i] == p[j]))
			} else {
				f[i][j] = f[i-1][j-1] && (p[j] == '.' || s[i] == p[j])
			}
		}
	}
	return f[n][m]
}

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	// s = string(' ') + s
	// p = string(' ') + p

	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true

	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[n][m]
}

//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func maxArea(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

func mergeAlternately(word1, word2 string) string {
	n, m := len(word1), len(word2)
	ans := make([]byte, n+m)
	l := 0
	for i := 0; i < n || i < m; i++ {
		if i < n {
			ans[l] = word1[i]
			l++
		}
		if i < m {
			ans[l] = word2[i]
			l++
		}
	}
	return string(ans)
}

func gcdOfStrings(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	if n < m {
		return gcdOfStrings(str2, str1)
	}
	if str2 == "" {
		return str1
	}
	if str1[:m] != str2 {
		return ""
	}
	return gcdOfStrings(str1[m:], str2)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies, n := 0, len(candies)
	res := make([]bool, n)
	for _, v := range candies {
		if v >= maxCandies {
			maxCandies = v
		}
	}
	for k, v := range candies {
		res[k] = v+extraCandies >= maxCandies
	}
	return res
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if length == 0 {
		return false
	}
	count := 0
	// 遍历
	for i := 0; i < length; i++ {
		if flowerbed[i] == 0 {
			// 第i个位置没有种花
			if i == 0 && length == 1 {
				flowerbed[i] = 1
				count++
			} else if i == 0 && length > 1 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				count++
			} else if i == length-1 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				count++
			} else if i > 0 && i < length && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				count++
			}
		}
		if count >= n {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	n := len(s)
	i, j := 0, n-1

	res := []byte(s)
	for i < j {
		for i < n && !strings.Contains("aeiouAEIOU", string(s[i])) {
			i++
		}
		for j > 0 && !strings.Contains("aeiouAEIOU", string(s[i])) {
			j--
		}
		if i < j {
			res[i], res[j] = res[j], res[i]
			i++
			j--
		}
	}
	return string(res)
}

func reverseWords(s string) string {
	t := strings.Fields(s)
	n := len(t)
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
	res := strings.Join(t, " ")
	return res
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	// 先求得前缀之积
	tmp := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			res[i] = 1
		} else {
			res[i] = nums[i-1] * res[i-1]
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			tmp = 1
			res[i] *= 1
		} else {
			tmp *= nums[i+1]
			res[i] *= tmp
		}
		fmt.Println(tmp)
	}

	return res
}

func compress0(chars []byte) int {
	n := len(chars)
	if n < 2 {
		return n
	}
	s := ""
	Len, j := 0, 0
	for i := 0; i < n-1; i += Len {
		ch := chars[i]
		j = i + 1
		Len = 1
		for j < n && ch == chars[j] {
			Len++
			j++
		}
		if Len == 1 {
			s += string(chars[i])
		} else {
			s += string(chars[i]) + strconv.Itoa(Len)
		}
	}
	fmt.Println(s)
	chars = []byte(s)
	fmt.Println(chars)
	return len(s)
}

func moveZeroes(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}
	left, right := 0, 0
	for right < n {

		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	return
}

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	j := 0
	for i := 0; i < m; i++ {
		if t[i] == s[j] && j < n {
			j++
		}
	}
	return j == n
}

func maxOperations(nums []int, k int) int {
	n := len(nums)
	ans := 0
	i, j := 0, n-1
	sort.Ints(nums)
	for i < j {
		temp := nums[i] + nums[j]
		if temp > k {
			j--
		} else if temp < k {
			i++
		} else {
			i++
			j--
			ans++
		}
	}
	return ans
}

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	maxSum := sum
	for i := k; i < n; i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return float64(sum) / float64(k)
}

// abciiidef
func maxVowels(s string, k int) int {
	Max := 0
	i, j := 0, k-1
	for j < len(s) {
		fmt.Println(s[i : j+1])
		count := Calc(s[i:j+1], k)
		if Max < count {
			Max = count
		}
		i++
		j++
	}
	return Max
}

func Calc(res string, k int) (count int) {
	count = 0
	for i := 0; i < k; i++ {
		if strings.Contains("aeiou", string(res[i])) {
			count++
		}
	}
	return
}

func longestOnes(nums []int, k int) int {
	res := 0
	left, right := 0, 0
	cnt := 0
	for right < len(nums) {
		if nums[right] == 0 {
			cnt++
		}
		right++

		for cnt > k {
			fmt.Println("cnt = ", cnt)
			if nums[left] == 0 {
				cnt--
				nums[left] = 1
			}
			left++
			fmt.Println(nums)
		}
		fmt.Println("right-left = ", right-left, "\nres = ", res)
		res = max(res, right-left)
	}
	return res
}
func longestSubarray(nums []int) int {
	left, right := 0, 0
	n := len(nums)
	maxLen := 0
	cnt := 0
	for right < n {
		if nums[right] == 0 {
			cnt++
		}
		for cnt > 1 {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}
		right++
		maxLen = max(maxLen, right-left)
	}
	return maxLen
}
func pivotIndex(nums []int) int {
	n := len(nums)
	prefix, suffix := nums[0], nums[n-1]
	res := 0
	i, j := 1, len(nums)-2
	for i < j {

		if prefix < suffix {
			prefix += nums[i]
			i++
		}

		if prefix > suffix {
			suffix += nums[j]
			j--
		}
		if prefix == suffix {
			res = i
		}
	}

	return res
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1, set2 := map[int]bool{}, map[int]bool{}

	for _, v := range nums1 {
		set1[v] = true
	}

	for _, v := range nums2 {
		set2[v] = true
	}

	var a, b []int
	for v := range set1 {
		if !set2[v] {
			a = append(a, v)
		}
	}

	for v := range set2 {
		if !set1[v] {
			b = append(b, v)
		}
	}
	return [][]int{a, b}
}

func uniqueOccurrences(arr []int) bool {
	hash := map[int]int{}
	for _, v := range arr {
		hash[v]++
	}
	fmt.Println(hash)
	times := map[int]struct{}{}
	for _, c := range hash {
		times[c] = struct{}{}
	}
	fmt.Println(times)
	return len(times) == len(hash)
}

func closeStrings(word1 string, word2 string) bool {
	n, m := len(word1), len(word2)
	if n != m {
		return false
	}

	for _, v := range word1 {
		if !strings.Contains(word2, string(v)) {
			return false
		}
	}

	hash1, hash2 := map[rune]int{}, map[rune]int{}

	arr1, arr2 := []int{}, []int{}

	//记录每个字符出现的次数
	for _, v := range word1 {
		_, ok := hash1[v]
		if ok {
			hash1[v]++
		} else {
			hash1[v] = 1
		}
	}

	for _, v := range hash1 {
		arr1 = append(arr1, v)
	}
	sort.Ints(arr1)

	for _, v := range word2 {
		_, ok := hash2[v]
		if ok {
			hash2[v]++
		} else {
			hash2[v] = 1
		}
	}

	for _, v := range hash2 {
		arr2 = append(arr2, v)
	}
	sort.Ints(arr2)

	if reflect.DeepEqual(arr1, arr2) {
		return true
	}
	return false

}

func equalPairs(grid [][]int) int {
	n := len(grid)
	cnt := make(map[string]int)
	for _, row := range grid {
		cnt[fmt.Sprint(row)]++
	}

	res := 0
	for j := 0; j < n; j++ {
		var arr []int
		for i := 0; i < n; i++ {
			arr = append(arr, grid[i][j])
		}
		if val, ok := cnt[fmt.Sprint(arr)]; ok {
			res += val
		}
	}
	return res
}

func decodeString(s string) string {
	strStack := make([]string, 0)
	numStack := make([]int, 0)
	currentStr := ""
	currentNum := 0
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if char >= "0" && char <= "9" { // 数字
			num, _ := strconv.Atoi(char)
			currentNum = currentNum*10 + num
			fmt.Println(currentNum)
		} else if char == "[" { // 左括号
			// 入栈并重置
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentStr)
			currentNum = 0
			currentStr = ""
		} else if char == "]" { // 右括号
			// 出栈并解码
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			currentStr = prevStr + strings.Repeat(currentStr, num)
		} else {
			currentStr += char
		}
	}
	return currentStr
}

func deleteMiddle(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre, slow, fast := dummyHead, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = slow.Next
	return dummyHead.Next
}
func reverseList(head *ListNode) *ListNode {
	var preNode *ListNode
	curNode := head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
	}
	return preNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	fast, slow := head, head
	// get middle node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow.Next
	slow.Next = nil
	fast = Reverse(fast)

	res := head.Val + fast.Val
	head = head.Next
	fast = fast.Next
	for head != nil && fast != nil {
		tmp := head.Val + fast.Val
		res = max(res, tmp)
		head = head.Next
		fast = fast.Next
	}
	return res

}

func Reverse(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		p := head
		head = head.Next
		p.Next = res
		res = p
	}
	return res
}

func main() {
	//// 创建两个逆序链表
	//l1 := convertToList(1000000000000000000000000000001) // 表示数字 342
	//l2 := convertToList(465)                             // 表示数字 465
	//
	//fmt.Print("链表 l1: ")
	//printList(l1)
	//
	//fmt.Print("链表 l2: ")
	//printList(l2)
	//
	//// 计算链表和
	//result := addTwoNumbers(l1, l2)
	//
	//fmt.Print("链表相加结果: ")
	//printList(result)
	//fmt.Println(subtractProductAndSum(123))
	// fmt.Println(median([]int{1, 2, 3, 4}))
	// fmt.Println(findMedianSortedArrays1([]int{1, 4}, []int{2, 3, 5, 6}))
	// fmt.Println(longestPalindrome2("19966991"))
	// fmt.Println(convert0("PAYPALISHIRING", 3))
	//fmt.Println(reverse0(123))
	//fmt.Println(reverse0(-123))
	//fmt.Println(reverse0(120))
	//fmt.Println(reverse0(0))
	// fmt.Println(myAtoi("9223372036854775808"))
	//fmt.Println(isPalindrome0(1001))
	//fmt.Println(isPalindrome2(10))
	//fmt.Println(isMatch("aab", "c*a*b"))
	//fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	// fmt.Println(intToRoman(1234))
	//fmt.Println(mergeAlternately("abcd", "aaccb"))
	//fmt.Println(gcdOfStrings("abcd", "abc"))
	//fmt.Println(kidsWithCandies([]int{1, 2, 3, 4, 6}, 3))
	// fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 2))
	//fmt.Println(reverseVowels("hello"))
	//fmt.Println(reverseWords("a good   example"))
	//fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	//fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}), "\n")
	//fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
	//test := []int{1, 2, 3, 0, 15, 16}
	//moveZeroes(test)
	//fmt.Println(test)
	//fmt.Println(isSubsequence("b", "t"))
	//fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
	//nums := []int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}
	//fmt.Println(findMaxAverage(nums, 93))
	// fmt.Println(maxVowels("abciiidef", 3))
	//fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	//fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
	//fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	//fmt.Println(findDifference([]int{1, 7, 3, 6, 5, 6}, []int{2, 4, 6}))
	// fmt.Println(uniqueOccurrences([]int{1, 7, 3, 6, 5, 6}))
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
	// fmt.Println(closeStrings("abc", "bca"))
	//fmt.Println(equalPairs([][]int{
	//	{3, 2, 1},
	//	{1, 7, 6},
	//	{2, 7, 7}}))
	//
	//stu1 := &stu{
	//	name:   "wcl",
	//	age:    20,
	//	gender: 0,
	//}
	//
	//stu2 := stu{
	//	name:   "wcl",
	//	age:    20,
	//	gender: 0,
	//}
	//
	//if flag := reflect.DeepEqual(stu1, stu2); flag {
	//	fmt.Println("equal")
	//} else {
	//	fmt.Println("not equal")
	//}
	// fmt.Println(decodeString("3[a]2[bc]"))
	// fmt.Println(decode("3[a2[c]]"))
	//obj := Constructor()
	//fmt.Println(obj.Ping(1))
	//fmt.Println(obj.Ping(100))
	//fmt.Println(obj.Ping(3001))
	//fmt.Println(obj.Ping(3002))
	//valList := []int{1, 2, 4, 5, 6}
	//head := &ListNode{Val: valList[0]}
	//tail := head
	//for i := 1; i < len(valList); i++ {
	//	tail.Next = &ListNode{Val: valList[i]}
	//	tail = tail.Next
	//}
	//test := deleteMiddle(head)
	//printLinklist(test)

	valList := []int{4, 2, 2, 3, 6, 7, 9, 120}
	head := &ListNode{Val: valList[0]}
	tail := head
	for i := 1; i < len(valList); i++ {
		tail.Next = &ListNode{Val: valList[i]}
		tail = tail.Next
	}
	//test := reverseList(head)
	//printLinklist(test)
	//test := pairSum(head)
	//fmt.Println(test)
	midNode(head)
}

func midNode(head *ListNode) {
	// 这里的fast必须是head.Next
	fast, slow := head.Next, head
	// get middle node
	for fast != nil && fast.Next != nil {
		fmt.Println(*slow, *fast)
		slow = slow.Next
		fast = fast.Next.Next
	}
	fmt.Println("Last Result:")
	fmt.Println(*slow, *fast)
}

func printLinklist(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type RecentCounter struct {
	request []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (c *RecentCounter) Ping(t int) int {
	// 将新的时间戳加入队列
	c.request = append(c.request, t)

	// 删除超出时间范围的时间戳
	for c.request[0] < t-3000 {
		c.request = c.request[1:]
	}
	return len(c.request)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
