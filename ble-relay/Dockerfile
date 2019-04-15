FROM arm64v8/golang:alpine as ssh_bin

ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64

RUN apk update && apk add git make

RUN go get github.com/yudai/gotty
RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/jteeuwen/go-bindata/...
RUN go get github.com/tools/godep

RUN echo $PWD

WORKDIR /go/github.com/yudai/gotty

RUN go build -o gotty main.go

COPY getSSH.go .

RUN go build -o getSSH getSSH.go

FROM arm64v8/python

COPY --from=ssh_bin /go/bin/gotty /usr/local/bin/gotty
COPY --from=ssh_bin /go/bin/gotty /usr/local/bin/gotty

CMD ["gotty", "-w", "/bin/sh"]