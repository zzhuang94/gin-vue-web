# 常见问题

<div align="right">

[English](FAQ_EN.md) | [中文](FAQ.md)

</div>

本文档收集了 gin-vue-web 框架使用过程中的常见问题和解决方案。

## 📋 目录

- [CRUD 相关](#crud-相关)
- [规则配置相关](#规则配置相关)
- [路由相关](#路由相关)
- [前端相关](#前端相关)
- [部署相关](#部署相关)
- [性能优化](#性能优化)

---

## CRUD 相关

### Q1: 如何自定义列表页？

**A**: 可以重写 `ActionIndex` 方法，使用自定义模板：

```go
func (a *Article) ActionIndex(c *gin.Context) {
    data := gin.H{
        "custom": "data",
        "extra": "info",
    }
    a.RenderDataPage(c, data, "modules/example/article/index")
}
```

然后创建对应的前端模板文件 `frontend/src/modules/example/article/index.vue`。

### Q2: 如何添加自定义验证？

**A**: 在 `ActionSave` 方法中添加验证逻辑：

```go
func (a *Article) ActionSave(c *gin.Context) {
    // 读取请求数据
    payload, _ := io.ReadAll(c.Request.Body)
    var article models.Article
    json.Unmarshal(payload, &article)
    
    // 自定义验证
    if article.Title == "" {
        a.JsonFail(c, fmt.Errorf("标题不能为空"))
        return
    }
    
    if len(article.Title) > 100 {
        a.JsonFail(c, fmt.Errorf("标题长度不能超过100个字符"))
        return
    }
    
    // 调用父类保存逻辑
    c.Request.Body = io.NopCloser(bytes.NewReader(payload))
    a.X.ActionSave(c)
}
```

### Q3: 如何实现关联查询？

**A**: 重写 `wrapData` 方法处理关联数据：

```go
func (a *Article) wrapData(data []map[string]any) []map[string]any {
    // 获取所有分类 ID
    categoryIds := []int{}
    for _, item := range data {
        if id, ok := item["category_id"].(int); ok {
            categoryIds = append(categoryIds, id)
        }
    }
    
    if len(categoryIds) > 0 {
        // 批量查询分类
        categories := []models.Category{}
        g.DB("base").In("id", categoryIds).Find(&categories)
        
        // 构建映射
        categoryMap := make(map[int]string)
        for _, cat := range categories {
            categoryMap[cat.Id] = cat.Name
        }
        
        // 添加到数据中
        for i := range data {
            if id, ok := data[i]["category_id"].(int); ok {
                data[i]["category_name"] = categoryMap[id]
            }
        }
    }
    
    return data
}
```

### Q4: 如何自定义搜索逻辑？

**A**: 重写 `buildCondition` 方法：

```go
import (
    "xorm.io/builder"
    "fmt"
)

func (a *Article) buildCondition(arg map[string]any) builder.Cond {
    cond := builder.NewCond()
    
    // 自定义搜索逻辑
    if title, ok := arg["title"]; ok && title != "" {
        cond = cond.And(builder.Like{"title", fmt.Sprintf("%%%s%%", title)})
    }
    
    // 日期范围搜索
    if startDate, ok := arg["start_date"]; ok {
        cond = cond.And(builder.Gte{"created", startDate})
    }
    if endDate, ok := arg["end_date"]; ok {
        cond = cond.And(builder.Lte{"created", endDate})
    }
    
    return cond
}
```

### Q5: 如何实现软删除？

**A**: 在 Model 的 `Delete` 方法中实现软删除：

```go
func (a *Article) Delete(sess *xorm.Session) error {
    // 软删除：更新 deleted 字段
    a.Deleted = 1
    a.DeletedAt = time.Now()
    _, err := sess.ID(a.Id).Update(a)
    return err
    
    // 或者硬删除
    // return a.DeleteBean(sess, a)
}
```

然后在查询时添加过滤条件：

```go
func NewArticle() *Article {
    a := &Article{X: g.NewX(&models.Article{})}
    
    // 只查询未删除的记录
    a.AndWheres = []map[string]any{
        {"deleted": "0"},
    }
    
    return a
}
```

### Q6: 如何实现数据权限（只查看自己的数据）？

**A**: 在 `ActionFetch` 方法中添加用户过滤条件：

```go
func (a *Article) ActionFetch(c *gin.Context) {
    // 获取当前用户
    user := a.GetUser(c)
    
    // 添加用户过滤条件
    a.AndWheres = []map[string]any{
        {"user_id": user.Id},
    }
    
    // 调用父类方法
    a.X.ActionFetch(c)
}
```

### Q7: 如何实现批量操作？

**A**: 使用 `g.XB` 替代 `g.X`：

```go
type Article struct {
    *g.XB[*models.Article]
}

func NewArticle() *Article {
    a := &Article{XB: g.NewXB(&models.Article{})}
    
    // 批量操作按钮
    a.Tool = []*g.Tool{
        {"新 增", "plus", "edit", "modal", "primary"},
        {"批量删除", "delete", "delete", "async", []string{"ids"}},
    }
    
    return a
}
```

---

## 规则配置相关

### Q8: 如何配置字段的默认值？

**A**: 在 Model 的 `Save` 方法中设置默认值：

```go
func (a *Article) Save(sess *xorm.Session) error {
    // 设置默认值
    if a.Status == "" {
        a.Status = "0"  // 默认草稿
    }
    if a.Views == 0 && a.Id == 0 {
        a.Views = 0  // 新增时默认浏览量
    }
    
    return a.SaveBean(sess, a)
}
```

### Q9: 如何实现字段的动态选项（从数据库加载）？

**A**: 使用 `trans` 配置的数据库查询方式：

```json
{
  "article": [
    {
      "key": "category_id",
      "name": "分类",
      "search": 1,
      "trans": {
        "table": "category",
        "key": "id",
        "val": "name"
      }
    }
  ]
}
```

或者使用 AJAX 方式：

```json
{
  "article": [
    {
      "key": "category_id",
      "name": "分类",
      "trans": {
        "ajax": true,
        "key": "id",
        "val": "name",
        "url": "/api/category/list"
      }
    }
  ]
}
```

### Q10: 如何实现日期范围搜索？

**A**: 在 `rule.json` 中配置日期字段，然后在 `buildCondition` 中处理：

```json
{
  "article": [
    {
      "key": "created",
      "name": "创建时间",
      "search": 0
    }
  ]
}
```

在 Controller 中：

```go
func (a *Article) buildCondition(arg map[string]any) builder.Cond {
    cond := builder.NewCond()
    
    // 日期范围搜索
    if startDate, ok := arg["start_date"]; ok && startDate != "" {
        cond = cond.And(builder.Gte{"created", startDate})
    }
    if endDate, ok := arg["end_date"]; ok && endDate != "" {
        cond = cond.And(builder.Lte{"created", endDate + " 23:59:59"})
    }
    
    return cond
}
```

### Q11: 如何隐藏某些字段在列表页显示？

**A**: 在 `rule.json` 中不配置该字段，或者设置 `"list": false`（如果支持）：

```json
{
  "article": [
    {
      "key": "content",
      "name": "内容",
      "list": false  // 不在列表页显示
    }
  ]
}
```

### Q12: 如何实现字段的条件显示（根据其他字段值）？

**A**: 在前端模板中自定义处理，或者在后端 `ActionEdit` 中传递条件数据：

```go
func (a *Article) ActionEdit(c *gin.Context) {
    // 获取规则
    rules := g.GetRule("article")
    
    // 根据条件过滤规则
    // ...
    
    data := gin.H{
        "rules": filteredRules,
    }
    a.RenderDataPage(c, data, "templates/index")
}
```

---

## 路由相关

### Q13: 如何添加自定义路由？

**A**: 在 `Route` 函数中，在 `BindActions` 之后手动添加：

```go
func Route(rg *gin.RouterGroup) {
    g.RegController("base", "user", NewUser())
    g.BindActions(rg)
    
    // 自定义路由
    rg.POST("/custom/route", customHandler)
    rg.GET("/api/custom", customAPIHandler)
}
```

### Q14: 路由冲突怎么办？

**A**: 检查是否有重复注册的 Controller，或者调整模块名和控制器名：

```go
// 避免冲突
g.RegController("base", "user", NewUser())
g.RegController("admin", "user", NewUser())  // 不同的模块
```

---

## 前端相关

### Q15: 前端如何调用后端 API？

**A**: 使用 `lib.curl` 方法：

```typescript
import lib from '@libs/lib.ts'

// GET 请求
const data = await lib.curl('/example/article/index')

// POST 请求
const result = await lib.curl('/example/article/save', {
  title: '标题',
  content: '内容'
})

// 错误处理
try {
  const result = await lib.curl('/example/article/save', data)
  console.log('保存成功', result)
} catch (error) {
  console.error('保存失败', error)
  // 显示错误提示
}
```

### Q16: 如何自定义表格列？

**A**: 创建自定义列表页模板，使用 Ant Design Vue 的 Table 组件：

```vue
<template>
  <a-table
    :columns="columns"
    :data-source="dataSource"
    :pagination="pagination"
    @change="handleTableChange"
  >
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'action'">
        <a-button @click="handleEdit(record)">编辑</a-button>
        <a-button @click="handleDelete(record)">删除</a-button>
      </template>
    </template>
  </a-table>
</template>
```

---

## 部署相关

### Q17: 生产环境如何配置？

**A**: 参考 [部署文档](DEPLOYMENT.md)，主要注意：

1. 修改 `cfg.json` 中的环境配置
2. 设置 `debug: false`
3. 配置正确的数据库连接
4. 配置日志路径
5. 使用进程管理工具（systemd/supervisor）

### Q18: 如何配置 HTTPS？

**A**: 使用 Nginx 反向代理，配置 SSL 证书：

```nginx
server {
    listen 443 ssl http2;
    server_name your_domain.com;
    
    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;
    
    # 其他配置...
}
```

### Q19: 前端路由刷新后 404？

**A**: 配置 Nginx 的 `try_files`：

```nginx
location / {
    root /path/to/frontend/dist;
    try_files $uri $uri/ /index.html;
}
```

---

## 性能优化

### Q20: 如何优化数据库查询性能？

**A**: 

1. **添加索引**: 在经常查询的字段上添加索引
2. **使用连接池**: 配置数据库连接池参数
3. **避免 N+1 查询**: 使用批量查询和关联查询
4. **分页查询**: 始终使用分页，避免一次性加载大量数据

```go
// 批量查询关联数据
func (a *Article) wrapData(data []map[string]any) []map[string]any {
    // 收集所有 ID
    ids := []int{}
    for _, item := range data {
        ids = append(ids, item["category_id"].(int))
    }
    
    // 批量查询
    categories := []models.Category{}
    g.DB("base").In("id", ids).Find(&categories)
    
    // 构建映射
    // ...
}
```

### Q21: 如何优化前端性能？

**A**:

1. **代码分割**: 使用路由懒加载
2. **图片优化**: 使用 WebP 格式，添加懒加载
3. **缓存策略**: 配置静态资源缓存
4. **Gzip 压缩**: 启用 Nginx Gzip

---

返回 [README.md](../README.md)
