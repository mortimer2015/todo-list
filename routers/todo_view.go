package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-list/models"
)

func getTodo(c *gin.Context) {
	id := c.Query("id")
	intId, _ := strconv.Atoi(id)
	todo := models.GetTodo(intId)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":       todo.Id,
		"name":     todo.Name,
		"value":    todo.Value,
		"describe": todo.Describe,
	})
}

func TodoList(c *gin.Context) {
	//id := c.Query("id")
	todos := models.TodoList()
	ret := make([]map[string]interface{}, 0)
	for _, todo := range todos {
		one := map[string]interface{}{
			"id":       todo.Id,
			"name":     todo.Name,
			"value":    todo.Value,
			"describe": todo.Describe,
		}
		ret = append(ret, one)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": ret,
	})
}

type TodoForm struct {
	Name     string
	Value    string
	Describe string
}

func addTodo(c *gin.Context) {
	var todoForm TodoForm
	c.Bind(&todoForm)
	todo := models.TodoModels{
		Name:     todoForm.Name,
		Value:    todoForm.Value,
		Describe: todoForm.Describe,
		Status:   "used",
	}
	_ = todo.Create()
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success",
	})
}

func editTodo(c *gin.Context) {
	var todo models.TodoModels
	_ = c.Bind(&todo)
	_ = todo.Update()
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success",
	})
}

type todoDeleteForm struct {
	Id int
}

func deleteTodo(c *gin.Context) {
	var deleteForm todoDeleteForm
	c.Bind(&deleteForm)
	//intId, _:=strconv.Atoi(deleteForm.Id)
	_ = models.DeleteTodo(deleteForm.Id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success",
	})
}
