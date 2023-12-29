package judge_python

import (
	"Judger/judger"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func JudgePy(code, username, problemNum string) int {

	fmt.Println("Code:", code)
	fmt.Println("Username:", username)
	fmt.Println("ProblemNum:", problemNum)
	fmt.Println()

	err := CreateSourceFile(code, username, problemNum)
	if err != nil {
		return judger.SystemError
	}

	path := judger.Dir + "\\" + "Problem_" + problemNum
	fileName := problemNum + "_" + username

	//if err = Compiler(path, fileName); err != nil {
	//	fmt.Println("err:", err)
	//	return judger.ComplierError
	//} //编译错误
	//
	defer os.Remove(judger.Dir + "\\" + "Problem_" + problemNum + "\\" + fileName + ".py")
	//删除文件

	if files, err := ioutil.ReadDir(path); err != nil {
		//fmt.Println("do not have ...")
		return judger.SystemError //没有测试数据集，系统错误！
	} else {
		for _, fileInfo := range files {
			if fileInfo.IsDir() {
				continue
			}
			if strings.HasSuffix(fileInfo.Name(), ".in") {

				if !RunAndComparePy(path+"\\"+fileName+".py",
					path+"\\"+strings.ReplaceAll(fileInfo.Name(), ".in", ""), username) {
					return judger.WrongAnswer
				}
			}
		}
	}

	return judger.Accepted

}

func RunAndComparePy(path string, dataPath string, username string) bool {
	cmd := exec.Command("python", path)
	inputFile, err := os.Open(dataPath + ".in")
	if err != nil {
		return false
	}
	defer inputFile.Close()
	cmd.Stdin = inputFile
	// 执行命令并获取标准输出
	// 创建一个字节缓冲区以捕获标准输出
	var outputBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer

	// 创建一个字节缓冲区以捕获标准错误输出
	var errorBuffer bytes.Buffer
	cmd.Stderr = &errorBuffer
	cmd.Run()

	ansFile, err := os.OpenFile(dataPath+username+".out", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return false
	}
	// 写入文件
	output := outputBuffer.String()
	_, err = ansFile.WriteString(output)
	if err != nil {
		// 写入文件失败时，关闭文件并返回错误
		ansFile.Close()
		return false
	}

	//fmt.Println("user file:", dataPath+"."+username+".out")

	ok := judger.Compare(dataPath+".out", dataPath+username+".out")
	ansFile.Close()
	defer os.Remove(dataPath + username + ".out")
	return ok
}

func CreateSourceFile(code, username, problemNum string) error {
	path := judger.Dir + "\\" + "Problem_" + problemNum
	fileName := problemNum + "_" + username
	filePath := path + "\\" + fileName + ".py"

	// 创建或截断py文件
	problemFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	// 写入py文件
	_, err = problemFile.WriteString(code)
	if err != nil {
		// 写入文件失败时，关闭文件并返回错误
		problemFile.Close()
		return err
	}

	// 关闭文件
	if err := problemFile.Close(); err != nil {
		return err
	}

	// 读取文件内容并打印
	_, err = ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return nil
}
