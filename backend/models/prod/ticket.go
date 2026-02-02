package prod

import (
	"backend/g"
	"time"
)

const (
	StatusInit      = "INIT"      // 初始
	StatusPlanned   = "PLANNED"   // 已编排
	StatusPrepared  = "PREPARED"  // 已准备
	StatusRunning   = "RUNNING"   // 生产中
	StatusFinished  = "FINISHED"  // 已完成
	StatusFailed    = "FAILED"    // 已失败
	StatusCancelled = "CANCELLED" // 已取消
)

type Ticket struct {
	g.Model `xorm:"extends"`

	Title            string `xorm:"title" json:"title"`                                // 标题
	Sku              string `xorm:"sku" json:"sku"`                                    // 产品型号
	Material         string `xorm:"material" json:"material"`                          // 材质
	Color            string `xorm:"color" json:"color"`                                // 颜色
	QuantityExpected int    `xorm:"quantity_expected" json:"quantity_expected,string"` // 期望数量
	QuantityActual   int    `xorm:"quantity_actual" json:"quantity_actual,string"`     // 交付数量
	LeadTime         string `xorm:"lead_time" json:"lead_time"`                        // 交付时间
	Image            string `xorm:"image" json:"image"`                                // 图片
	MachinePlan      string `xorm:"machine_plan" json:"machine_plan"`                  // 机器编排
	MachineList      string `xorm:"machine_list" json:"machine_list"`                  // 机器列表
	Status           string `xorm:"status" json:"status"`                              // 状态
	Remark           string `xorm:"remark" json:"remark"`                              // 备注

	ApplyUser   string `xorm:"apply_user" json:"apply_user"`     // 申请人
	ApplyTime   string `xorm:"apply_time" json:"apply_time"`     // 申请时间
	PlanUser    string `xorm:"plan_user" json:"plan_user"`       // 编排人
	PlanTime    string `xorm:"plan_time" json:"plan_time"`       // 编排时间
	PrepareUser string `xorm:"prepare_user" json:"prepare_user"` // 准备人
	PrepareTime string `xorm:"prepare_time" json:"prepare_time"` // 准备时间
	RunUser     string `xorm:"run_user" json:"run_user"`         // 生产人
	RunTime     string `xorm:"run_time" json:"run_time"`         // 生产时间
}

func (Ticket) TableName() string {
	return "ticket"
}

func (Ticket) New() g.ModelX {
	return &Ticket{}
}

func (t *Ticket) Save(sess *g.Sess) error {
	if t.Id == 0 {
		t.ApplyUser = sess.Ctx.GetString("username")
		t.ApplyTime = time.Now().Format("2006-01-02 15:04:05")
		t.Status = StatusInit
	} else if t.Status == StatusInit && t.MachinePlan != "" {
		t.PlanUser = sess.Ctx.GetString("username")
		t.PlanTime = time.Now().Format("2006-01-02 15:04:05")
		t.Status = StatusPlanned
	} else if t.Status == StatusPlanned && t.MachineList != "" {
		t.PrepareUser = sess.Ctx.GetString("username")
		t.PrepareTime = time.Now().Format("2006-01-02 15:04:05")
		t.Status = StatusPrepared
	}
	return t.SaveBean(sess, t)
}

func (t *Ticket) Delete(sess *g.Sess) error {
	return t.DeleteBean(sess, t)
}
