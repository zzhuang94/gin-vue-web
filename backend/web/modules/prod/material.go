package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"

	"github.com/gin-gonic/gin"
)

type Material struct {
	*frm.XB[*prod.Material]
}

func NewMaterial() *Material {
	r := &Material{XB: frm.NewXB(&prod.Material{})}
	r.DB = g.CoreDB
	r.Dump = true
	r.TopMenu = [][]string{{"批量新增", "plus", "batch-add-modal"}}
	return r
}

func (r *Material) ActionBatchAddModal(c *gin.Context) {
	r.BatchAddModal(c)
}

func (r *Material) ActionBatchAdd(c *gin.Context) {
	r.BatchAdd(c)
}
