package models

import (
	"fmt"
	"log"
)

type TodoModels struct {
	Id       int
	Name     string
	Value    string
	Describe string
	Status   string
}

func (todo *TodoModels) Create() error {
	return db.Create(&todo).Error
}

func DeleteTodo(id int) error {
	todo := TodoModels{}
	todo.Id = id
	return db.Debug().Delete(&todo).Error
}

func GetTodo(id int) TodoModels {
	var t TodoModels
	u := db.Where("id = ?", id).First(&t)
	if u.Error != nil {
		log.Fatal(u)
	}
	fmt.Println(t.Name)
	return t
}

func (todo *TodoModels) Update() error {
	return db.Save(todo).Error
}

func TodoList() []*TodoModels {
	todos := make([]*TodoModels, 0)
	u := db.Find(&todos)
	if u.Error != nil {
		log.Fatal(u)
	}
	//fmt.Println(t.Name)
	return todos
}
