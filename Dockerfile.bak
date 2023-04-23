# build golang application
FROM golang:latest as builder
WORKDIR /go/src/github.com/niubir/area-service
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# compress application
FROM alpine:latest as compresser
RUN apk --no-cache add ca-certificates
WORKDIR /apps
COPY --from=builder /go/src/github.com/niubir/area-service/app ./user-service
COPY --from=builder /go/src/github.com/niubir/area-service/docs ./docs
EXPOSE 10011
EXPOSE 10012

ENTRYPOINT ["./user-service"]