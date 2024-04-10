# 校园社区项目

# 还在制作中。。。。。。

**此项目使用`Gin`+`Gorm` ，基于`RESTful API`实现的一个备忘录**。


## 项目主要依赖：

**Golang V1.15**

- Gin
- Gorm
- mysql
- redis
- ini
- jwt-go
- logrus
- go-swagger

## 项目结构

```shell
TodoList/
├── api
├── cmd
├── conf
├── consts
├── docs
├── middleware
├── pkg
│  ├── e
│  └── util
├── routes
├── repository
│  ├── cache
│  └── db
│     ├── dao
│     └── model
├── routes
├── service
└── types
```

- api : 用于定义接口函数,也就是controller层
- cmd : 程序启动
- conf : 用于存储配置文件
- middleware : 应用中间件
- pkg/e : 封装错误码
- pkg/logging : 日志打印
- pkg/util : 工具函数
- repository: 仓库放置所有存储
- repository/cache: 放置redis缓存
- repository/db: 持久层MySQL仓库
- repository/db/dao: 对db进行操作的dao层
- repository/db/model: 定义所有持久层数据库表结构的model层
- routes : 路由逻辑处理
- service : 接口函数的实现
- types : 放置所有的定义的结构体

## 配置文件
配置文件在conf/config.ini.example中，把.example去掉，然后根据自己的情况配置就好了

**conf/config.ini**
```ini
# debug开发模式,release生产模式
[service]
AppMode = debug
HttpPort = :3000
# 运行端口号 3000端口

[redis]
RedisDb = redis
RedisAddr = 
# redis ip地址和端口号
RedisPw = 
# redis 密码
RedisDbName = 2
# redis 名字

[mysql]
Db = mysql
DbHost =
# mysql ip地址
DbPort = 
# mysql 端口号
DbUser = 
# mysql 用户名
DbPassWord = 
# mysql 密码
DbName = 
# mysql 名字
```

# 