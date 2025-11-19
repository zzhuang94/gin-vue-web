package g

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Rule struct {
	Key       string            `json:"key"`                  // 数据表字段名
	Name      string            `json:"name"`                 // 显示中文名称
	Default   string            `json:"default,omitempty"`    // 默认值
	Readonly  bool              `json:"readonly,omitempty"`   // 是否只读
	Required  bool              `json:"required,omitempty"`   // 是否必填
	Describe  string            `json:"describe,omitempty"`   // 字段描述，一般在form表单时提示
	Textarea  bool              `json:"textarea,omitempty"`   // 是否多行文本
	Json      bool              `json:"json,omitempty"`       // 是否JSON
	Bold      bool              `json:"bold,omitempty"`       // 是否加粗展示
	SplitSep  string            `json:"split_sep,omitempty"`  // 分割符
	Limit     []*Limit          `json:"limit,omitempty"`      // 下拉选项
	LimitList []string          `json:"limit_list,omitempty"` // 下拉选项列表
	LimitMap  map[string]*Limit `json:"limit_map,omitempty"`  // 下拉选项映射
	Search    int               `json:"search,omitempty"`     // 搜索匹配方式 0: none, 1: eq, 2: like, 3: in
	NoSort    bool              `json:"no_sort,omitempty"`    // 是否不排序
	Hide      bool              `json:"hide,omitempty"`       // 是否在table中隐藏
	Width     string            `json:"width,omitempty"`      // 宽度
	Trans     *Trans            `json:"trans,omitempty"`      // 转译，自动将 外键id转换为外键值
}

type Limit struct {
	Key   string `json:"key"`
	Label string `json:"label"`
	Badge string `json:"badge,omitempty"`
}

type Trans struct {
	Ajax  bool   `json:"ajax,omitempty"`
	DB    string `json:"db,omitempty"`
	SQL   string `json:"sql,omitempty"`
	Table string `json:"table,omitempty"`
	Key   string `json:"key"`
	Val   string `json:"val"`
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
		Key:      r.Key,
		Name:     r.Name,
		Default:  r.Default,
		Readonly: r.Readonly,
		Required: r.Required,
		Describe: r.Describe,
		Textarea: r.Textarea,
		Json:     r.Json,
		Bold:     r.Bold,
		SplitSep: r.SplitSep,
		Search:   r.Search,
		NoSort:   r.NoSort,
		Hide:     r.Hide,
		Width:    r.Width,
		Trans:    r.Trans,
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
