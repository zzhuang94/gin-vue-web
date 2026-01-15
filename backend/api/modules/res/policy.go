package res

import (
	"backend/api/tool"
	"backend/g"
	"backend/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetPolicyList 获取 Policy 列表
// GET /api/res/policy
func GetPolicyList(c *gin.Context) {
	var policies []res.Policy
	query := g.CoreDB.NewSession()

	// 解析分页参数
	page, pageSize := tool.ParsePagination(c)

	// 支持搜索过滤
	if name := c.Query("name"); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if serviceId := c.Query("service_id"); serviceId != "" {
		query = query.Where("service_id = ?", serviceId)
	}
	if vidcId := c.Query("vidc_id"); vidcId != "" {
		query = query.Where("vidc_id = ?", vidcId)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	total, err := query.Count(&res.Policy{})
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}

	// 应用分页并查询
	query = tool.ApplyPagination(query, page, pageSize)
	err = query.Find(&policies)
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}

	// 返回分页响应
	tool.SuccessWithPagination(c, policies, total, page, pageSize)
}

// GetPolicyById 根据 ID 获取 Policy
// GET /api/res/policy/:id
func GetPolicyById(c *gin.Context) {
	// 解析 ID
	id, err := tool.ParseID(c)
	if err != nil {
		tool.BadRequest(c, "无效的 ID 参数")
		return
	}

	// 查询数据
	var policy res.Policy
	has, err := g.CoreDB.ID(id).Get(&policy)
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}
	if !has {
		tool.NotFound(c, "Policy 不存在")
		return
	}

	// 返回成功响应
	tool.Success(c, policy)
}
