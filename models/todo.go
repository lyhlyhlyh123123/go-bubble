package models

import "bubble/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodo() (todolist []Todo, err error) {
	err = dao.DB.Debug().Find(&todolist).Error
	return
}
func GetATodo(id string) (todo *Todo, err error) {
	err = dao.DB.Debug().Where("id = ?", id).First(&todo).Error
	return
}
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Debug().Save(todo).Error
	return
}
func DeleteATodo(id string) (err error) {
	err = dao.DB.Debug().Where("id = ?", id).Delete(Todo{}).Error
	return
}
