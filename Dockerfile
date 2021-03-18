FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/liao
COPY . /go/src/liao
RUN go build -o liao .
RUN cat config.yaml

EXPOSE 8088

ENTRYPOINT ["./liao"]
