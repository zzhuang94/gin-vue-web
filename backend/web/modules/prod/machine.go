package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"

	"github.com/gin-gonic/gin"
)

type Machine struct {
	*frm.XB[*prod.Machine]
}

func NewMachine() *Machine {
	r := &Machine{XB: frm.NewXB(&prod.Machine{})}
	r.DB = g.CoreDB
	r.Dump = true
	r.Tool = [][]string{{"批量新增", "plus", "batch-add-modal"}}
	return r
}

func (r *Machine) ActionBatchAddModal(c *gin.Context) {
	r.BatchAddModal(c)
}

func (r *Machine) ActionBatchAdd(c *gin.Context) {
	r.BatchAdd(c)
}
