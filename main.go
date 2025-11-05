package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	// 创建数据库

	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()

	r.Run(":8080")
}
