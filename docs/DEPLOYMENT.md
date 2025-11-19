# 部署指南

<div align="right">

[English](DEPLOYMENT_EN.md) | [中文](DEPLOYMENT.md)

</div>

本文档说明如何在开发环境和生产环境中部署 gin-vue-web 项目。

## 📋 目录

- [生产环境部署](#生产环境部署)
- [Nginx 配置](#nginx-配置)
- [其他部署方式](#其他部署方式)

---

## 生产环境部署

### 后端部署

#### 1. 编译项目

```bash
cd backend
go build -o backend main.go
```

#### 2. 配置生产环境

创建生产环境配置文件 `backend/cfg.json`：


#### 3. 启动服务

直接运行二进制文件：

```bash
./backend -c ./cfg.prod.json
```

#### 4. 使用进程管理工具（推荐）

推荐使用 **Ansible + Supervisor** 进行自动化部署和进程管理。

### 前端部署

#### 1. 构建生产版本

```bash
cd frontend
npm run build
```

构建产物在 `frontend/dist` 目录。

#### 2. 使用 Nginx 启动服务

将 `frontend/dist` 目录部署到服务器，然后使用 Nginx 配置指向该目录。具体配置见 [Nginx 配置](#nginx-配置) 章节。

---

## Nginx 配置

### 配置示例

参考项目根目录的 `nginx.conf` 文件，配置示例如下：

```nginx
server {
    listen 80;
    server_name your_domain.com;
    
    # 前端静态文件
    location / {
        root /path/to/frontend/dist;
        try_files $uri $uri/ /index.html;
    }
    
    # 后端 代理
    location /web {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
```

### 配置说明

- `root /path/to/frontend/dist`: 指向前端构建产物目录
- `try_files $uri $uri/ /index.html`: 支持前端路由（history 模式）
- `location /web`: 将请求代理到后端服务（端口 3000）

---

## 其他部署方式

### Docker 部署

项目支持 Docker 容器化部署。可以构建后端和前端 Docker 镜像，使用 Docker Compose 编排服务。具体配置可根据实际需求进行定制。

### 云部署

项目可以部署到各种云平台（如 AWS、阿里云、腾讯云等）。建议使用云平台提供的容器服务或虚拟机服务，结合自动化部署工具（如 Ansible）进行部署。

---

返回 [README.md](../README.md)
