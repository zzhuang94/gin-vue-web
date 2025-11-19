package g

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type node struct {
	Id     string  `json:"-"`
	Name   string  `json:"name"`
	Path   string  `json:"path,omitempty"`
	Icon   string  `json:"icon,omitempty"`
	Active bool    `json:"active,omitempty"`
	Subs   []*node `json:"subs,omitempty"`
}

func (w *Web) GetLayout(c *gin.Context) map[string]any {
	path := w.GetPath(c)
	user := w.GetUser(c)

	var l1, l2 []*node
	tmp := make(map[string]*node)
	subs1, subs2 := make(map[string][]*node), make(map[string][]*node)
	paths := make(map[string]string)
	actives := make(map[string]bool)
	sql := `
SELECT DISTINCT
	t.id, t.level, t.parent_id, t.key_path, t.name, t.icon,
	t.action_id, a.path
FROM navtree t
JOIN navtree tp ON t.parent_id = tp.id
LEFT JOIN action a ON t.action_id = a.id
WHERE t.status = 1 AND t.level > 0 AND t.level < 4
ORDER BY t.level, tp.rank, t.rank, t.id`
	rows, _ := BaseDB.SQL(sql).QueryString()
	for _, r := range rows {
		id := r["id"]
		pid := r["parent_id"]
		tmp[id] = &node{Name: r["name"], Path: r["path"]}
		switch r["level"] {
		case "1":
			l1 = append(l1, &node{Id: id, Name: r["name"], Icon: r["icon"]})
		case "2":
			subs1[pid] = append(subs1[pid], &node{Id: id, Name: r["name"], Icon: r["icon"]})
		default:
			if r["path"] == "" || !user.IsAdmin && !user.AccPaths[r["path"]] {
				continue
			}
			pids := strings.Split(r["key_path"], "_")
			id1, id2 := pids[1], pids[2]
			if paths[id1] == "" {
				paths[id1] = r["path"]
			}
			if paths[id2] == "" {
				paths[id2] = r["path"]
			}
			subs2[id2] = append(subs2[id2], &node{Id: id, Name: r["name"], Path: r["path"]})
		}
	}

	title := []*node{}
	nt := getCurrNt(path)
	if nt != nil {
		pids := strings.Split(nt["key_path"], "_")
		pids = append(pids, nt["id"])
		for i := 1; i < 4; i++ {
			id := pids[i]
			actives[id] = true
			if t, ok := tmp[id]; ok {
				if t.Path == "" {
					t.Path = paths[id]
				}
				title = append(title, t)
			}
		}
		if nt["level"] > "3" {
			title = append(title, &node{Name: nt["name"], Path: path})
		}
		l2 = subs1[pids[1]]
	} else if len(l1) > 0 {
		l2 = subs1[l1[0].Id]
	}

	for _, n := range l1 {
		n.Active = actives[n.Id]
		n.Path = paths[n.Id]
	}
	for _, n := range l2 {
		n.Active = actives[n.Id]
		n.Path = paths[n.Id]
		ss := subs2[n.Id]
		if ss == nil {
			ss = []*node{}
		}
		for _, s := range ss {
			s.Active = actives[s.Id]
		}
		n.Subs = ss
	}

	return map[string]any{
		"l1":    l1,
		"l2":    l2,
		"title": title,
		"fold":  user.User["fold"] == "1",
		"env":   C.Env,
		"name":  C.Name,
		"user":  user.User,
	}
}

func getCurrNt(path string) map[string]string {
	sql := `
SELECT nt.* FROM navtree nt 
JOIN action a ON nt.action_id = a.id 
WHERE a.path = ? AND nt.status = 1 AND nt.level > 2`
	rows, err := BaseDB.SQL(sql, path).QueryString()
	if err != nil {
		logrus.Error(err)
	}
	if len(rows) == 0 {
		return nil
	}
	return rows[0]
}
