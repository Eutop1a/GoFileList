package judge_cpp

import (
	"Judger/judger"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func JudgeCpp(code, username, problemNum string) int {

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

	if err = Compiler(path, fileName); err != nil {
		fmt.Println("err:", err)
		return judger.ComplierError
	} //编译错误
	Delpath := judger.Dir + "\\" + "Problem_" + problemNum + "\\"
	defer judger.RemoveFilesWithSuffix(Delpath, "java")
	defer judger.RemoveFilesWithSuffix(Delpath, "exe")
	defer judger.RemoveFilesWithSuffix(Delpath, "cpp")
	defer judger.RemoveFilesWithSuffix(Delpath, "go")
	defer judger.RemoveFilesWithSuffix(Delpath, "py")
	//删除文件

	if files, err := ioutil.ReadDir(path); err != nil {
		return judger.SystemError //没有测试数据集，系统错误！
	} else {
		for _, fileInfo := range files {
			if fileInfo.IsDir() {
				continue
			}
			if strings.HasSuffix(fileInfo.Name(), ".in") {
				if !RunAndCompare(path+"\\"+fileName+".exe", path+"\\"+strings.ReplaceAll(fileInfo.Name(), ".in", ""), username) {
					return judger.WrongAnswer
				}
			}
		}
	}

	return judger.Accepted

}

func RunAndCompare(path string, dataPath string, username string) bool {

	c := fmt.Sprintf("/C %s < %s > %s", path, dataPath+".in", dataPath+"."+username+".out")
	cmd := exec.Command("cmd", strings.Split(c, " ")...)

	err := cmd.Run()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err != nil {
		return false
	}

	defer os.Remove(dataPath + "." + username + ".out")
	return judger.Compare(dataPath+".out", dataPath+"."+username+".out")

}

func CreateSourceFile(code, username, problemNum string) error {
	path := judger.Dir + "\\" + "Problem_" + problemNum
	fileName := problemNum + "_" + username
	filePath := path + "\\" + fileName + ".cpp"

	// 创建或截断CPP文件
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
