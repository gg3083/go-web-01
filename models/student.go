package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Student struct {
	Id       int
	Name     string
	BirthDay string
	Sex      bool
}

func init() {
	orm.RegisterModel(new(Student))
}

func GetAllStudents() []*Student {
	o := orm.NewOrm()
	o.Using("default")
	var students []*Student
	q := o.QueryTable("student")
	q.All(&students)
	return students
}

func GetStudentById(id int) Student {
	fmt.Println("start query")
	u := Student{Id: id}
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(&u)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
	return u
}

func AddStudent(s *Student) int {
	fmt.Println("start add")
	o := orm.NewOrm()
	o.Using("default")
	o.Insert(s)
	return s.Id
}
func UpdateStudent(s *Student) {
	fmt.Println("start update")
	o := orm.NewOrm()
	o.Using("default")
	o.Update(s)

}
func DeleteStudent(id int) {
	fmt.Println("start delete")
	o := orm.NewOrm()
	o.Using("default")
	o.Delete(&Student{Id: id})

}
