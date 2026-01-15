package res

import (
	"backend/api/tool"
	"backend/g"
	"backend/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetServiceList 获取 Service 列表
// GET /api/res/service
func GetServiceList(c *gin.Context) {
	var services []res.Service
	query := g.CoreDB.NewSession()

	// 解析分页参数
	page, pageSize := tool.ParsePagination(c)

	// 支持搜索过滤
	if name := c.Query("name"); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if business := c.Query("business"); business != "" {
		query = query.Where("business = ?", business)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	total, err := query.Count(&res.Service{})
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}

	// 应用分页并查询
	query = tool.ApplyPagination(query, page, pageSize)
	err = query.Find(&services)
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}

	// 返回分页响应
	tool.SuccessWithPagination(c, services, total, page, pageSize)
}

// GetServiceById 根据 ID 获取 Service
// GET /api/res/service/:id
func GetServiceById(c *gin.Context) {
	// 解析 ID
	id, err := tool.ParseID(c)
	if err != nil {
		tool.BadRequest(c, "无效的 ID 参数")
		return
	}

	// 查询数据
	var service res.Service
	has, err := g.CoreDB.ID(id).Get(&service)
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}
	if !has {
		tool.NotFound(c, "Service 不存在")
		return
	}

	// 返回成功响应
	tool.Success(c, service)
}
