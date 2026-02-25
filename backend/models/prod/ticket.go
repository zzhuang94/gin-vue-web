package prod

import (
	"backend/g"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	StatusInit     = "INIT"     // 初始
	StatusPlanned  = "PLANNED"  // 已编排
	StatusPrepared = "PREPARED" // 已准备
	StatusRunning  = "RUNNING"  // 生产中
	StatusFinished = "FINISHED" // 已完成
	StatusStopped  = "STOPPED"  // 已停止
)

type Ticket struct {
	g.Model `xorm:"extends"`

	Category    string `xorm:"category" json:"category"`         // 类型
	Material    string `xorm:"material" json:"material"`         // 材质
	Color       string `xorm:"color" json:"color"`               // 颜色
	Quantity    int    `xorm:"quantity" json:"quantity,string"`  // 期望数量
	Progress    string `xorm:"progress" json:"progress"`         // 生产进度
	LeadTime    string `xorm:"lead_time" json:"lead_time"`       // 交付时间
	MachinePlan string `xorm:"machine_plan" json:"machine_plan"` // 机器编排
	MachineList string `xorm:"machine_list" json:"machine_list"` // 机器列表
	Status      string `xorm:"status" json:"status"`             // 状态
	Remark      string `xorm:"remark" json:"remark"`             // 备注

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
		t.Progress = fmt.Sprintf("0/0/%d", t.Quantity)
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

func (t *Ticket) Receive(sess *g.Sess, count int, remark string) error {
	if err := t.updateProgress(sess, count, 0); err != nil {
		return err
	}

	store := &Store{}
	has, err := sess.Where(
		"category = ? AND material = ? AND color = ?",
		t.Category, t.Material, t.Color,
	).Get(store)
	if err != nil {
		return err
	}
	if !has {
		store.Category = t.Category
		store.Material = t.Material
		store.Color = t.Color
		if err := store.Save(sess); err != nil {
			return err
		}
	}
	return store.Plus(sess, t, count, remark)
}

func (t *Ticket) Reject(sess *g.Sess, count int, remark string) error {
	if err := t.updateProgress(sess, 0, count); err != nil {
		return err
	}

	reject := &Reject{
		TicketId: t.Id,
		Count:    count,
		Remark:   remark,
		User:     sess.Ctx.GetString("username"),
	}
	return reject.Save(sess)
}

func (t *Ticket) updateProgress(sess *g.Sess, good, rej int) error {
	ss := strings.Split(t.Progress, "/")
	s0, _ := strconv.Atoi(ss[0])
	s1, _ := strconv.Atoi(ss[1])
	good = s0 + good
	rej = s1 + rej
	if good+rej > t.Quantity {
		return fmt.Errorf(
			"该工单良品数[%d]+劣品数[%d]不能大于期望数量[%d]",
			good, rej, t.Quantity,
		)
	}
	t.Progress = fmt.Sprintf("%d/%d/%d", good, rej, t.Quantity)

	// 如果良品数等于期望数量，则状态为已完成
	if good == t.Quantity {
		t.Status = StatusFinished
	}

	return t.Save(sess)
}
