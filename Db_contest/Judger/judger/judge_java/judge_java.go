package judge_java

import (
	"Judger/judger"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func JudgeJava(code, username, problemNum string) int {

	fmt.Println("Code:", code)
	fmt.Println("Username:", username)
	fmt.Println("ProblemNum:", problemNum)
	fmt.Println()

	err := CreateSourceFile(code, problemNum)
	if err != nil {
		return judger.SystemError
	}

	path := judger.Dir + "\\" + "Problem_" + problemNum

	fileName := "Main"

	if err = Compiler(path, fileName+".java"); err != nil {
		fmt.Println("err:", err)
		return judger.ComplierError
	} //编译错误

	Delpath := judger.Dir + "\\" + "Problem_" + problemNum + "\\"
	defer judger.RemoveFilesWithSuffix(Delpath, "java")
	defer judger.RemoveFilesWithSuffix(Delpath, "exe")
	defer judger.RemoveFilesWithSuffix(Delpath, "cpp")
	defer judger.RemoveFilesWithSuffix(Delpath, "go")
	defer judger.RemoveFilesWithSuffix(Delpath, "py")
	defer judger.RemoveFilesWithSuffix(Delpath, "class")
	//删除文件

	if files, err := ioutil.ReadDir(path); err != nil {
		return judger.SystemError //没有测试数据集，系统错误！
	} else {
		for _, fileInfo := range files {
			if fileInfo.IsDir() {
				continue
			}
			if strings.HasSuffix(fileInfo.Name(), ".in") {
				if !RunAndCompare(Delpath, path+"\\"+strings.ReplaceAll(fileInfo.Name(), ".in", ""), username) {
					return judger.WrongAnswer
				}
			}
		}
	}

	return judger.Accepted

}
func RunAndCompare(path string, dataPath string, username string) bool {
	// java -cp . Main > output.txt && del Main.class
	c := fmt.Sprintf("/C java -cp %s Main < %s > %s", path, dataPath+".in", dataPath+"."+username+".out")
	cmd := exec.Command("cmd", strings.Split(c, " ")...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	err := cmd.Run()

	// 在执行完命令后删除临时文件
	defer os.Remove(dataPath + "." + username + ".out")
	if err != nil {
		return false
	}

	return judger.Compare(dataPath+".out", dataPath+"."+username+".out")
}

func CreateSourceFile(code, problemNum string) error {
	path := judger.Dir + "\\" + "Problem_" + problemNum
	filePath := path + "\\" + "Main.java"
	// 创建或截断Java文件
	problemFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	// 写入CPP文件
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
