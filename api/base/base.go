package base

import (
	"api/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Resp(c *gin.Context, data interface{}, err error) {
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "msg": consts.REQUEST_SUCCESS_MSG})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": -1, "data": "", "msg": err.Error()})
	return
}
