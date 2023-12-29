package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Score    int
}

type UserSolved struct {
	gorm.Model
	Username string
	Solved   string
}

type ProblemList struct {
	gorm.Model
	ProblemID          int    //问题ID
	ProblemScore       int    //问题分值
	ProblemTitle       string //问题标题
	ProblemLore        string //问题内容
	ProblemStandardIn  string //问题输入样例
	ProblemStandardOut string //问题输出样例
	ProblemTips        string //问题提示
	Author             string //发布人
}

type Problems struct {
	gorm.Model
	ProblemID    int    //问题ID
	ProblemTitle string //问题标题
	Author       string //发布人
}

type Admins struct {
	gorm.Model
	Username string
}
