package user

import (
	"api/api/base"
	"github.com/JobNing/user-rpc/client/user"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userInfo, err := user.GetUser(c, 1)
	if err != nil {
		base.Resp(c, nil, err)
		return
	}

	base.Resp(c, userInfo, err)
	return
}
