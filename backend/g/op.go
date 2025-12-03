package g

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Op struct {
	Name string         `json:"name"`
	Db   string         `json:"db"`
	Show []string       `json:"show"`
	Rela map[string]FKs `json:"rela"`
}

type FKs map[string]string

func initOps() error {
	bytes, err := os.ReadFile("op.json")
	if err != nil {
		return err
	}
	ops := make(map[string]*Op)
	if err = json.Unmarshal(bytes, &ops); err != nil {
		return fmt.Errorf("parse op file failed: %v", err)
	}
	Ops = ops
	return nil
}

type Event struct {
	Model  `xorm:"extends"`
	User   string `xorm:"user" json:"user"`
	Path   string `xorm:"path" json:"path"`
	Remark string `xorm:"remark" json:"remark"`
}

func (Event) TableName() string {
	return "op_event"
}

func (Event) New() ModelX {
	return &Event{}
}

func (Event) Save(*Sess) error {
	return nil
}

func (Event) Delete(*Sess) error {
	return nil
}

type Log struct {
	Model     `xorm:"extends"`
	Eid       int    `xorm:"eid" json:"eid,string"`
	Uuid      string `xorm:"uuid" json:"uuid"`
	Op        int    `xorm:"op" json:"op,string"`
	DataTable string `xorm:"data_table" json:"data_table"`
	DataId    int    `xorm:"data_id" json:"data_id,string"`
	DataOld   string `xorm:"data_old" json:"data_old"`
	DataNew   string `xorm:"data_new" json:"data_new"`
}

func (Log) TableName() string {
	return "op_log"
}

func recordUserLog(user, path string) {
	sql := "INSERT INTO user_log (username, path) VALUES (?, ?)"
	BaseDB.Exec(sql, user, path)
}

func recordOp(c *gin.Context) {
	log := c.GetBool("log_exists")
	if !log {
		return
	}
	ok := c.GetBool("op_ok")
	uuid := c.GetString("op_uuid")
	if !ok {
		BaseDB.Where("uuid = ?", uuid).Delete(new(Log))
		return
	}
	event := &Event{
		User:   c.GetString("username"),
		Path:   c.GetString("path"),
		Remark: c.GetString("op_remark"),
	}
	BaseDB.Insert(event)
	sql := "UPDATE op_log SET eid = ?, uuid = '' WHERE uuid = ?"
	BaseDB.Exec(sql, event.Id, uuid)
}

func (l *Log) CalcDiffs() []map[string]any {
	op := Ops[l.DataTable]
	rules := Rules[l.DataTable]
	od, nd := l.GetOdNd()
	keys := l.CalcDiffKeys(od, nd, op, rules)
	ans := make([]map[string]any, 0)
	for _, r := range rules {
		if _, ok := keys[r.Key]; !ok {
			continue
		}
		if r.Trans != nil {
			r.Translate([]map[string]string{od, nd})
		}
		ans = append(ans, map[string]any{
			"rule": r,
			"old":  od[r.Key],
			"new":  nd[r.Key],
		})
	}
	return ans
}

func (l *Log) GetOdNd() (map[string]string, map[string]string) {
	od := make(map[string]string)
	json.Unmarshal([]byte(l.DataOld), &od)
	nd := make(map[string]string)
	json.Unmarshal([]byte(l.DataNew), &nd)
	return od, nd
}

func (l *Log) CalcDiffKeys(od, nd map[string]string, op *Op, rules []*Rule) map[string]bool {
	ans := make(map[string]bool)
	for _, k := range op.Show {
		ans[k] = true
	}
	if l.Op != 0 {
		if len(ans) > 0 {
			return ans
		}
		for _, r := range rules {
			ans[r.Key] = true
		}
		return ans
	}
	for k, v := range od {
		if nd[k] != v {
			ans[k] = true
		}
	}
	return ans
}

func (l *Log) Rollback(sess *Sess) error {
	_, err := sess.Exec(l.buildRollbackArgs()...)
	if err != nil {
		return fmt.Errorf("执行回滚SQL失败: %v", err)
	}

	rollbackLog := &Log{
		Uuid:      sess.Ctx.GetString("op_uuid"),
		Op:        -l.Op,
		DataTable: l.DataTable,
		DataId:    l.DataId,
		DataOld:   l.DataNew,
		DataNew:   l.DataOld,
	}
	_, err = BaseDB.Insert(rollbackLog)
	if err != nil {
		return fmt.Errorf("插入回滚log失败: %v", err)
	}

	sess.Ctx.Set("log_exists", true)
	return nil
}

func (l *Log) buildRollbackArgs() []any {
	if l.Op == 1 {
		sql := fmt.Sprintf("DELETE FROM `%s` WHERE id = ?", l.DataTable)
		return []any{sql, l.DataId}
	}

	od, _ := l.GetOdNd()

	if l.Op == 0 {
		sets := make([]string, 0)
		vals := make([]any, 0)
		for k, v := range od {
			sets = append(sets, fmt.Sprintf("`%s` = ?", k))
			vals = append(vals, v)
		}
		vals = append(vals, l.DataId)
		sql := fmt.Sprintf("UPDATE `%s` SET %s WHERE id = ?", l.DataTable, strings.Join(sets, ", "))
		return append([]any{sql}, vals...)
	}

	var args, keys []string
	var vals []any
	for k, v := range od {
		args = append(args, "?")
		keys = append(keys, fmt.Sprintf("`%s`", k))
		vals = append(vals, v)
	}
	sql := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%s)", l.DataTable, strings.Join(keys, ", "), strings.Join(args, ", "))
	return append([]any{sql}, vals...)
}

type relyChecker struct {
	Eids []string
}

func CheckRely(eids []string) []string {
	rc := &relyChecker{Eids: eids}
	return rc.check()
}

func (rc *relyChecker) check() []string {
	// TODO check op rely
	// example:
	// op1: a -> b; op2: b -> c;  ------>  op2 rely op1
	// if rollback op1, must rollback op2 first
	return []string{}
}
