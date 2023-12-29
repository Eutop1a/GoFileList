package main

import (
	"Judger/judger"
	"Judger/judger/judge_cpp"
	"Judger/judger/judge_go"
	"Judger/judger/judge_java"
	"Judger/judger/judge_python"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os"
)

func init() {
	os.Mkdir(judger.Dir, 777)
}

type Judge int

// JudgeCpp cpp
func (c *Judge) JudgeCpp(args map[string]interface{}, result *int) error {

	*result = judge_cpp.JudgeCpp(args["code"].(string), args["username"].(string), args["problem"].(string))
	return nil
}

// JudgePy python
func (c *Judge) JudgePy(args map[string]interface{}, result *int) error {

	*result = judge_python.JudgePy(args["code"].(string), args["username"].(string), args["problem"].(string))
	return nil
}

// JudgeJava java
func (c *Judge) JudgeJava(args map[string]interface{}, result *int) error {

	*result = judge_java.JudgeJava(args["code"].(string), args["username"].(string), args["problem"].(string))
	return nil
}

// JudgeGo go
func (c *Judge) JudgeGo(args map[string]interface{}, result *int) error {

	*result = judge_go.JudgeGo(args["code"].(string), args["username"].(string), args["problem"].(string))
	return nil
}

func main() {
	rpc.Register(new(Judge))
	rpc.HandleHTTP()

	fmt.Println("############################################################")
	fmt.Println("#       This Judger RPC is working on localhost:65534       #")
	fmt.Println("############################################################")

	if err := http.ListenAndServe(":65534", nil); err != nil {
		log.Fatalf("Error serve")
	}

}
