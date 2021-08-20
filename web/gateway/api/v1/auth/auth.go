package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
)

func RegisterRouter(rg *fizz.RouterGroup, handlerFunc ...gin.HandlerFunc) {
	for _, v := range handlerFunc {
		rg.Use(v)
	}
}
