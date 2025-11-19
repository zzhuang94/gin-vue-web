package api

import (
	"backend/g"

	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {
	rg.Use(g.ApiAuth)
	rg.GET("/test", Test)
}
