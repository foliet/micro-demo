# Title

## 目录结构

```tree
├─common
│  └─errorx
└─service
    ├─account
    │  ├─model
    │  │  └─sql
    │  └─rpc
    │      ├─account
    │      ├─etc
    │      ├─internal
    │      │  ├─config
    │      │  ├─logic
    │      │  ├─server
    │      │  └─svc
    │      └─pb
    ├─gateway
    │  └─api
    │      ├─doc
    │      │  ├─account
    │      │  └─price
    │      ├─etc
    │      └─internal
    │          ├─config
    │          ├─handler
    │          │  └─account
    │          ├─logic
    │          │  └─account
    │          ├─svc
    │          └─types
    └─price
        ├─cron
        ├─model
        │  └─sql
        └─rpc
```
common为通用公共库

service中每一个文件夹都是一个微服务

对任意一个微服务，可能有以下文件夹:
 - model为数据库模型
 - api为http服务
 - rpc为rpc服务
 - cron为定时任务
 - rmq为消息队列任务
 - scripts为手动脚本