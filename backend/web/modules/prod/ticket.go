package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
	"fmt"
	"strings"
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
	r.BuildTool = func(c *gin.Context) []*frm.Tool {
		return r.buildTool(c)
	}
	r.BuildOption = func(c *gin.Context) []*frm.Option {
		return r.buildOption(c)
	}
	r.WrapData = func(data []map[string]string) {
		r.wrapData(data)
	}
	return r
}

func (t *Ticket) buildTool(c *gin.Context) []*frm.Tool {
	user := t.GetUser(c)
	if user.IsSales || user.IsManager {
		return []*frm.Tool{{
			Title: "新增需求",
			Icon:  "pencil",
			URL:   "apply",
			Type:  "modal",
			Color: "primary",
		}}
	}
	return []*frm.Tool{}
}

func (t *Ticket) buildOption(c *gin.Context) []*frm.Option {
	user := t.GetUser(c)
	ans := []*frm.Option{}
	ans = append(ans, t.WrapOption([]any{
		"查看详情", "eye", "read", "modal", []string{"id"},
	}))
	if user.IsSales || user.IsManager {
		ans = append(ans, t.WrapOption([]any{
			"修改需求", "pencil", "apply", "modal", []string{"id"},
			[]string{"status", "EQ", "INIT"},
		}))
	}
	if user.IsManager {
		ans = append(ans, t.WrapOption([]any{
			"编排机器", "grip-vertical", "plan", "modal", []string{"id"},
			[]any{"status", "IN", []string{"INIT", "PLANNED"}},
		}))
	}
	if user.IsWorker || user.IsManager {
		ans = append(ans, t.WrapOption([]any{
			"准备机器", "list", "prepare", "modal", []string{"id"},
			[]any{"status", "IN", []string{"PLANNED", "PREPARED"}},
		}))
	}
	if user.IsManager {
		ans = append(ans, t.WrapOption([]any{
			"开始生产", "play", "run", "async", []string{"id"},
			[]string{"status", "EQ", "PREPARED"},
		}))
	}
	if user.IsStorekeeper || user.IsManager {
		ans = append(ans, t.WrapOption([]any{
			"库存信息", "box", "/prod/store/detail", "link",
			[]map[string]string{{"k": "store_id", "v": "id"}},
		}))
		ans = append(ans, t.WrapOption([]any{
			"良品入库", "download", "/prod/store/plus", "modal",
			[]map[string]string{{"k": "store_id", "v": "id"}},
			[]string{"status", "EQ", "RUNNING"},
		}))
		ans = append(ans, t.WrapOption([]any{
			"劣品上报", "warning", "/prod/store/reject", "modal",
			[]map[string]string{{"k": "store_id", "v": "id"}},
			[]string{"status", "EQ", "RUNNING"},
		}))
	}
	if user.IsManager {
		ans = append(ans, t.WrapOption([]any{
			"终止生产", "stop", "stop", "async", []string{"id"},
			[]string{"status", "EQ", "RUNNING"},
		}))
		ans = append(ans, t.WrapOption([]any{
			"完成生产", "check", "finish", "async", []string{"id"},
			[]string{"status", "EQ", "RUNNING"},
		}))
		ans = append(ans, t.WrapOption([]any{
			"修改工单", "edit", "edit", "modal", []string{"id"},
		}))
		ans = append(ans, t.WrapOption([]any{
			"删除工单", "trash", "delete", "async", []string{"id"},
		}))
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
		ss := strings.Split(d["progress"], "/")
		if ss[0] == ss[2] {
			ss[0] = "<b class='text-success'>" + ss[2] + "</b>"
		} else {
			ss[0] = "<b class='text-info'>" + ss[0] + "</b>"
		}
		ss[2] = "<b class='text-primary'>" + ss[2] + "</b>"
		if ss[1] != "0" {
			ss[1] = "<b class='text-danger'>" + ss[1] + "</b>"
		}
		d["progress"] = strings.Join(ss, "/")
	}
}
