package DataBase

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"modulename/pkg/StructPackage"
	"sort"
)

// 定义数据库连接信息
const (
	DBUsername = "root"
	DBPassword = "123456"
	DBHost     = "localhost"
	DBPort     = 3306
	DBName     = "users"
)

// ConnectToDB 连接到数据库
func ConnectToDB() (*sql.DB, error) {
	// 构建数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateTableIfNotExists 如果表不存在就创建一个
func CreateTableIfNotExists(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nickName VARCHAR(8) NOT NULL,
        userName VARCHAR(10) NOT NULL,
        passWord VARCHAR(32) NOT NULL,
        score int NOT NULL
    )
    `
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// Login 登录
func Login(username, password string, db *sql.DB) int {
	query0 := `SELECT count(*) FROM users WHERE userName = ?`
	row0 := db.QueryRow(query0, username)
	var count int
	row0.Scan(&count)
	// 如果有相同的用户名就返回error
	if count == 0 {
		return 1 //fmt.Errorf("Do not have this UserName\n")
	}
	query := `SELECT * FROM users WHERE  userName = ? AND passWord = ?`
	row := db.QueryRow(query, username, password)
	var (
		id       int
		NickName string
		UserName string
		PassWord string
		Score    int
	)
	Err := row.Scan(&id, &NickName, &UserName, &PassWord, &Score)
	if Err != nil {
		if Err == sql.ErrNoRows {
			// 查询结果为空，用户名和密码不匹配
			return 2 //fmt.Errorf("Password error\n")
		}
	}
	return 0
}

// Register 注册
func Register(nickname, username, password string, db *sql.DB, token string) error {
	var Score = 0
	err := Check(username, db)
	// 用户名重复
	if err != nil {
		return err
	}
	// 没有就开始注册
	// 执行插入语句
	query := `INSERT INTO users (nickName, userName, passWord, score) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(query, nickname, username, password, Score)
	if err != nil {
		return err
	}
	return nil
}

// Check 检查用户名是否重复
func Check(username string, db *sql.DB) error {
	// 检查表是否存在，如果不存在则创建表
	if err := CreateTableIfNotExists(db); err != nil {
		return err
	}
	// 先检测是否有相同的用户名
	query := `SELECT COUNT(*) FROM users WHERE userName = ?`
	row := db.QueryRow(query, username)
	var count int
	row.Scan(&count)
	// 如果有相同的用户名就返回error
	if count != 0 {
		return fmt.Errorf("Exist UserName\n")
	}
	return nil
}

// GetNickName 获得昵称
func GetNickName(Username string) (string, error) {
	fmt.Println(Username)
	db, err := ConnectToDB()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer db.Close()
	query := `SELECT * FROM users WHERE  userName = ?`
	row := db.QueryRow(query, Username)
	var (
		id       int
		NickName string
		UserName string
		PassWord string
		Score    int
	)
	Err := row.Scan(&id, &NickName, &UserName, &PassWord, &Score)
	fmt.Println(id, NickName, UserName, Score)
	if Err != nil {
		fmt.Println(Err)
		return "", Err
	}
	return NickName, nil
}

// ChangeScore 改变分数
func ChangeScore(Username string, db *sql.DB) int {
	query := `UPDATE users SET score = ? + score WHERE userName = ? `
	result, err := db.Exec(query, 10, Username)
	if err != nil {
		return 1 //fmt.Errorf("数据库操作错误：%v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 2 //fmt.Errorf("获取受影响行数错误：%v", err)
	}

	if rowsAffected == 0 {
		return 3 //fmt.Errorf("未找到匹配的成员")
	}
	return 0
}

// GetUsersByScore 从数据库中按照 score 从大到小获取所有用户
func GetUsersByScore(db *sql.DB) ([]StructPackage.User, error) {
	query := `SELECT username, nickname,  score FROM users ORDER BY score DESC`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("数据库查询错误：%v", err)
	}
	defer rows.Close()

	users := make([]StructPackage.User, 0)

	for rows.Next() {
		var user StructPackage.User
		if err := rows.Scan(&user.Username, &user.Nickname, &user.Score); err != nil {
			return nil, fmt.Errorf("读取行错误：%v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历行错误：%v", err)
	}

	// 按照 score 从大到小排序
	sort.Slice(users, func(i, j int) bool {
		return users[i].Score > users[j].Score
	})

	return users, nil
}
