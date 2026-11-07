package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Ticket struct {
	*frm.X
}

func NewTicket() *Ticket {
	r := &Ticket{X: frm.NewX(&prod.Ticket{})}
	r.DB = g.CoreDB
	r.NoID = true
	r.SearchID = true
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
	ans = append(ans, t.WrapTableMenu([]string{"打印工单", "print", "read?print=1"}))
	if user.IsSales || user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"修改需求", "pencil", "apply"}),
			Args:  []*frm.MenuArg{{Key: "id", Val: "id"}},
			Conds: []*frm.MenuCond{{Key: "status", Val: "INIT", Comp: "EQ"}},
		})
	}
	if user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"编排机器", "grip-vertical", "plan"}),
			Args:  []*frm.MenuArg{{Key: "id", Val: "id"}},
			Conds: []*frm.MenuCond{{Key: "status", Val: []string{"INIT", "PLANNED"}, Comp: "IN"}},
		})
	}
	if user.IsWorker || user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"准备机器", "list", "prepare"}),
			Args:  []*frm.MenuArg{{Key: "id", Val: "id"}},
			Conds: []*frm.MenuCond{{Key: "status", Val: []string{"PLANNED", "PREPARED"}, Comp: "IN"}},
		})
	}
	if user.IsManager {
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"开始生产", "play", "run", "async"}),
			Args:  []*frm.MenuArg{{Key: "id", Val: "id"}},
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
			Args:  []*frm.MenuArg{{Key: "id", Val: "id"}},
			Conds: []*frm.MenuCond{{Key: "status", Val: "RUNNING", Comp: "EQ"}},
		})
		ans = append(ans, &frm.TableMenu{
			Menu:  t.WrapMenu([]string{"完成生产", "check", "finish", "async"}),
			Args:  []*frm.MenuArg{{Key: "id", Val: "id"}},
			Conds: []*frm.MenuCond{{Key: "status", Val: "RUNNING", Comp: "EQ"}},
		})
		for _, action := range [][]string{
			{"修改工单", "edit", "edit"},
			{"删除工单", "trash", "delete", "async"},
		} {
			menu := t.WrapTableMenu(action)
			menu.Conds = []*frm.MenuCond{{
				Key: "status", Val: []string{prod.StatusFinished, prod.StatusStopped}, Comp: "NIN",
			}}
			ans = append(ans, menu)
		}
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
	if c.Query("print") == "1" {
		ticket, err := t.loadTicket(c.Query("id"))
		if err != nil {
			t.JsonFail(c, err)
			return
		}
		printFields := append(fields, "lead_time", "progress", "rate")
		t.ModalPage(c, gin.H{
			"ticket":           ticket,
			"rules":            t.RulesFilter(rules, printFields),
			"host":             g.C.Host,
			"closingOperation": t.closingOperation(ticket),
		}, "modules/prod/ticket/print")
		return
	}
	rules = t.RulesReadonly(t.RulesFilter(rules, fields), fields)
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
		ticket, err := t.loadTicket(id)
		if err != nil {
			t.JsonFail(c, err)
			return
		}
		props["data"] = ticket
		props["action"] = "save?id=" + id
	}
	t.ModalPage(c, props, "components/edit")
}

func (t *Ticket) loadTicket(id string) (*prod.Ticket, error) {
	ticketID, err := strconv.Atoi(id)
	if err != nil || ticketID <= 0 {
		return nil, fmt.Errorf("工单编号无效")
	}
	ticket := &prod.Ticket{}
	has, err := t.DB.ID(ticketID).Get(ticket)
	if err != nil || !has {
		return nil, fmt.Errorf("数据不存在")
	}
	return ticket, nil
}

func (t *Ticket) closingOperation(ticket *prod.Ticket) map[string]string {
	var path string
	switch ticket.Status {
	case prod.StatusFinished:
		path = "/prod/ticket/finish"
	case prod.StatusStopped:
		path = "/prod/ticket/stop"
	default:
		return nil
	}

	const query = `
SELECT e.user AS operator_name, DATE_FORMAT(l.created, '%Y-%m-%d %H:%i:%s') AS operation_time
FROM op_log l JOIN op_event e ON e.id = l.eid
WHERE l.data_table = 'ticket' AND l.data_id = ? AND l.op = 0 AND e.path = ?
  AND JSON_UNQUOTE(JSON_EXTRACT(l.data_new, '$.status')) = ?
  AND JSON_UNQUOTE(JSON_EXTRACT(l.data_old, '$.status')) <> ?
ORDER BY l.id DESC LIMIT 1`
	rows, err := g.BaseDB.SQL(query, ticket.Id, path, ticket.Status, ticket.Status).QueryString()
	if err != nil {
		logrus.WithError(err).Warn("查询工单结束记录失败")
		return nil
	}
	if len(rows) == 0 {
		return nil
	}
	return rows[0]
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
