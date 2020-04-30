FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/hub/src/app
COPY . $GOPATH/src/hub/src/app
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]
