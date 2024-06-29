# 校园社区项目
ee
# 还在制作中。。。。。。

# 项目主要依赖：
 what can i say
 daoqiu
**Golang V1.18**

- Gin
- Gorm
- mysql
- go-micro
- protobuf
- grpc
- amqp
- ini
- hystrix
- jwt-go
- crypto





## 项目结构


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

# 项目结构
## 1.micro_todolist 项目总体
```
micro-todolist/
├── app                   // 各个微服务
│   ├── gateway           // 网关
│   ├── task              // 任务模块微服务
│   └── user              // 用户模块微服务
├── bin                   // 编译后的二进制文件模块
├── config                // 配置文件
├── consts                // 定义的常量
├── doc                   // 接口文档
├── idl                   // protoc文件
│   └── pb                // 放置生成的pb文件
├── logs                  // 放置打印日志模块
├── pkg                   // 各种包
│   ├── ctl               // 用户操作
│   ├── e                 // 统一错误状态码
│   ├── logger            // 日志
│   └── util              // 各种工具、JWT等等..
└── types                 // 定义各种结构体
```

## 2.gateway 网关部分
```
gateway/
├── cmd                   // 启动入口
├── http                  // HTTP请求头
├── handler               // 视图层
├── logs                  // 放置打印日志模块
├── middleware            // 中间件
├── router                // http 路由模块
├── rpc                   // rpc 调用
└── wrappers              // 熔断
```

## 3.user && task 用户与任务模块
```
task/
├── cmd                   // 启动入口
├── service               // 业务服务
├── repository            // 持久层
│    ├── db               // 视图层
│    │    ├── dao         // 对数据库进行操作
│    │    └── model       // 定义数据库的模型
│    └── mq               // 放置 mq
├── script                // 监听 mq 的脚本
└── service               // 服务
```


# 4.test git push dao qiu ee
