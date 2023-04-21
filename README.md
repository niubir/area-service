English | [简体中文](https://github.com/niubir/area-service/blob/main/helper/README-cn.md)

## Usage

### Start using it

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

### http docs

```
http://localhost:10011/swagger/index.html
```

### grpc docs

TODO

## Configuration option

> config/config.yml

| Option | Default | Description |
| - | - | - |
| mode | PRODUCTION | system mode(PRODUCTION,TEST,DEVELOPMENT) |
| debug | false | is debug enabled |
| amapKey | - | [Amap App](https://console.amap.com/dev/key/app) |
| autoFreshTime | - | at what time does the automatic fresh of area every day, for example:00:00:00 |
