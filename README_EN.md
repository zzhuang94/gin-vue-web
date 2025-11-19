# gin-vue-web

<div align="right">

[English](README_EN.md) | [中文](README.md)

</div>

A **minimal MVC web framework** built with **Gin + Vue 3 + Ant Design Vue + Vite**, focused on building admin consoles fast.

> 🚀 **Take off even as a beginner!** Just a pinch of Go syntax (being able to define a `struct` is enough), a little frontend knowledge (`html`, `css`, `js` basics), plus a basic understanding of `MySQL`, and you are ready for a full-stack journey.  
>
> 🎓 **No Vue background?** This project is your best practice ground. Learn Vue 3 while doing real work. The framework already wraps the complex logic so you only focus on business.  
>
> ⚙️ **Configuration free!** Tedious frontend compilation and bundling are already handled. Ant Design Vue, ECharts, Font Awesome, and other common components are pre-integrated. Just concentrate on core development and forget the annoying config files.  
>
> 💪 **CRUD in three steps!** Development? Three quick steps. Deployment? One command. From zero to production faster than writing Hello World.  
>
> ⚡ **Build in 5 minutes, run in 10, finish the first module in 30!** No hype—this is the real onboarding experience. 🎯
>
> 🌟 **Embrace open source with an all-star tech stack!** Gin (Go’s hottest web framework) on the backend, Vue 3 (the benchmark progressive framework) on the frontend, and Vite (lightning-fast build tool). All widely adopted, battle-tested technologies. Add whatever you want: plug in ChatGPT, integrate WeChat/Alipay payments, enable real-time chat with WebSocket, create data dashboards with built-in ECharts, or add third-party login via OAuth 2.0. The open-source ecosystem is your toolbox—build anything you imagine. 🚀

---

## Sneak Peek

> 😎 See the experience before diving into the internals—feel the instant “what you see is what you get” with gin-vue-web.

### Feature Demo Video

<div align="center" style="margin: 16px 0 32px;">
  <a href="docs/screenshots/gin-vue-show.webm" style="display:inline-block;padding:14px 36px;border-radius:999px;background:linear-gradient(135deg,#ff7a18,#ffb347);color:#fff;font-size:18px;font-weight:600;text-decoration:none;box-shadow:0 12px 30px rgba(255,122,24,0.35);">
    🎬 Click to download the demo video (.webm)
  </a>
  <p style="margin-top:12px;color:#7f8aa7;font-size:14px;">
    Your browser will download the video directly. Wait a few seconds and double-click to play.
  </p>
</div>

### Feature Screenshots

<table>
  <tr>
    <td align="center">
      <strong>List page · search/sort/pagination</strong><br>
      <img src="docs/screenshots/x-list.png" width="240" />  
    </td>
    <td align="center">
      <strong>Create modal</strong><br>
      <img src="docs/screenshots/x-add.png" width="240" />
    </td>
    <td align="center">
      <strong>Edit modal</strong><br>
      <img src="docs/screenshots/x-edit.png" width="240" />
    </td>
    <td align="center">
      <strong>Relation management modal</strong><br>
      <img src="docs/screenshots/x-rela.png" width="240" />
    </td>
  </tr>
  <tr>
    <td align="center">
      <strong>Navigation tree</strong><br>
      <img src="docs/screenshots/navtree.png" width="240" />
    </td>
    <td align="center">
      <strong>Route list</strong><br>
      <img src="docs/screenshots/route.png" width="240" />
    </td>
    <td align="center">
      <strong>Role management</strong><br>
      <img src="docs/screenshots/role.png" width="240" />
    </td>
    <td align="center">
      <strong>Role authorization</strong><br>
      <img src="docs/screenshots/role-access.png" width="240" />
    </td>
    <td align="center">
      <strong>Profile center</strong><br>
      <img src="docs/screenshots/user-edit.png" width="240" />
    </td>
  </tr>
  <tr>
    <td align="center">
      <strong>Shared component library</strong><br>
      <img src="docs/screenshots/widget.png" width="240" />
    </td>
    <td align="center">
      <strong>ECharts sample</strong><br>
      <img src="docs/screenshots/echarts.png" width="240" />
    </td>
    <td align="center">
      <strong>Icon library</strong><br>
      <img src="docs/screenshots/icon.png" width="240" />
    </td>
  </tr>
</table>

---

## 📋 Table of Contents

