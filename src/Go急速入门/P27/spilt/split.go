package split

import "strings"

// split.go

// Split 把s按照sep进行分割，返回一个字符串切片
// Split("我爱你", "爱") => ['我', '你']
func Split(s, sep string) (ret []string) {
	// 减少内存申请次数
	ret = make([]string, 0, strings.Count(s, sep)+1)
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
