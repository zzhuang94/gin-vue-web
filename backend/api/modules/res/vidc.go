package res

import (
	"backend/api/tool"
	"backend/g"
	"backend/models/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetVidcList 获取 Vidc 列表
// GET /api/res/vidc
func GetVidcList(c *gin.Context) {
	var vidcs []res.Vidc
	query := g.CoreDB.NewSession()

	// 解析分页参数
	page, pageSize := tool.ParsePagination(c)

	// 支持搜索过滤
	if name := c.Query("name"); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if vidcType := c.Query("type"); vidcType != "" {
		query = query.Where("type = ?", vidcType)
	}
	if location := c.Query("location"); location != "" {
		query = query.Where("location LIKE ?", "%"+location+"%")
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	total, err := query.Count(&res.Vidc{})
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}

	// 应用分页并查询
	query = tool.ApplyPagination(query, page, pageSize)
	err = query.Find(&vidcs)
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}

	// 返回分页响应
	tool.SuccessWithPagination(c, vidcs, total, page, pageSize)
}

// GetVidcById 根据 ID 获取 Vidc
// GET /api/res/vidc/:id
func GetVidcById(c *gin.Context) {
	// 解析 ID
	id, err := tool.ParseID(c)
	if err != nil {
		tool.BadRequest(c, "无效的 ID 参数")
		return
	}

	// 查询数据
	var vidc res.Vidc
	has, err := g.CoreDB.ID(id).Get(&vidc)
	if err != nil {
		tool.InternalServerError(c, fmt.Sprintf("查询失败: %v", err))
		return
	}
	if !has {
		tool.NotFound(c, "Vidc 不存在")
		return
	}

	// 返回成功响应
	tool.Success(c, vidc)
}
