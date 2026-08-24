# golang+vue3

基于 codectl 工具生成的 Go + Vue3 全栈应用脚手架。本项目演示了如何使用 codectl 从数据库表结构自动生成 CRUD 代码，并包含一套完整的模板系统，可用于二次开发。

---

## 目录

- [环境准备](#环境准备)
- [快速开始](#快速开始)
- [Project / App / Table 概念](#project--app--table-概念)
- [项目结构](#项目结构)
- [Manifest 三段式结构](#manifest-三段式结构)
- [编程规范](#编程规范)
- [模板变量参考](#模板变量参考)
- [模板函数参考](#模板函数参考)
- [常见问题](#常见问题)

---

## 环境准备

### 1. 安装 codectl

```bash
go install github.com/turingdance/codectl@latest
```

或从源码编译：

```bash
git clone https://github.com/turingdance/codectl.git
cd codectl
go install .
```

### 2. 准备数据库

```bash
mysql -u root -p -e "CREATE DATABASE mydb DEFAULT CHARSET utf8mb4;"
```

> 表结构会在 `codectl create` 时通过 `project.post_create.runsql` 自动初始化。

---

## 快速开始

### 第一步：创建项目

```bash
codectl create myapp \
  --template ./golang-vue3 \
  --dsn "mysql://root:123456@tcp(127.0.0.1:3306)/mydb?parseTime=True&loc=Local" \
  --package github.com/me/myapp \
  --force
```

必填参数：`<projectname>` + `--template` + `--dsn`

创建时自动完成：
- 拷贝项目骨架（server + web + .template）
- 执行 `project.post_create` 钩子：
  - **replace**：包名从 `turingdance.com/turing` 替换为 `github.com/me/myapp`
  - **runsql**：执行 `init.sql` 初始化数据库表
  - **exec**：生成路由注册代码
- 保存项目配置到 `~/.codectl/projects/myapp.yaml` 并设为当前激活项目

`--template` 支持四种形态：

```bash
# 1. 相对路径
codectl create myapp -t ./golang-vue3 -s "mysql://..." --force

# 2. 绝对路径
codectl create myapp -t D:\tpl\golang-vue3 -s "mysql://..."

# 3. Git URL
codectl create myapp -t https://github.com/turingdance/golang-vue3.git -s "mysql://..."

# 4. 已注册模板名
codectl create myapp -t golang+vue3 -s "mysql://..." -g github.com/me/myapp --force
```

### 第二步：生成代码

```bash
cd myapp
codectl gen --tables *
```

或只生成部分表：

```bash
# 通配符 + 剥离前缀
codectl gen -n let --tables let_* --strip let_ --force

# 预览（只打印不写文件）
codectl gen --tables let_* --dry-run

# 强制覆盖已有文件
codectl gen --tables let_user --force
```

gen 时的生命周期顺序：

```
app.pre_create
  → 渲染 app 级模板（conf.go, api.go, app-dev.yaml 等）
  → table.pre_create
  → 渲染 table 级模板（每张表的 model/rest/vo/sql/web）
  → table.post_create
  → app.post_create    ← 最后执行 codectl router + codectl runsql
```

> **注意**：`app.post_create` 在所有 table 模板渲染之后才执行，确保 `codectl router` 能扫描到全部 `*.gen.go` 文件。

### 第三步：启动服务

```bash
cd server
go run . -c app-dev.yaml
```

服务默认监听 `http://localhost:8080`。

---

## Project / App / Table 概念

codectl 采用三层架构：

```
Project (项目)          一个落盘骨架 + 一份 projects/<name>.yaml 配置
  └─ App (应用)         项目下的一个子应用，独立 package / dsn / template
       └─ Table (表)    app 下的一张业务表，每次 gen 批处理
```

| 层级 | 创建方式 | 配置存储 |
|---|---|---|
| Project | `codectl create` | `~/.codectl/projects/<name>.yaml` |
| App | create 时默认创建（name = project name）；可添加多个 | `projects/<name>.yaml` 的 `apps` 字段 |
| Table | `codectl gen` 从数据库读取 | 不持久化，每次 gen 从 DB 动态读取 |

当前激活的 project 和 app 存储在 `~/.codectl/session.yaml`。

---

## 项目结构

```
myapp/
├── .template/                    # codectl 模板目录（核心）
│   ├── manifest.yaml             # 三段式模板清单（project / app / table）
│   ├── assets/                   # 静态资源（原样复制，不渲染）
│   │   └── init.sql              # 数据库初始化 SQL
│   ├── app/                      # 应用级模板（每个应用生成一次）
│   │   ├── golang/
│   │   │   ├── init.gen.go.tpl   # 应用初始化入口
│   │   │   ├── conf.go.tpl       # 配置定义
│   │   │   ├── api.go.tpl        # API 启动入口
│   │   │   ├── app-dev.yaml.tpl  # 开发环境配置
│   │   │   ├── app-prod.yaml.tpl # 生产环境配置
│   │   │   ├── logic.gen.go.tpl  # logic 层基类
│   │   │   ├── router.gen.go.tpl # 路由注册模板
│   │   │   └── vars.gen.go.tpl   # 模型注册
│   │   └── web/
│   │       ├── index.ts.tpl      # API 请求基类
│   │       └── types.ts.tpl      # 通用类型定义
│   └── table/                    # 表级模板（每张表生成一份）
│       ├── golang/
│       │   ├── model.gen.go.tpl  # 数据模型
│       │   ├── rest.gen.go.tpl   # REST 控制器
│       │   └── vo.gen.go.tpl     # 请求/响应 VO
│       ├── sql/
│       │   └── rights.sql.tpl    # 权限初始化 SQL
│       └── web/
│           ├── api.ts.tpl        # API 调用
│           └── model.vue.tpl     # Vue 页面
├── server/                       # 生成的 Go 后端代码
│   ├── api/
│   ├── internal/
│   │   ├── app/myapp/
│   │   │   ├── model/            # 数据模型
│   │   │   ├── logic/            # 业务逻辑
│   │   │   ├── rest/             # REST 控制器
│   │   │   └── vo/               # 视图对象
│   │   └── conf/
│   ├── app-dev.yaml
│   └── go.mod
├── web/                          # 生成的前端代码
│   └── src/
│       ├── api/myapp/            # API 调用
│       └── views/myapp/          # Vue 页面
└── sql/                          # 生成的 SQL 脚本
    └── myapp/                    # 按应用名分目录
        ├── init.sql              # 初始化 SQL（create 时已执行）
        └── let_*.sql             # 每张表的权限 SQL
```

---

## Manifest 三段式结构

`manifest.yaml` 采用 **project / app / table** 三段式结构，每段支持 `pre_create` / `post_create` 生命周期钩子。

### 生命周期执行顺序

**`codectl create` 时：**

```
project.pre_create
  → 拷贝/克隆骨架
  → project.post_create
    → replace: 包名替换（turingdance.com/turing → {{.App.Package}}）
    → runsql:  执行 init.sql 初始化数据库
    → exec:    codectl router 生成路由
  → app.pre_create
  → 渲染 app 级模板
  → app.post_create
```

**`codectl gen` 时：**

```
app.pre_create
  → 渲染 app 级模板（conf.go, api.go, app-dev.yaml 等）
  → table.pre_create
  → 渲染 table 级模板（每张表循环）
  → table.post_create
  → app.post_create    ← 最后执行，确保 router 能扫描全部 *.gen.go
    → exec: codectl router -s ./server/internal/app/{app.name}/rest
    → exec: codectl runsql --dir ./sql/{app.name}
```

### manifest.yaml 结构

```yaml
# 模板元信息
meta:
  name: golang-vue3
  lang: golang
  dbtype: mysql
  author: codectl

# 全局默认配置
defaults:
  replace:
    exts:                 # replace 扫描的文件扩展名；留空则扫描所有非二进制文件
      - .go
      - .ts
      - .vue
      - .yaml
      - .mod
      - .sum

# 字段类型映射
mapper:
  mysql-golang:
    DATETIME: types.DateTime
    VARCHAR: string
    TINYINT: int8
    BIGINT: int64
    TEXT: string
    # ...

# Project 段（create 一次性）
project:
  pre_create: []
  templates:
    - tpl: golang/init.gen.go.tpl
      out: server/internal/app/{app.name}/init.gen.go
  post_create:
    replace:
      - from: "turingdance.com/turing"
        to: "{{.App.Package}}"
    runsql:
      - init.sql
    exec:
      - codectl router -s server/internal/app/{app.name}/rest

# App 段（create + 每次 gen）
app:
  pre_create: []
  templates:
    - tpl: golang/conf.go.tpl
      out: server/internal/conf/conf.go
    - tpl: golang/api.go.tpl
      out: server/api/api.go
    - tpl: golang/app-dev.yaml.tpl
      out: server/app-dev.yaml
    # ...
  post_create:
    exec:
      - codectl router -s ./server/internal/app/{app.name}/rest
      - codectl runsql --dir ./sql/{app.name}

# Table 段（每次 gen 批处理）
table:
  pre_create: []
  templates:
    - tpl: golang/model.gen.go.tpl
      out: server/internal/app/{app.name}/model/{module}.model.gen.go
    - tpl: golang/rest.gen.go.tpl
      out: server/internal/app/{app.name}/rest/{module}.gen.go
    - tpl: golang/vo.gen.go.tpl
      out: server/internal/app/{app.name}/vo/{module}.vo.gen.go
    - tpl: sql/rights.sql.tpl
      out: sql/{app.name}/{module}.sql
    - tpl: web/api.ts.tpl
      out: web/src/api/{app.name}/{module}.ts
    - tpl: web/model.vue.tpl
      out: web/src/views/{app.name}/{module}/index.vue
  post_create: []
```

### 生命周期钩子动作

每个 `pre_create` / `post_create` 支持三种通用动作：

| 动作 | 说明 | 示例 |
|---|---|---|
| `replace` | 文本替换（from → to） | 替换 package name |
| `runsql` | 执行 SQL 脚本 | 初始化表结构、权限数据 |
| `exec` | 执行命令行 | `codectl router`、`codectl runsql` |

`replace` 的 `to` 字段和 `exec` 的命令支持模板渲染：`{{.App.Package}}`、`{app.name}` 等。

### `defaults.replace.exts`

- 指定后只扫描这些扩展名的文件进行 replace
- 留空（或不设）则扫描所有非二进制文件（通过 NUL 字节探测自动跳过二进制文件）

---

## 编程规范

### `.gen.go` vs `.ext.go` 文件分离原则

codectl 生成的代码文件统一使用 `.gen.go` 后缀。这些文件**不可手动修改**——每次执行 `codectl gen --force` 都会重新生成并覆盖。

如果需要扩展某个表的功能，**必须新建 `.ext.go` 文件**：

| 文件 | 用途 | 是否可修改 |
|---|---|---|
| `user.model.gen.go` | codectl 自动生成的数据模型 | **不可修改** |
| `user.model.ext.go` | 用户手动编写的模型扩展 | **可修改** |
| `user.rest.gen.go` | codectl 自动生成的 REST 控制器 | **不可修改** |
| `user.rest.ext.go` | 用户手动编写的控制器扩展 | **可修改** |

### 示例：给 model 添加自定义方法

```go
// user.model.ext.go - 手动维护,不会被 codectl 覆盖
package model

import "errors"

// ValidatePassword 验证密码强度
func (obj *User) ValidatePassword() error {
    if len(obj.Password) < 8 {
        return errors.New("密码长度不能少于8位")
    }
    return nil
}
```

### 示例：给 rest 控制器添加自定义接口

```go
// user.rest.ext.go - 手动维护,不会被 codectl 覆盖
package rest

import (
    "net/http"
    "github.com/turingdance/infra/wraper"
)

// ChangePassword 修改密码
func (ctrl *User) ChangePassword(w http.ResponseWriter, req *http.Request) {
    wraper.OkMsg("密码修改成功").Encode(w)
}
```

---

## 模板变量参考

### 应用级模板可用的变量

| 变量 | 类型 | 说明 | 示例 |
|---|---|---|---|
| `.App` | App | 应用信息对象 | - |
| `.App.Name` | string | 应用名称 | `myapp` |
| `.App.Title` | string | 应用标题 | `我的应用` |
| `.App.Author` | string | 作者 | `winlion` |
| `.App.Package` | string | Go 包名 | `github.com/me/myapp` |
| `.App.Dsn` | string | 数据库连接串 | `mysql://...` |
| `.App.DbType` | string | 数据库类型 | `mysql` |
| `.Package` | string | Go 包名（同 `.App.Package`） | - |

### 表级模板可用的变量

包含应用级的所有变量，另加：

| 变量 | 类型 | 说明 | 示例 |
|---|---|---|---|
| `.Module` | string | 模块名（表名去前缀后驼峰化） | `userOrder` |
| `.Name` | string | 原始表名 | `sys_user` |
| `.Title` | string | 表标题（来自 COMMENT） | `用户表` |
| `.Columns` | []Column | 列信息列表 | - |
| `.Primary` | Column | 主键列信息 | - |
| `.Types` | []string | 模板中用到的数据类型列表 | `["string","types.DateTime"]` |

### Column 结构体字段

| 字段 | 类型 | 说明 | 示例 |
|---|---|---|---|
| `.DataColumn` | string | 列名（原始） | `user_name` |
| `.DataType` | string | 映射后的数据类型 | `string` |
| `.DataSize` | int | 字段长度 | `255` |
| `.Title` | string | 列注释 | `用户名` |
| `.IsPrimaryKey` | bool | 是否主键 | `true` |
| `.AutoIncrement` | bool | 是否自增 | `true` |
| `.RawData.DataType` | string | 数据库原始类型 | `VARCHAR(255)` |

---

## 模板函数参考

| 函数 | 说明 | 示例 |
|---|---|---|
| `ucfirst` | 首字母大写 | `{{.Module\|ucfirst}}` → `User` |
| `lcfirst` | 首字母小写 | `{{.Module\|lcfirst}}` → `user` |
| `camel` | 下划线转小驼峰 | `{{.Name\|camel}}` → `userName` |
| `upercamel` | 下划线转大驼峰 | `{{.Name\|upercamel}}` → `UserName` |
| `lower` | 转小写 | |
| `upper` | 转大写 | |
| `has` | 切片包含 | `{{has .Types "time.Time"}}` |
| `contains` | 字符串包含 | `{{contains .DataType "int"}}` |
| `mysqltogorm` | MySQL 类型转 GORM tag | |
| `gormtagtype` | 智能类型映射（varchar 超限降级） | |

---

## 常见问题

### Q: create 时报错 "missing required flag: --dsn"

create 必须指定 `--dsn`：

```bash
codectl create myapp -t ./golang-vue3 -s "mysql://root:123456@tcp(127.0.0.1:3306)/mydb"
```

### Q: create 时报错 "missing required flag: --template"

`--template` 是必填项，支持四种形态：

```bash
codectl create myapp -t ./golang-vue3           # 相对路径
codectl create myapp -t /opt/tpl/golang-vue3    # 绝对路径
codectl create myapp -t https://github.com/turingdance/golang-vue3.git  # Git URL
codectl create myapp -t golang+vue3              # 已注册模板名
```

### Q: 如何只生成部分表？

```bash
codectl gen --tables let_*
codectl gen --tables let_user,let_role
codectl gen --tables let_* --exclude let_log*
```

### Q: 如何覆盖已有文件？

```bash
codectl gen --tables let_user --force
```

`--force` 会在覆盖前自动备份到 `.bak/<时间戳>/` 目录。

### Q: 如何剥离表名前缀？

```bash
codectl gen --tables let_* --strip let_
```

`let_user` → 模块名 `user`，生成 `user.model.gen.go`。

### Q: 修改模板后如何生效？

```bash
# 预览
codectl gen --tables let_user --dry-run

# 确认后强制覆盖
codectl gen --tables let_user --force
```

### Q: 如何添加自定义模板？

1. 在 `.template/app/` 或 `.template/table/` 下创建 `.tpl` 文件
2. 在 `manifest.yaml` 的对应段（app/table）注册
3. 执行 `codectl gen --force`

### Q: 配置文件在哪里？

```
~/.codectl/
├── session.yaml            # 当前 project + app
├── config.yaml             # 默认配置（向后兼容）
└── projects/
    └── myapp.yaml           # 项目配置（含 apps 多应用）
```

### Q: 如何切换项目？

```bash
codectl project ls           # 列出所有项目
codectl project use myapp2   # 切换到 myapp2
```

### Q: 生成的代码有编译错误怎么办？

1. 检查 `manifest.yaml` 中的字段类型映射
2. 用 `--dry-run` 预览生成的代码
3. 常见问题：
   - 缺少 import：模板中的 `{{if has .Types "xxx"}}` 条件不匹配
   - 类型不匹配：`mapper` 中映射的类型与模板期望的不一致
   - 包名错误：检查 `--package` 参数

---

## 作者

**微信：** huwinlion

**团队：** 湖南云立方智能科技有限公司
