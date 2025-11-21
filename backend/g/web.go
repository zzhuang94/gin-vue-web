package g

import (
	"strconv"

	"github.com/gin-gonic/gin"
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

func (w *Web) GetPageSize(c *gin.Context) int {
	ans, _ := strconv.Atoi(w.GetUser(c).User["page_size"])
	if ans == 0 {
		ans = 10
	}
	return ans
}

type Sess struct {
	*xorm.Session
	Ctx *gin.Context
}

func (w *Web) BeginSess(engine *xorm.Engine, c *gin.Context) *Sess {
	sess := engine.NewSession()
	sess.Begin()
	return &Sess{Session: sess, Ctx: c}
}
