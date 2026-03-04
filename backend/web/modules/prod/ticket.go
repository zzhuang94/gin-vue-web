package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Ticket struct {
	*frm.X
}

func NewTicket() *Ticket {
	r := &Ticket{X: frm.NewX(&prod.Ticket{})}
	r.DB = g.CoreDB
	r.NoID = true
	r.Dump = true
	r.AndWheres = []map[string]any{
		{
			"status": []string{
				prod.StatusInit,
				prod.StatusPlanned,
				prod.StatusPrepared,
				prod.StatusRunning,
			},
		},
	}
	r.BuildTopMenu = func(c *gin.Context) []*frm.Menu {
		return r.buildTopMenu(c)
	}
	r.BuildTableMenu = func(c *gin.Context) []*frm.TableMenu {
		return r.buildTableMenu(c)
	}
	r.WrapData = func(data []map[string]string) {
		r.wrapData(data)
	}
	return r
}

func (t *Ticket) buildTopMenu(c *gin.Context) []*frm.Menu {
	user := t.GetUser(c)
	if user.IsSales || user.IsManager {
		return []*frm.Menu{{
			Title: "新增需求",
			Icon:  "pencil",
			URL:   "apply",
			Type:  "modal",
			Color: "primary",
		}}
	}
	return []*frm.Menu{}
}

func (t *Ticket) buildTableMenu(c *gin.Context) []*frm.TableMenu {
	user := t.GetUser(c)
	ans := []*frm.TableMenu{}
	ans = append(ans, t.WrapTableMenu([]string{"查看详情", "eye", "read"}))
	if user.IsSales || user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"修改需求", "pencil", "apply"}),
			Conds: []*frm.MenuCond{{Key: "status", Val: "INIT", Comp: "EQ"}},
		})
	}
	if user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"编排机器", "grip-vertical", "plan"}),
			Conds: []*frm.MenuCond{{Key: "status", Val: []string{"INIT", "PLANNED"}, Comp: "IN"}},
		})
	}
	if user.IsWorker || user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"准备机器", "list", "prepare"}),
			Conds: []*frm.MenuCond{{Key: "status", Val: []string{"PLANNED", "PREPARED"}, Comp: "IN"}},
		})
	}
	if user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"开始生产", "play", "run", "async"}),
			Conds: []*frm.MenuCond{{Key: "status", Val: "PREPARED", Comp: "EQ"}},
		})
	}
	if user.IsStorekeeper || user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"库存信息", "box", "/prod/store/detail", "link"}),
			Args:  []*frm.MenuArg{{Key: "store_id", Val: "id"}},
			Conds: []*frm.MenuCond{{Key: "status", Val: "RUNNING", Comp: "EQ"}},
		})
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"良品入库", "download", "/prod/store/plus"}),
			Args:  []*frm.MenuArg{{Key: "store_id", Val: "id"}},
			Conds: []*frm.MenuCond{{Key: "status", Val: "RUNNING", Comp: "EQ"}},
		})
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"劣品上报", "warning", "/prod/store/reject"}),
			Args:  []*frm.MenuArg{{Key: "store_id", Val: "id"}},
			Conds: []*frm.MenuCond{{Key: "status", Val: "RUNNING", Comp: "EQ"}},
		})
	}
	if user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"终止生产", "stop", "stop", "async"}),
			Conds: []*frm.MenuCond{{Key: "status", Val: "RUNNING", Comp: "EQ"}},
		})
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"完成生产", "check", "finish", "async"}),
			Conds: []*frm.MenuCond{{Key: "status", Val: "RUNNING", Comp: "EQ"}},
		})
		ans = append(ans, t.WrapTableMenu([]string{"修改工单", "edit", "edit"}))
		ans = append(ans, t.WrapTableMenu([]string{"删除工单", "trash", "delete", "async"}))
	}
	return ans
}

func (t *Ticket) ActionRead(c *gin.Context) {
	rules := t.GetRules()
	fields := []string{
		"category", "material", "color",
		"quantity", "machine_plan", "machine_list",
		"status", "remark",
	}
	rules = t.RulesFilter(rules, fields)
	rules = t.RulesReadonly(rules, fields)
	title := "<i class='fa fa-eye'></i>&nbsp;&nbsp;查看详情"
	t.editModel(c, title, rules)
}

