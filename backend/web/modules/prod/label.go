package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"

	"github.com/gin-gonic/gin"
)

type Label struct {
	*frm.XB[*prod.Label]
}

func NewLabel() *Label {
	r := &Label{XB: frm.NewXB(&prod.Label{})}
	r.DB = g.CoreDB
	r.Dump = true
	r.Tool = [][]string{{"批量新增", "plus", "batch-add-modal"}}
	return r
}

func (r *Label) ActionBatchAddModal(c *gin.Context) {
	r.BatchAddModal(c)
}

func (r *Label) ActionBatchAdd(c *gin.Context) {
	r.BatchAdd(c)
}
