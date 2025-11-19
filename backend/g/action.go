package g

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type reg struct {
	module     string
	controller string
	instance   any
}

var regs []reg

func RegController(module, controller string, instance any) {
	regs = append(regs, reg{module, controller, instance})
}

func BindActions(rg *gin.RouterGroup) {
	paths := []string{}
	for _, r := range regs {
		paths = append(paths, bindReg(rg, r)...)
	}
	updateActions(paths)
}

func updateActions(paths []string) error {
	var adds, dels []string
	news := make(map[string]string)
	for _, path := range paths {
		if _, ok := news[path]; ok {
			return fmt.Errorf("path %s is repeated", path)
		}
		news[path] = path
	}
	olds := make(map[string]string)
	sql := "SELECT id, path FROM action"
	rows, _ := BaseDB.SQL(sql).QueryString()
	for _, r := range rows {
		olds[r["path"]] = r["id"]
		if _, ok := news[r["path"]]; !ok {
			dels = append(dels, r["id"])
		}
	}
	for path := range news {
		if _, ok := olds[path]; !ok {
			adds = append(adds, fmt.Sprintf("('%s')", path))
		}
	}
	if len(adds) > 0 {
		sql = "INSERT INTO action (path) VALUES " + strings.Join(adds, ",")
		BaseDB.Exec(sql)
	}
	if len(dels) > 0 {
		sql = fmt.Sprintf("DELETE FROM action WHERE id IN (%s)", strings.Join(dels, ","))
		BaseDB.Exec(sql)
	}
	return nil
}

func bindReg(rg *gin.RouterGroup, r reg) []string {
	paths := []string{}
	val := reflect.ValueOf(r.instance)
	typ := val.Type()
	for i := 0; i < typ.NumMethod(); i++ {
		m := typ.Method(i)
		if !strings.HasPrefix(m.Name, "Action") {
			continue
		}
		if m.Type.NumIn() != 2 || m.Type.In(1) != reflect.TypeOf((*gin.Context)(nil)) {
			continue
		}
		action := toKebab(strings.TrimPrefix(m.Name, "Action"))
		if action == "" {
			continue
		}

		path := "/" + r.module + "/" + r.controller + "/" + action
		rg.POST(path, makeHandler(val.Method(i)))
		paths = append(paths, path)
	}
	return paths
}

func makeHandler(method reflect.Value) gin.HandlerFunc {
	if fn, ok := method.Interface().(func(*gin.Context)); ok {
		return func(c *gin.Context) {
			defer func() {
				if r := recover(); r != nil {
					logrus.Errorf("panic: %v\n%s", r, debug.Stack())
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}()
			fn(c)
		}
	}

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("panic: %v\n%s", r, debug.Stack())
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		method.Call([]reflect.Value{reflect.ValueOf(c)})
	}
}

var kebabRegex = regexp.MustCompile("[A-Z][^A-Z]*")

func toKebab(s string) string {
	parts := kebabRegex.FindAllString(s, -1)
	for i := range parts {
		parts[i] = strings.ToLower(parts[i])
	}
	return strings.Join(parts, "-")
}
