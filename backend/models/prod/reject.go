package prod

import (
	"backend/g"
)

type Reject struct {
	g.Model `xorm:"extends"`

	TicketId int    `xorm:"ticket_id" json:"ticket_id,string"` // 工单ID
	Count    int    `xorm:"count" json:"count,string"`         // 数量
	Remark   string `xorm:"remark" json:"remark"`              // 备注
	User     string `xorm:"user" json:"user"`                  // 操作人
}

func (Reject) TableName() string {
	return "reject"
}

func (Reject) New() g.ModelX {
	return &Reject{}
}

func (r *Reject) Save(sess *g.Sess) error {
	return r.SaveBean(sess, r)
}

func (r *Reject) Delete(sess *g.Sess) error {
	return r.DeleteBean(sess, r)
}
