package admin

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(context *gin.Context) {
	con.success(context)
}