- [I. Project Introduction](#i-project-introduction)
  - [1.1 Minimal MVC Architecture](#11-minimal-mvc-architecture)
  - [1.2 Intelligent Route Auto-Registration](#12-intelligent-route-auto-registration)
  - [1.3 Universal CRUD Framework](#13-universal-crud-framework)
  - [1.4 Rule-Driven Development](#14-rule-driven-development)
  - [1.5 Built-in Enterprise-Grade Features](#15-built-in-enterprise-grade-features)
- [II. Quick Start](#ii-quick-start)
  - [2.1 Environment Setup](#21-environment-setup)
  - [2.2 Development Mode](#22-development-mode)
  - [2.3 Production Deployment Guide](#23-production-deployment-guide)
- [III. How to Read the Code](#iii-how-to-read-the-code)
  - [3.1 Project Structure Overview](#31-project-structure-overview)
  - [3.2 Backend Code Structure](#32-backend-code-structure)
  - [3.3 Frontend Code Structure](#33-frontend-code-structure)
  - [3.4 Code Reading Suggestions](#34-code-reading-suggestions)
- [IV. Core Design Philosophy](#iv-core-design-philosophy)
  - [4.1 X Struct - The Core of the Framework](#41-x-struct---the-core-of-the-framework)
  - [4.2 XB Struct - Extension for Batch Operations](#42-xb-struct---extension-for-batch-operations)
  - [4.3 Rule Configuration - The Power of Configuration-Driven Development](#43-rule-configuration---the-power-of-configuration-driven-development)
  - [4.4 Route Auto-Registration - The Magic of Zero Configuration](#44-route-auto-registration---the-magic-of-zero-configuration)
  - [4.5 Data Flow Process - Understanding the Whole System](#45-data-flow-process---understanding-the-whole-system)
- [V. Debugging Methods](#v-debugging-methods)
  - [5.1 Locating Backend Code from Browser Requests](#51-locating-backend-code-from-browser-requests)
  - [5.2 Locating Vue Components from Page Elements](#52-locating-vue-components-from-page-elements)
  - [5.3 Common Debugging Scenarios](#53-common-debugging-scenarios)
- [VI. Development Tutorial](#vi-development-tutorial)
  - [6.1 Adding a Regular Page](#61-adding-a-regular-page)
  - [6.2 Adding a CRUD Feature Set](#62-adding-a-crud-feature-set)
  - [6.3 Custom Feature Extensions](#63-custom-feature-extensions)
- [VII. FAQ](#vii-faq)
- [VIII. License](#viii-license)

---

## I. Project Introduction

gin-vue-web is an **enterprise-grade full-stack web development framework** designed for rapidly building modern admin platforms. It adopts a frontend-backend separation architecture, based on the Go and Vue 3 ecosystem, and follows the **convention over configuration** design philosophy, improving traditional CRUD development efficiency by more than 10x.

### 1.1 Minimal MVC Architecture

The framework adopts the classic **MVC (Model-View-Controller) architecture pattern**, reducing boilerplate code in traditional development by more than 80% through intelligent conventions and auto-assembly. Developers only need to focus on business logic, while the framework automatically handles routing, data binding, form validation, and other tedious tasks.

**Core Advantages:**
- **Zero-config startup**: No configuration files needed, works out of the box
- **Convention over configuration**: Follows best practices, reduces decision-making overhead
- **Highly extensible**: Supports custom middleware, interceptors, and extension points
- **Excellent performance**: Built on the high-performance Gin framework, easily handles high-concurrency scenarios

### 1.2 Intelligent Route Auto-Registration

The framework automatically scans `Action*` methods in Controllers through **Go reflection mechanism**, achieving zero-configuration automatic route registration. Say goodbye to manually maintaining route tables—just define methods, and routes are automatically generated and registered.

**Technical Highlights:**
- **Reflection scanning**: Automatically identifies methods like `ActionIndex`, `ActionEdit`, `ActionSave`
- **RESTful style**: Automatically generates routes conforming to RESTful standards
- **Type safety**: Compile-time checks to avoid runtime errors
- **Auto documentation**: Supports automatic API documentation generation

### 1.3 Universal CRUD Framework

The framework builds a complete universal CRUD solution around **three core files**, achieving fully automatic generation from data models to frontend pages.

**Core File Architecture:**

| File | Responsibility | Tech Stack |
|------|---------------|-----------|
| `backend/g/x.go` | CRUD core logic engine | Go + XORM |
| `frontend/src/templates/index.vue` | Universal list page template | Vue 3 + Ant Design Vue |
| `backend/g/rule.go` | Field rule configuration engine | JSON configuration-driven |

**Feature Matrix:**

| Feature Module | Supported Features | Description |
|----------------|-------------------|-------------|
| **List Page** | Search, sort, pagination, filter | Supports multi-field combined queries |
| **Data Operations** | Create, edit, delete | Single and batch operations |
| **Data Validation** | Frontend + backend dual validation | Type checking, required validation, format validation |
| **Access Control** | Role-based access control | Fine-grained permission management |

### 1.4 Rule-Driven Development

Through the `rule.json` configuration file, field rules are defined using a **declarative programming** approach, achieving true configuration-driven development. Configure once, works across the full stack.

**Rule Configuration Capabilities:**

- **Display rules**: Field labels, display order, visibility, formatting methods
- **Validation rules**: Required validation, type validation, length limits, regular expressions
- **Search rules**: Exact match, fuzzy search, range queries, IN queries, relational queries
- **Interaction rules**: Dropdown options, cascading selection, date ranges, file uploads

**Advantages:**
- **Zero-code configuration**: No need to write validation logic
- **Frontend-backend sync**: Configure once, automatically syncs across frontend and backend
- **Business focus**: Focus on business logic, not technical details
- **Rapid iteration**: Modify configuration to adjust functionality

### 1.5 Built-in Enterprise-Grade Features

The framework includes **out-of-the-box** enterprise-grade base feature modules, covering 90% of common admin system requirements, allowing you to start building business features from day one.

**Feature List:**

| Feature Module | Technical Implementation | Feature Description |
|----------------|-------------------------|---------------------|
| **User Authentication** | Session | Supports login, registration, account switching |
| **Permission Management** | RBAC + Route Guards | Fine-grained permission control based on roles and routes |
| **Menu Management** | Dynamic menu tree | Supports multi-level menus, icons, permission binding |
| **Session Management** | Redis distributed sessions | Supports cluster deployment, session sharing |
| **Component Library** | Ant Design Vue | 60+ high-quality components, ready to use |
| **Chart Library** | ECharts | 20+ chart types, supports data visualization |
| **Icon Library** | Font Awesome | 1000+ icons, meets various scenario needs |

**Extension Capabilities:**
- **Plugin mechanism**: Supports custom plugin extensions
- **Theme customization**: Supports theme switching and custom styles
- **Internationalization**: Built-in i18n support, easily implement multi-language
- **Responsive design**: Perfectly adapts to PC, tablet, and mobile devices

---

## II. Quick Start

### 2.1 Environment Setup

#### 2.1.1 Development Environment Setup

**Go Installation:**
- Download: https://golang.org/dl/ or https://golang.google.cn/dl/
- Recommended version: Go 1.25.0 and above
- Verify after installation: `go version`

**Node.js Installation:**
- Download: https://nodejs.org/
- Recommended version: Node.js v24.9.0 and above (includes npm)
- Verify after installation: `node -v` and `npm -v`

#### 2.1.2 Database Setup

**MySQL Installation:**
- Download: https://dev.mysql.com/downloads/mysql/
- Or use Docker: `docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=yourpassword mysql:8.0`
- Create two database instances `base` and `core` for storing base data and business data respectively
- Execute initialization SQL (see `docs/sql/init_base.sql` and `docs/sql/init_core.sql`)

#### 2.1.3 Redis Setup

**Redis Installation:**
- Windows: Download https://github.com/microsoftarchive/redis/releases or use WSL
- Linux/Mac: `sudo apt-get install redis-server` or `brew install redis`
- Or use Docker: `docker run -d -p 6379:6379 redis:latest`
- Verify after startup: `redis-cli ping` (should return `PONG`)

#### 2.1.4 Configuration File

Edit `backend/cfg.json` and adjust the database and Redis connection to your actual addresses.

### 2.2 Development Mode

#### 2.2.1 Backend Startup

```bash
cd backend
go mod tidy
go build
./backend
```

The backend runs on `http://localhost:3000` by default.

#### 2.2.2 Frontend Startup

```bash
cd frontend
npm install
npm run dev -- --host 0.0.0.0
```

The frontend runs on `http://localhost:5173` and `http://{real-server-IP}:5173` by default.

#### 2.2.3 Login to System

It is recommended to use `http://{real-server-IP}:5173` to log in. Admin credentials: admin/Admin321!

### 2.3 Production Deployment Guide

For detailed deployment instructions, including production environment configuration, Nginx configuration, Docker deployment, etc., please refer to the [Deployment Documentation](docs/DEPLOYMENT_EN.md).

---

## III. How to Read the Code

### 3.1 Project Structure Overview

gin-vue-web adopts a classic frontend-backend separation architecture. Understanding this architecture is the first step in reading the code.

**What does frontend-backend separation mean?**

Simply put, the frontend (Vue 3) and backend (Go) are two completely independent applications. They communicate through HTTP APIs, with the frontend responsible for display and interaction, and the backend responsible for data processing and business logic. The benefit of this architecture is that frontend and backend can be developed, deployed, and maintained independently.

**Project Directory Structure**

```
gin-vue-web/
├── backend/                 # Backend code (Go)
│   ├── g/                   # Core framework code
│   │   ├── x.go            # CRUD core logic (X struct)
│   │   ├── xb.go           # Batch operation extension (XB struct)
│   │   ├── rule.go         # Rule configuration processing
│   │   ├── action.go       # Route auto-registration
│   │   ├── web.go          # Web base functionality
│   │   └── model.go        # Base model
│   ├── modules/            # Business modules
│   │   ├── base/           # Base module (users, roles, etc.)
│   │   │   ├── controllers/ # Controllers
│   │   │   └── models/      # Data models
│   │   └── res/            # Resource module (example)
│   │       ├── controllers/
│   │       └── models/
│   ├── libs/               # Utility libraries
│   ├── main.go             # Entry file
│   └── rule.json           # Field rule configuration
│
├── frontend/               # Frontend code (Vue 3)
│   ├── src/
│   │   ├── components/     # Common components
│   │   │   ├── searcher.vue    # Searcher
│   │   │   ├── table.vue       # Data table
│   │   │   ├── pager.vue       # Pagination
│   │   │   ├── edit.vue        # Edit form
│   │   │   └── ...
│   │   ├── modules/        # Business pages
│   │   ├── templates/      # Page templates
│   │   │   ├── index.vue   # Universal list page template
│   │   │   ├── layout.vue  # Layout template
│   │   │   └── ...
│   │   ├── libs/           # Utility libraries
│   │   │   └── lib.ts      # Core utility functions
│   │   └── app.vue         # Root component
│   └── package.json
│
└── docs/                   # Documentation
```

### 3.2 Backend Code Structure

#### 3.2.1 Core Framework Code (`backend/g/`)

This directory contains all the core functionality of the framework. Let's understand the role of each file:

**`x.go` - CRUD Core Logic**

This is the most important file in the entire framework. The `X` struct encapsulates complete CRUD (Create, Read, Update, Delete) functionality. When you create a Controller and inherit the `X` struct, you automatically get all functions like list query, data retrieval, create, edit, delete, etc.

The `X` struct contains many fields, each with a specific purpose:
- `DB`: Database connection, used to execute SQL queries
- `Model`: Data model, defines the structure of the data table to be operated on
- `Rules`: Field rules, loaded from `rule.json`, controls field display, validation, search, and other behaviors
- `Tool`: Toolbar button configuration, such as the "Add" button
- `Option`: Row operation button configuration, such as "Edit" and "Delete" buttons for each row
- `AndWheres`: Fixed query conditions, such as only showing published data
- `WrapData`: Data processing function, can perform custom processing before returning data

**`xb.go` - Batch Operation Extension**

`XB` is an extended version of `X`, specifically designed to support batch operations. If you need batch edit or batch delete functionality, use `XB` instead of `X`. It will automatically add "Batch Modify" and "Batch Delete" buttons to the toolbar.

**`rule.go` - Rule Configuration Processing**

This file is responsible for loading and parsing the `rule.json` configuration file. `rule.json` defines field rules for each table, such as whether a field is required, searchable, what dropdown options are available, etc. The framework automatically generates search forms, data tables, and edit forms based on these rules.

**`action.go` - Route Auto-Registration**

This is where the framework's "magic" happens. Through Go's reflection mechanism, the framework automatically scans all methods starting with `Action` in all Controllers, then automatically registers them as routes. This means you don't need to manually write route tables—just define methods, and routes will be automatically generated.

**`web.go` - Web Base Functionality**

Provides base functionality such as page rendering, JSON responses, Session management, etc. These are the fundamental capabilities of web applications, inherited and used by the `X` struct.

#### 3.2.2 Business Module Code (`backend/modules/`)

Business code is organized by modules, with each module representing a business domain. For example, the `base` module contains base functionality like users and roles, while the `res` module might contain resource management-related functionality.

Each module typically contains two directories:
- **`models/`**: Data models, define the structure of database tables
- **`controllers/`**: Controllers, handle HTTP requests, call models for data operations

**A Typical Module Structure**

Taking the Service module as an example, it contains:
- `models/service.go`: Defines the Service data model, containing fields like name, business type, status, etc.
- `controllers/service.go`: Defines the Service controller, inherits `g.XB`, automatically gets all CRUD functionality

This structure is clear and straightforward: models are responsible for data structure, controllers are responsible for business logic.

### 3.3 Frontend Code Structure

#### 3.3.1 Core Files

**`app.vue` - Root Component**

This is the entry point of the entire frontend application. It listens for route changes, and when users access different pages, it calls backend APIs to get page configuration, then dynamically loads the corresponding Vue components. This design allows pages to be completely controlled by the backend, with the frontend only responsible for rendering.

**`libs/lib.ts` - Core Utility Functions**

This file provides the most commonly used utility functions in the frontend:
- `curl()`: Calls backend APIs, automatically adds `/web` prefix, handles requests and responses
- `loadModal()`: Loads modal components, used for edit, create, and other popup operations
- `loadComponent()`: Dynamically loads Vue components
- `smartUrl()`: Intelligently processes URLs, converts relative paths to absolute paths

These functions encapsulate the details of frontend-backend interaction, allowing you to focus on business logic.

**`templates/index.vue` - Universal List Page Template**

This is the universal list page template provided by the framework. It automatically generates search forms, data tables, pagination components, etc., based on the rule configuration returned by the backend. You don't need to write this code for every list page—the framework has already done it for you.

#### 3.3.2 Component System

The framework provides a series of reusable common components:

- **`searcher.vue`**: Searcher component, automatically generates search forms based on rules
- **`table.vue`**: Data table component, supports sorting, selection, row operations
- **`pager.vue`**: Pagination component
- **`edit.vue`**: Edit form component, automatically generates form fields based on rules
- **`button.vue`**: Button component, supports different types of button styles

These components are all built on Ant Design Vue, providing a unified UI style and interaction experience.

### 3.4 Code Reading Suggestions

Reading code is a gradual process. It's recommended to follow this order:

**Step 1: Start from the Entry Point**

Understanding the application startup flow is important. The backend starts from `main.go`, which initializes the database, loads configuration, and registers routes. The frontend starts from `main.ts`, which creates the Vue application, registers routes, and mounts to the DOM.

**Step 2: Understand the Core Framework**

Focus on reading `backend/g/x.go`, which is the core of the entire framework. Understanding how the `X` struct works and what each of its methods does is key to understanding the entire framework.

At the same time, understand how `rule.json` drives frontend forms and backend validation. This is the embodiment of the framework's "configuration-driven" philosophy.

**Step 3: Trace the Data Flow**

Choose a specific feature (such as list query) and trace how data flows from the frontend request:
1. Frontend initiates request
2. Backend receives request
3. Build query conditions
4. Execute database query
5. Process returned data
6. Return to frontend
7. Frontend updates interface

Understanding this flow will help you understand how the entire system works.

**Step 4: Review Example Code**

Refer to example code in `modules/base/` and `modules/res/`. These are actual working code that can help you understand how to use the framework in real projects.

---

## IV. Core Design Philosophy

### 4.1 X Struct - The Core of the Framework

The `X` struct is the core of the entire framework. Understanding it is key to understanding the entire framework.

#### 4.1.1 Design Philosophy: Convention over Configuration

The framework adopts the "convention over configuration" design philosophy. This means the framework defines a set of conventions, and as long as you follow these conventions, the framework can automatically complete most of the work for you.

For example, you only need to:
- Define data models (Model)
- Configure field rules (rule.json)
- Create a Controller and inherit `X`

The framework will automatically generate for you:
- List page
- Search functionality
- Create functionality
- Edit functionality
- Delete functionality

You don't need to write a lot of boilerplate code—the framework has already done it for you.

#### 4.1.2 Core Methods of the X Struct

The `X` struct provides five core methods, corresponding to the five CRUD operations:

**`ActionIndex` - List Page Rendering**

When a user accesses the list page, this method is called. It will:
1. Load field rules from `rule.json`
2. Build toolbar buttons and row operation buttons
3. Get URL parameters (for search conditions)
4. Render the frontend template, passing all configuration information

The frontend template will automatically generate search forms, data tables, and other components based on these configurations.

**`ActionFetch` - Data Retrieval**

This is the core method for retrieving data on the list page. When users search, sort, or paginate, the frontend calls this method.

Its workflow is:
1. Receive search parameters, sort parameters, and pagination parameters from the frontend
2. Build SQL query conditions based on the `search` configuration in `rule.json`
3. Execute database queries to get total count and data list
4. Call `WrapData` to process data (e.g., translate foreign key IDs to display names)
5. Return data in JSON format to the frontend

**`ActionEdit` - Edit Page Rendering**

When a user clicks the "Add" or "Edit" button, this method is called. It will:
1. Determine if it's create or edit (via the `id` in URL parameters)
2. If editing, load existing data from the database
3. Load field rules from `rule.json`
4. Render the edit form template, passing data and rules

The frontend template will automatically generate form fields based on rules, including input boxes, dropdown selections, multi-line text, etc.

**`ActionSave` - Data Saving**

When a user submits a form, this method is called. It will:
1. Determine if it's create or edit
2. Parse data from the request body
3. Validate required fields based on the `required` configuration in `rule.json`
4. Handle special fields (e.g., JSON fields, delimiter fields)
5. Save to the database (using transactions to ensure data consistency)
6. Return success or failure information

**`ActionDelete` - Data Deletion**

When a user clicks the "Delete" button, this method is called. It will:
1. Get the data ID to delete from URL parameters
2. Query the database to check if the data exists
3. Call the model's `Delete` method to delete data (using transactions)
4. Return success information

#### 4.1.3 Building Query Conditions

The `buildCondition` method is responsible for building SQL query conditions based on search parameters. This is a typical example of the framework's "rule-driven" approach.

The framework will decide how to build query conditions based on each field's `search` configuration in `rule.json`:
- `search: 0`: No search, ignore this field
- `search: 1`: Exact match, generates `WHERE field = value`
- `search: 2`: Fuzzy search, generates `WHERE field LIKE '%value%'`
- `search: 3`: IN query, generates `WHERE field IN (value1, value2, ...)`

This way, you only need to configure the search type in `rule.json`, and the framework will automatically handle it without writing SQL code.

### 4.2 XB Struct - Extension for Batch Operations

`XB` is an extended version of `X`, specifically designed to support batch operations.

#### 4.2.1 Why XB is Needed

In real business scenarios, it's often necessary to perform batch operations on multiple records. For example, batch status modification, batch deletion, etc. If you use `X`, you need to implement these features yourself. `XB` has already implemented them for you.

#### 4.2.2 XB Design

`XB` uses Go's generics feature to ensure type safety. It inherits from `X`, so it has all the functionality of `X`, while adding three new methods:

- `ActionBatchEdit`: Batch edit, displays a form where you can select fields to modify
- `ActionBatchSave`: Batch save, applies modifications to selected multiple records
- `ActionBatchDelete`: Batch delete, deletes selected multiple records

#### 4.2.3 Benefits of Using XB

When you use `XB`, the framework will automatically:
1. Add "Batch Modify" and "Batch Delete" buttons to the toolbar
2. Enable multi-select functionality in the table
3. Register routes for batch operations

You don't need to write any additional code—just change `g.X` to `g.XB`.

### 4.3 Rule Configuration - The Power of Configuration-Driven Development

`rule.json` is the core configuration file of the framework, defining field rules for each table. This embodies the framework's "configuration-driven" philosophy.

#### 4.3.1 Role of Rule Configuration

Through `rule.json`, you can define:
- Field display names (for form labels, table column headers)
- Whether fields are required (for form validation)
- Field search types (for building query conditions)
- Field dropdown options (for generating dropdown select boxes)
- Field translation rules (for converting foreign key IDs to display names)

The framework will automatically generate frontend forms and backend validation logic based on these rules—you don't need to write code.

#### 4.3.2 Field Descriptions in Rule Configuration

**Basic Fields**

- `key`: Field name, corresponding to the database table column name
- `name`: Display name, used for form labels and table column headers
- `required`: Whether required, `true` means this field must be filled when saving
- `readonly`: Whether readonly, `true` means this field cannot be edited in the edit page
- `default`: Default value, automatically filled when creating new data

**Search Configuration**

The `search` field controls the search behavior of this field on the list page:
- `0`: No search, this field will not appear in the search form
- `1`: Exact match, uses `WHERE field = value` when searching
- `2`: Fuzzy search, uses `WHERE field LIKE '%value%'` when searching
- `3`: IN query, supports selecting multiple values, uses `WHERE field IN (value1, value2, ...)` when searching

**Display Configuration**

- `textarea`: Whether multi-line text input, `true` uses `<textarea>`, `false` uses `<input>`
- `json`: Whether JSON field, needs to set `textarea: true` at the same time, framework will automatically validate and format JSON
- `bold`: Whether bold display, used for table columns
- `hide`: Whether hidden in table, `true` means this field is not displayed in the table
- `width`: Column width, used for table columns

**Dropdown Option Configuration**

The `limit` field is used to define dropdown option lists. Each option contains:
- `key`: Option value (value stored in database)
- `label`: Option display text (text seen by users)
- `badge`: Option badge style (for table display, such as "success", "danger", etc.)

**Translation Configuration (Trans)**

The `trans` field is used to convert foreign key IDs to display names. This is very useful in relational queries.

For example, an article table has a `category_id` field that stores the category ID. But when displaying, we want to show the category name instead of the ID.

By configuring `trans`, the framework will automatically query the category table and convert IDs to names. If the data volume is large, you can set `ajax: true` to use AJAX for dynamically loading options.

#### 4.3.3 Rule Loading Mechanism

The framework calls `initRules()` at startup to load the `rule.json` file and parse it into an in-memory data structure. When creating an `X` struct, it will look up the corresponding configuration from the rules based on the model's table name (`TableName()`).

When needed, the framework calls the `SelfWrap()` method to process rules, such as:
- Processing dropdown options, generating option lists and mappings
- Processing translation configuration, if in non-AJAX mode, queries the database to generate option lists

### 4.4 Route Auto-Registration - The Magic of Zero Configuration

Route auto-registration is a major feature of the framework, eliminating the need to manually maintain route tables.

#### 4.4.1 How It Works

The framework uses Go's reflection mechanism to automatically scan all methods starting with `Action` in all Controllers, then automatically registers them as routes.

The workflow is:
1. You call `RegController()` to register a Controller
2. The framework uses reflection to get all methods of the Controller
3. Filters methods starting with `Action`
4. Converts method names to route paths (camelCase to kebab-case)
5. Registers as POST routes

#### 4.4.2 Route Rules

The route format is: `/{module}/{controller}/{action}`

For example:
- `ActionIndex` → `/res/service/index`
- `ActionFetch` → `/res/service/fetch`
- `ActionBatchEdit` → `/res/service/batch-edit`

#### 4.4.3 Method Naming Conventions

Method names must follow these conventions:
- Must start with `Action`
- Use camelCase naming (e.g., `ActionBatchEdit`)
- Accept `*gin.Context` parameter

As long as you follow these conventions, the framework can automatically identify and register routes.

#### 4.4.4 Route Table Management

The framework also automatically saves registered routes to the `action` table in the database. This table is used for permission management, controlling which users can access which routes.

### 4.5 Data Flow Process - Understanding the Whole System

Understanding how data flows between frontend and backend is key to understanding the entire system.

#### 4.5.1 Complete Data Flow (Taking List Page as Example)

Let's trace a complete request flow:

**Step 1: User Accesses Page**

User enters a URL in the browser, such as `/res/service/index`.

**Step 2: Frontend Route Handling**

Vue Router captures route changes, and the `app.vue` component listens for route changes.

**Step 3: Frontend Calls Backend API**

`app.vue` calls `lib.curl('/res/service/index')`, and `lib.ts` automatically adds the `/web` prefix, ultimately requesting `POST /web/res/service/index`.

**Step 4: Backend Route Matching**

The Gin framework matches the corresponding route and finds the `ActionIndex` method of the `Service` Controller.

**Step 5: Backend Processes Request**

The `X.ActionIndex()` method is called:
1. Load field rules from `rule.json`
2. Build toolbar buttons and row operation buttons
3. Get URL parameters (for search conditions)
4. Render the frontend template `templates/index.vue`, passing all configurations

**Step 6: Frontend Renders Page**

The frontend receives page configuration and dynamically loads the `templates/index.vue` template. The template automatically generates based on rules:
- Search form (based on `search` configuration)
- Data table (based on field rules)
- Toolbar buttons
- Row operation buttons

**Step 7: Frontend Retrieves Data**

After the page is rendered, `templates/index.vue` automatically calls `lib.curl('fetch', params)` to get data. Requests `POST /web/res/service/fetch`.

**Step 8: Backend Queries Data**

The `X.ActionFetch()` method is called:
1. Parse request parameters (search, sort, pagination)
2. Build query conditions based on the `search` configuration in `rule.json`
3. Execute database query
4. Call `WrapData` to process data (translate foreign keys, etc.)
5. Return data in JSON format

**Step 9: Frontend Updates Interface**

The frontend receives data, updates reactive variables, and Vue automatically re-renders the table to display data.

#### 4.5.2 Details of Search Flow

When a user enters content in the search box and clicks search:

1. Frontend collects values of all search fields, forming an object
2. Calls `lib.curl('fetch', {arg: searchParams, sort: sortParams, page: pageParams})`
3. Backend `ActionFetch` receives parameters
4. Iterates through search parameters, for each field:
   - Finds corresponding rule configuration
   - Builds query conditions based on `search` type
   - Exact match uses `WHERE field = value`
   - Fuzzy search uses `WHERE field LIKE '%value%'`
   - IN query uses `WHERE field IN (value1, value2, ...)`
5. Executes SQL query
6. Returns result data

#### 4.5.3 Details of Save Flow

When a user fills out a form and clicks save:

1. Frontend collects form data, forming an object
2. Calls `lib.curl('save', formData)` or `lib.curl('save?id=123', formData)` (when editing)
3. Backend `ActionSave` receives request
4. Determines if it's create or edit (via `id` in URL parameters)
5. If editing, first loads existing data from database
6. Parses request body, deserializes JSON data to model
7. Validates required fields based on `required` configuration in `rule.json`
8. Handles special fields:
   - JSON fields: Validate and format JSON
   - Delimiter fields: Process multi-line text, remove delimiters
9. Calls model's `Save` method to save to database (using transactions)
10. Returns success or failure information

---

## V. Debugging Methods

### 5.1 Locating Backend Code from Browser Requests

In actual development, you often need to find the corresponding backend code based on browser requests. This is a very important debugging skill.

#### 5.1.1 Step 1: View Network Requests

Open browser developer tools (press F12), switch to the Network tab. When you perform operations (such as search, save), you'll see a list of network requests. Find the corresponding request and click to view details.

#### 5.1.2 Step 2: Parse Request URL

The framework's request format is: `POST /web/{module}/{controller}/{action}`

For example, if you see `POST /web/res/service/fetch`, you can parse:
- `module`: `res`
- `controller`: `service`
- `action`: `fetch`

#### 5.1.3 Step 3: Locate Backend Code

Based on the parsed information, you can quickly locate the code:

1. **Find Controller Registration**

Search for `RegController("res", "service"` in `backend/modules/router.go` to find the registration code.

2. **Find Controller File**

Based on module and controller names, find the file: `backend/modules/res/controllers/service.go`

3. **Find Action Method**

If the Controller inherits from `g.X` or `g.XB`, the `ActionFetch` method is in `backend/g/x.go`. If it's a custom method, it's in the Controller file.

#### 5.1.4 Step 4: Add Debug Code

After finding the code, you can add logs or breakpoints for debugging:

**Using Logs for Debugging**

Add `logrus.Infof()` in the code to output key information, such as parameter values, query conditions, etc. This is the simplest and most direct debugging method.

**Using Debugger**

If using GoLand or VS Code, you can set breakpoints for debugging. This is the most powerful debugging method, allowing you to view variable values, step through execution, etc.

### 5.2 Locating Vue Components from Page Elements

Sometimes you need to find the corresponding Vue component code based on elements on the page.

#### 5.2.1 Step 1: Use Browser Developer Tools

Open developer tools, use the element selector (Ctrl+Shift+C or click the icon in the top left), then click on elements on the page.

#### 5.2.2 Step 2: View Vue Components

If you have the Vue DevTools extension installed (highly recommended), you can:
1. Switch to the Vue tab
2. Select an element to view the corresponding component
3. View the component's props, data, methods, etc.

Vue DevTools is the best tool for debugging Vue applications, clearly showing component trees, reactive data, event triggers, etc.

#### 5.2.3 Step 3: Locate Component Files

**By Component Name**

If you see a component name in Vue DevTools, such as `Searcher`, the corresponding file is `frontend/src/components/searcher.vue`.

**By Template Path**

In `templates/index.vue`, you can see which components are used, such as `<Searcher>`, `<Table>`, etc.

**By Source Maps**

In the browser's Sources tab, you can find files starting with `webpack://` or `vite://`, which correspond to the `frontend/src/` directory.

#### 5.2.4 Step 4: Add Debug Code

Add `console.log()` in the component to output key information, or use Vue DevTools' breakpoint functionality.

### 5.3 Common Debugging Scenarios

#### 5.3.1 Scenario 1: Search Not Working

If search functionality doesn't work, check the following steps:

1. **Check rule.json Configuration**

Ensure the field's `search` configuration is not 0. If `search: 0`, this field will not appear in the search form.

2. **Check Frontend Request Parameters**

Add `console.log()` in the `fetchData` method of `templates/index.vue` to check if search parameters are correctly passed.

3. **Check Backend Query Conditions**

Add logs in the `buildCondition` method of `x.go` to check if the built query conditions are correct.

#### 5.3.2 Scenario 2: Save Failed

If save functionality fails, check the following:

1. **Check Required Fields**

Ensure all fields marked as `required: true` in `rule.json` have values.

2. **Check Backend Validation**

Add logs in the `parsePayload` method of `x.go` to view parsed data and validation errors.

3. **Check Database Constraints**

Use SQL tools to view table structure, check if there are NOT NULL constraints or other constraints causing save failures.

#### 5.3.3 Scenario 3: Dropdown Options Not Displaying

If dropdown options don't display, check:

1. **Check rule.json Configuration**

Ensure `limit` or `trans` configuration is correct.

2. **Check Translation Configuration**

If using `trans`, ensure table names and field names are correct, and there is data in the database.

3. **Check Database Data**

Use SQL queries to confirm there is data in the related table.

#### 5.3.4 Scenario 4: Route 404

If accessing a page returns 404, check:

1. **Check Controller Registration**

Ensure `RegController()` and `BindActions()` are called in `router.go`.

2. **Check Method Name**

Ensure method names start with `Action` and use camelCase naming.

3. **Check Route Table**

Query the `action` table in the database to confirm routes are registered.

---

## VI. Development Tutorial

### 6.1 Adding a Regular Page

Sometimes you need to create a simple page that doesn't involve database operations, such as help pages, about pages, etc.

**Important: Route Convention**

The framework requires a three-level route structure: `{module}/{controller}/{action}`. Frontend file paths must also follow this convention.

#### 6.1.1 Step 1: Create Frontend Page

Frontend file paths must follow the three-level structure: `frontend/src/modules/{module}/{controller}/{action}.vue`

For example, to create route `/example/help/index`, you need to create the file:
- `frontend/src/modules/example/help/index.vue`

The page content is just a regular Vue component and can contain any content.

#### 6.1.2 Step 2: Create Backend Controller

Create a controller file under `backend/modules/{module}/controllers/`.

For example, for route `/example/help/index`, you need to create:
- `backend/modules/example/controllers/help.go`

The Controller needs to inherit `g.Web` (not `g.X`, because CRUD functionality is not needed).

Define an `ActionIndex` method that calls `Render()` to render the frontend page.

#### 6.1.3 Step 3: Register Route

Call `RegController()` in `backend/modules/router.go` to register the Controller.

When registering, you need to specify three parameters: module, controller, and instance.

For example: `g.RegController("example", "help", controllers.NewHelp())`

Then call `BindActions()` to bind routes.

#### 6.1.4 Step 4: Access Page

After starting the service, access the corresponding URL to see the page. The route address is `/example/help/index`.

### 6.2 Adding a CRUD Feature Set

This is the core functionality of the framework. You can complete a full CRUD feature set in just three steps.

#### 6.2.1 Step 1: Define Data Model

Create a Model file and define the data structure. The Model needs to:
- Inherit `g.Model` (includes id, created, updated fields)
- Define business fields
- Implement four methods of the `ModelX` interface: `TableName()`, `New()`, `Save()`, `Delete()`

#### 6.2.2 Step 2: Configure Field Rules

Add table configuration in `rule.json`. Configure for each field:
- `key`: Field name
- `name`: Display name
- `required`: Whether required
- `search`: Search type
- Other needed configurations (such as dropdown options, translations, etc.)

#### 6.2.3 Step 3: Create Controller

Create a Controller file that inherits `g.X` (or `g.XB` if batch operations are needed). The framework will automatically provide all CRUD functionality.

If customization is needed, you can configure `Tool` (toolbar buttons) and `Option` (row operation buttons).

#### 6.2.4 Step 4: Register Route and Create Database Table

Register the Controller in `router.go`, then execute SQL to create the database table.

Done! Now you have a complete CRUD feature, including list page, search, create, edit, delete, and all other functionality.

### 6.3 Custom Feature Extensions

The framework provides many extension points for customizing various features.

#### 6.3.1 Extension 1: Using Batch Operations

If you need batch edit or batch delete, just change `g.X` to `g.XB`. The framework will automatically add batch operation buttons and functionality.

#### 6.3.2 Extension 2: Adding Fixed Query Conditions

By configuring `AndWheres`, you can add fixed query conditions. For example, only show published data, only show current user's data, etc.

#### 6.3.3 Extension 3: Custom Data Processing

Override the `WrapData` method to perform custom processing before returning data. For example, add calculated fields, format data, etc.

#### 6.3.4 Extension 4: Custom Search Logic

Override the `buildCondition` method to implement custom search logic. For example, date range search, complex condition combinations, etc.

#### 6.3.5 Extension 5: Custom Save Logic

Override the `ActionSave` method to perform custom validation and processing before saving. For example, set default values, calculate field values, etc.

#### 6.3.6 Extension 6: Relational Queries

Implement relational queries in `WrapData` to convert foreign key IDs to related data. For example, display category names and author names in article lists.

---

## VII. FAQ

For answers to common questions, including custom list pages, adding validation, relational queries, custom search, etc., please refer to the [FAQ Documentation](docs/FAQ_EN.md).

---

## VIII. License

MIT License

---

<div align="center">

**gin-vue-web** - Rapid Web Application Development

Made with ❤️ by the gin-vue-web Team

</div>
