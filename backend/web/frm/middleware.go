package frm

import (
	"backend/g"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Middleware(c *gin.Context) {
	path := strings.TrimPrefix(c.Request.URL.Path, "/web")
	// 完整路径（包含查询参数），用于登录后跳转回原始请求
	fullPath := strings.TrimPrefix(c.Request.URL.RequestURI(), "/web")
	c.Set("path", path)
	c.Set("fullPath", fullPath)

	if path == "/base/user/join" || path == "/base/user/sign-up" || path == "/base/user/log-in" {
		c.Next()
		return
	}

	session := sessions.Default(c)
	username := session.Get("username")
	if username == nil {
		web := NewWeb()
		// 未登录时将完整请求路径传给登录页，登录成功后可带参数跳转回原页面
		web.RenderDataPage(c, gin.H{"path": fullPath}, "modules/base/user/join")
		c.Abort()
		return
	}

	user := buildUser(username.(string))
	if !user.IsAdmin && !user.AccPaths[path] {
		c.JSON(200, gin.H{
			"code": -1,
			"data": "抱歉，您无权限使用此功能",
		})
		c.Abort()
		return
	}

	c.Set("username", user.Name)
	c.Set("user", user)
	c.Set("op_uuid", uuid.New().String())

	g.RecordUserLog(user.Name, path)

	c.Next()

	g.RecordOp(c)
}
