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

type Tool struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	URL   string `json:"url"`
	Type  string `json:"type"`
	Color string `json:"color"`
}

type Option struct {
	Title string         `json:"title"`
	Icon  string         `json:"icon"`
	URL   string         `json:"url"`
	Type  string         `json:"type"`
	Args  []string       `json:"args"`
	Cond  map[string]any `json:"cond,omitempty"`
}

type X struct {
	*Web
	DB          *xorm.Engine
	Model       g.ModelX
	Rules       []*g.Rule
	NoID        bool
	WrapTime    bool
	HideFields  []string
	AndWheres   []map[string]any
	HeaderHint  string
	BatchEdit   bool
	BatchDelete bool
	Dump        bool
	Tool        []*Tool
	Option      [][]any
	WrapData    func([]map[string]string)
	BuildQuery  func(cond builder.Cond, sele bool) *xorm.Session
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
		Tool:     []*Tool{{"新 增", "plus", "edit", "modal", "primary"}},
		Option: [][]any{
			{"编 辑", "edit", "edit"},
			{"删 除", "trash", "delete", "async"},
		},
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
		ans = append(ans, &g.Rule{
			Key:  "created",
			Name: "创建时间",
		})
		ans = append(ans, &g.Rule{
			Key:  "updated",
			Name: "更新时间",
		})
	}
	return ans
}

func (x *X) getRuleByKey(key string) *g.Rule {
	for _, r := range x.Rules {
		if r.Key == key {
			return r
		}
	}
	return nil
}

func (x *X) buildTool() []*Tool {
	ans := make([]*Tool, 0)
	ans = append(ans, x.Tool...)
	if x.BatchEdit {
		ans = append(ans, &Tool{"批量修改", "edit", "", "batch-edit", "warning"})
	}
	if x.BatchDelete {
		ans = append(ans, &Tool{"批量删除", "trash", "", "batch-delete", "danger"})
	}
	return ans
}

func (x *X) buildOption() []*Option {
	ans := make([]*Option, 0)
	for _, r := range x.Option {
		ans = append(ans, x.wrapOption(r))
	}
	return ans
}

func (x *X) wrapOption(r []any) *Option {
	opt := &Option{
		Title: r[0].(string),
		Icon:  r[1].(string),
		URL:   r[2].(string),
	}
	if len(r) > 3 {
		opt.Type = r[3].(string)
	} else {
		opt.Type = "modal"
	}
	if len(r) > 4 {
		opt.Args = r[4].([]string)
	} else {
		opt.Args = []string{"id"}
	}
	if len(r) > 5 {
		opt.Cond = r[5].(map[string]any)
	}
	return opt
}

func (x *X) initSort(c *gin.Context) map[string]string {
	return map[string]string{
		"key":   c.DefaultQuery("sort_key", "id"),
		"order": c.DefaultQuery("sort_order", "DESC"),
	}
}

func (x *X) ActionIndex(c *gin.Context) {
	data := gin.H{
		"headerHint": x.HeaderHint,
		"batch":      x.BatchEdit || x.BatchDelete,
		"dump":       x.Dump,
		"tool":       x.buildTool(),
		"option":     x.buildOption(),
		"sort":       x.initSort(c),
		"rules":      x.getTableRules(),
		"arg":        x.GetUriArg(c),
		"page_size":  x.GetPageSize(c),
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

	fields := make([]string, 0)
	if !x.NoID {
		fields = append(fields, "`id`")
	}
	for _, rule := range x.Rules {
		if !rule.Hide && rule.Key != "" {
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
		"data":   gin.H{},
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
