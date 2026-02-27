package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"

	"github.com/gin-gonic/gin"
)

type Color struct {
	*frm.X
}

func NewColor() *Color {
	r := &Color{X: frm.NewX(&prod.Color{})}
	r.DB = g.CoreDB
	r.Dump = true
	r.Tool = [][]string{{"批量新增", "plus", "batch-add-modal"}}
	return r
}

func (r *Color) ActionBatchAddModal(c *gin.Context) {
	r.BatchAddModal(c)
}

func (r *Color) ActionBatchAdd(c *gin.Context) {
	r.BatchAdd(c)
}
