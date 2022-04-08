FROM golang:1.17-alpine3.15 AS builder

COPY . /grpc/
WORKDIR /grpc/

RUN go mod download
RUN go build -o ./bin/app cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /grpc/bin/app .
COPY --from=builder /grpc/configs configs/
COPY --from=builder /grpc/migrations migrations/

EXPOSE 50082

CMD ["./app"]