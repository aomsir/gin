package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (con BaseController) success(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"success": "baseController调用成功",
	})
}
