package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/rpc"
)

type Judge struct {
	UserName     string `form:"username" json:"username"`
	Problem      string `form:"problem" json:"problem"`
	Code         string `form:"code" json:"code"`
	LanguageType string `form:"type" json:"type"`
}

const (
	Accepted      = 1000
	WrongAnswer   = 1001
	ComplierError = 1002
	TimeLimited   = 1003
	RuntimeError  = 1004
	MemoryLimited = 1005
	SystemError   = 1010
)

var flag = map[int]string{
	1000: "Accepted",
	1001: "WrongAnswer",
	1002: "ComplierError",
	1003: "TimeLimited",
	1004: "RuntimeError",
	1005: "MemoryLimited",
	1010: "SystemError",
}

func (judge *Judge) JudgeCode() (int, string) {
	// 从请求中获取执行操作的用户名
	username := judge.UserName
	code := judge.Code
	problem := judge.Problem
	languageType := judge.LanguageType

	// 连接 RPC 服务器进行代码评测
	client, _ := rpc.DialHTTP("tcp", "localhost:65534")
	defer client.Close()

	var res int

	if languageType == "C++" {
		cpp := client.Go("Judge.JudgeCpp", gin.H{
			"code":     code,
			"username": username,
			"problem":  problem,
		}, &res, nil)
		<-cpp.Done
	} else if languageType == "Python" {
		python := client.Go("Judge.JudgePy", gin.H{
			"code":     code,
			"username": username,
			"problem":  problem,
		}, &res, nil)

		<-python.Done
	} else if languageType == "Java" {
		java := client.Go("Judge.JudgeJava", gin.H{
			"code":     code,
			"username": username,
			"problem":  problem,
		}, &res, nil)

		<-java.Done
	} else if languageType == "Go" {
		Go := client.Go("Judge.JudgeGo", gin.H{
			"code":     code,
			"username": username,
			"problem":  problem,
		}, &res, nil)

		<-Go.Done
	}

	fmt.Println("res =", res)
	//// 如果评测结果为 1010，返回状态码 403
	if res == 1010 {
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, flag[res]
}