func (t *Ticket) ActionApply(c *gin.Context) {
	rules := t.GetRules()
	rules = t.RulesFilter(rules, []string{
		"category", "material", "color", "quantity", "lead_time", "remark",
	})
	title := "<i class='fa fa-pencil'></i>&nbsp;&nbsp;生产需求"
	t.editModel(c, title, rules)
}

func (t *Ticket) ActionPlan(c *gin.Context) {
	rules := t.GetRules()
	rules = t.RulesFilter(rules, []string{
		"apply_user", "apply_time",
		"category", "material", "color", "quantity", "lead_time",
		"machine_plan", "remark",
	})
	rules = t.RulesReadonly(rules, []string{
		"category", "material", "color", "quantity", "lead_time",
	})
	title := "<i class='fa fa-grip-vertical'></i>&nbsp;&nbsp;编排机器"
	t.editModel(c, title, rules)
}

func (t *Ticket) ActionPrepare(c *gin.Context) {
	rules := t.GetRules()
	rules = t.RulesFilter(rules, []string{
		"plan_user", "plan_time", "machine_plan", "machine_list", "remark",
	})
	rules = t.RulesReadonly(rules, []string{"machine_plan"})
	title := "<i class='fa fa-list'></i>&nbsp;&nbsp;准备机器"
	t.editModel(c, title, rules)
}

func (t *Ticket) ActionEdit(c *gin.Context) {
	rules := t.GetRules()
	rules = t.RulesFilter(rules, []string{
		"category", "material", "color", "quantity",
		"apply_user", "plan_user", "prepare_user",
		"lead_time", "image", "machine_plan", "machine_list", "remark",
	})
	title := "<i class='fa fa-edit'></i>&nbsp;&nbsp;修改工单"
	t.editModel(c, title, rules)
}

func (t *Ticket) ActionRun(c *gin.Context) {
	if err := t.updateStatus(c, prod.StatusRunning); err != nil {
		t.JsonFail(c, err)
	} else {
		t.JsonSucc(c, "开始生产成功")
	}
}

func (t *Ticket) ActionStop(c *gin.Context) {
	if err := t.updateStatus(c, prod.StatusStopped); err != nil {
		t.JsonFail(c, err)
	} else {
		t.JsonSucc(c, "停止生产成功")
	}
}

func (t *Ticket) ActionFinish(c *gin.Context) {
	if err := t.updateStatus(c, prod.StatusFinished); err != nil {
		t.JsonFail(c, err)
	} else {
		t.JsonSucc(c, "完成生产成功")
	}
}

func (t *Ticket) editModel(c *gin.Context, title string, rules []*g.Rule) {
	props := gin.H{
		"action": "save",
		"title":  title,
		"data":   gin.H{},
		"rules":  rules,
	}
	id := c.DefaultQuery("id", "")
	if id != "" {
		m := t.Model.New()
		has, err := t.DB.ID(id).Get(m)
		if err != nil || !has {
			t.JsonFail(c, fmt.Errorf("数据不存在"))
			return
		}
		props["data"] = m
		props["action"] = "save?id=" + id
	}
	t.ModalPage(c, props, "components/edit")
}

func (t *Ticket) updateStatus(c *gin.Context, status string) error {
	m := &prod.Ticket{}
	has, err := t.DB.ID(c.DefaultQuery("id", "")).Get(m)
	if err != nil || !has {
		return fmt.Errorf("数据不存在")
	}
	m.Status = status
	if status == prod.StatusRunning {
		m.RunUser = t.GetUsername(c)
		m.RunTime = time.Now().Format("2006-01-02 15:04:05")
	}
	sess := t.BeginSess(t.DB, c)
	if err := m.Save(sess); err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}

func (t *Ticket) wrapData(data []map[string]string) {
	today := time.Now()
	for _, d := range data {
		lead, _ := time.Parse("2006-01-02", d["lead_time"])
		if lead.Before(today) {
			d["lead_time"] = "<b class='text-danger'>" + d["lead_time"] + "</b>"
		} else if lead.Sub(today).Hours()/24 < 3 {
			d["lead_time"] = "<b class='text-warning'>" + d["lead_time"] + "</b>"
		}
	}
}
