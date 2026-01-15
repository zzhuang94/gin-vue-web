package res

import (
	"backend/g"
	"backend/models/res"
	"math"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type Ip struct {
	*g.XB[*res.Ip]
}

func NewIp() *Ip {
	r := &Ip{XB: g.NewXB(&res.Ip{})}
	r.Option = append([][]any{{"流量统计", "chart-line", "flow"}}, r.Option...)
	r.DB = g.CoreDB
	r.Dump = true
	return r
}

func (r *Ip) ActionDashboard(c *gin.Context) {
	r.Render(c)
}

func (r *Ip) ActionFlow(c *gin.Context) {
	id := c.Query("id")
	ip := new(res.Ip)
	r.DB.ID(id).Get(ip)
	option := gin.H{
		"tooltip": map[string]string{"trigger": "axis"},
		"grid": map[string]any{
			"left":         "10",
			"right":        "10",
			"containLabel": true,
		},
		"dataZoom": []gin.H{
			{
				"start":  80,
				"end":    100,
				"height": 30,
			},
			{
				"type": "inside",
			},
		},
		"xAxis": gin.H{
			"type": "category",
			"data": r.mockFlowX(),
		},
		"yAxis": gin.H{
			"type": "value",
		},
		"series": []gin.H{
			{
				"name":   "流量",
				"type":   "line",
				"smooth": true,
				"symbol": "none",
				"data":   r.mockFlowY(),
			},
		},
	}
	r.Modal(c, gin.H{"ip": ip.Ip, "option": option})
}

func (r *Ip) mockFlowX() []string {
	ans := make([]string, 0)
	now := time.Now()
	// 从7天前开始，对齐到整5分钟
	startTime := now.AddDate(0, 0, -7)
	// 对齐到最近的整5分钟（向下取整）
	startMinute := startTime.Minute() / 5 * 5
	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(),
		startTime.Hour(), startMinute, 0, 0, startTime.Location())

	// 每5分钟一个点，7天共 7*24*12 = 2016个点
	for i := 0; i < 7*24*12; i++ {
		t := startTime.Add(time.Duration(i*5) * time.Minute)
		// 格式化为完整时间戳：2006-01-02 15:04:05
		ans = append(ans, t.Format("2006-01-02 15:04:05"))
	}
	return ans
}

func (r *Ip) mockFlowY() []int {
	ans := make([]int, 0)
	// 7天，每5分钟一个点，共 7*24*12 = 2016个点
	pointsPerDay := 24 * 12 // 288个点/天

	for day := 0; day < 7; day++ {
		for point := 0; point < pointsPerDay; point++ {
			// 计算当前是第几分钟（0-1439）
			minuteOfDay := point * 5
			hour := minuteOfDay / 60

			// 基础流量模式（模拟真实流量曲线）
			var baseFlow float64
			switch {
			case hour >= 0 && hour < 6: // 夜间：低流量 30-40
				baseFlow = 30 + float64(hour)*1.5
			case hour >= 6 && hour < 9: // 早晨：逐渐上升 40-85
				baseFlow = 40 + float64(hour-6)*15
			case hour >= 9 && hour < 12: // 上午：稳定 85-100
				baseFlow = 85 + float64(hour-9)*5
			case hour >= 12 && hour < 14: // 午高峰：峰值 140-180
				baseFlow = 140 + float64(hour-12)*20
			case hour >= 14 && hour < 19: // 下午：稳定下降 120-80
				baseFlow = 120 - float64(hour-14)*8
			case hour >= 19 && hour < 22: // 晚高峰：最高峰 180-210
				baseFlow = 180 + float64(hour-19)*15
			case hour >= 22 && hour < 24: // 晚上：逐渐下降 180-50
				baseFlow = 180 - float64(hour-22)*65
			}

			// 添加平滑的波动（使用正弦波模拟自然波动）
			wave := math.Sin(float64(minuteOfDay)*math.Pi/720) * 10

			// 添加随机波动（±15%）
			randomFactor := 1.0 + (rand.Float64()-0.5)*0.3

			// 周末流量略低（假设第6、7天是周末）
			weekendFactor := 1.0
			if day >= 5 {
				weekendFactor = 0.85
			}

			flow := int((baseFlow + wave) * randomFactor * weekendFactor)
			// 确保流量不为负
			if flow < 10 {
				flow = 10
			}
			ans = append(ans, flow)
		}
	}
	return ans
}
