package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"

	"github.com/gin-gonic/gin"
)

type Category struct {
	*frm.X
}

func NewCategory() *Category {
	r := &Category{X: frm.NewX(&prod.Category{})}
	r.DB = g.CoreDB
	r.Dump = true
	r.Tool = [][]string{{"批量新增", "plus", "batch-add-modal"}}
	return r
}

func (r *Category) ActionBatchAddModal(c *gin.Context) {
	r.BatchAddModal(c)
}

func (r *Category) ActionBatchAdd(c *gin.Context) {
	r.BatchAdd(c)
}
