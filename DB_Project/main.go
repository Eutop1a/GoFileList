package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"modulename/pkg/DataBase"
	"modulename/pkg/RpcFunc"
	"modulename/pkg/Security"
	"modulename/pkg/StructPackage"
	"net/http"
	"net/rpc"
	"sort"
	"time"
)

// 处理注册请求
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// 处理OPTIONS请求
		optionsHandler(w, r)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体中的 JSON 数据
	var data StructPackage.RegMsg
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 将密码进行md5加密之后再存储到数据库中
	pwd := Security.MD5(data.Password)

	// 连接数据库
	db, err := DataBase.ConnectToDB()
	if err != nil {
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	err = DataBase.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	// 存储到 MySQL 数据库
	if err := DataBase.Register(data.Nickname, data.Username, pwd, db, ""); err != nil {
		log.Printf("该用户名已经存在：%v\n", err)
		http.Error(w, "UserName already exist", http.StatusConflict) // 409
		return
	}
	// 创建token
	// 先设置过期时间
	expirationTime := time.Now().Add(10 * time.Hour)
	token, err := Security.GenerateToken(data.Username, expirationTime)
	if err != nil {
		fmt.Println("Generate token error")
		return
	}

	fmt.Println("token: ", token)
	// 将token和EndTime打包发送给前端
	SendData := StructPackage.TokenData{
		EndTime: expirationTime,
		Token:   token,
	}
	// 将TokenData对象转换为JSON格式
	jsonData, err := json.Marshal(SendData)
	if err != nil {
		log.Println("JSON encoding error:", err)
		return
	}
	// 返回响应
	w.WriteHeader(http.StatusOK) // 200
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("Error sending JSON response:", err)
	}
}

// 处理登录请求
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// 处理OPTIONS请求
		optionsHandler(w, r)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println(r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体中的 JSON 数据
	var data StructPackage.LogMsg
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 将密码进行md5加密之后再存储到数据库中
	pwd := Security.MD5(data.Password)

	// 连接数据库
	db, err := DataBase.ConnectToDB()
	if err != nil {
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	// 登录
	if err := DataBase.Login(data.Username, pwd, db); err != 0 {
		switch err {
		case 1:
			w.WriteHeader(http.StatusNotFound)
		case 2:
			w.WriteHeader(http.StatusUnauthorized)
		}
		return
	}

	// 创建token
	// 过期时间
	expirationTime := time.Now().Add(10 * time.Hour)
	token, err := Security.GenerateToken(data.Username, expirationTime)
	if err != nil {
		fmt.Println("Generate token error")
		return
	}
	fmt.Printf("token: %s\n", token)
	// 返回响应
	// 将token和EndTime打包发送给前端
	SendData := StructPackage.TokenData{
		EndTime: expirationTime,
		Token:   token,
	}

	// 将TokenData对象转换为JSON格式
	jsonData, err := json.Marshal(SendData)
	if err != nil {
		log.Println("JSON encoding error:", err)
		return
	}
	// 返回响应
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("Error sending JSON response:", err)
	}
}

// 检查是否有相同的用户名
func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// 处理OPTIONS请求
		optionsHandler(w, r)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体中的 JSON 数据
	var data StructPackage.CheckMsg
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Printf("账号：%s\n", data.Username)
	// 连接数据库
	db, err := DataBase.ConnectToDB()
	if err != nil {
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	err = DataBase.CreateTableIfNotExists(db)
	if err != nil {
		return
	}
	// 检查是否有重复
	if err := DataBase.Check(data.Username, db); err != nil {
		response := "1"
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(response))

		return
	}

	// 返回响应
	response := "0"
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))

}

// 返回nickname
func getNickNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// 处理OPTIONS请求
		optionsHandler(w, r)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// 解析请求体中的 JSON 数据
	var data StructPackage.CheckMsg
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println(data.Username)
	NickName, err := DataBase.GetNickName(data.Username)
	if err != nil {
		fmt.Println("Get NickName error")
		return
	}
	response := NickName
	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte(response))
	if err != nil {
		fmt.Println(err)
	}
}

