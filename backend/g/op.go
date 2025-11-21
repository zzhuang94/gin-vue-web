package g

import (
	"encoding/json"
	"fmt"
	"os"

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
	ok := c.GetBool("op_ok")
	uuid := c.GetString("op_uuid")
	if !ok {
		BaseDB.Where("uuid = ?", uuid).Delete(new(Log))
		return
	}
	has, _ := BaseDB.Where("uuid = ?", uuid).Exist(new(Log))
	if !has {
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
