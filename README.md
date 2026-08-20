# golang+vue3

基于 codectl 工具生成的 Go + Vue3 全栈应用脚手架。本项目演示了如何使用 codectl 从数据库表结构自动生成 CRUD 代码，并包含一套完整的模板系统，可用于二次开发。

---

## 目录

- [环境准备](#环境准备)
- [快速开始](#快速开始)
- [项目结构](#项目结构)
- [编程规范](#编程规范)
- [codectl 工作流](#codectl-工作流)
- [模板系统详解](#模板系统详解)
- [如何构建自定义模板](#如何构建自定义模板)
- [如何修改已有模板](#如何修改已有模板)
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
go build -o codectl .
```

### 2. 准备数据库

testapp 默认使用 MySQL。创建数据库并导入初始数据：

```bash
mysql -u root -p -e "CREATE DATABASE testapp DEFAULT CHARSET utf8mb4;"
mysql -u root -p testapp < .template/assets/init.sql
```

> `init.sql` 位于 `.template/assets/`，是模板的静态资源，`codectl gen` 时会原样复制。

---

## 快速开始

### 第一步：配置应用

```bash
cd testapp
codectl set \
  --dsn mysql://root:123456@tcp(127.0.0.1:3306)/testapp \
  --appname testapp \
  --prefix sys_ \
  --package github.com/myorg/testapp \
  --author yourname \
  --title "测试应用"
```

参数会被缓存到 `~/.codectl/config.yaml`，后续命令无需重复输入。

查看当前配置：

```bash
codectl set -l
```

### 第二步：生成代码

```bash
codectl gen --tables *
```

工具会连接数据库，读取所有表结构，按 `.template/` 中的模板生成全套 CRUD 代码。

**预览模式**（只打印不写文件）：

```bash
codectl gen --tables * --dry-run
```

**只生成部分表**：

```bash
codectl gen --tables sys_user,sys_role
codectl gen --tables sys_*
codectl gen --tables sys_* --exclude sys_log*
```

**覆盖已有文件**（默认跳过已存在文件，加 `--overwrite` 会备份后覆盖）：

```bash
codectl gen --tables sys_user --overwrite
```

### 第三步：启动服务

```bash
cd server
go run . -c app-dev.yaml
```

服务默认监听 `http://localhost:8080`。

---

## 项目结构

```
testapp/
├── .template/                    # codectl 模板目录（核心）
│   ├── manifest.yaml             # 模板清单（唯一配置入口）
│   ├── assets/                   # 静态资源（原样复制，不渲染）
│   │   └── init.sql
│   ├── app/                      # 应用级模板（每个应用生成一次）
│   │   ├── golang/               # Go 后端模板
│   │   │   ├── init.gen.go.tpl   # 应用初始化入口
│   │   │   ├── conf.go.tpl       # 配置定义
│   │   │   ├── api.go.tpl        # API 启动入口
│   │   │   ├── app-dev.yaml.tpl  # 开发环境配置
│   │   │   ├── app-prod.yaml.tpl # 生产环境配置
│   │   │   ├── logic.gen.go.tpl  # logic 层基类
│   │   │   ├── router.gen.go.tpl # 路由注册
│   │   │   └── vars.gen.go.tpl   # 模型注册
│   │   └── web/                  # 前端模板
│   │       ├── index.ts.tpl      # API 请求基类
│   │       └── types.ts.tpl      # 通用类型定义
│   └── table/                    # 表级模板（每张表生成一份）
│       ├── golang/               # Go 后端每表的模板
│       │   ├── model.gen.go.tpl  # 数据模型
│       │   ├── rest.gen.go.tpl   # REST 控制器
│       │   └── vo.gen.go.tpl     # 请求/响应 VO
│       ├── sql/
│       │   └── rights.sql.tpl    # 权限初始化 SQL
│       └── web/                  # 前端每表的模板
│           ├── api.ts.tpl        # API 调用
│           └── model.vue.tpl     # Vue 页面
├── server/                       # 生成的 Go 后端代码
│   ├── api/
│   ├── internal/
│   │   ├── app/testapp/
│   │   │   ├── model/            # 数据模型
│   │   │   ├── logic/            # 业务逻辑
│   │   │   ├── rest/             # REST 控制器
│   │   │   └── vo/               # 视图对象
│   │   └── conf/
│   └── app-dev.yaml
├── web/                          # 生成的前端代码
│   └── src/
│       ├── api/testapp/          # API 调用
│       └── views/testapp/        # Vue 页面
└── sql/                          # 生成的 SQL 脚本
```

---

## 编程规范

### `.gen.go` vs `.ext.go` 文件分离原则

codectl 生成的代码文件统一使用 `.gen.go` 后缀（如 `user.model.gen.go`、`user.rest.gen.go`）。这些文件**不可手动修改**——每次执行 `codectl gen` 都会重新生成并覆盖。

如果需要扩展某个表的功能（添加自定义方法、实现额外接口等），**必须新建 `.ext.go` 文件**，而不是修改 `.gen.go` 文件。

### 文件命名规范

| 文件 | 用途 | 是否可修改 |
|---|---|---|
| `user.model.gen.go` | codectl 自动生成的数据模型 | **不可修改** |
| `user.model.ext.go` | 用户手动编写的模型扩展 | **可修改** |
| `user.rest.gen.go` | codectl 自动生成的 REST 控制器 | **不可修改** |
| `user.rest.ext.go` | 用户手动编写的控制器扩展 | **可修改** |
| `user.vo.gen.go` | codectl 自动生成的 VO | **不可修改** |
| `user.vo.ext.go` | 用户手动编写的 VO 扩展 | **可修改** |

命名规则：`<module>.<layer>.ext.go`

### 为什么不能修改 `.gen.go`？

```
第一次 codectl gen --tables sys_user --overwrite
  → 生成 user.model.gen.go（自动）

用户修改 user.model.gen.go（手动添加方法）
  → 文件内容被修改

第二次 codectl gen --tables sys_user --overwrite
  → user.model.gen.go 被重新生成
  → 用户的修改全部丢失！
```

### 正确做法：使用 `.ext.go` 扩展

**场景 1：给 model 添加自定义方法**

不要修改 `user.model.gen.go`，而是新建 `user.model.ext.go`：

```go
// user.model.ext.go - 手动维护,不会被 codectl 覆盖
package model

import "errors"

// ValidatePassword 验证密码强度
// 这个方法不会被 codectl gen 覆盖
func (obj *User) ValidatePassword() error {
    if len(obj.Password) < 8 {
        return errors.New("密码长度不能少于8位")
    }
    return nil
}

// FullName 返回完整名称
func (obj *User) FullName() string {
    return obj.FirstName + " " + obj.LastName
}
```

**场景 2：给 rest 控制器添加自定义接口**

新建 `user.rest.ext.go`：

```go
// user.rest.ext.go - 手动维护,不会被 codectl gen 覆盖
package rest

import (
    "net/http"
    "github.com/turingdance/infra/wraper"
)

// ChangePassword 修改密码
// 注意:方法名不能和 .gen.go 中的 Create/Update/Delete 等重名
func (ctrl *User) ChangePassword(w http.ResponseWriter, req *http.Request) {
    // 自定义业务逻辑
    wraper.OkMsg("密码修改成功").Encode(w)
}
```

**场景 3：添加额外的 VO 结构体**

新建 `user.vo.ext.go`：

```go
// user.vo.ext.go - 手动维护,不会被 codectl gen 覆盖
package vo

// UserLoginReq 登录请求(自定义,不在生成的 VO 中)
type UserLoginReq struct {
    Username string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
}

// UserProfileResp 用户详情响应(自定义)
type UserProfileResp struct {
    User
    RoleNames []string `json:"roleNames"`
    LastLogin string   `json:"lastLogin"`
}
```

**场景 4：在 logic 层添加自定义业务逻辑**

新建 `user.logic.ext.go`：

```go
// user.logic.ext.go - 手动维护,不会被 codectl gen 覆盖
package logic

import (
    "errors"
    "golang.org/x/crypto/bcrypt"
)

// Login 用户登录(自定义业务逻辑)
func Login(username, password string) (token string, err error) {
    // 自定义登录逻辑
    return "", nil
}

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPassword 校验密码
func CheckPassword(password, hash string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
```

### 目录结构示例

```
internal/app/testapp/
├── model/
│   ├── user.model.gen.go    # codectl 生成(自动)
│   ├── user.model.ext.go    # 用户扩展(手动)
│   ├── role.model.gen.go    # codectl 生成(自动)
│   └── role.model.ext.go    # 用户扩展(手动)
├── logic/
│   ├── logic.gen.go         # codectl 生成(自动)
│   ├── user.logic.ext.go    # 用户扩展(手动)
│   └── role.logic.ext.go    # 用户扩展(手动)
├── rest/
│   ├── router.gen.go        # codectl 生成(自动)
│   ├── user.rest.gen.go     # codectl 生成(自动)
│   ├── user.rest.ext.go     # 用户扩展(手动)
│   └── role.rest.gen.go     # codectl 生成(自动)
└── vo/
    ├── user.vo.gen.go       # codectl 生成(自动)
    └── user.vo.ext.go       # 用户扩展(手动)
```

### 路由注册

自定义的控制器方法需要手动注册路由。在 `user.rest.ext.go` 中添加路由注册函数：

```go
// user.rest.ext.go
package rest

import "net/http"

// RegisterUserExtRoutes 注册用户扩展路由
// 在 init.gen.go 的 CreateRouter 中调用此函数
func RegisterUserExtRoutes(mux *MuxHandler) {
    ctrl := &User{}
    mux.HandleFunc("/user/changePassword", ctrl.ChangePassword)
    mux.HandleFunc("/user/profile", ctrl.GetProfile)
}
```

然后在 `init.gen.go`（如果它是应用级模板生成的一次性文件，不会被覆盖）或 `init.ext.go` 中调用：

```go
// init.ext.go - 手动维护
package testapp

import "github.com/myorg/testapp/internal/app/testapp/rest"

func init() {
    // 注册扩展路由
    rest.RegisterUserExtRoutes(rest.Mux)
}
```

### 总结

| 规则 | 说明 |
|---|---|
| `.gen.go` 文件 | codectl 自动生成，**绝对不要手动修改** |
| `.ext.go` 文件 | 用户手动创建，用于扩展功能 |
| 同名同包 | `.ext.go` 和 `.gen.go` 在同一个 package 中 |
| 方法名不冲突 | `.ext.go` 中的方法名不能和 `.gen.go` 中的重名 |
| 自定义路由 | 在 `.ext.go` 中定义路由注册函数 |

---

## codectl 工作流

```
┌──────────────────────────────────────────────────────────────┐
│                    codectl 三步工作流                         │
├──────────────────────────────────────────────────────────────┤
│                                                              │
│  1. codectl set        2. codectl gen         3. go run     │
│  ┌─────────┐          ┌─────────────┐        ┌──────────┐   │
│  │ 配置数据库│         │ 读取表结构  │        │ 启动服务  │   │
│  │ 缓存参数  │──┐      │ 渲染模板    │──┐     │          │   │
│  └─────────┘  │      └─────────────┘  │     └──────────┘   │
│               ▼                       ▼                      │
│      ~/.codectl/             server/  web/  sql/             │
│      config.yaml                                                   │
│                                                              │
└──────────────────────────────────────────────────────────────┘
```

### 命令速查

```bash
# 配置
codectl set --dsn mysql://... --appname myapp

# 生成
codectl gen --tables *                    # 全部表
codectl gen --tables sys_* --dry-run      # 预览
codectl gen --tables sys_user --overwrite # 覆盖

# 查看数据库表
codectl tables --tables sys_* --detail

# 路由生成
codectl router -s ./server/internal/app/testapp/rest
```

---

## 模板系统详解

### 核心概念

codectl 的模板系统基于 **manifest 清单驱动**，核心思想是：

1. **目录分层表达语义**：`app/` 放应用级模板，`table/` 放表级模板
2. **manifest.yaml 承载元数据**：输出路径、字段映射、静态资源都在清单文件中声明
3. **模板文件只管渲染逻辑**：不再在模板内容里藏输出路径

### manifest.yaml 结构

```yaml
# 模板元信息
meta:
  name: golang-vue3          # 模板名称
  lang: golang               # 主语言
  dbtype: mysql              # 目标数据库类型
  author: codectl            # 默认作者

# 字段类型映射规则
# key 格式: 数据库类型-目标语言
mapper:
  mysql-golang:
    TINYINT: int8            # MySQL TINYINT → Go int8
    VARCHAR: string          # MySQL VARCHAR → Go string
    DATETIME: types.DateTime # MySQL DATETIME → Go types.DateTime
    TEXT: string
    BIGINT: int64
    # ... 完整映射见 manifest.yaml

# 静态资源（原样复制到输出目录，不经过模板渲染）
assets:
  - src: init.sql            # 源文件（位于 assets/ 目录下）
    dst: init.sql            # 目标路径（相对于输出目录）

# 应用级模板（整个应用只生成一次）
app:
  - tpl: golang/init.gen.go.tpl     # 模板文件（位于 app/ 目录下）
    out: server/internal/app/{app.name}/init.gen.go   # 输出路径
  - tpl: golang/conf.go.tpl
    out: server/internal/conf/conf.go
  # ...

# 表级模板（每张表生成一份）
table:
  - tpl: golang/model.gen.go.tpl
    out: server/internal/app/{app.name}/model/{module}.model.gen.go
  - tpl: golang/rest.gen.go.tpl
    out: server/internal/app/{app.name}/rest/{module}.gen.go
  # ...
```

### 输出路径占位符

`out` 字段支持以下占位符，渲染时自动替换：

| 占位符 | 说明 | 示例值 |
|---|---|---|
| `{app.name}` | 应用名称（小写） | `testapp` |
| `{module}` | 模块名（表名去除前缀后驼峰化） | `userOrder` |
| `{pkgpath}` | go 包名（点号转斜杠） | `github.com/myorg/testapp` |

**示例：**

表名 `sys_user`，应用名 `testapp`，前缀 `sys_`：

| out 模板 | 实际输出路径 |
|---|---|
| `server/internal/app/{app.name}/model/{module}.model.gen.go` | `server/internal/app/testapp/model/user.model.gen.go` |
| `web/src/api/{app.name}/{module}.ts` | `web/src/api/testapp/user.ts` |

### app 级 vs table 级模板

| 维度 | app 级模板 | table 级模板 |
|---|---|---|
| 目录 | `.template/app/` | `.template/table/` |
| 生成频率 | 每个应用生成一次 | 每张表生成一次 |
| 数据上下文 | 只有 `.App` | `.App` + `.Module` + `.Columns` 等 |
| 典型产物 | `init.go`、`conf.go`、`router.go` | `user.model.gen.go`、`user.rest.gen.go` |
| 覆盖行为 | `--overwrite` 时备份后覆盖 | `--overwrite` 时备份后覆盖 |

---

## 如何构建自定义模板

### 场景：给 testapp 添加一个 Python 模板

#### 1. 创建目录结构

```
.template/
├── app/
│   └── python/                    # 新增 Python 应用级模板
│       └── main.py.tpl
└── table/
    └── python/                    # 新增 Python 表级模板
        └── model.py.tpl
```

#### 2. 编写模板文件

**app/python/main.py.tpl**：

```
# gen by codectl
# @author {{.App.Author}}
from fastapi import FastAPI

app = FastAPI(title="{{.App.Title}}")

# 路由注册
{{- range .Tables}}
from .routes import {{.Module|lcfirst}}_router
app.include_router({{.Module|lcfirst}}_router)
{{- end}}
```

**table/python/model.py.tpl**：

```
# gen by codectl
# @author {{.App.Author}}
from pydantic import BaseModel
from typing import Optional

class {{.Module|ucfirst}}(BaseModel):
    """{{.Title}}"""
    {{- range .Columns}}
    {{.DataColumn|lcfirst}}: Optional[{{.DataType|pythontype}}] = None
    {{- end}}
```

#### 3. 注册到 manifest.yaml

在 `manifest.yaml` 中添加：

```yaml
app:
  # ... 已有的 golang/web 模板 ...
  - tpl: python/main.py.tpl
    out: server/{app.name}/main.py

table:
  # ... 已有的 golang/web/sql 模板 ...
  - tpl: python/model.py.tpl
    out: server/{app.name}/models/{module}.py
```

#### 4. 添加字段映射

在 `manifest.yaml` 的 `mapper` 中添加 Python 映射：

```yaml
mapper:
  mysql-golang:
    # ... 已有 ...
  mysql-python:                    # 新增
    TINYINT: int
    INT: int
    BIGINT: int
    VARCHAR: str
    TEXT: str
    DATETIME: str
    # ...
```

#### 5. 使用

```bash
# 只用 Python 模板需要修改 manifest.yaml 的 app/table 列表
# 或通过 --template 指定包含 Python 模板的目录
codectl gen --template ./.template --tables *
```

### 场景：添加全新的模板仓库

#### 1. 创建模板目录

```bash
mkdir -p my-template/{app,table,assets}
```

#### 2. 编写 manifest.yaml

```yaml
meta:
  name: my-template
  lang: golang
  dbtype: mysql

mapper:
  mysql-golang:
    VARCHAR: string
    INT: int32
    # ...

assets:
  - src: init.sql
    dst: init.sql

app:
  - tpl: golang/main.go.tpl
    out: cmd/main.go

table:
  - tpl: golang/model.go.tpl
    out: internal/model/{module}.go
```

#### 3. 编写模板文件

```
{{- /* app/golang/main.go.tpl */ -}}
package main

import "{{.App.Package}}/internal/model"

func main() {
    {{- range .Tables}}
    _ = model.{{.Module|ucfirst}}{}
    {{- end}}
}
```

```
{{- /* table/golang/model.go.tpl */ -}}
package model

type {{.Module|ucfirst}} struct {
    {{- range .Columns}}
    {{.DataColumn|upercamel}} {{.DataType}}
    {{- end}}
}
```

#### 4. 注册到 codectl 模板库

```bash
codectl tpl add /path/to/my-template
```

或使用 git 地址：

```bash
codectl tpl add https://github.com/myorg/my-template.git
```

#### 5. 使用

```bash
codectl gen --template my-template --tables *
```

---

## 如何修改已有模板

### 修改模板的步骤

1. **定位模板文件**：根据 manifest.yaml 的 `tpl` 字段找到对应文件
2. **修改模板内容**：使用 Go template 语法编辑
3. **预览**：`codectl gen --dry-run` 查看渲染结果
4. **生成**：`codectl gen --tables xxx --overwrite` 覆盖生成

### 示例 1：给 model 模板添加注释

**目标**：在每个字段上方添加中文注释。

**修改前**（`.template/table/golang/model.gen.go.tpl`）：

```
type {{.Module|ucfirst}} struct{
	{{- range $i,$v := .Columns}}
        {{ $v.DataColumn | upercamel }} {{ $v.DataType }} `json:"..."`
    {{end -}}
}
```

**修改后**：

```
type {{.Module|ucfirst}} struct{
	{{- range $i,$v := .Columns}}
        // {{$v.Title}}
        {{ $v.DataColumn | upercamel }} {{ $v.DataType }} `json:"..."`
    {{end -}}
}
```

**验证**：

```bash
codectl gen --tables sys_user --dry-run
```

### 示例 2：修改输出路径

**目标**：把 model 文件从 `model/` 目录移到 `models/` 目录。

**修改 manifest.yaml**：

```yaml
table:
  - tpl: golang/model.gen.go.tpl
    out: server/internal/app/{app.name}/models/{module}.gen.go  # model/ → models/
```

无需修改模板文件本身，只改 manifest.yaml 即可。

### 示例 3：添加新的表级模板

**目标**：为每张表生成一个 Markdown 文档。

**1. 创建模板文件** `.template/table/docs/readme.md.tpl`：

```markdown
# {{.Title}}

> 表名: `{{.Name}}` | 模块: `{{.Module}}`

## 字段说明

| 列名 | 类型 | 说明 |
|---|---|---|
{{- range .Columns}}
| {{.DataColumn}} | {{.DataType}} | {{.Title}} |
{{- end}}
```

**2. 注册到 manifest.yaml**：

```yaml
table:
  # ... 已有模板 ...
  - tpl: docs/readme.md.tpl
    out: docs/{app.name}/{module}.md
```

**3. 生成**：

```bash
codectl gen --tables * --overwrite
```

### 示例 4：添加新的应用级模板

**目标**：生成一个 `Makefile`。

**1. 创建模板** `.template/app/golang/Makefile.tpl`：

```makefile
# gen by codectl
# @author {{.App.Author}}

.PHONY: build run test

build:
	go build -o bin/{{.App.Name}} ./server

run:
	go run ./server -c server/app-dev.yaml

test:
	go test ./...
```

**2. 注册到 manifest.yaml**：

```yaml
app:
  # ... 已有模板 ...
  - tpl: golang/Makefile.tpl
    out: Makefile
```

**3. 生成**：

```bash
codectl gen --tables * --overwrite
```

### 示例 5：修改字段类型映射

**目标**：把 `DATETIME` 映射为 `time.Time` 而不是 `types.DateTime`。

**修改 manifest.yaml**：

```yaml
mapper:
  mysql-golang:
    DATETIME: time.Time    # 原来是 types.DateTime
    TIMESTAMP: time.Time   # 原来是 types.DateTime
```

**重新生成**：

```bash
codectl gen --tables * --overwrite
```

> 注意：模板中的 import 逻辑通常根据 `.Types` 列表条件引入，修改映射后可能需要同步调整模板中的 import 判断。

---

## 模板变量参考

### 应用级模板可用的变量

| 变量 | 类型 | 说明 | 示例 |
|---|---|---|---|
| `.App` | App | 应用信息对象 | - |
| `.App.Name` | string | 应用名称 | `testapp` |
| `.App.Title` | string | 应用标题 | `测试应用` |
| `.App.Author` | string | 作者 | `winlion` |
| `.App.Package` | string | go 包名 | `github.com/myorg/testapp` |
| `.App.Dsn` | string | 数据库连接串 | `mysql://...` |
| `.App.Prefix` | string | 表名前缀 | `sys_` |
| `.App.DbType` | string | 数据库类型 | `mysql` |
| `.App.Lang` | string | 语言 | `golang` |
| `.App.Dirsave` | string | 输出目录 | `./` |
| `.Package` | string | go 包名（同 `.App.Package`） | - |

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

遍历 `.Columns` 时，每列可访问：

| 字段 | 类型 | 说明 | 示例 |
|---|---|---|---|
| `.DataColumn` | string | 列名（原始） | `user_name` |
| `.DataType` | string | 映射后的数据类型 | `string`、`types.DateTime` |
| `.DataSize` | int | 字段长度 | `255` |
| `.Title` | string | 列注释 | `用户名` |
| `.IsPrimaryKey` | bool | 是否主键 | `true` |
| `.IsIndex` | bool | 是否索引 | `false` |
| `.AutoIncrement` | bool | 是否自增 | `true` |
| `.SuportCreate` | bool | 是否参与创建 | `true` |
| `.RawData` | ColumnInfo | 原始数据库信息 | - |
| `.RawData.ColumnName` | string | 数据库原始列名 | `user_name` |
| `.RawData.DataType` | string | 数据库原始类型 | `VARCHAR(255)` |

---

## 模板函数参考

### 字符串函数

| 函数 | 说明 | 输入 → 输出示例 |
|---|---|---|
| `ucfirst` | 首字母大写 | `user` → `User` |
| `lcfirst` | 首字母小写 | `User` → `user` |
| `camel` | 下划线转小驼峰 | `user_name` → `userName` |
| `upercamel` | 下划线转大驼峰 | `user_name` → `UserName` |
| `lower` | 转小写 | `TestApp` → `testapp` |
| `upper` | 转大写 | `testapp` → `TESTAPP` |
| `js` | 类型转 TypeScript 类型 | `User` → `IUserRecord` |

### 集合函数

| 函数 | 说明 | 示例 |
|---|---|---|
| `has` | 切片是否包含某元素 | `{{has .Types "time.Time"}}` |
| `contains` | 字符串是否包含子串 | `{{contains .DataType "int"}}` |

### 数据库函数

| 函数 | 说明 | 示例 |
|---|---|---|
| `mysqltogorm` | MySQL 类型转 GORM tag | `VARCHAR(255)` → `string` |

### 使用示例

```
{{- /* 模块名大驼峰 */ -}}
{{.Module|ucfirst}}              → UserOrder

{{- /* 列名小驼峰 */ -}}
{{$v.DataColumn|camel}}          → userName

{{- /* 条件引入 import */ -}}
{{if has .Types "time.Time"}}
	"time"
{{end}}

{{- /* 根据主键类型生成不同代码 */ -}}
{{if eq .Primary.DataType "string"}}
	pkId, _ := wraper.MuxStringVar(req, "pkId", "")
{{else}}
	pkId, _ := wraper.MuxIntVar(req, "pkId", int32(0))
{{end}}

{{- /* 遍历列生成结构体字段 */ -}}
{{- range $i,$v := .Columns}}
	{{$v.DataColumn|upercamel}} {{$v.DataType}} `json:"{{$v.DataColumn|camel}}"`
{{- end}}
```

---

## 常见问题

### Q: 修改模板后如何生效？

模板修改后需要重新执行 `codectl gen`。建议先用 `--dry-run` 预览：

```bash
codectl gen --tables sys_user --dry-run
```

确认无误后再加 `--overwrite` 生成：

```bash
codectl gen --tables sys_user --overwrite
```

### Q: 生成的文件被覆盖了怎么办？

`--overwrite` 会在覆盖前自动备份到 `.bak/<时间戳>/` 目录：

```
./
├── .bak/
│   └── 20260820120000/
│       └── server/internal/app/testapp/model/user.model.gen.go
└── server/
```

### Q: 如何只生成某一层（如只生成 model）？

codectl 不提供按层过滤的 flag（因为不同模板的"层"概念不统一）。可以通过以下方式实现：

1. **临时修改 manifest.yaml**：只保留 `table` 中需要的条目
2. **复制一份精简 manifest**：创建 `manifest-model.yaml`，只含 model 条目

### Q: 如何为不同的表使用不同模板？

目前一次 `codectl gen` 只能用一套模板。可以通过多次执行实现：

```bash
# 第一批表用默认模板
codectl gen --tables sys_* --overwrite

# 第二批表用自定义模板
codectl gen --tables biz_* --template ./custom-template --overwrite
```

### Q: 配置文件在哪里？

- **全局配置**：`~/.codectl/config.yaml`
- **自定义配置**：`codectl -c myconfig.yaml gen`
- **codectl 数据库**：`~/.codectl/codectl.db`（存储模板注册信息）

### Q: 如何调试模板渲染？

使用 `--dry-run` 预览渲染结果。如果模板语法错误，codectl 会报错并指出错误位置。

### Q: 支持哪些数据库？

目前支持 MySQL。数据库连接串格式：

```
mysql://[username]:[password]@tcp([host]:[port])/[dbname]?charset=utf8mb4&parseTime=true&loc=Local
```

### Q: 如何修改项目目录名（如 server → backend）？

codectl 生成的目录结构由 `manifest.yaml` 的 `out` 字段决定。修改输出路径即可重命名目录。

**场景：把 `server/` 改成 `backend/`，`web/` 改成 `website/`**

**第一步：修改 manifest.yaml 的输出路径**

```yaml
# .template/manifest.yaml

app:
  - tpl: golang/init.gen.go.tpl
    out: backend/internal/app/{app.name}/init.gen.go      # server/ → backend/
  - tpl: golang/conf.go.tpl
    out: backend/internal/conf/conf.go                   # server/ → backend/
  - tpl: golang/api.go.tpl
    out: backend/api/api.go                              # server/ → backend/
  - tpl: golang/app-dev.yaml.tpl
    out: backend/app-dev.yaml                            # server/ → backend/
  - tpl: golang/app-prod.yaml.tpl
    out: backend/app-prod.yaml                           # server/ → backend/
  - tpl: golang/logic.gen.go.tpl
    out: backend/internal/app/{app.name}/logic/logic.gen.go  # server/ → backend/
  # ... 其他 app 级模板同理

table:
  - tpl: golang/model.gen.go.tpl
    out: backend/internal/app/{app.name}/model/{module}.model.gen.go  # server/ → backend/
  - tpl: golang/rest.gen.go.tpl
    out: backend/internal/app/{app.name}/rest/{module}.gen.go         # server/ → backend/
  - tpl: golang/vo.gen.go.tpl
    out: backend/internal/app/{app.name}/vo/{module}.vo.gen.go         # server/ → backend/
  - tpl: web/api.ts.tpl
    out: website/src/api/{app.name}/{module}.ts          # web/ → website/
  - tpl: web/model.vue.tpl
    out: website/src/views/{app.name}/{module}/index.vue  # web/ → website/
  # ... 其他 table 级模板同理
```

**第二步：重命名已有的项目目录（如果代码已经生成过）**

```bash
# 在项目根目录执行
mv server backend
mv web website
```

**第三步：重新生成代码**

```bash
codectl gen --tables * --overwrite
```

**第四步：修改启动命令**

原来：
```bash
cd server && go run . -c app-dev.yaml
```

现在：
```bash
cd backend && go run . -c app-dev.yaml
```

> **提示**：如果模板中硬编码了 `server/` 路径（比如 conf.go.tpl 中的 import 路径或配置文件路径），也需要同步检查模板文件内容。可用以下命令检查：
>
> ```bash
> # 搜索模板中硬编码的 server 字符串
> grep -r "server/" .template/
> ```

### Q: 如何修改前端目录名（如 web → website）？

见上一问，把 `manifest.yaml` 中所有 `web/` 开头的 `out` 路径改成 `website/` 即可。

**完整修改对照：**

| 模板路径 out 字段 | 修改前 | 修改后 |
|---|---|---|
| `web/src/api/{app.name}/index.ts` | `web/src/...` | `website/src/api/{app.name}/index.ts` |
| `web/src/api/types.ts` | `web/src/...` | `website/src/api/types.ts` |
| `web/src/api/{app.name}/{module}.ts` | `web/src/...` | `website/src/api/{app.name}/{module}.ts` |
| `web/src/views/{app.name}/{module}/index.vue` | `web/src/...` | `website/src/views/{app.name}/{module}/index.vue` |

然后重命名目录 + 重新生成：

```bash
mv web website
codectl gen --tables * --overwrite
```

### Q: 如何只生成某张表的代码？

```bash
# 精确匹配
codectl gen --tables sys_user

# 通配符匹配
codectl gen --tables sys_*

# 多张表
codectl gen --tables sys_user,sys_role,sys_menu
```

### Q: 如何排除某些表？

```bash
# 生成所有 sys_ 表,但排除以 _log 结尾的
codectl gen --tables sys_* --exclude sys_*_log
```

### Q: 如何为不同的表使用不同模板？

```bash
# sys_ 表用默认模板
codectl gen --tables sys_* --overwrite

# biz_ 表用自定义模板
codectl gen --tables biz_* --template ./custom-template --overwrite
```

### Q: 生成的代码有编译错误怎么办？

1. 检查 `manifest.yaml` 中的字段类型映射是否正确
2. 检查模板中的 import 逻辑是否完整
3. 用 `--dry-run` 预览生成的代码
4. 常见问题：
   - 缺少 import：模板中的 `{{if has .Types "xxx"}}` 条件不匹配
   - 类型不匹配：`mapper` 中映射的类型与模板期望的不一致
   - 包名错误：检查 `--package` 参数或配置文件中的 `package` 字段

### Q: 如何更新已生成表的代码（数据库结构变更后）？

```bash
# 1. 先预览变更
codectl gen --tables sys_user --dry-run

# 2. 确认后覆盖生成（会自动备份旧文件）
codectl gen --tables sys_user --overwrite
```

`--overwrite` 会在覆盖前备份到 `.bak/<时间戳>/` 目录。

> **注意**：只覆盖 `.gen.go` 文件，用户编写的 `.ext.go` 文件不受影响。

### Q: 如何在多个项目中复用模板？

**方式 1：注册到 codectl 模板库**

```bash
# 注册本地模板
codectl tpl add /path/to/your-template

# 注册 git 模板
codectl tpl add https://github.com/myorg/my-template.git

# 查看已注册的模板
codectl tpl list

# 使用
codectl gen --template my-template --tables *
```

**方式 2：直接指定模板路径**

```bash
codectl gen --template /path/to/your-template --tables *
```

**方式 3：通过 git 地址**

```bash
codectl gen --template https://github.com/myorg/my-template.git --tables *
```

### Q: 配置文件在哪里？

- **全局配置**：`~/.codectl/config.yaml`（codectl set 缓存的参数）
- **自定义配置**：`codectl -c myconfig.yaml gen`
- **codectl 数据库**：`~/.codectl/codectl.db`（存储模板注册信息）
- **项目配置**：`backend/app-dev.yaml`（生成的应用配置）

### Q: 如何调试模板渲染？

```bash
# 预览单张表的渲染结果
codectl gen --tables sys_user --dry-run
```

如果模板语法错误，codectl 会报错并指出错误位置。常见错误：

- `template: syntax error` — 模板语法错误，检查 `{{}}` 是否闭合
- `undefined function` — 使用了不存在的模板函数，参考[模板函数参考](#模板函数参考)
- `no entry for key` — manifest.yaml 中 `tpl` 指定的文件不存在

---

## 作者

**微信：** huwinlion

**团队：** 湖南云立方智能科技有限公司
