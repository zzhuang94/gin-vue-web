package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"

	"github.com/gin-gonic/gin"
)

type Material struct {
	*frm.X
}

func NewMaterial() *Material {
	r := &Material{X: frm.NewX(&prod.Material{})}
	r.DB = g.CoreDB
	r.Dump = true
	r.Tool = [][]string{{"批量新增", "plus", "batch-add-modal"}}
	return r
}

func (r *Material) ActionBatchAddModal(c *gin.Context) {
	r.BatchAddModal(c)
}

func (r *Material) ActionBatchAdd(c *gin.Context) {
	r.BatchAdd(c)
}
