package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, errr%v\n", err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

func (u user) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// 查询单条
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err :%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多条
func queryMultiRowDemo() {
	sql := "select id, name, age from user where id>?"
	var users []user
	err := db.Select(&users, sql, 0)
	if err != nil {
		fmt.Printf("query failed, err :%v\n", err)
		return
	}
	fmt.Printf("users: %#v\n", users)
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func insertUserDemo() (err error) {
	_, err = db.NamedExec(`insert into user (name, age) values (:name, :age)`,
		map[string]interface{}{
			"name": "utop1a",
			"age":  20,
		})
	return
}

func namedQuery() {
	sqlStr := "SELECT * FROM user WHERE name=:name"
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "q1mi"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

	u := user{
		Name: "q1mi",
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

func transactionDemo2() (err error) {
	tx, err := db.Beginx()
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Println("Commit")
		}
	}()
	sqlStr1 := "Update user set age=20 where id=?"
	rs, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		return err
	}
	n, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	sqlStr2 := "Update user set age=50 where id=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr2 failed")
	}
	return err
}

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUser(users []*user) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err := db.Exec(stmt, valueArgs...)
	return err
}

func BatchInsertUser2(users []interface{}) error {
	query, args, _ := sqlx.In(
		"insert into user (name, age) values (?), (?), (?)",
		users...,
	)
	fmt.Println(query)
	fmt.Println(args)
	_, err := db.Exec(query, args...)
	return err
}

func BatchInsertUser3(users []*user) error {
	_, err := db.NamedExec("insert into user (name, age) values (:name, :age)", users)
	return err
}

func QueryByIDs(ids []int) (users []user, err error) {
	query, args, err := sqlx.In(
		"select name, age from user where id in (?)", ids,
	)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}

func QueryAndOrderByIDs(ids []int) (users []user, err error) {
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In(
		"select name, age from user where id in (?) order by find_in_set(id, ?)",
		ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&users, query, args...)
	return
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//u1 := user{Name: "七米", Age: 18}
	//u2 := user{Name: "q1mi", Age: 28}
	//u3 := user{Name: "小王子", Age: 38}

	// queryRowDemo()
	// queryMultiRowDemo()
	// insertRowDemo()
	// updateRowDemo()
	// deleteRowDemo()
	// insertUserDemo()
	// namedQuery()
	// transactionDemo2()
	// BatchInsertUser2()

	//// 方法1
	//users := []*user{&u1, &u2, &u3}
	//err = BatchInsertUser(users)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	//}
	//
	//// 方法2
	//users2 := []interface{}{u1, u2, u3}
	//err = BatchInsertUser2(users2)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	//}
	//
	//// 方法3
	//users3 := []*user{&u1, &u2, &u3}
	//err = BatchInsertUser3(users3)
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	//}

	//users, err := QueryByIDs([]int{4, 5, 6})
	//if err != nil {
	//	fmt.Printf("QueryByIDs failed, err:%#v\n", err)
	//	return
	//}
	//for _, user := range users {
	//	fmt.Printf("user:%#v\n", user)
	//}
	// 1、用代码排序
	// 2、让MySQL排序
	users, err := QueryAndOrderByIDs([]int{7, 5, 6, 1})
	if err != nil {
		fmt.Printf("QueryAndOrderByIDs failed, err:%#v\n", err)
		return
	}
	for _, user := range users {
		fmt.Printf("user:%#v\n", user)
	}
}
