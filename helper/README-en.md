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


3. http
> swgger http://localhost:10011/swagger/index.html
>
> client https://github.com/niubir/area-service-client

4. grpc
> proto https://github.com/niubir/area-service/blob/main/grpc/area.proto
>
> client https://github.com/niubir/area-service-client


## Configuration option

> config/config.yml

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
