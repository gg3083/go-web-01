package controllers

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type StudentController struct {
	beego.Controller
}

func (S StudentController) List() {
	student := models.GetAllStudents()
	S.Data["json"] = student
	S.ServeJSON()

}

func (S StudentController) Add() {
	var student models.Student
	json.Unmarshal(S.Ctx.Input.RequestBody, &student)
	//sex, error := S.GetBool("sex")
	//if error != nil {
	//	fmt.Println(error)
	//}
	//student.Sex = sex
	sid := models.AddStudent(&student)
	S.Data["json"] = map[string]int{"sid": sid}
	S.ServeJSON()

}

func (S StudentController) Del() {

	id, error := S.GetInt("id")
	if error != nil {
		fmt.Println(error)
	}
	models.DeleteStudent(id)
	S.Data["json"] = map[string]int{"sid": id}
	S.ServeJSON()

}
func (S StudentController) UpdateData() {

	var student models.Student
	json.Unmarshal(S.Ctx.Input.RequestBody, &student)
	//sex, error := S.GetBool("sex")
	//if error != nil {
	//	fmt.Println(error)
	//}
	//student.Sex = sex
	models.UpdateStudent(&student)
	S.Data["json"] = student
	S.ServeJSON()

}
