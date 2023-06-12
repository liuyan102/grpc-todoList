# grpc-todoList
采用gin、gorm、grpc、etcd、mysql、docker实现备忘录微服务demo


| 文件                              | 说明              |
|---------------------------------|-----------------|
| api-gateway                     | 路由接口服务          |
| api-gateway/cmd                 | main函数程序启动入口    |
| api-gateway/config              | 配置文件            |
| api-gateway/discovery           | etcd服务发现        |
| api-gateway/internal            | 程序主体            |
| api-gateway/internal/handler    | controller层     |
| api-gateway/internal/service    | 业务逻辑层           |
| api-gateway/internal/service/pb | protoBuf文件      |
| api-gateway/middleware          | 中间件             |
| api-gateway/pkg                 | 工具类             |
| api-gateway/routes              | router路由        |
| data                            | 用于docker创建mysql |
| task                            | 备忘录模块           |
| task/cmd                        | main函数程序启动入口    |
| task/config                     | 配置文件            |
| task/discovery                  | etcd服务发现        |
| task/internal                   | 程序主体            |
| task/internal/handler           | controller层     |
| task/internal/repository        | 数据库实体类          |
| task/internal/service           | 业务逻辑层           |
| task/internal/service/pb        | protoBuf文件      |
| task/pkg                        | 工具类             |
| user                            | 用户模块            |
| user/cmd                        | main函数程序启动入口    |
| user/config                     | 配置文件            |
| user/discovery                  | etcd服务发现        |
| user/internal                   | 程序主体            |
| user/internal/handler           | controller层     |
| user/internal/repository        | 数据库实体类          |
| user/internal/service           | 业务逻辑层           |
| user/internal/service/pb        | protoBuf文件      |
| user/pkg                        | 工具类             |


> docker-compose.yml 创建了五个docker容器：
> 1. mysql容器
> 2. etcd容器
> 3. gateway路由容器
> 4. user模块容器
> 5. task模块容器
> > 程序启动时需将各个模块中的配置文件中的IP地址改为容器名如：127.0.0.1：3306 -> mysql:3306
>
> 程序启动命令：docker-compose up