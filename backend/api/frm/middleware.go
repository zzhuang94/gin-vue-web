package frm

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Middleware(c *gin.Context) {
	path := strings.TrimPrefix(c.Request.URL.Path, "/api")
	logrus.Infof("api middleware: %s", path)

	c.Next()
}
