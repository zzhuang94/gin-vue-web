package g

import (
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func WebAuth(c *gin.Context) {
	path := strings.TrimPrefix(c.Request.URL.Path, "/web")
	c.Set("path", path)

	if path == "/base/user/join" || path == "/base/user/sign-up" || path == "/base/user/log-in" {
		c.Next()
		return
	}

	session := sessions.Default(c)
	username := session.Get("username")
	if username == nil {
		web := NewWeb()
		web.RenderDataPage(c, gin.H{"path": path}, "modules/base/user/join")
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

	// 记录用户访问日志
	sql := "INSERT INTO user_log (username, path) VALUES (?, ?)"
	BaseDB.Exec(sql, user.Name, path)

	c.Next()
}

func ApiAuth(c *gin.Context) {
	path := strings.TrimPrefix(c.Request.URL.Path, "/api")
	logrus.Infof("api middleware: %s", path)

	c.Next()
}
