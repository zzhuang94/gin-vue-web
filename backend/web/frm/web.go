package frm

import (
	"backend/g"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

type Web struct {
}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Render(c *gin.Context) {
	w.RenderDataPage(c, nil, "")
}

func (w *Web) RenderData(c *gin.Context, data any) {
	w.RenderDataPage(c, data, "")
}

func (w *Web) RenderDataPage(c *gin.Context, data any, page string) {
	path := w.GetPath(c)
	if page == "" {
		page = "modules" + path
	}
	c.JSON(200, gin.H{
		"page":   page,
		"data":   data,
		"layout": w.GetLayout(c),
	})
}

func (w *Web) RenderPage(c *gin.Context, data any, page string) {
	c.JSON(200, gin.H{
		"page": page,
		"data": data,
	})
}

func (w *Web) Modal(c *gin.Context, props any) {
	w.ModalPage(c, props, "modules"+w.GetPath(c))
}

func (w *Web) ModalPage(c *gin.Context, props any, page string) {
	c.JSON(200, gin.H{
		"props": props,
		"page":  page,
	})
}

func (w *Web) JsonFail(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": err.Error(),
	})
}

func (w *Web) JsonSucc(c *gin.Context, data any) {
	c.Set("op_ok", true)

	c.JSON(200, gin.H{
		"code": 1,
		"data": data,
	})
}

func (w *Web) GetPath(c *gin.Context) string {
	return c.GetString("path")
}

func (w *Web) GetUsername(c *gin.Context) string {
	return c.GetString("username")
}

func (w *Web) GetUser(c *gin.Context) *User {
	user, ok := c.Get("user")
	if ok {
		return user.(*User)
	}
	return buildUser(w.GetUsername(c))
}

func (w *Web) GetUriArg(c *gin.Context) map[string]any {
	arg := map[string]any{}
	for k, vs := range c.Request.URL.Query() {
		if len(vs) == 1 {
			arg[k] = vs[0]
		} else {
			arg[k] = vs
		}
	}
	return arg
}

func (w *Web) GetIds(c *gin.Context) ([]string, error) {
	arg := new(struct {
		Ids []string `json:"ids"`
	})
	if err := c.ShouldBindJSON(arg); err != nil {
		return nil, err
	}
	return arg.Ids, nil
}

func (w *Web) GetPageSize(c *gin.Context) int {
	ans, _ := strconv.Atoi(w.GetUser(c).User["page_size"])
	if ans == 0 {
		ans = 10
	}
	return ans
}

func (w *Web) BeginSess(engine *xorm.Engine, c *gin.Context) *g.Sess {
	sess := engine.NewSession()
	sess.Begin()
	return &g.Sess{Session: sess, Ctx: c}
}

func (w *Web) RulesFilter(rules []*g.Rule, keys []string) []*g.Rule {
	ans := make([]*g.Rule, 0)
	for _, r := range rules {
		if slices.Contains(keys, r.Key) {
			ans = append(ans, r)
		}
	}
	return ans
}

func (w *Web) RulesUnset(rules []*g.Rule, keys []string) []*g.Rule {
	ans := make([]*g.Rule, 0)
	for _, r := range rules {
		if !slices.Contains(keys, r.Key) {
			ans = append(ans, r)
		}
	}
	return ans
}

func (w *Web) RulesReadonly(rules []*g.Rule, keys []string) []*g.Rule {
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
