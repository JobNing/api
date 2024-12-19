package user

import (
	"context"
	"fmt"
	"github.com/JobNing/user-rpc/client/user"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	fmt.Println("收到请求")
	userInfo, err := user.GetUser(context.Background(), 1)

	fmt.Println(userInfo)
	fmt.Println(err)
}
