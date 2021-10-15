FROM golang:1.16.0-alpine3.13 AS builder
RUN go env -w GO111MODULE=auto \
  && go env -w GOPROXY=https://goproxy.cn,direct  \
  && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && cat /etc/apk/repositories \
  && apk add --no-cache bash git openssh 
WORKDIR /build
COPY ./ .
RUN go build -o app main.go 

FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && cat /etc/apk/repositories
WORKDIR /data
RUN apk add --no-cache tzdata curl
COPY --from=builder /build/app /usr/bin/go-aliyundriver-webdav
RUN chmod +x /usr/bin/go-aliyundriver-webdav
VOLUME /data
ENTRYPOINT ["/usr/bin/go-aliyundrive-webdav"]
