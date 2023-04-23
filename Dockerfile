FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /area-service

EXPOSE 10011

CMD ["/area-service"]