# step1 安装codectl

```bash
go install github.com/turingdance/codectl
```

# 初始化项目

```bash
codectl init appname -d  /path/to/save
```
# 配置项目
本项目以farmbase为列子
```bash
cd server
codectl set

appname:the name of app,must be english ,will use as service name
default:  farmbase
>
author: author of this app
default:  winlion
>
title: the title of project
default:  智慧农场
>
dsn: required, eg:appdata.db or mysql://user:password@tcp(host:port)/db?quest
default:  admfarmbase:admfarmbase@123@tcp(127.0.0.1:3306)/farmbase
>mysql://admfarmbase:rootme@123@tcp(localhost:3306)/farmbase?charset=utf8mb4&parseTime=true&loc=Local 
prefix:prefix of table,such as sys_, or kf_>
>fb_
package:,eg turingdance.com/turing/app>
>turingdance.com/turing/app 
2025/10/29 16:51:09 std.go:49: update default file C:\Users\Administrator\AppData\Roaming/.codectl/config.yaml

```

# 配置APP
## 生成APP代码
```base
codectl reverse
```
## 路由处理
注意其中 $(app) 换成应用名称 
如
codectl router  -s ./internal/app/farmbase/rest
```bash
#cd server
#codectl router  -s ./internal/app/$(app)/rest
```

## 开启路由支持

在文件 api/cmd/app.go
60行左右,参考
```golang
// ecls.Initialize(appConf)
// httpserver.Handle("/ecls/", ecls.CreateRouter("/ecls/"), auth.NewAuthorize(appConf.Auth.WhiteList).Handle)
// 添加配置如下
farmbase.Initialize(appConf)
httpserver.Handle("/farmbase/", farmbase.CreateRouter("/farmbase/"), auth.NewAuthorize(appConf.Auth.WhiteList).Handle)
//顶部需要导入相关的包
import "turingdance.com/turing/internal/app/farmbase"
```
## 添加数据库配置文件

app-dev.yaml/app-prod.yaml/ 中参考sys 添加数据库配置文件

```yaml
farmbaseConf:
  dsn: mysql://[username]:[password]]@tcp([ip]:[port])/[dbname]?charset=utf8mb4&parseTime=true&loc=Local
  prefix: fb_
  engin: InnoDB
  charset: utf8mb4
  loglevel: 6
  connMaxIdleTime: 3600s
  connMaxLifeTime: 80s
  conMaxOpen: 20
  conMaxIdle: 2
```

## 启动
```
go run . -c app-dev.yaml
```

