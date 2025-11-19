package g

import (
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

type XB[T ModelX] struct {
	*X
}

func NewXB[T ModelX](m T) *XB[T] {
	x := NewX(m)
	xb := &XB[T]{
		X: x,
	}
	xb.BatchEdit = true
	xb.BatchDelete = true
	return xb
}

func (xb *XB[T]) ActionBatchEdit(c *gin.Context) {
	props := gin.H{
		"action":   "batch-save?ids=" + c.Query("ids"),
		"title":    `<i class="fa fa-edit"></i>&nbsp;&nbsp;批量修改`,
		"subtitle": fmt.Sprintf("您将修改 %s 条数据，请先勾选要修改的属性", c.Query("count")),
		"data":     gin.H{},
		"rules":    xb.GetRules(),
		"check":    true,
	}
	xb.ModalPage(c, props, "components/edit")
}

func (xb *XB[T]) ActionBatchSave(c *gin.Context) {
	list, err := xb.getList(c)
	if err != nil {
		xb.JsonFail(c, err)
		return
	}

	payload, _ := io.ReadAll(c.Request.Body)

	sess := xb.BeginSess(xb.DB)
	for _, m := range list {
		if err := xb.saveModel(m, payload, sess); err != nil {
			sess.Rollback()
			xb.JsonFail(c, err)
			return
		}
	}
	sess.Commit()
	xb.JsonSucc(c, "批量修改成功")
}

func (xb *XB[T]) ActionBatchDelete(c *gin.Context) {
	list, err := xb.getList(c)
	if err != nil {
		xb.JsonFail(c, err)
		return
	}
	sess := xb.BeginSess(xb.DB)
	for _, m := range list {
		if err := m.Delete(sess); err != nil {
			sess.Rollback()
			xb.JsonFail(c, fmt.Errorf("删除失败: %v", err))
			return
		}
	}
	sess.Commit()
	xb.JsonSucc(c, "批量删除成功")
}

func (xb *XB[T]) getList(c *gin.Context) ([]T, error) {
	ids := c.Query("ids")
	if strings.TrimSpace(ids) == "" {
		return nil, fmt.Errorf("参数错误: 缺少ids")
	}
	list := make([]T, 0)
	if err := xb.DB.In("id", strings.Split(ids, ",")).Find(&list); err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	return list, nil
}
