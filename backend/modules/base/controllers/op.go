package controllers

import (
	"backend/g"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		"rules":     r.Rules,
		"arg":       r.GetUriArg(c),
		"page_size": r.GetPageSize(c),
		"log_rules": g.Rules["op_log"],
	}
	r.RenderData(c, data)
}

func (r *Op) ActionLog(c *gin.Context) {
	id := c.Query("id")
	log := new(g.Log)
	g.BaseDB.Where("id = ?", id).Get(log)
	c.JSON(200, gin.H{
		"log":   log,
		"diffs": log.CalcDiffs(),
		"name":  g.Ops[log.DataTable].Name,
		"time":  log.Created.Format("2006-01-02 15:04:05"),
	})
}

func (r *Op) ActionConfirm(c *gin.Context) {
	ids, _ := r.GetIds(c)
	relys := g.CheckRely(ids)
	if len(relys) == 0 {
		r.JsonSucc(c, "ok")
		return
	}
	c.JSON(200, gin.H{
		"ids": ids,
	})
}

func (r *Op) ActionRollback(c *gin.Context) {
	ids, _ := r.GetIds(c)
	logs := make([]*g.Log, 0)
	g.BaseDB.In("eid", ids).OrderBy("id DESC").Find(&logs)
	sess := r.BeginSess(g.CoreDB, c)
	for _, log := range logs {
		if err := log.Rollback(sess); err != nil {
			logrus.Errorf("rollback log %d failed: %v", log.Id, err)
			sess.Rollback()
			r.JsonFail(c, err)
			return
		}
	}
	sess.Commit()
	r.JsonSucc(c, "ok")
}
