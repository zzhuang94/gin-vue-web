# Deployment Guide

<div align="right">

[English](DEPLOYMENT_EN.md) | [中文](DEPLOYMENT.md)

</div>

This document explains how to deploy the gin-vue-web project in both development and production environments.

## 📋 Table of Contents

- [Production Deployment](#production-deployment)
- [Nginx Configuration](#nginx-configuration)
- [Other Deployment Methods](#other-deployment-methods)

---

## Production Deployment

### Backend Deployment

#### 1. Build Project

```bash
cd backend
go build -o backend main.go
```

#### 2. Configure Production Environment

Create production environment configuration file `backend/cfg.json`:

#### 3. Start Service

Run the binary directly:

```bash
./backend -c ./cfg.prod.json
```

#### 4. Use Process Management Tools (Recommended)

It is recommended to use **Ansible + Supervisor** for automated deployment and process management.

### Frontend Deployment

#### 1. Build Production Version

```bash
cd frontend
npm run build
```

The build output is in the `frontend/dist` directory.

#### 2. Start Service with Nginx

Deploy the `frontend/dist` directory to the server, then configure Nginx to point to that directory. For specific configuration, see the [Nginx Configuration](#nginx-configuration) section.

---

## Nginx Configuration

### Configuration Example

Refer to the `nginx.conf` file in the project root directory. Configuration example:

```nginx
server {
    listen 80;
    server_name your_domain.com;
    
    # Frontend static files
    location / {
        root /path/to/frontend/dist;
        try_files $uri $uri/ /index.html;
    }
    
    # Backend proxy
    location /web {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### Configuration Notes

- `root /path/to/frontend/dist`: Points to the frontend build output directory
- `try_files $uri $uri/ /index.html`: Supports frontend routing (history mode)
- `location /web`: Proxies requests to the backend service (port 3000)

---

## Other Deployment Methods

### Docker Deployment

The project supports Docker containerized deployment. You can build backend and frontend Docker images and use Docker Compose to orchestrate services. Specific configurations can be customized according to actual needs.

### Cloud Deployment

The project can be deployed to various cloud platforms (such as AWS, Alibaba Cloud, Tencent Cloud, etc.). It is recommended to use container services or virtual machine services provided by cloud platforms, combined with automated deployment tools (such as Ansible) for deployment.

---

Return to [README_EN.md](../README_EN.md)

