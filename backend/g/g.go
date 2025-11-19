package g

import (
	"xorm.io/xorm"
)

var (
	C      *cfg
	Rules  map[string][]*Rule
	BaseDB *xorm.Engine
	CoreDB *xorm.Engine
)

func Init() error {
	if err := initCfg(); err != nil {
		return err
	}
	if err := initRules(); err != nil {
		return err
	}

	C.Log.InitLogrus()

	if err := initDB(); err != nil {
		return err
	}
	return nil
}
