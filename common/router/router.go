package router

import (
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

func init() {
	r = gin.New()
}

// Router --
func Router() *gin.Engine {
	return r
}
