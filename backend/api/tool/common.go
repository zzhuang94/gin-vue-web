package tool

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

// ApiResponse 标准 RESTful API 响应结构
type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// PaginationResponse 分页响应结构
type PaginationResponse struct {
	List     any   `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
}

// Success 成功响应
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ApiResponse{
		Code: 1,
		Data: data,
	})
}

// SuccessWithPagination 分页成功响应
func SuccessWithPagination(c *gin.Context, list any, total int64, page, pageSize int) {
	Success(c, PaginationResponse{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// Error 错误响应
func Error(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, ApiResponse{
		Code:    0,
		Message: message,
	})
}

// BadRequest 400 错误响应
func BadRequest(c *gin.Context, message string) {
	Error(c, message, http.StatusBadRequest)
}

// NotFound 404 错误响应
func NotFound(c *gin.Context, message string) {
	Error(c, message, http.StatusNotFound)
}

// InternalServerError 500 错误响应
func InternalServerError(c *gin.Context, message string) {
	Error(c, message, http.StatusInternalServerError)
}

// ParsePagination 解析分页参数
func ParsePagination(c *gin.Context) (page, pageSize int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100 // 限制最大页面大小
	}
	return page, pageSize
}

// ParseID 解析路径中的 ID 参数
func ParseID(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

// ApplyPagination 应用分页到查询
func ApplyPagination(query *xorm.Session, page, pageSize int) *xorm.Session {
	offset := (page - 1) * pageSize
	return query.Limit(pageSize, offset)
}
