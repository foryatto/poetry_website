# 中华诗词网站

基于`Go`语言和`Vue 3`构建的前后端分离项目。( A poetry website created by Go and Vue. )

中华诗词收录了上至先秦，下至当代共计五万余首诗词，完全免费并开放所有诗词内容和程序源代码，您可自由复制、修改、传播诗词内容和程序源码。

# 特色

- 前端采用`Vue 3`实现响应式布局，支持手机端，PC端访问；
- 后端采用`GoFrame v2`框架，提供`REST API`接口；
- 数据库使用`SQLite3`；
- 可使用`docker`打包并生成网站；(todo)

# 运行
```bash
go mod tidy
go run main.go
```
本地访问: http://localhost:8080

接口文档: http://localhost:8080/swagger

# Docker部署

todo.

# 示例

![pc](https://github.com/foryatto/poetry/blob/main/example/pc.png)

![mobile](https://github.com/foryatto/poetry/blob/main/example/mobile.png)

![swagger](https://github.com/foryatto/poetry/blob/main/example/swagger.png)



