package routes

import (
	"todo_list/api"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("some thing secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleWare.JWT())	//中间件鉴权
	}

	return r
}
