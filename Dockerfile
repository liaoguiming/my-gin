FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /$GOPATH/src/liao/
COPY . /$GOPATH/src/liao/
COPY /docker/config.yaml ./config.yaml

#增加缺失的包，移除没用的包
RUN go mod tidy

EXPOSE 8088