package g

import (
	"backend/libs"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"xorm.io/xorm"
)

var (
	C      *cfg
	Ops    map[string]*Op
	Rules  map[string][]*Rule
	BaseDB *xorm.Engine
	CoreDB *xorm.Engine
)

type cfg struct {
	Env   string                    `json:"env"`
	Name  string                    `json:"name"`
	Debug bool                      `json:"debug"`
	Port  int                       `json:"port"`
	Log   *libs.LogConf             `json:"log"`
	DBs   map[string]*libs.MysqlCfg `json:"dbs"`
	Redis *libs.RedisCfg            `json:"redis"`
}

func Init() error {
	if err := initCfg(); err != nil {
		return err
	}
	if err := initRules(); err != nil {
		return err
	}
	if err := initOps(); err != nil {
		return err
	}

	C.Log.InitLogrus()

	if err := initDB(); err != nil {
		return err
	}
	return nil
}

func initCfg() error {
	var cf = flag.String("c", "./cfg.json", "ths path of config file")
	flag.Parse()

	bytes, err := os.ReadFile(*cf)
	if err != nil {
		return fmt.Errorf("read cfg file failed: %v", err)
	}

	c := new(cfg)
	if err = json.Unmarshal(bytes, c); err != nil {
		return fmt.Errorf("parse cfg file [%s] failed: %v", *cf, err)
	}

	C = c
	return nil
}

func initDB() error {
	var err error
	if BaseDB, err = libs.NewMysqlEngine(C.DBs["base"]); err != nil {
		return err
	}
	if CoreDB, err = libs.NewMysqlEngine(C.DBs["core"]); err != nil {
		return err
	}
	return nil
}
