package frm

import (
	"backend/g"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zzhuang94/go-kit/str"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type Menu struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	URL   string `json:"url"`
	Type  string `json:"type"`
	Color string `json:"color"`
}

type TableMenu struct {
	*Menu
	Alone bool        `json:"alone"`
	Args  []*MenuArg  `json:"args"`
	Conds []*MenuCond `json:"conds"`
}

type MenuArg struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type MenuCond struct {
	Key  string `json:"key"`
	Val  any    `json:"val"`
	Comp string `json:"comp"` // EQ, NQ, LT, LE, GT, GE, IN, NIN
}
type X struct {
	*Web
	DB             *xorm.Engine
	Model          g.ModelX
	Rules          []*g.Rule
	NoID           bool
	WrapTime       bool
	AndWheres      []map[string]any
	Header         string
	BatchEdit      bool
	BatchDelete    bool
	Dump           bool
	TopMenu        [][]string
	TableMenu      [][]string
	BuildTopMenu   func(*gin.Context) []*Menu
	BuildTableMenu func(*gin.Context) []*TableMenu
	WrapData       func([]map[string]string)
	BuildQuery     func(cond builder.Cond, sele bool) *xorm.Session
}

func NewX(m g.ModelX) *X {
	tableName := m.TableName()
	rules, ok := g.Rules[tableName]
	if !ok {
		logrus.Warnf("rule not found for table: %s", tableName)
	}

	x := &X{
		Web:      NewWeb(),
		DB:       g.BaseDB,
		Model:    m,
		Rules:    rules,
		WrapTime: true,
		TopMenu:  [][]string{{"新 增", "plus", "edit", "modal", "primary"}},
		TableMenu: [][]string{
			{"编 辑", "edit", "edit"},
			{"删 除", "trash", "delete", "async"},
		},
	}

	x.BuildTopMenu = func(c *gin.Context) []*Menu {
		return x.BuildTopMenuX(c)
	}
	x.BuildTableMenu = func(c *gin.Context) []*TableMenu {
		return x.BuildTableMenuX(c)
	}
	x.BuildQuery = func(cond builder.Cond, withSelect bool) *xorm.Session {
		return x.BuildQueryX(cond, withSelect)
	}
	x.WrapData = func(data []map[string]string) {
		x.WrapDataX(data)
	}

	return x
}

func (x *X) GetRules() []*g.Rule {
	ans := make([]*g.Rule, 0)
	for _, r := range x.Rules {
		ans = append(ans, r.SelfWrap())
	}
	return ans
}

func (x *X) getTableRules() []*g.Rule {
	ans := make([]*g.Rule, 0)
	if !x.NoID {
		ans = append(ans, &g.Rule{
			Key:  "id",
			Name: "ID",
		})
	}
	for _, r := range x.Rules {
		if !r.Hide {
			ans = append(ans, r.SelfWrap())
		}
	}
	if x.WrapTime {
		ans = append(ans, x.getTimeRule("created"))
		ans = append(ans, x.getTimeRule("updated"))
	}
	return ans
}

func (x *X) getRuleByKey(key string) *g.Rule {
	if key == "created" || key == "updated" {
		return x.getTimeRule(key)
	}
	for _, r := range x.Rules {
		if r.Key == key {
			return r
		}
	}
	return nil
}

func (x *X) getTimeRule(key string) *g.Rule {
	name := ""
	switch key {
	case "created":
		name = "创建时间"
	case "updated":
		name = "更新时间"
	}
	return &g.Rule{
		Key:      key,
		Name:     name,
		Search:   2,
		AutoHide: "l1",
	}
}

func (x *X) BuildTopMenuX(c *gin.Context) []*Menu {
	ans := make([]*Menu, 0)
	for _, r := range x.TopMenu {
		ans = append(ans, x.WrapMenu(r))
	}
	if x.BatchEdit {
		ans = append(ans, &Menu{"批量修改", "edit", "", "batch-edit", "warning"})
	}
	if x.BatchDelete {
		ans = append(ans, &Menu{"批量删除", "trash", "", "batch-delete", "danger"})
	}
	return ans
}

