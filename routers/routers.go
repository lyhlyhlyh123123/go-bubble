package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	// v1
	v1Group := r.Group("/v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)

		// 修改
		v1Group.PUT("/todo/:id", controller.ModifiyTodo)

		// 查询
		v1Group.GET("/todo", controller.GetTodo)

		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
	}
	return r
}
