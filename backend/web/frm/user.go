package frm

import "backend/g"

type User struct {
	Name          string
	User          map[string]string
	IsAdmin       bool
	IsDba         bool
	IsSales       bool
	IsManager     bool
	IsWorker      bool
	IsStorekeeper bool
	AccPaths      map[string]bool
}

func buildUser(username string) *User {
	user := &User{
		Name:     username,
		User:     getUser(username),
		AccPaths: map[string]bool{},
	}
	user.loadRoles()
	if !user.IsAdmin {
		user.AccPaths = getAccPaths(username)
	}
	return user
}

func (user *User) loadRoles() {
	sql := `
SELECT r.name
FROM role r
JOIN role_user ru ON r.id = ru.role_id
WHERE ru.username = ? AND ru.status = 1`
	rows, _ := g.BaseDB.SQL(sql, user.Name).QueryString()
	roles := map[string]bool{}
	for _, r := range rows {
		roles[r["name"]] = true
	}
	user.IsAdmin = roles["admin"]
	user.IsDba = roles["dba"]
	user.IsSales = roles["销售"]
	user.IsManager = roles["生产管理"]
	user.IsWorker = roles["操作员"]
	user.IsStorekeeper = roles["仓库管理员"]
}

func getUser(username string) map[string]string {
	sql := `
SELECT username, email, cn_name, fold, page_size,
CASE WHEN avatar IS NOT NULL AND LENGTH(avatar) > 0 THEN '1' ELSE '0' END as has_avatar
FROM user WHERE username = ?`
	rows, _ := g.BaseDB.SQL(sql, username).QueryString()
	if len(rows) == 0 {
		return map[string]string{}
	}
	return rows[0]
}

func getAccPaths(username string) map[string]bool {
	sql := `
SELECT a.path
FROM action a
LEFT JOIN role_action ra ON a.id = ra.action_id
LEFT JOIN role_user ru ON ru.role_id = ra.role_id
WHERE a.green = 1 OR ru.username = ?`
	rows, _ := g.BaseDB.SQL(sql, username).QueryString()
	ans := map[string]bool{}
	for _, r := range rows {
		ans[r["path"]] = true
	}
	return ans
}
