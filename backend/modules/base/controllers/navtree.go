package controllers

import (
	"backend/g"
	"backend/modules/base/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Navtree struct {
	*g.X
}

func NewNavtree() *Navtree {
	return &Navtree{X: g.NewX(&models.Navtree{})}
}

func (n *Navtree) ActionIndex(c *gin.Context) {
	n.RenderData(c, map[string]any{"data": n.getData()})
}

func (n *Navtree) ActionFetch(c *gin.Context) {
	c.JSON(200, map[string]any{"data": n.getData()})
}

func (n *Navtree) getData() map[int]map[string]any {
	ans := make(map[int]map[string]any)
	rows := []*models.Navtree{}
	g.BaseDB.Find(&rows)
	for _, r := range rows {
		color := "success"
		if r.Status == 0 {
			color = "danger"
		}
		ans[r.ID] = map[string]any{
			"name":      fmt.Sprintf("%s [No. %d]", r.Name, r.Rank),
			"parent_id": r.ParentId,
			"expanded":  r.Level < 2,
			"rank":      r.Rank,
			"color":     color,
		}
	}
	return ans
}

func (n *Navtree) ActionAdd(c *gin.Context) {
	props := gin.H{
		"action": "save",
		"title":  `<i class="fa fa-plus"></i>&nbsp;&nbsp;新增子节点`,
		"data":   gin.H{"parent_id": c.Query("id")},
		"rules":  n.GetRules(),
	}
	n.ModalPage(c, props, "components/edit")
}
