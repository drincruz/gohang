FROM golang:1.18

ENV GO111MODULE=on
WORKDIR /go/src/app

COPY . /go/src/app

RUN env GOOS=linux go build gohang.go
RUN go install .

ENTRYPOINT /go/src/app/gohang

EXPOSE 5555
