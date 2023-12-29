package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func RunPythonScript(scriptPath, inputPath string) (string, error) {
	cmd := exec.Command("python", scriptPath)

	// 打开输入文件并将其内容重定向到标准输入
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return "", fmt.Errorf("error opening input file: %v", err)
	}
	defer inputFile.Close()
	cmd.Stdin = inputFile

	// 创建一个字节缓冲区以捕获标准输出
	var outputBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer

	// 创建一个字节缓冲区以捕获标准错误输出
	var errorBuffer bytes.Buffer
	cmd.Stderr = &errorBuffer

	// 执行命令并等待完成
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing Python script: %v\nStderr: %s", err, errorBuffer.String())
	}

	// 从缓冲区获取输出
	output := outputBuffer.String()
	return output, nil
}

func main() {
	scriptPath := "./script.py"
	inputPath := "./input.txt"

	output, err := RunPythonScript(scriptPath, inputPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", output)
}
