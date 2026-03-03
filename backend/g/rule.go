package g

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Rule struct {
	Key  string `json:"key"`  // 数据表字段名
	Name string `json:"name"` // 显示中文名称

	Default  string `json:"default"`  // 默认值
	Readonly bool   `json:"readonly"` // 是否只读
	Required bool   `json:"required"` // 是否必填
	Describe string `json:"describe"` // 字段描述，一般在form表单时提示

	Textarea   bool        `json:"textarea"`   // 是否多行文本
	Json       bool        `json:"json"`       // 是否JSON
	Date       bool        `json:"date"`       // 是否日期
	Datetime   bool        `json:"datetime"`   // 是否日期时间
	Trans      *Trans      `json:"trans"`      // 转译，自动将 外键id转换为外键值
	Validation *Validation `json:"validation"` // 验证规则

	Bold      bool   `json:"bold"`      // 是否加粗展示
	Textcolor string `json:"textcolor"` // 文本颜色 primary/success/warning/danger/info/metal
	Suffix    string `json:"suffix"`    // 后缀
	Prefix    string `json:"prefix"`    // 前缀

	SplitSep string `json:"split_sep"` // 分割符

	Limit     []*Limit          `json:"limit"`      // 下拉选项
	LimitList []string          `json:"limit_list"` // 下拉选项列表
	LimitMap  map[string]*Limit `json:"limit_map"`  // 下拉选项映射

	Search int `json:"search"` // 搜索匹配方式 0: 不支持搜索, 1: eq, 2: like, 3: in

	NoSort   bool   `json:"no_sort"`   // 是否不排序
	Hide     bool   `json:"hide"`      // 是否在table中隐藏
	OpUse    bool   `json:"op_use"`    // 是否在操作中使用
	AutoHide string `json:"auto_hide"` // 自动隐藏的屏幕宽度，l1/l2/l3/l4/l5
	Width    string `json:"width"`     // 宽度

}

type Limit struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Badge string `json:"badge"`
}

type Trans struct {
	Ajax  bool   `json:"ajax"`
	DB    string `json:"db"`
	SQL   string `json:"sql"`
	Table string `json:"table"`
	Key   string `json:"key"`
	Val   string `json:"val"`
}

type Validation struct {
	IsInt    bool `json:"is_int"`
	IntRange bool `json:"int_range"`
	IntMin   int  `json:"int_min"`
	IntMax   int  `json:"int_max"`

	IsFloat    bool    `json:"is_float"`
	FloatRange bool    `json:"float_range"`
	FloatMin   float64 `json:"float_min"`
	FloatMax   float64 `json:"float_max"`

	IsIP   bool `json:"is_ip"`
	IsIPv4 bool `json:"is_ipv4"`
	IsIPv6 bool `json:"is_ipv6"`

	Regex string `json:"regex"`
}

func initRules() error {
	bytes, err := os.ReadFile("./rule.json")
	if err != nil {
		return fmt.Errorf("read rule file failed: %v", err)
	}

	rules := make(map[string][]*Rule)
	if err = json.Unmarshal(bytes, &rules); err != nil {
		return fmt.Errorf("parse rule file failed: %v", err)
	}

	Rules = rules
	return nil
}

func (r *Rule) SelfWrap() *Rule {
	ans := &Rule{
		Key:        r.Key,
		Name:       r.Name,
		Default:    r.Default,
		Readonly:   r.Readonly,
		Required:   r.Required,
		Describe:   r.Describe,
		Textarea:   r.Textarea,
		Date:       r.Date,
		Datetime:   r.Datetime,
		Json:       r.Json,
		Bold:       r.Bold,
		Suffix:     r.Suffix,
		Prefix:     r.Prefix,
		SplitSep:   r.SplitSep,
		Search:     r.Search,
		NoSort:     r.NoSort,
		Hide:       r.Hide,
		OpUse:      r.OpUse,
		AutoHide:   r.AutoHide,
		Width:      r.Width,
		Trans:      r.Trans,
		Validation: r.Validation,
	}
	ans.Limit, ans.LimitMap = r.smartLimit()
	return ans
}

func (r *Rule) smartLimit() ([]*Limit, map[string]*Limit) {
	limits := make([]*Limit, 0)
	for i, limit := range r.Limit {
		badge := limit.Badge
		if badge == "" {
			badge = fmt.Sprintf("rand-%d", i%40)
		}
		limits = append(limits, &Limit{
			Key:   limit.Key,
			Label: limit.Label,
			Badge: badge,
		})
	}

	limits = append(limits, r.getTrans()...)

	for i, limit := range r.LimitList {
		limits = append(limits, &Limit{
			Key:   limit,
			Label: limit,
			Badge: fmt.Sprintf("rand-%d", i%40),
		})
	}

	lm := make(map[string]*Limit)
	for _, limit := range limits {
		lm[limit.Key] = limit
	}
	if len(lm) == 0 {
		return nil, nil
	}
	return limits, lm
}

func (r *Rule) getTrans() []*Limit {
	ans := make([]*Limit, 0)
	if r.Trans == nil || r.Trans.Ajax {
		return ans
	}
	db := BaseDB
	if r.Trans.DB == "core" {
		db = CoreDB
	}
	var sql string
	if r.Trans.SQL == "" {
		sql = fmt.Sprintf("SELECT %s, %s FROM %s", r.Trans.Key, r.Trans.Val, r.Trans.Table)
	} else {
		sql = r.Trans.SQL
	}
	rows, err := db.SQL(sql).QueryString()
	if err != nil {
		logrus.Error(err)
		return ans
	}
	for _, row := range rows {
		ans = append(ans, &Limit{
			Key:   row[r.Trans.Key],
			Label: row[r.Trans.Val],
		})
	}
	return ans
}

func (r *Rule) Translate(data []map[string]string) []map[string]string {
	keys := []string{}
	for _, d := range data {
		keys = append(keys, d[r.Key])
	}
	db := BaseDB
	if r.Trans.DB == "core" {
		db = CoreDB
	}
	sql := r.Trans.SQL
	if sql == "" {
		sql = fmt.Sprintf(
			"SELECT `%s`, `%s` FROM `%s` WHERE `%s` IN (%s)",
			r.Trans.Key, r.Trans.Val, r.Trans.Table, r.Trans.Key,
			"'"+strings.Join(keys, "','")+"'",
		)
	}
	rows, err := db.SQL(sql).QueryString()
	if err != nil {
		logrus.Error(err)
		return data
	}
	lm := make(map[string]string)
	for _, row := range rows {
		lm[row[r.Trans.Key]] = row[r.Trans.Val]
	}
	for _, d := range data {
		if v, ok := lm[d[r.Key]]; ok {
			d[r.Key] = v
		}
	}
	return data
}

func (r *Rule) Transone(d map[string]string) {
	r.Translate([]map[string]string{d})
}
