package user

import (
	"api/api/user"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	u := r.Group("/user")
	u.POST("/info", user.GetUserInfo)

	add := r.Group("/user/address")
	add.POST("/info", user.GetAddressInfo)
}
