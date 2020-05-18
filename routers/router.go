package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"todo-list/pkg/setting"
)

type HandlerFunc func(c *gin.Context) error

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)
		err = handler(c)
		if err != nil {
			//var apiException *APIException
			//if h,ok := err.(*APIException); ok {
			//	apiException = h
			//}else if e, ok := err.(error); ok {
			//	if gin.Mode() == "debug" {
			//		// 错误
			//		apiException = UnknownError(e.Error())
			//	}else{
			//		// 未知错误
			//		apiException = UnknownError(e.Error())
			//	}
			//}else{
			//	apiException = ServerError()
			//}
			//apiException.Request = c.Request.Method + " "+ c.Request.URL.String()
			//c.JSON(apiException.Code,apiException)
			return
		}
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CorsMiddleware())
	gin.SetMode(setting.RunMode)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
			"code":    200,
		})
	})
	r.GET("/todo/describe", getTodo)
	r.GET("/todo/list", wrapper(TodoList))
	r.POST("/todo/add", addTodo)
	r.POST("/todo/delete", deleteTodo)
	r.POST("/todo/edit", editTodo)
	return r
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var filterHost = [...]string{"http://localhost.*", "http://*.hfjy.com"}
		// filterHost 做过滤器，防止不合法的域名访问
		var isAccess = false
		for _, v := range filterHost {
			match, _ := regexp.MatchString(v, origin)
			if match {
				isAccess = true
			}
		}
		if isAccess {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
