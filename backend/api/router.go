package api

import (
	"backend/g"

	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {
	rg.Use(g.ApiAuth)
	rg.GET("/test", Test)

	// RESTful API for res module (查询接口)
	resGroup := rg.Group("/res")
	{
		// Ip API
		resGroup.GET("/ip", GetIpList)
		resGroup.GET("/ip/:id", GetIpById)

		// Service API
		resGroup.GET("/service", GetServiceList)
		resGroup.GET("/service/:id", GetServiceById)

		// Policy API
		resGroup.GET("/policy", GetPolicyList)
		resGroup.GET("/policy/:id", GetPolicyById)

		// Vidc API
		resGroup.GET("/vidc", GetVidcList)
		resGroup.GET("/vidc/:id", GetVidcById)

		// VidcIp API
		resGroup.GET("/vidc-ip", GetVidcIpList)
		resGroup.GET("/vidc-ip/:id", GetVidcIpById)
	}
}
