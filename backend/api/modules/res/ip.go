package res

import (
	"backend/api/tool"
	"backend/g"
	"backend/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetIpList 获取 IP 列表
// GET /api/res/ip
func GetIpList(c *gin.Context) {
	var ips []res.Ip
	query := g.CoreDB.NewSession()

	// 解析分页参数
	page, pageSize := tool.ParsePagination(c)

	// 支持搜索过滤
	if ip := c.Query("ip"); ip != "" {
		query = query.Where("ip LIKE ?", "%"+ip+"%")
	}
	if ipType := c.Query("type"); ipType != "" {
		query = query.Where("type = ?", ipType)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	total, err := query.Count(&res.Ip{})
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}

	// 应用分页并查询
	query = tool.ApplyPagination(query, page, pageSize)
	err = query.Find(&ips)
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}

	// 返回分页响应
	tool.SuccessWithPagination(c, ips, total, page, pageSize)
}

// GetIpById 根据 ID 获取 IP
// GET /api/res/ip/:id
func GetIpById(c *gin.Context) {
	// 解析 ID
	id, err := tool.ParseID(c)
	if err != nil {
		tool.BadRequest(c, "无效的 ID 参数")
		return
	}

	// 查询数据
	var ip res.Ip
	has, err := g.CoreDB.ID(id).Get(&ip)
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}
	if !has {
		tool.NotFound(c, "IP 不存在")
		return
	}

	// 返回成功响应
	tool.Success(c, ip)
}
