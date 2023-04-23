[English](https://github.com/niubir/area-service/blob/main/helper/README-en.md) | 简体中文

# 什么是区域服务

区域服务是行政区域代码服务.

基于[高德地图](https://console.amap.com)实现.

# 快速开始

## 配置

[config.yml](https://github.com/niubir/area-service/blob/main/config/config.yml)
```
mode: TEST
debug: true
amapKey: <set you amap key>
autoFreshTime: 00:00:00
http:
  enable: true
  port: 10011
grpc:
  enable: true
  port: 10012
```

## Golang运行

```sh
git clone https://github.com/niubir/area-service && cd area-service
CONFIG_DIR='you config.yml path' go run main.go
```

## Docker运行

```sh
docker pull niubir/area-service:latest
docker run -p 10011:10011 -p 10012:10012 -v <you config.yml path>/config:/config -d niubir/area-service:latest
```

# 使用方式

## http

```
swgger http://localhost:10011/swagger/index.html
```

## grpc

```
proto https://github.com/niubir/area-service/blob/main/grpc/area.proto
```

## cli

```
git clone https://github.com/niubir/area-service-client
```

# 配置选项([config.yml](https://github.com/niubir/area-service/blob/main/config/config.yml))

| Option | Default | Description |
| - | - | - |
| mode | PRODUCTION | system mode(PRODUCTION,TEST,DEVELOPMENT) |
| debug | false | is debug enabled |
| amapKey | - | [Amap App](https://console.amap.com/dev/key/app) |
| autoFreshTime | - | at what time does the automatic fresh of area every day, for example:00:00:00 |
| http.enable | true | enable http |
| http.port | 10011 | http port |
| grpc.enable | true | enable grpc |
| grpc.port | 10012 | grpc port |

# 环境变量选项
| Option | Default | Description |
| - | - | - |
| CONFIG_DIR | /config | config dir |
| CONFIG_FILENAME | config.yml | config filename |
| LOG_DIR | /logs | log dir |
| AREA_DIR | /data | area dir |
