package controllers

import (
	"encoding/json"

	"../models"
	"github.com/astaxie/beego"
)

type StudentController struct {
	beego.Controller
}

// @Title addStudent
// @Description create students
// @Param       body            body    models.Student  true            "body for student content"
// @Success 200 {int} models.Student.Id
// @router / [post]
func (u *StudentController) Post() {
	var student models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &student)
	uid := models.AddStudent(&student)
	u.Data["json"] = map[string]int{"uid": uid}
	u.ServeJSON()
}

// @Title 获得所有学生
// @Description 返回所有的学生数据
// @Success 200 {object} models.Student
// @router / [get]
func (u *StudentController) GetAllStudents() {
	students := models.GetAllStudents()
	u.Data["json"] = students
	u.ServeJSON()
}

// @Title updateStudent
// @Description update the student
// @Param       uid             path    string  true            "The uid you want to update"
// @Param       body            body    models.Student  true            "body for student content"
// @Success 200 {object} models.Student
// @Failure 403 :uid is not int
// @router / [put]
func (u *StudentController) UpdateStudent() {
	var student models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &student)
	models.UpdateStudent(&student)
	u.Data["json"] = student
	u.ServeJSON()
}

// @Title deleteStudent
// @Description delete the student
// @Param       uid             path    string  true            "The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:id [delete]
func (u *StudentController) DeleteStudent() {
	uid, _ := u.GetInt(":id")
	models.DeleteStudent(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
