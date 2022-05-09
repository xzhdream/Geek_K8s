FROM golang:1.17 AS builder
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /build
COPY . .
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM scratch
COPY --from=builder /build/httpserver /
EXPOSE 80
ENTRYPOINT ["/httpserver"]



