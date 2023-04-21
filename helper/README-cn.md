[English](https://github.com/niubir/area-service/blob/main/helper/README-en.md) | 简体中文

## 用法

### 开始使用

1. Go

```sh
go get -u github.com/niubir/area-service
go mod tidy
go run main.go
```

2. Docker

```sh
docker pull xxxx/niubir/area-service
docker run xxxx/niubir/area-service
```

### http 文档

```
http://localhost:10011/swagger/index.html
```

### grpc 文档

暂无

## 配置选项

> config/config.yml

| 选项 | 默认值 | 描述 |
| - | - | - |
| mode | PRODUCTION | 系统模式(PRODUCTION,TEST,DEVELOPMENT) |
| debug | false | 是否开启debug |
| amapKey | - | [高德地图 App](https://console.amap.com/dev/key/app) |
| autoFreshTime | - | 每天几点执行自动刷新区域信息, 例如:00:00:00 |
