package roter

import (
	"api/roter/good"
	"api/roter/user"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	good.Register(r)
	user.Register(r)
}
