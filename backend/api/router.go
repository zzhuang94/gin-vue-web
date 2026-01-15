package api

import (
	"backend/api/modules/res"
	"backend/g"

	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {
	rg.Use(g.ApiAuth)

	// RESTful API for res module (查询接口)
	resGroup := rg.Group("/res")
	{
		// Ip API
		resGroup.GET("/ip", res.GetIpList)
		resGroup.GET("/ip/:id", res.GetIpById)

		// Service API
		resGroup.GET("/service", res.GetServiceList)
		resGroup.GET("/service/:id", res.GetServiceById)

		// Policy API
		resGroup.GET("/policy", res.GetPolicyList)
		resGroup.GET("/policy/:id", res.GetPolicyById)

		// Vidc API
		resGroup.GET("/vidc", res.GetVidcList)
		resGroup.GET("/vidc/:id", res.GetVidcById)
	}
}
