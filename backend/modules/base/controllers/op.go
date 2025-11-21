package controllers

import (
	"backend/g"
	"strings"

	"github.com/gin-gonic/gin"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type Op struct {
	*g.X
}

func NewOp() *Op {
	r := &Op{X: g.NewX(&g.Event{})}
	r.initRules()

	r.BuildQuery = func(cond builder.Cond, withSelect bool) *xorm.Session {
		return r.buildQuery(cond, withSelect)
	}
	return r
}

func (r *Op) initRules() {
	r.Rules = append(r.Rules, g.Rules["op_log"]...)
	tableR := &g.Rule{Key: "data_table", Name: "数据表", Search: 1}
	for tableName, op := range g.Ops {
		tableR.Limit = append(tableR.Limit, &g.Limit{
			Key:   tableName,
			Label: op.Name,
		})
	}
	r.Rules = append(r.Rules, tableR)
	r.Rules = append(r.Rules, &g.Rule{Key: "key_str", Name: "关键字", Search: 2})
	r.Rules = append(r.Rules, &g.Rule{Key: "e.created", Name: "操作时间", Search: 2})
}

func (r *Op) buildQuery(cond builder.Cond, withSelect bool) *xorm.Session {
	ans := g.BaseDB.Table("op_event e").Join("LEFT", "op_log l", "e.id = l.eid")
	ans.Where(cond).GroupBy("e.id")
	if withSelect {
		fields := []string{"e.id", "e.user", "e.path"}
		fields = append(fields, "DATE_FORMAT(e.created, '%Y-%m-%d %H:%i:%s') AS created")
		fields = append(fields, "GROUP_CONCAT(l.id ORDER BY l.id) AS lids")
		ans.OrderBy("e.id DESC")
		ans.Select(strings.Join(fields, ", "))
	}
	return ans
}

func (r *Op) ActionIndex(c *gin.Context) {
	data := gin.H{
		"ops":       g.Ops,
		"rules":     r.Rules,
		"arg":       r.GetUriArg(c),
		"page_size": r.GetPageSize(c),
	}
	r.RenderData(c, data)
}

func (r *Op) ActionLog(c *gin.Context) {
	id := c.Query("id")
	log := new(g.Log)
	g.BaseDB.Where("id = ?", id).Get(log)
	c.JSON(200, gin.H{
		"log":   log,
		"rules": g.Rules[log.DataTable],
	})
}
