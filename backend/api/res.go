package api

import (
	"backend/g"
	"backend/models/res"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RESTful API 响应结构
type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// 成功响应
func jsonSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Code: 1,
		Data: data,
	})
}

// 失败响应
func jsonError(c *gin.Context, message string, statusCode ...int) {
	code := http.StatusBadRequest
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	c.JSON(code, ApiResponse{
		Code:    0,
		Message: message,
	})
}

// ==================== Ip API ====================

// GetIpList 获取 IP 列表
// GET /api/res/ip
func GetIpList(c *gin.Context) {
	var ips []res.Ip
	query := g.CoreDB.NewSession()

	// 支持分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 支持搜索
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
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Limit(pageSize, offset).Find(&ips)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	jsonSuccess(c, gin.H{
		"list":      ips,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetIpById 根据 ID 获取 IP
// GET /api/res/ip/:id
func GetIpById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonError(c, "无效的 ID")
		return
	}

	var ip res.Ip
	has, err := g.CoreDB.ID(id).Get(&ip)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}
	if !has {
		jsonError(c, "IP 不存在", http.StatusNotFound)
		return
	}

	jsonSuccess(c, ip)
}

// ==================== Service API ====================

// GetServiceList 获取 Service 列表
// GET /api/res/service
func GetServiceList(c *gin.Context) {
	var services []res.Service
	query := g.CoreDB.NewSession()

	// 支持分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 支持搜索
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
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Limit(pageSize, offset).Find(&services)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	jsonSuccess(c, gin.H{
		"list":      services,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetServiceById 根据 ID 获取 Service
// GET /api/res/service/:id
func GetServiceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonError(c, "无效的 ID")
		return
	}

	var service res.Service
	has, err := g.CoreDB.ID(id).Get(&service)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}
	if !has {
		jsonError(c, "Service 不存在", http.StatusNotFound)
		return
	}

	jsonSuccess(c, service)
}

// ==================== Policy API ====================

// GetPolicyList 获取 Policy 列表
// GET /api/res/policy
func GetPolicyList(c *gin.Context) {
	var policies []res.Policy
	query := g.CoreDB.NewSession()

	// 支持分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 支持搜索
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
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Limit(pageSize, offset).Find(&policies)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	jsonSuccess(c, gin.H{
		"list":      policies,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetPolicyById 根据 ID 获取 Policy
// GET /api/res/policy/:id
func GetPolicyById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonError(c, "无效的 ID")
		return
	}

	var policy res.Policy
	has, err := g.CoreDB.ID(id).Get(&policy)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}
	if !has {
		jsonError(c, "Policy 不存在", http.StatusNotFound)
		return
	}

	jsonSuccess(c, policy)
}

// ==================== Vidc API ====================

// GetVidcList 获取 Vidc 列表
// GET /api/res/vidc
func GetVidcList(c *gin.Context) {
	var vidcs []res.Vidc
	query := g.CoreDB.NewSession()

	// 支持分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 支持搜索
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
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Limit(pageSize, offset).Find(&vidcs)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	jsonSuccess(c, gin.H{
		"list":      vidcs,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetVidcById 根据 ID 获取 Vidc
// GET /api/res/vidc/:id
func GetVidcById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonError(c, "无效的 ID")
		return
	}

	var vidc res.Vidc
	has, err := g.CoreDB.ID(id).Get(&vidc)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}
	if !has {
		jsonError(c, "Vidc 不存在", http.StatusNotFound)
		return
	}

	jsonSuccess(c, vidc)
}

// ==================== VidcIp API ====================

// GetVidcIpList 获取 VidcIp 列表
// GET /api/res/vidc-ip
func GetVidcIpList(c *gin.Context) {
	var vidcIps []res.VidcIp
	query := g.CoreDB.NewSession()

	// 支持分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 支持搜索
	if vidcId := c.Query("vidc_id"); vidcId != "" {
		query = query.Where("vidc_id = ?", vidcId)
	}
	if ipId := c.Query("ip_id"); ipId != "" {
		query = query.Where("ip_id = ?", ipId)
	}
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	total, err := query.Count(&res.VidcIp{})
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Limit(pageSize, offset).Find(&vidcIps)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}

	jsonSuccess(c, gin.H{
		"list":      vidcIps,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetVidcIpById 根据 ID 获取 VidcIp
// GET /api/res/vidc-ip/:id
func GetVidcIpById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		jsonError(c, "无效的 ID")
		return
	}

	var vidcIp res.VidcIp
	has, err := g.CoreDB.ID(id).Get(&vidcIp)
	if err != nil {
		jsonError(c, fmt.Sprintf("查询失败: %v", err), http.StatusInternalServerError)
		return
	}
	if !has {
		jsonError(c, "VidcIp 不存在", http.StatusNotFound)
		return
	}

	jsonSuccess(c, vidcIp)
}
