FROM arm64v8/golang:alpine as ssh_bin

ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64

RUN apk update && apk add git

RUN go get github.com/yudai/gotty

FROM arm64v8/python

COPY --from=ssh_bin /go/bin/gotty /usr/local/bin/gotty

CMD ["gotty", "-w", "/bin/sh"]