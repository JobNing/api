package main

import (
	"api/roter"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	roter.Register(r)
	r.Run(":8082")
}
