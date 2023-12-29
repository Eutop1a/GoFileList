package main

import "fmt"

type student struct {
	id    int //唯一
	name  string
	class string
}

// student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学员管理的类型
type studentMgr struct {
	allStudents []*student
}

// studentMgr类型的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

//1、添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

//2、编辑学生
func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id { // 学号相同时，就表示找到了要修改的学生
			s.allStudents[i] = newStu // 根据切片的索引，直接把新学生赋值进去
			return
		}
	}
	fmt.Printf("输入的学生学号%d不存在", newStu.id)
}

//3、展示学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号： %d 姓名： %s 班级： %s\n", v.id, v.name, v.class)
	}
}
