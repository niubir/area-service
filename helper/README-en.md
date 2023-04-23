English | [简体中文](https://github.com/niubir/area-service/blob/main/helper/README-cn.md)

# What is area service

Area service is administrative region code service.

Based on [Amap](https://console.amap.com) Implementation.

# Quick start

## Configure

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

## Run by Golang

```sh
git clone https://github.com/niubir/area-service && cd area-service
CONFIG_DIR='you config.yml path' go run main.go
```

## Run by Docker

```sh
docker pull niubir/area-service:latest
docker run -p 10011:10011 -p 10012:10012 -v <you config.yml path>/config:/config -d niubir/area-service:latest
```

# Usage

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

# Configuration option([config.yml](https://github.com/niubir/area-service/blob/main/config/config.yml))

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

# Env option
| Option | Default | Description |
| - | - | - |
| CONFIG_DIR | /config | config dir |
| CONFIG_FILENAME | config.yml | config filename |
| LOG_DIR | /logs | log dir |
| AREA_DIR | /data | area dir |
