package prod

import (
	"backend/g"
	"backend/models/prod"
	"backend/web/frm"
)

type Stats struct {
	*frm.X
}

func NewStats() *Stats {
	r := &Stats{X: frm.NewX(&prod.Ticket{})}
	r.DB = g.CoreDB
	r.Dump = true
	r.AndWheres = []map[string]any{{"status": []string{prod.StatusFinished, prod.StatusStopped}}}
	r.initRules()
	r.Tool = []*frm.Tool{}
	r.Option = [][]any{{"查看详情", "eye", "/prod/ticket/read", "modal", []string{"id"}}}
	return r
}

func (r *Stats) initRules() {
	r.Rules = r.RulesFilter(r.GetRules(), []string{
		"title", "lead_time", "progress",
		"apply_user", "plan_user", "prepare_user",
		"run_user", "remark", "status",
	})
	for _, r := range r.Rules {
		r.Hide = false
	}
}