func (x *X) BuildTableMenuX(c *gin.Context) []*TableMenu {
	ans := make([]*TableMenu, 0)
	for _, r := range x.TableMenu {
		ans = append(ans, x.WrapTableMenu(r))
	}
	return ans
}

func (x *X) WrapMenu(r []string) *Menu {
	menu := &Menu{
		Title: r[0],
		Icon:  r[1],
		URL:   r[2],
	}
	if len(r) > 3 {
		menu.Type = r[3]
	} else {
		menu.Type = "modal"
	}
	if len(r) > 4 {
		menu.Color = r[4]
	} else {
		menu.Color = "primary"
	}
	return menu
}

func (x *X) WrapTableMenu(r []string) *TableMenu {
	return &TableMenu{
		Menu: x.WrapMenu(r),
		Args: []*MenuArg{{Key: "id", Val: "id"}},
	}
}

func (x *X) initSort(c *gin.Context) map[string]string {
	return map[string]string{
		"key":   c.DefaultQuery("sort_key", "id"),
		"order": c.DefaultQuery("sort_order", "DESC"),
	}
}

func (x *X) ActionIndex(c *gin.Context) {
	data := gin.H{
		"header":     x.Header,
		"batch":      x.BatchEdit || x.BatchDelete,
		"dump":       x.Dump,
		"rules":      x.getTableRules(),
		"topMenus":   x.BuildTopMenu(c),
		"tableMenus": x.BuildTableMenu(c),
		"sort":       x.initSort(c),
		"arg":        x.GetUriArg(c),
		"pageSize":   x.GetPageSize(c),
	}
	x.RenderDataPage(c, data, "templates/index")
}

