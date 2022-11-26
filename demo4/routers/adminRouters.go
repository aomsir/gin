package routers

import (
	"demo4/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(router *gin.Engine) {
	adminRouters := router.Group("/admin")
	{
		// /admin/userList
		adminRouters.GET("/list", admin.UserController{}.Index)
	}
}
