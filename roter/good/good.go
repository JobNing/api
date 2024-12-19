package good

import (
	"api/api/good"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	g := r.Group("/good")
	g.POST("/info", good.GetGoodInfo)
}
