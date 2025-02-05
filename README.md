# 11_4_micro_captcha

## 项目简介

`11_4_micro_captcha` 是一个基于 Go 语言开发的微服务项目，主要功能包括用户认证、验证码生成与验证、短信验证码发送等。项目使用了 Gin 框架作为 Web 服务器，使用 Redis 作为会话存储，并通过 Consul 进行服务注册与发现。

## 目录结构

```
11_4_micro_captcha/
├── README.md
├── server/
│   └── user/
│       ├── handler/
│       ├── model/
│       ├── proto/
│       └── main.go
└── web/
    ├── controller/
    ├── static/
    ├── view/
    └── main.go
```

## 主要功能

- 用户认证
- 验证码生成与验证
- 短信验证码发送
- 用户信息获取
- 会话管理

## 安装与运行

### 前置条件

- Go 1.16 及以上版本
- Redis
- Consul

### 安装依赖

在项目根目录下运行以下命令安装依赖：

```sh
go mod tidy
```

### 配置 Redis

确保 Redis 服务正在运行，并在 `web/main.go` 中配置 Redis 地址：

```go
store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
```

### 配置 Consul

确保 Consul 服务正在运行，并在 `server/user/main.go` 中配置 Consul 地址：

```go
consulReg := consul.NewRegistry()
```

### 运行 Web 服务

在 `web` 目录下运行以下命令启动 Web 服务：

```sh
go run main.go
```

Web 服务将监听 `8080` 端口。

### 运行用户服务

在 `server/user` 目录下运行以下命令启动用户服务：

```sh
go run main.go
```

用户服务将监听 `12342` 端口。

## API 接口

### Web 服务接口

- `GET /home` - 首页路由，需要验证登录
- `GET /login` - 登录页面路由
- `GET /user` - 用户页面路由
- `GET /captcha/:uuid` - 获取验证码图片
- `GET /api/smscode/:phoneNumber` - 发送短信验证码
- `GET /getUserInfo` - 获取用户信息
- `GET /deleteSession` - 删除会话
- `POST /form` - 表单验证

### 用户服务接口

- `POST /user` - 用户注册
- `POST /user/login` - 用户登录

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助改进这个项目。

## 许可证

本项目使用 MIT 许可证，详情请参阅 LICENSE 文件。
