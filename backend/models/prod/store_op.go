package prod

import (
	"backend/g"
)

const (
	OpPlus  = "PLUS"  // 入库
	OpMinus = "MINUS" // 出库
	OpEdit  = "EDIT"  // 盘点
)

type StoreOp struct {
	g.Model `xorm:"extends"`

	StoreId  int    `xorm:"store_id" json:"store_id,string"`   // 仓库ID
	TicketId int    `xorm:"ticket_id" json:"ticket_id,string"` // 工单ID
	Op       string `xorm:"op" json:"op"`                      // 操作
	Count    int    `xorm:"count" json:"count,string"`         // 数量
	Remark   string `xorm:"remark" json:"remark"`              // 备注
	User     string `xorm:"user" json:"user"`                  // 操作人
}

func (StoreOp) TableName() string {
	return "store_op"
}

func (StoreOp) New() g.ModelX {
	return &StoreOp{}
}

func (s *StoreOp) Save(sess *g.Sess) error {
	return s.SaveBean(sess, s)
}

func (s *StoreOp) Delete(sess *g.Sess) error {
	return s.DeleteBean(sess, s)
}
