package routers

import (
	"github.com/gin-gonic/gin"
	"todo-list/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/heath", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.GET("/todo/describe", getTodo)
	r.GET("/todo/list", TodoList)
	r.POST("/todo/add", addTodo)
	r.POST("/todo/delete", deleteTodo)
	r.POST("/todo/edit", editTodo)
	return r
}
