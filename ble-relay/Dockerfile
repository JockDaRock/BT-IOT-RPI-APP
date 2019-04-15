FROM arm64v8/golang:alpine as ssh_bin

ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64

RUN apk update && apk add git

RUN go get github.com/yudai/gotty
RUN go get github.com/julienschmidt/httprouter
#RUN go get github.com/jteeuwen/go-bindata/...
#RUN go get github.com/tools/godep

RUN echo $PWD

RUN wget https://github.com/yudai/gotty/releases/download/v2.0.0-alpha.3/gotty_2.0.0-alpha.3_linux_arm.tar.gz && tar -xvzf gotty_2.0.0-alpha.3_linux_arm.tar.gz && chmod +x gotty

COPY getSSH.go .

RUN go build -o getSSH getSSH.go

FROM arm64v8/python

ENV INITSYSTEM on

COPY --from=ssh_bin /go/gotty /usr/local/bin/gotty
COPY --from=ssh_bin /go/getSSH /usr/local/bin/getSSH

COPY loop.sh .

RUN chmod +x loop.sh

CMD ["./loop.sh"]