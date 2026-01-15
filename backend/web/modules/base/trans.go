package base

import (
	"backend/g"
	"backend/web/frm"
	"fmt"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type Trans struct {
	*frm.Web
}

func NewTrans() *Trans {
	return &Trans{Web: frm.NewWeb()}
}

func (t *Trans) ActionInit(c *gin.Context) {
	sql, db, ts, err := t.parseSqlAndDB(c)
	if err != nil {
		t.JsonFail(c, err)
		return
	}
	sql += fmt.Sprintf(" WHERE `%s` = '%s'", ts.Key, c.Query("val"))
	rows, _ := db.SQL(sql).QueryString()
	if len(rows) == 0 {
		t.JsonFail(c, fmt.Errorf("数据不存在"))
		return
	}
	t.JsonSucc(c, rows[0])
}

func (t *Trans) ActionLoad(c *gin.Context) {
	sql, db, ts, err := t.parseSqlAndDB(c)
	if err != nil {
		t.JsonFail(c, err)
		return
	}
	term := c.Query("term")
	sql += fmt.Sprintf(
		" WHERE `%s` like '%%%s%%' ORDER BY length(`%s`), `%s` LIMIT 10",
		ts.Val, term, ts.Val, ts.Val,
	)
	rows, _ := db.SQL(sql).QueryString()
	t.JsonSucc(c, rows)
}

func (t *Trans) parseSqlAndDB(c *gin.Context) (string, *xorm.Engine, *g.Trans, error) {
	ts := new(g.Trans)
	if err := c.ShouldBindJSON(ts); err != nil {
		return "", nil, nil, err
	}
	db := g.BaseDB
	if ts.DB == "core" {
		db = g.CoreDB
	}
	sql := ts.SQL
	if sql == "" {
		sql = fmt.Sprintf("SELECT `%s` as `key`, `%s` as `val` FROM `%s`", ts.Key, ts.Val, ts.Table)
	}
	return sql, db, ts, nil
}
