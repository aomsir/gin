package main

import (
	"demo4/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.DefaultRouterInit(router) // 分组路由注册
	routers.AdminRoutersInit(router)

	router.Run()
}
