package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//store := cookie.NewStore([]byte("secret11"))      // 配置session中间件,参数为加密的密钥
	//router.Use(sessions.Sessions("mySession", store)) // store是存储引擎

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "123456", []byte("secret111"))
	router.Use(sessions.Sessions("mySession", store))

	// 设置session的值
	router.GET("/", func(context *gin.Context) {
		session := sessions.Default(context)
		session.Set("user", "Aomsir") // 设置Session的值
		session.Save()                // 保存Session的值,必须调用
	})

	// 获取session的值
	router.GET("/getSession", func(context *gin.Context) {
		session := sessions.Default(context)
		user := session.Get("user")
		context.JSON(http.StatusOK, gin.H{
			"sessionId":    session.ID(),
			"sessionValue": user,
		})
	})

	router.Run()
}
