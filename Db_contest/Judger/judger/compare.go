package judger

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Compare(pathA, pathB string) bool {
	contentA, errA := ioutil.ReadFile(pathA)
	if errA != nil {
		fmt.Println("Error reading file:", errA)
		return false
	}

	contentB, errB := ioutil.ReadFile(pathB)
	if errB != nil {
		fmt.Println("Error reading file:", errB)
		return false
	}

	strContentA := strings.TrimSpace(string(contentA))
	strContentB := strings.TrimSpace(string(contentB))

	return strContentA == strContentB
}
