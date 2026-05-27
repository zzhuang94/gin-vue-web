package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Machine struct {
	*frm.XB[*prod.Machine]
}

func NewMachine() *Machine {
	r := &Machine{XB: frm.NewXB(&prod.Machine{})}
	r.DB = g.CoreDB
	r.Dump = true
	r.TopMenu = [][]string{
		{"批量新增", "plus", "batch-add-modal"},
	}
	return r
}

func (r *Machine) ActionBatchAddModal(c *gin.Context) {
	r.BatchAddModal(c)
}

func (r *Machine) ActionBatchAdd(c *gin.Context) {
	r.BatchAdd(c)
}

func (r *Machine) ActionBoard(c *gin.Context) {
	r.Render(c)
}

func (r *Machine) ActionBoardData(c *gin.Context) {
	machines := make([]prod.Machine, 0)
	if err := r.DB.Asc("name").Find(&machines); err != nil {
		r.JsonFail(c, err)
		return
	}

	tickets := make([]prod.Ticket, 0)
	if err := r.DB.Where("status = ?", prod.StatusRunning).Find(&tickets); err != nil {
		r.JsonFail(c, err)
		return
	}

	running := make(map[string][]gin.H)
	for i := range tickets {
		t := &tickets[i]
		label := fmt.Sprintf("%s / %s / %s", t.Category, t.Material, t.Color)
		for _, name := range strings.Split(t.MachineList, ",") {
			name = strings.TrimSpace(name)
			if name == "" {
				continue
			}
			running[name] = append(running[name], gin.H{
				"id":    t.Id,
				"label": label,
			})
		}
	}

	items := make([]gin.H, 0, len(machines))
	for _, m := range machines {
		ts := running[m.Name]
		if ts == nil {
			ts = []gin.H{}
		}
		items = append(items, gin.H{
			"id":      m.Id,
			"name":    m.Name,
			"remark":  m.Remark,
			"health":  m.Health,
			"running": len(ts) > 0,
			"tickets": ts,
		})
	}
	r.JsonSucc(c, items)
}
