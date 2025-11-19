# FAQ

<div align="right">

[English](FAQ_EN.md) | [中文](FAQ.md)

</div>

This document collects common questions and solutions when using the gin-vue-web framework.

## 📋 Table of Contents

- [CRUD Related](#crud-related)
- [Rule Configuration Related](#rule-configuration-related)
- [Routing Related](#routing-related)
- [Frontend Related](#frontend-related)
- [Deployment Related](#deployment-related)
- [Performance Optimization](#performance-optimization)

---

## CRUD Related

### Q1: How to customize the list page?

**A**: You can override the `ActionIndex` method and use a custom template:

```go
func (a *Article) ActionIndex(c *gin.Context) {
    data := gin.H{
        "custom": "data",
        "extra": "info",
    }
    a.RenderDataPage(c, data, "modules/example/article/index")
}
```

Then create the corresponding frontend template file `frontend/src/modules/example/article/index.vue`.

### Q2: How to add custom validation?

**A**: Add validation logic in the `ActionSave` method:

```go
func (a *Article) ActionSave(c *gin.Context) {
    // Read request data
    payload, _ := io.ReadAll(c.Request.Body)
    var article models.Article
    json.Unmarshal(payload, &article)
    
    // Custom validation
    if article.Title == "" {
        a.JsonFail(c, fmt.Errorf("title cannot be empty"))
        return
    }
    
    if len(article.Title) > 100 {
        a.JsonFail(c, fmt.Errorf("title length cannot exceed 100 characters"))
        return
    }
    
    // Call parent class save logic
    c.Request.Body = io.NopCloser(bytes.NewReader(payload))
    a.X.ActionSave(c)
}
```

### Q3: How to implement relational queries?

**A**: Override the `wrapData` method to handle related data:

```go
func (a *Article) wrapData(data []map[string]any) []map[string]any {
    // Get all category IDs
    categoryIds := []int{}
    for _, item := range data {
        if id, ok := item["category_id"].(int); ok {
            categoryIds = append(categoryIds, id)
        }
    }
    
    if len(categoryIds) > 0 {
        // Batch query categories
        categories := []models.Category{}
        g.DB("base").In("id", categoryIds).Find(&categories)
        
        // Build mapping
        categoryMap := make(map[int]string)
        for _, cat := range categories {
            categoryMap[cat.Id] = cat.Name
        }
        
        // Add to data
        for i := range data {
            if id, ok := data[i]["category_id"].(int); ok {
                data[i]["category_name"] = categoryMap[id]
            }
        }
    }
    
    return data
}
```

### Q4: How to customize search logic?

**A**: Override the `buildCondition` method:

```go
import (
    "xorm.io/builder"
    "fmt"
)

func (a *Article) buildCondition(arg map[string]any) builder.Cond {
    cond := builder.NewCond()
    
    // Custom search logic
    if title, ok := arg["title"]; ok && title != "" {
        cond = cond.And(builder.Like{"title", fmt.Sprintf("%%%s%%", title)})
    }
    
    // Date range search
    if startDate, ok := arg["start_date"]; ok {
        cond = cond.And(builder.Gte{"created", startDate})
    }
    if endDate, ok := arg["end_date"]; ok {
        cond = cond.And(builder.Lte{"created", endDate})
    }
    
    return cond
}
```

### Q5: How to implement soft delete?

**A**: Implement soft delete in the Model's `Delete` method:

```go
func (a *Article) Delete(sess *xorm.Session) error {
    // Soft delete: update deleted field
    a.Deleted = 1
    a.DeletedAt = time.Now()
    _, err := sess.ID(a.Id).Update(a)
    return err
    
    // Or hard delete
    // return a.DeleteBean(sess, a)
}
```

Then add filter conditions when querying:

```go
func NewArticle() *Article {
    a := &Article{X: g.NewX(&models.Article{})}
    
    // Only query non-deleted records
    a.AndWheres = []map[string]any{
        {"deleted": "0"},
    }
    
    return a
}
```

### Q6: How to implement data permissions (only view own data)?

**A**: Add user filter conditions in the `ActionFetch` method:

```go
func (a *Article) ActionFetch(c *gin.Context) {
    // Get current user
    user := a.GetUser(c)
    
    // Add user filter conditions
    a.AndWheres = []map[string]any{
        {"user_id": user.Id},
    }
    
    // Call parent method
    a.X.ActionFetch(c)
}
```

### Q7: How to implement batch operations?

**A**: Use `g.XB` instead of `g.X`:

```go
type Article struct {
    *g.XB[*models.Article]
}

func NewArticle() *Article {
    a := &Article{XB: g.NewXB(&models.Article{})}
    
    // Batch operation buttons
    a.Tool = []*g.Tool{
        {"Add", "plus", "edit", "modal", "primary"},
        {"Batch Delete", "delete", "delete", "async", []string{"ids"}},
    }
    
    return a
}
```

---

## Rule Configuration Related

### Q8: How to configure field default values?

**A**: Set default values in the Model's `Save` method:

```go
func (a *Article) Save(sess *xorm.Session) error {
    // Set default values
    if a.Status == "" {
        a.Status = "0"  // Default draft
    }
    if a.Views == 0 && a.Id == 0 {
        a.Views = 0  // Default page views when creating
    }
    
    return a.SaveBean(sess, a)
}
```

### Q9: How to implement dynamic field options (loaded from database)?

**A**: Use the database query method with `trans` configuration:

```json
{
  "article": [
    {
      "key": "category_id",
      "name": "Category",
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

Or use AJAX method:

```json
{
  "article": [
    {
      "key": "category_id",
      "name": "Category",
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

### Q10: How to implement date range search?

**A**: Configure the date field in `rule.json`, then handle it in `buildCondition`:

```json
{
  "article": [
    {
      "key": "created",
      "name": "Created Time",
      "search": 0
    }
  ]
}
```

In the Controller:

```go
func (a *Article) buildCondition(arg map[string]any) builder.Cond {
    cond := builder.NewCond()
    
    // Date range search
    if startDate, ok := arg["start_date"]; ok && startDate != "" {
        cond = cond.And(builder.Gte{"created", startDate})
    }
    if endDate, ok := arg["end_date"]; ok && endDate != "" {
        cond = cond.And(builder.Lte{"created", endDate + " 23:59:59"})
    }
    
    return cond
}
```

### Q11: How to hide certain fields from the list page?

**A**: Don't configure the field in `rule.json`, or set `"list": false` (if supported):

```json
{
  "article": [
    {
      "key": "content",
      "name": "Content",
      "list": false  // Not displayed on list page
    }
  ]
}
```

### Q12: How to implement conditional field display (based on other field values)?

**A**: Customize in the frontend template, or pass conditional data in the backend `ActionEdit`:

```go
func (a *Article) ActionEdit(c *gin.Context) {
    // Get rules
    rules := g.GetRule("article")
    
    // Filter rules based on conditions
    // ...
    
    data := gin.H{
        "rules": filteredRules,
    }
    a.RenderDataPage(c, data, "templates/index")
}
```

---

## Routing Related

### Q13: How to add custom routes?

**A**: Manually add them in the `Route` function after `BindActions`:

```go
func Route(rg *gin.RouterGroup) {
    g.RegController("base", "user", NewUser())
    g.BindActions(rg)
    
    // Custom routes
    rg.POST("/custom/route", customHandler)
    rg.GET("/api/custom", customAPIHandler)
}
```

### Q14: What to do about route conflicts?

**A**: Check for duplicate Controller registrations, or adjust module and controller names:

```go
// Avoid conflicts
g.RegController("base", "user", NewUser())
g.RegController("admin", "user", NewUser())  // Different module
```

---

## Frontend Related

### Q15: How does the frontend call backend APIs?

**A**: Use the `lib.curl` method:

```typescript
import lib from '@libs/lib.ts'

// GET request
const data = await lib.curl('/example/article/index')

// POST request
const result = await lib.curl('/example/article/save', {
  title: 'Title',
  content: 'Content'
})

// Error handling
try {
  const result = await lib.curl('/example/article/save', data)
  console.log('Save successful', result)
} catch (error) {
  console.error('Save failed', error)
  // Show error message
}
```

### Q16: How to customize table columns?

**A**: Create a custom list page template using Ant Design Vue's Table component:

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
        <a-button @click="handleEdit(record)">Edit</a-button>
        <a-button @click="handleDelete(record)">Delete</a-button>
      </template>
    </template>
  </a-table>
</template>
```

---

## Deployment Related

### Q17: How to configure for production environment?

**A**: Refer to the [Deployment Documentation](DEPLOYMENT.md), mainly note:

1. Modify environment configuration in `cfg.json`
2. Set `debug: false`
3. Configure correct database connection
4. Configure log path
5. Use process management tools (systemd/supervisor)

### Q18: How to configure HTTPS?

**A**: Use Nginx reverse proxy and configure SSL certificate:

```nginx
server {
    listen 443 ssl http2;
    server_name your_domain.com;
    
    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;
    
    # Other configurations...
}
```

### Q19: Frontend route returns 404 after refresh?

**A**: Configure Nginx's `try_files`:

```nginx
location / {
    root /path/to/frontend/dist;
    try_files $uri $uri/ /index.html;
}
```

---

## Performance Optimization

### Q20: How to optimize database query performance?

**A**: 

1. **Add indexes**: Add indexes on frequently queried fields
2. **Use connection pool**: Configure database connection pool parameters
3. **Avoid N+1 queries**: Use batch queries and relational queries
4. **Pagination**: Always use pagination to avoid loading large amounts of data at once

```go
// Batch query related data
func (a *Article) wrapData(data []map[string]any) []map[string]any {
    // Collect all IDs
    ids := []int{}
    for _, item := range data {
        ids = append(ids, item["category_id"].(int))
    }
    
    // Batch query
    categories := []models.Category{}
    g.DB("base").In("id", ids).Find(&categories)
    
    // Build mapping
    // ...
}
```

### Q21: How to optimize frontend performance?

**A**:

1. **Code splitting**: Use route lazy loading
2. **Image optimization**: Use WebP format, add lazy loading
3. **Cache strategy**: Configure static resource caching
4. **Gzip compression**: Enable Nginx Gzip

---

Return to [README_EN.md](../README_EN.md)

