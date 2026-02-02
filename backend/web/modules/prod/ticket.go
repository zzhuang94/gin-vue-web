package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
	"fmt"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Ticket struct {
	*frm.X
}

func NewTicket() *Ticket {
	r := &Ticket{X: frm.NewX(&prod.Ticket{})}
	r.DB = g.CoreDB
	r.Tool = []*frm.Tool{{
		Title: "新增生产需求",
		Icon:  "pencil",
		URL:   "apply",
		Type:  "modal",
		Color: "primary",
	}}
	r.Option = [][]any{
		{"修改生产需求", "pencil", "apply", "modal", []string{"id"}},
		{"编排机器", "grip-vertical", "plan", "modal", []string{"id"}},
		{"准备机器", "list", "prepare", "modal", []string{"id"}},
		{"开始生产", "play", "run", "async", []string{"id"}},
		{"更新收货", "truck", "receive", "modal", []string{"id"}},
	}
	return r
}

func (t *Ticket) ActionApply(c *gin.Context) {
	rules := t.GetRules()
	rules = t.filterRules(rules, []string{
		"title", "sku", "material", "color", "quantity_expected", "lead_time", "image", "remark",
	})
	title := "<i class='fa fa-pencil'></i>&nbsp;&nbsp;生产需求"
	t.editModel(c, title, rules)
}

func (t *Ticket) ActionPlan(c *gin.Context) {
	rules := t.GetRules()
	rules = t.filterRules(rules, []string{
		"title", "sku", "material", "color", "quantity_expected", "lead_time", "image",
		"apply_user", "apply_time", "machine_plan",
		"remark",
	})
	rules = t.readonlyRules(rules, []string{
		"title", "sku", "material", "color", "quantity_expected", "lead_time", "image",
	})
	title := "<i class='fa fa-grip-vertical'></i>&nbsp;&nbsp;编排机器"
	t.editModel(c, title, rules)
}

func (t *Ticket) ActionPrepare(c *gin.Context) {
	rules := t.GetRules()
	rules = t.filterRules(rules, []string{
		"title", "plan_user", "plan_time", "machine_plan", "machine_list", "remark",
	})
	rules = t.readonlyRules(rules, []string{
		"title", "machine_plan",
	})
	title := "<i class='fa fa-list'></i>&nbsp;&nbsp;准备机器"
	t.editModel(c, title, rules)
}

func (t *Ticket) ActionReceive(c *gin.Context) {
	rules := t.GetRules()
	rules = t.filterRules(rules, []string{
		"title", "lead_time", "quantity_expected", "quantity_actual",
		"run_user", "run_time",
		"remark",
	})
	rules = t.readonlyRules(rules, []string{
		"title", "lead_time", "quantity_expected",
	})
	title := "<i class='fa fa-truck'></i>&nbsp;&nbsp;更新收货"
	t.editModel(c, title, rules)
}

func (t *Ticket) ActionRun(c *gin.Context) {
	t.JsonSucc(c, "开始生产")
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

func (t *Ticket) filterRules(rules []*g.Rule, keys []string) []*g.Rule {
	ans := make([]*g.Rule, 0)
	for _, r := range rules {
		if slices.Contains(keys, r.Key) {
			ans = append(ans, r)
		}
	}
	return ans
}

func (t *Ticket) unsetRules(rules []*g.Rule, keys []string) []*g.Rule {
	ans := make([]*g.Rule, 0)
	for _, r := range rules {
		if !slices.Contains(keys, r.Key) {
			ans = append(ans, r)
		}
	}
	return ans
}

func (t *Ticket) readonlyRules(rules []*g.Rule, keys []string) []*g.Rule {
	ans := make([]*g.Rule, 0)
	for _, r := range rules {
		if slices.Contains(keys, r.Key) {
			r.Readonly = true
			logrus.Infof("readonly rule: %s", r.Key)
		}
		ans = append(ans, r)
	}
	return ans
}
