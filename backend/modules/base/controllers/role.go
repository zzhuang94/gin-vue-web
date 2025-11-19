package controllers

import (
	"backend/g"
	"backend/modules/base/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Role struct {
	*g.XB[*models.Role]
}

func NewRole() *Role {
	r := &Role{XB: g.NewXB(&models.Role{})}
	r.Option = append([][]any{{"特权管理", "folder-tree", "access"}}, r.Option...)
	r.WrapData = func(data []map[string]string) { r.wrapData(data) }
	return r
}

func (r *Role) wrapData(data []map[string]string) {
	for _, d := range data {
		d["icon"] = fmt.Sprintf("<i class='fa fa-%s'></i> %s", d["icon"], d["icon"])
	}
}

func (r *Role) ActionAccess(c *gin.Context) {
	id := c.Query("id")
	role := new(models.Role)
	has, err := r.DB.ID(id).Get(role)
	if err != nil || !has {
		r.JsonFail(c, fmt.Errorf("数据不存在"))
		return
	}
	props := map[string]any{
		"role": role,
		"data": r.buildTree(),
		"ids":  role.GetActionIds(),
	}
	r.Modal(c, props)
}

func (r *Role) buildTree() map[string]map[string]any {
	ans := map[string]map[string]any{"root": {
		"name":     "根节点",
		"icon":     "cog",
		"size":     "1.4",
		"expanded": true,
	}}
	rows := []*models.Action{}
	r.DB.Find(&rows)
	for _, r := range rows {
		ss := strings.Split(r.Path, "/")
		m, c, a := ss[1], ss[2], ss[3]
		ans["m-"+m] = map[string]any{
			"parent_id": "root",
			"name":      m,
			"icon":      "book",
			"size":      "1.3",
		}
		ans["c-"+c] = map[string]any{
			"parent_id": "m-" + m,
			"name":      c,
			"icon":      "bookmark",
			"size":      "1.15",
		}
		id := strconv.Itoa(r.ID)
		icon := "circle"
		color := "success"
		if r.Green == 0 {
			icon = "circle-dot"
			color = ""
		}
		ans[id] = map[string]any{
			"parent_id": "c-" + c,
			"name":      a,
			"icon":      icon,
			"color":     color,
		}
	}
	return ans
}

func (r *Role) ActionAccessSave(c *gin.Context) {
	type arg struct {
		Id  string   `json:"id"`
		Ids []string `json:"ids"`
	}
	a := new(arg)
	if err := c.ShouldBindJSON(a); err != nil {
		r.JsonFail(c, err)
		return
	}
	sess := r.BeginSess(r.DB)
	olds := []*models.RoleAction{}
	news := []*models.Action{}
	sess.Where("role_id = ?", a.Id).Find(&olds)
	sess.In("id", a.Ids).Find(&news)
	oldMap := make(map[int]bool)
	newMap := make(map[int]bool)
	for _, o := range olds {
		oldMap[o.ActionId] = true
	}
	for _, n := range news {
		logrus.Infof("new %d", n.ID)
		if !oldMap[n.ID] {
			ra := new(models.RoleAction)
			ra.RoleId, _ = strconv.Atoi(a.Id)
			ra.ActionId = n.ID
			ra.Save(sess)
		}
		newMap[n.ID] = true
	}
	for _, o := range olds {
		if !newMap[o.ActionId] {
			o.Delete(sess)
		}
	}
	sess.Commit()
	r.JsonSucc(c, "保存成功")
}