// Check Token
func checkTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// 处理OPTIONS请求
		optionsHandler(w, r)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// 解析请求体中的 JSON 数据
	var data StructPackage.TokenData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println(data.Token)
	ErrorType := Security.ParseToken(data.Token)
	switch ErrorType {
	case 0:
		// token解析成功
		w.WriteHeader(http.StatusOK) //200
		fmt.Println("token解析成功")
		break
	case 1:
		// token过期
		w.WriteHeader(http.StatusUnauthorized) //401
		fmt.Println("token过期")
		break
	case 2:
		// token解析失败
		w.WriteHeader(http.StatusBadRequest) //400
		fmt.Println("token解析失败")
		break
	}
	return
}

//// 改成绩
//func changeScoreHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//
//	w.WriteHeader(http.StatusOK)
//
//	if r.Method != http.MethodPost {
//		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//		return
//	}
//	// 解析请求体中的 JSON 数据
//	var data StructPackage.ChangeScore
//	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
//		http.Error(w, "Bad Request", http.StatusBadRequest)
//		return
//	}
//	// 连接数据库
//	db, err := DataBase.ConnectToDB()
//	if err != nil {
//		return
//	}
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//
//		}
//	}(db)
//	res := DataBase.ChangeScore(data.Username, data.Score, db)
//	if res == 1 {
//		// 数据库操作错误
//		fmt.Println("error 1")
//	} else if res == 2 {
//		// 获取受影响行数错误
//		fmt.Println("error 2")
//	} else if res == 3 {
//		// 未找到匹配的成员
//		fmt.Println("error 3")
//		//w.WriteHeader(http.StatusNotFound)
//	}
//	fmt.Println("success")
//	//w.WriteHeader(http.StatusOK)
//	return
//}

// 处理OPTIONS请求
func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// 改成绩
func changeScoreHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// 处理OPTIONS请求
		optionsHandler(w, r)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// 解析请求体中的 JSON 数据
	var data StructPackage.ChangeScore
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("Bad Request")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Println(r.Body)
	fmt.Println(data.Username)
	// 关闭请求体
	defer r.Body.Close()
	// 连接数据库
	db, err := DataBase.ConnectToDB()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// fmt.Println("Success Connect db")
	defer db.Close()
	//scoreStr := data.Score
	//score, err := strconv.Atoi(scoreStr)
	res := DataBase.ChangeScore(data.Username, db)
	if res == 1 {
		// 数据库操作错误
		http.Error(w, "Database operate error", http.StatusInternalServerError)
		return
	} else if res == 2 {
		// 获取受影响行数错误
		http.Error(w, "effect row error", http.StatusInternalServerError)
		return
	} else if res == 3 {
		// 未找到匹配的成员
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	fmt.Println("Change score success")
	// 成功
	w.WriteHeader(http.StatusOK)
	return
}

// 按照成绩排序
func sortByScoreHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 连接数据库
	db, err := DataBase.ConnectToDB()
	if err != nil {
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	users, err := DataBase.GetUsersByScore(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 按照 score 从大到小排序
	sort.Slice(users, func(i, j int) bool {
		return users[i].Score > users[j].Score
	})

	// 将用户数组转换为 JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 发送 JSON 响应
	w.Write(jsonData)
}

func main() {
	// 异步启动
	// 注册rpc函数
	rpc.Register(new(RpcFunc.Token))

	// 启动RPC服务器
	go func() {
		rpc.HandleHTTP()
		log.Printf("Serving RPC server on port %d", 33335)
		if err := http.ListenAndServe(":65532", nil); err != nil {
			log.Fatal("Error serving RPC server: ", err)
		}
	}()

	// 启动HTTP服务器
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/get", getNickNameHandler)
	http.HandleFunc("/open", checkTokenHandler)
	http.HandleFunc("/changeScore", changeScoreHandler)
	http.HandleFunc("/sort", sortByScoreHandler)

	http.Handle("/", http.FileServer(http.Dir("static")))
	//log.Println("HTTP服务器启动，监听端口 8080...")
	log.Println("HTTP服务器启动，监听端口 65533...")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":65530", nil))

}