func (x *X) ActionFetch(c *gin.Context) {
	var params struct {
		Arg  map[string]any    `json:"arg"`
		Sort map[string]string `json:"sort"`
		Page struct {
			Curr int `json:"curr"`
			Size int `json:"size"`
		} `json:"page"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		x.JsonFail(c, err)
		return
	}

	cond := x.buildCondition(params.Arg)
	query := x.BuildQuery(cond, false)

	total, err := query.Count()
	if err != nil {
		x.JsonFail(c, err)
		return
	}

	page := params.Page
	if total <= int64((page.Curr-1)*page.Size) {
		page.Curr = 1
	}

	query = x.BuildQuery(cond, true)

	if params.Sort != nil && params.Sort["key"] != "" {
		query = query.OrderBy(params.Sort["key"] + " " + params.Sort["order"])
	}

	query = query.Limit(page.Size, (page.Curr-1)*page.Size)

	data := make([]map[string]string, 0)
	err = query.Find(&data)
	if err != nil {
		x.JsonFail(c, err)
		return
	}

	x.WrapData(data)

	c.JSON(200, gin.H{
		"page": gin.H{
			"total": total,
			"curr":  page.Curr,
			"size":  page.Size,
		},
		"data": data,
	})
}

func (x *X) BuildQueryX(cond builder.Cond, withSelect bool) *xorm.Session {
	ans := x.DB.Table(x.Model.TableName()).Where(cond)
	if !withSelect {
		return ans
	}

	fields := []string{"`id`"}
	for _, rule := range x.Rules {
		if rule.OpUse || (!rule.Hide && rule.Key != "") {
			fields = append(fields, "`"+rule.Key+"`")
		}
	}
	if x.WrapTime {
		fields = append(fields, "DATE_FORMAT(`created`, '%Y-%m-%d %H:%i:%s') as `created`")
		fields = append(fields, "DATE_FORMAT(`updated`, '%Y-%m-%d %H:%i:%s') as `updated`")
	}
	ans = ans.Select(strings.Join(fields, ", "))
	return ans
}

func (x *X) buildCondition(arg map[string]any) builder.Cond {
	cond := builder.NewCond()

	for _, where := range x.AndWheres {
		for k, v := range where {
			cond = cond.And(builder.Eq{k: v})
		}
	}

	for k, v := range arg {
		qv := fmt.Sprintf("%v", v)
		if strings.TrimSpace(qv) == "" {
			continue
		}
		r := x.getRuleByKey(k)
		if r == nil {
			continue
		}
		switch r.Search {
		case 1:
			cond = cond.And(builder.Eq{k: qv})
		case 2:
			cond = cond.And(builder.Like{k, qv})
		case 3:
			if vs, ok := v.(string); ok {
				cond = cond.And(builder.Eq{k: vs})
				continue
			}
			vvs := []string{}
			for _, vs := range v.([]any) {
				vvs = append(vvs, vs.(string))
			}
			cond = cond.And(builder.Eq{k: vvs})
		}
	}
	return cond
}

func (x *X) WrapDataX(data []map[string]string) {
	if len(data) == 0 {
		return
	}
	for _, r := range x.Rules {
		if r.Trans != nil && r.Trans.Ajax {
			r.Translate(data)
		}
	}
}

func (x *X) ActionEdit(c *gin.Context) {
	props := gin.H{
		"action": "save",
		"title":  `<i class="fa fa-plus"></i>&nbsp;&nbsp;新 增`,
		"rules":  x.GetRules(),
	}
	id := c.DefaultQuery("id", "")
	if id != "" {
		m := x.Model.New()
		has, err := x.DB.ID(id).Get(m)
		if err != nil || !has {
			x.JsonFail(c, fmt.Errorf("数据不存在"))
			return
		}
		props["data"] = m
		props["title"] = `<i class="fa fa-edit"></i>&nbsp;&nbsp;编 辑`
		props["action"] = "save?id=" + id
	}
	x.ModalPage(c, props, "components/edit")
}

func (x *X) ActionSave(c *gin.Context) {
	m := x.Model.New()
	id := c.DefaultQuery("id", "")
	if id != "" {
		has, err := x.DB.ID(id).Get(m)
		if !has || err != nil {
			x.JsonFail(c, fmt.Errorf("数据不存在"))
			return
		}
	}
	payload, _ := io.ReadAll(c.Request.Body)
	sess := x.BeginSess(x.DB, c)
	if err := x.saveModel(m, payload, sess); err != nil {
		sess.Rollback()
		x.JsonFail(c, err)
		return
	}
	sess.Commit()
	x.JsonSucc(c, "保存成功")
}

func (x *X) saveModel(m g.ModelX, payload []byte, sess *g.Sess) error {
	payload, err := x.parseCheckPayload(payload)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(payload, m); err != nil {
		return fmt.Errorf("参数错误: %v", err)
	}
	if err := m.Save(sess); err != nil {
		return fmt.Errorf("保存失败: %v", err)
	}
	return nil
}

func (x *X) parseCheckPayload(payload []byte) ([]byte, error) {
	args := make(map[string]string)
	json.Unmarshal(payload, &args)
	for _, r := range x.Rules {
		if _, ok := args[r.Key]; !ok {
			continue
		}
		if r.Required && args[r.Key] == "" {
			return nil, fmt.Errorf("参数[%s]不能为空", r.Name)
		}
		if r.Textarea && r.SplitSep != "" {
			lines := str.SplitLines(args[r.Key])
			for _, line := range lines {
				if strings.Contains(line, r.SplitSep) {
					return nil, fmt.Errorf("参数[%s]不能包含分隔符[%s]", r.Name, r.SplitSep)
				}
			}
			args[r.Key] = strings.Join(lines, r.SplitSep)
		}
		if r.Textarea && r.Json {
			v, err := str.ParseAndFormatJson(args[r.Key])
			if err != nil {
				return nil, fmt.Errorf("参数[%s]不是合法JSON: %v", r.Name, err)
			}
			args[r.Key] = v
		}
		if r.Validation != nil {
			err := x.checkValidation(r.Validation, r.Name, args[r.Key])
			if err != nil {
				return nil, err
			}
		}
	}
	return json.Marshal(args)
}

func (x *X) checkValidation(v *g.Validation, name, val string) error {
	if v.IsInt {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("参数 [%s: %s] 不是整数", name, val)
		}
		if v.IntRange && (valInt < v.IntMin || valInt > v.IntMax) {
			return fmt.Errorf("参数 [%s: %s] 不在范围[%d,%d]内", name, val, v.IntMin, v.IntMax)
		}
	}
	if v.IsFloat {
		valFloat, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return fmt.Errorf("参数 [%s: %s] 不是浮点数", name, val)
		}
		if v.FloatRange && (valFloat < v.FloatMin || valFloat > v.FloatMax) {
			return fmt.Errorf("参数 [%s: %s] 不在范围[%f,%f]内", name, val, v.FloatMin, v.FloatMax)
		}
	}
	if v.IsIP {
		if !govalidator.IsIP(val) {
			return fmt.Errorf("参数 [%s: %s] 不是合法IP", name, val)
		}
	}
	if v.IsIPv4 {
		if !govalidator.IsIPv4(val) {
			return fmt.Errorf("参数 [%s: %s] 不是合法IPv4", name, val)
		}
	}
	if v.IsIPv6 {
		if !govalidator.IsIPv6(val) {
			return fmt.Errorf("参数 [%s: %s] 不是合法IPv6", name, val)
		}
	}
	if v.Regex != "" {
		if !govalidator.Matches(val, v.Regex) {
			return fmt.Errorf("参数 [%s: %s] 不符合正则表达式%s", name, val, v.Regex)
		}
	}
	return nil
}

func (x *X) ActionDelete(c *gin.Context) {
	id := c.Query("id")
	m := x.Model.New()
	has, err := x.DB.ID(id).Get(m)
	if err != nil || !has {
		x.JsonFail(c, fmt.Errorf("数据不存在"))
		return
	}
	sess := x.BeginSess(x.DB, c)
	err = m.Delete(sess)
	if err != nil {
		sess.Rollback()
		x.JsonFail(c, err)
		return
	}
	sess.Commit()
	x.JsonSucc(c, "删除成功")
}

func (x *X) List(c *gin.Context, args map[string]string, title, width, order string, edit bool) {
	cond := builder.NewCond()
	for k, v := range args {
		cond = cond.And(builder.Eq{k: v})
	}
	query := x.DB.Table(x.Model.TableName()).Where(cond)
	if order != "" {
		query = query.OrderBy(order)
	}
	if width == "" {
		width = "80%"
	}
	data := make([]map[string]string, 0)
	query.Find(&data)
	path := x.GetPath(c)
	i := strings.LastIndex(path, "/")
	prefix := path[:i+1]
	props := gin.H{
		"title": title,
		"width": width,
		"args":  args,
		"data":  data,
		"rules": x.GetListRules(args),
		"action": map[string]any{
			"prefix": prefix,
			"add":    true,
			"del":    true,
			"edit":   edit,
			"sort":   order != "",
		},
	}
	x.ModalPage(c, props, "components/list")
}

func (x *X) GetListRules(args map[string]string) []*g.Rule {
	ans := make([]*g.Rule, 0)
	for _, r := range x.Rules {
		if _, ok := args[r.Key]; ok {
			continue
		}
		ans = append(ans, r.SelfWrap())
	}
	return ans
}

func (x *X) BatchAddModal(c *gin.Context) {
	x.ModalPage(c, gin.H{}, "components/batch-add")
}

// 注意，这种批量添加使用原生SQL，不会自动记录操作日志
func (x *X) BatchAdd(c *gin.Context) {
	arg := &struct {
		Names []string `json:"names"`
	}{}
	if err := c.ShouldBindJSON(arg); err != nil {
		x.JsonFail(c, err)
		return
	}
	sql := "SELECT name FROM " + x.Model.TableName()
	rows, err := x.DB.QueryString(sql)
	if err != nil {
		x.JsonFail(c, err)
		return
	}
	olds := make(map[string]bool)
	for _, row := range rows {
		olds[strings.ToLower(row["name"])] = true
	}
	news := []string{}
	for _, name := range arg.Names {
		if _, ok := olds[strings.ToLower(name)]; !ok {
			news = append(news, fmt.Sprintf("('%s')", name))
		}
	}
	if len(news) == 0 {
		x.JsonSucc(c, "没有需要新增的数据")
		return
	}
	sql = "INSERT INTO " + x.Model.TableName() + " (name) VALUES " + strings.Join(news, ",")
	if _, err = x.DB.Exec(sql); err != nil {
		x.JsonFail(c, err)
		return
	}
	x.JsonSucc(c, fmt.Sprintf("批量新增成功 %d 条", len(news)))
}
