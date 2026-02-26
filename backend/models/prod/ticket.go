package prod

import (
	"backend/g"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
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

	StoreId int `xorm:"store_id" json:"store_id,string"` // 仓库ID

	Category string `xorm:"category" json:"category"` // 类型
	Material string `xorm:"material" json:"material"` // 材质
	Color    string `xorm:"color" json:"color"`       // 颜色

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
		store, err := t.getOrCreateStore(sess)
		if err != nil {
			return err
		}
		t.StoreId = store.Id
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

func (t *Ticket) tryUpdateProgress(sess *g.Sess, goods, bads int) (int, int) {
	logrus.Infof("tryUpdateProgress: goods=%d, bads=%d", goods, bads)

	ss := strings.Split(t.Progress, "/")
	currGoods, _ := strconv.Atoi(ss[0])
	currBads, _ := strconv.Atoi(ss[1])
	newGoods, newBads := currGoods, currBads

	if goods > 0 {
		needGoods := t.Quantity - currGoods - currBads
		if needGoods < goods {
			goods -= needGoods
			newGoods += needGoods
		} else {
			newGoods += goods
			goods = 0
		}
	}
	if bads > 0 {
		needBads := t.Quantity - currGoods - currBads
		if needBads < bads {
			bads -= needBads
			newBads += needBads
		} else {
			newBads += bads
			bads = 0
		}
	}
	t.Progress = fmt.Sprintf("%d/%d/%d", newGoods, newBads, t.Quantity)
	if newGoods+newBads == t.Quantity {
		t.Status = StatusFinished
	}
	logrus.Infof("tryUpdateProgress: progress=%s", t.Progress)
	if err := t.SaveBean(sess, t); err != nil {
		logrus.Errorf("tryUpdateProgress: %v", err)
		return goods, bads
	}
	return goods, bads
}

func (t *Ticket) getOrCreateStore(sess *g.Sess) (*Store, error) {
	store := &Store{}
	has, err := sess.Where(
		"category = ? AND material = ? AND color = ?",
		t.Category, t.Material, t.Color,
	).Get(store)
	if err != nil {
		return nil, err
	}
	if !has {
		store.Category = t.Category
		store.Material = t.Material
		store.Color = t.Color
		if err := store.Save(sess); err != nil {
			return nil, err
		}
	}
	return store, nil
}
