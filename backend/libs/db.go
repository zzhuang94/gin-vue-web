package libs

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type MysqlConfig struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	URI     string   `json:"uri"`
	ShowSQL bool     `json:"show_sql"`
	MaxOpen int      `json:"max_open"`
	MaxIdle int      `json:"max_idle"`
	IsMgr   bool     `json:"is_mgr"`
	Baks    []string `json:"baks"`
}

func NewMysqlEngine(c *MysqlConfig) (*xorm.Engine, error) {
	engine, err := tryNewEngine(c.URI, c.IsMgr, c.Baks)
	if err != nil {
		errInfo := fmt.Sprintf("[MySQL] %s ping failed", c.Name)
		txtInfo := errInfo + " : " + err.Error()
		fmt.Println(txtInfo)
		return nil, err
	}
	engine.SetMaxOpenConns(c.MaxOpen)
	engine.SetMaxIdleConns(c.MaxIdle)
	engine.ShowSQL(c.ShowSQL)
	return engine, nil
}

func tryNewEngine(uri string, isMgr bool, baks []string) (*xorm.Engine, error) {
	xdb, err := xorm.NewEngine("mysql", uri)
	if err == nil && isMysqlOK(xdb, isMgr) {
		return xdb, nil
	}
	for _, bakURI := range baks {
		bakXDB, err := tryNewEngine(bakURI, isMgr, []string{})
		if err == nil && isMysqlOK(bakXDB, true) {
			fmt.Println("use bak: ", bakXDB)
			return bakXDB, nil
		}
	}
	return nil, fmt.Errorf("连接数据库 %s 失败: %s , 备用数据库也失败", uri, err)
}

func isMysqlOK(xdb *xorm.Engine, isMgr bool) bool {
	err := xdb.Ping()
	if err != nil {
		fmt.Println("Mysql Ping error: ", err)
		return false
	}
	if !isMgr {
		return true
	}
	mgrErr := isMgrOK(xdb)
	if mgrErr != nil {
		fmt.Println("Mysql MGR error: ", err)
		return false
	}
	return true
}

func isMgrOK(xdb *xorm.Engine) error {
	sql := "SELECT @@global.super_read_only AS read_only"
	rows := []struct {
		ReadOnly int `xorm:"read_only"`
	}{}
	err := xdb.SQL(sql).Find(&rows)
	if err != nil {
		return err
	}
	if len(rows) != 1 {
		return fmt.Errorf("查询MGR正常状态失败")
	}
	if rows[0].ReadOnly != 0 {
		return fmt.Errorf("MGR只读")
	}
	return nil
}
